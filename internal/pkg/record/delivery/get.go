package recorddelivery

import (
	// "fmt"
	"fmt"
	"hesh/internal/pkg/domain"
	"hesh/internal/pkg/utils/filesaver"
	"hesh/internal/pkg/utils/log"
	"hesh/internal/pkg/utils/sanitizer"

	"io/ioutil"
	"mime/multipart"

	"strconv"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
)

func validImageExtenstions(files []*multipart.FileHeader) bool {
	availableExtensions := map[string]struct{}{"jpeg": {}, "png": {}, "jpg": {}}
	for i := range files {
		extension, haveExtension := filesaver.GetExtension(files[i])
		if !haveExtension {
			return false
		}
		_, is := availableExtensions[extension]
		if !is {
			return false
		}

	}
	return true
}

func readMultipartDataImages(r *http.Request) ([]domain.RecordImageInfo, int, error) {

	err := r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	formdata := r.MultipartForm
	//get the *fileheaders
	files := formdata.File["images"] // grab the filenames
	imageInfo := []domain.RecordImageInfo{}
	// filePaths := []string{}
	for i, _ := range files { // loop through the files one by one
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			// TODO: add mapping from error to http code
			return []domain.RecordImageInfo{}, http.StatusInternalServerError, domain.Err.ErrObj.InternalServer
		}

		if !validImageExtenstions(files) {
			return []domain.RecordImageInfo{}, http.StatusBadRequest, domain.Err.ErrObj.BadFileExtension
		}
		// TODO: parse tags
		// tags := make([]string, 0)
		// for j := range (r.Form["tags"])[i]{
		// 	fmt.Sprintf("%v", (r.Form["title"])[0])
		// 	tags = append(tags, fmt.Sprintf("%v", (r.Form["tags"])[i][j]))
		// }
		if err != nil {
			return []domain.RecordImageInfo{}, http.StatusBadRequest, domain.Err.ErrObj.BadInput
		}
		imageInfo = append(imageInfo, domain.RecordImageInfo{ImageName: filesaver.ExtractName(files[i].Filename), Tags: nil})
	}
	return imageInfo, http.StatusCreated, nil
}

// MEDIC
func (handler *RecordHandler) CreateMedicRecord(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// sessionId, err := sessions.CheckSession(r)
	// if err == domain.Err.ErrObj.UserNotLoggedIn {
	// 	http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusForbidden)
	// 	return
	// }
	queryParameter := r.URL.Query().Get("vk_user_id")
	userId, err := strconv.ParseUint(queryParameter, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	diaryId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	RecordCreateRequest := new(domain.MedicRecordCreateRequest)
	RecordCreateRequest.SetDefault()
	readedImages, httpCode, err := readMultipartDataImages(r)
	RecordCreateRequest.Images = readedImages
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), httpCode)
		w.WriteHeader(httpCode)
		return
	}
	// basicInfo := domain.RecordBasicInfo{}
	// TODO change to easyjson
	// err = json.Unmarshal(([]byte)((r.Form["basicInfo"])[0]), RecordCreateRequest.BasicInfo)

	RecordCreateRequest.BasicInfo.Title = fmt.Sprintf("%v", (r.Form["title"])[0])
	RecordCreateRequest.BasicInfo.Treatment = fmt.Sprintf("%v", (r.Form["treatment"])[0])
	RecordCreateRequest.BasicInfo.Recommendations = fmt.Sprintf("%v", (r.Form["recommendations"])[0])
	RecordCreateRequest.BasicInfo.Details = fmt.Sprintf("%v", (r.Form["details"])[0])
	// RecordCreateRequest.Diarisation = fmt.Sprintf("%v", (r.Form["diarisation"])[0])

	// diarisation := make([]string, 0)
	// for i := range audioResponse{

	// }
	// RecordCreateRequest.Auido

	sanitizer.SanitizeMedicRecordCreateRequest(RecordCreateRequest)
	es, err := handler.RecordUsecase.CreateMedicRecord(diaryId, userId, *RecordCreateRequest)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// We check if file with this name alredy exist in usecase
	imageNames := []string{}
	for i := range RecordCreateRequest.Images {
		imageNames = append(imageNames, RecordCreateRequest.Images[i].ImageName)
	}
	//TODO ser response image valuse in repository
	filesaver.SaveMultipartDataFiles(imageNames, r.MultipartForm.File["images"])

	es.ImageList = make([]domain.RecordImageInfo, 0)
	for i := range imageNames {
		es.ImageList = append(es.ImageList, domain.RecordImageInfo{ImageName: imageNames[i], Tags: nil})
	}
	out, err := easyjson.Marshal(es)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(out)
}

func (handler *RecordHandler) GetMedicRecord(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// sessionId, err := sessions.CheckSession(r)
	// if err == domain.Err.ErrObj.UserNotLoggedIn {
	// 	http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusForbidden)
	// 	return
	// }
	queryParameter := r.URL.Query().Get("vk_user_id")
	userId, err := strconv.ParseUint(queryParameter, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	recordId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	es, err := handler.RecordUsecase.GetMedicRecord(userId, recordId)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	out, err := easyjson.Marshal(es)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(out)
}

func (handler *RecordHandler) GetMedicRecordDiarisations(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// sessionId, err := sessions.CheckSession(r)
	// if err == domain.Err.ErrObj.UserNotLoggedIn {
	// 	http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusForbidden)
	// 	return
	// }
	queryParameter := r.URL.Query().Get("vk_user_id")
	userId, err := strconv.ParseUint(queryParameter, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	recordId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	es, err := handler.RecordUsecase.GetMedicRecordDiarisations(userId, recordId)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	out, err := easyjson.Marshal(es)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(out)
}

func (handler *RecordHandler) UpdateTextMedicRecord(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// sessionId, err := sessions.CheckSession(r)
	// if err == domain.Err.ErrObj.UserNotLoggedIn {
	// 	http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusForbidden)
	// 	return
	// }
	queryParameter := r.URL.Query().Get("vk_user_id")
	medicId, err := strconv.ParseUint(queryParameter, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	recordId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	MedicRecordUpdateTextRequest := new(domain.MedicRecordBasicInfo)
	// DiaryCreateRequest.SetDefault()

	err = easyjson.Unmarshal(b, MedicRecordUpdateTextRequest)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// if cast.IntToStr(sessionId) != EventCreatingRequest.UserId {
	// 	http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
	// 	w.WriteHeader(http.StatusBadRequest)
	// }

	sanitizer.SanitizeMedicRecordBasicInfo(MedicRecordUpdateTextRequest)

	es, err := handler.RecordUsecase.UpdateMedicRecordText(medicId, recordId, *MedicRecordUpdateTextRequest)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// We check if file with this name alredy exist in usecase
	out, err := easyjson.Marshal(es)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(out)
}

func (handler *RecordHandler) UpdateImageMedicRecord(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// sessionId, err := sessions.CheckSession(r)
	// if err == domain.Err.ErrObj.UserNotLoggedIn {
	// 	http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusForbidden)
	// 	return
	// }
	queryParameter := r.URL.Query().Get("vk_user_id")
	medicId, err := strconv.ParseUint(queryParameter, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	recordId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	request := new(domain.RecordUpdateImageRequest)
	images, httpCode, err := readMultipartDataImages(r)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), httpCode)
		w.WriteHeader(httpCode)
		return
	}
	request.Images = images
	sanitizer.SanitizeRecordImages(request)

	medicRecordUpdateResponse, err := handler.RecordUsecase.UpdateRecordImage(true, medicId, recordId, *request)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// We check if file with this name alredy exist in usecase
	imageNames := []string{}
	for i := range medicRecordUpdateResponse.Images {
		imageNames = append(imageNames, medicRecordUpdateResponse.Images[i].ImageName)
	}
	filesaver.SaveMultipartDataFiles(imageNames, r.MultipartForm.File["images"])

	out, err := easyjson.Marshal(medicRecordUpdateResponse)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(out)
}

func (handler *RecordHandler) DeleteMedicRecord(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// sessionId, err := sessions.CheckSession(r)
	// if err == domain.Err.ErrObj.UserNotLoggedIn {
	// 	http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusForbidden)
	// 	return
	// }
	queryParameter := r.URL.Query().Get("vk_user_id")
	medicId, err := strconv.ParseUint(queryParameter, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	recordId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	err = handler.RecordUsecase.DeleteMedicRecord(medicId, recordId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// PATIENT
func (handler *RecordHandler) CreatePatientRecord(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// sessionId, err := sessions.CheckSession(r)
	// if err == domain.Err.ErrObj.UserNotLoggedIn {
	// 	http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusForbidden)
	// 	return
	// }
	queryParameter := r.URL.Query().Get("vk_user_id")
	userId, err := strconv.ParseUint(queryParameter, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	diaryId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	PatientRecordCreateRequest := new(domain.PatientRecordCreateRequest)
	PatientRecordCreateRequest.SetDefault()
	readedImages, httpCode, err := readMultipartDataImages(r)
	PatientRecordCreateRequest.Images = readedImages
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), httpCode)
		w.WriteHeader(httpCode)
		return
	}
	// basicInfo := domain.RecordBasicInfo{}
	// TODO change to easyjson
	// err = json.Unmarshal(([]byte)((r.Form["basicInfo"])[0]), RecordCreateRequest.BasicInfo)

	PatientRecordCreateRequest.BasicInfo.Title = fmt.Sprintf("%v", (r.Form["title"])[0])
	PatientRecordCreateRequest.BasicInfo.Treatment = fmt.Sprintf("%v", (r.Form["treatment"])[0])
	PatientRecordCreateRequest.BasicInfo.Complaints = fmt.Sprintf("%v", (r.Form["complaints"])[0])
	PatientRecordCreateRequest.BasicInfo.Details = fmt.Sprintf("%v", (r.Form["details"])[0])

	sanitizer.SanitizePatientRecordCreateRequest(PatientRecordCreateRequest)
	es, err := handler.RecordUsecase.CreatePatientRecord(userId, diaryId, *PatientRecordCreateRequest)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// We check if file with this name alredy exist in usecase
	imageNames := []string{}
	for i := range PatientRecordCreateRequest.Images {
		imageNames = append(imageNames, PatientRecordCreateRequest.Images[i].ImageName)
	}
	//TODO ser response image valuse in repository
	filesaver.SaveMultipartDataFiles(imageNames, r.MultipartForm.File["images"])

	es.ImageList = make([]domain.RecordImageInfo, 0)
	for i := range imageNames {
		es.ImageList = append(es.ImageList, domain.RecordImageInfo{ImageName: imageNames[i], Tags: nil})
	}
	out, err := easyjson.Marshal(es)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(out)
}

func (handler *RecordHandler) GetPatientRecord(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// sessionId, err := sessions.CheckSession(r)
	// if err == domain.Err.ErrObj.UserNotLoggedIn {
	// 	http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusForbidden)
	// 	return
	// }
	queryParameter := r.URL.Query().Get("vk_user_id")
	userId, err := strconv.ParseUint(queryParameter, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	recordId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	es, err := handler.RecordUsecase.GetPatientRecord(userId, recordId)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	out, err := easyjson.Marshal(es)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(out)
}

func (handler *RecordHandler) UpdateTextPatientRecord(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// sessionId, err := sessions.CheckSession(r)
	// if err == domain.Err.ErrObj.UserNotLoggedIn {
	// 	http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusForbidden)
	// 	return
	// }
	queryParameter := r.URL.Query().Get("vk_user_id")
	patientId, err := strconv.ParseUint(queryParameter, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	recordId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	PatientRecordUpdateTextRequest := new(domain.PatientRecordBasicInfo)
	// DiaryCreateRequest.SetDefault()

	err = easyjson.Unmarshal(b, PatientRecordUpdateTextRequest)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// if cast.IntToStr(sessionId) != EventCreatingRequest.UserId {
	// 	http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
	// 	w.WriteHeader(http.StatusBadRequest)
	// }

	sanitizer.SanitizePatientRecordBasicInfo(PatientRecordUpdateTextRequest)

	es, err := handler.RecordUsecase.UpdatePatientRecordText(patientId, recordId, *PatientRecordUpdateTextRequest)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// We check if file with this name alredy exist in usecase
	out, err := easyjson.Marshal(es)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(out)
}

func (handler *RecordHandler) UpdateImagePatientRecord(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// sessionId, err := sessions.CheckSession(r)
	// if err == domain.Err.ErrObj.UserNotLoggedIn {
	// 	http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusForbidden)
	// 	return
	// }
	queryParameter := r.URL.Query().Get("vk_user_id")
	patientId, err := strconv.ParseUint(queryParameter, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	recordId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	request := new(domain.RecordUpdateImageRequest)
	images, httpCode, err := readMultipartDataImages(r)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), httpCode)
		w.WriteHeader(httpCode)
		return
	}
	request.Images = images
	sanitizer.SanitizeRecordImages(request)

	patientRecordUpdateResponse, err := handler.RecordUsecase.UpdateRecordImage(false, patientId, recordId, *request)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// We check if file with this name alredy exist in usecase
	imageNames := []string{}
	for i := range patientRecordUpdateResponse.Images {
		imageNames = append(imageNames, patientRecordUpdateResponse.Images[i].ImageName)
	}
	filesaver.SaveMultipartDataFiles(imageNames, r.MultipartForm.File["images"])

	out, err := easyjson.Marshal(patientRecordUpdateResponse)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(out)
}

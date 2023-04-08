package recorddelivery

import (
	// "fmt"
	"bytes"
	"context"
	"fmt"
	"hesh/internal/pkg/domain"
	mlsgrpc "hesh/internal/pkg/mlservices/delivery/grpc"
	"hesh/internal/pkg/utils/config"
	"hesh/internal/pkg/utils/filesaver"
	"hesh/internal/pkg/utils/log"
	"hesh/internal/pkg/utils/sanitizer"

	"io"
	"io/ioutil"
	"mime/multipart"
	"path/filepath"

	"strconv"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/mailru/easyjson"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getExtension(file *multipart.FileHeader) (string, bool) {
	k := len(file.Filename) - 1
	extension := ""
	for k != 0 {
		if k == 0 {
			return "", false
		}
		if (file.Filename)[k] == '.' {
			extension = (file.Filename)[k+1:]
		}
		k = k - 1
	}
	return extension, true
}

func validImageExtenstions(files []*multipart.FileHeader) bool {
	availableExtensions := map[string]struct{}{"jpeg": {}, "png": {}, "jpg": {}}
	for i := range files {
		extension, haveExtension := getExtension(files[i])
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

func validAudioExtenstions(files []*multipart.FileHeader) bool {
	availableExtensions := map[string]struct{}{"mp3": {}}
	for i := range files {
		extension, haveExtension := getExtension(files[i])
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

func extractNames(filePaths []string) (fileName []string) {
	imageNames := []string{}
	for i := range filePaths {
		imageNames = append(imageNames, extractName(filePaths[i]))
	}
	return imageNames
}

func extractName(filePath string) (fileName string) {
	i := len(filePath) - 1
	for i >= 0 {
		if filePath[i] == '/' || filePath[i] == '\\' {
			fileName = filePath[i+1:]
		}
		if i == 0 {
			fileName = filePath[i:]

		}
		i--
	}
	return fileName
}

func readMultipartDataImages(r *http.Request) ([]domain.RecordImageInfo, int,  error) {

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
		imageInfo = append(imageInfo, domain.RecordImageInfo{ImageName: extractName(files[i].Filename), Tags: nil})
	}
	return imageInfo, http.StatusCreated, nil
}

func readMultipartDataAudio(r *http.Request) ([]string, int,  error) {

	err := r.ParseMultipartForm(1 << 28) // maxMemory 256MB
	if err != nil {
		return nil, http.StatusBadRequest, err 
	}

	formdata := r.MultipartForm
	//get the *fileheaders
	files := formdata.File["audio"] // grab the filenames
	audioInfo := make([]string, 0)
	for i, _ := range files { // loop through the files one by one
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			// TODO: add mapping from error to http code
			return nil, http.StatusInternalServerError, domain.Err.ErrObj.InternalServer
		}

		if !validAudioExtenstions(files) {
			return nil, http.StatusBadRequest, domain.Err.ErrObj.BadFileExtension
		}
		// TODO: parse tags
		// tags := make([]string, 0)
		// for j := range (r.Form["tags"])[i]{
		// 	fmt.Sprintf("%v", (r.Form["title"])[0])
		// 	tags = append(tags, fmt.Sprintf("%v", (r.Form["tags"])[i][j]))
		// }
		if err != nil {
			return nil, http.StatusBadRequest, domain.Err.ErrObj.BadInput
		}
		audioInfo = append(audioInfo, extractName(files[i].Filename))
	}
	return audioInfo, http.StatusCreated, nil
}

func saveMultipartDataFiles(fileNames []string, fileHeaders []*multipart.FileHeader) (error, int) {
	// TODO: add mapping from error to http code
	for i, _ := range fileNames {
		file, err := fileHeaders[i].Open()
		defer file.Close()
		if err != nil {
			return domain.Err.ErrObj.InternalServer, http.StatusInternalServerError
		}
		extension := filepath.Ext(fileNames[i])
		nameWithouExtension := fileNames[i][:len(fileNames[i])-len(extension)]
		_, err = filesaver.UploadFile(file, "", config.DevConfigStore.LoadedFilesPath, nameWithouExtension, filepath.Ext(fileNames[i]))
		if err != nil {
			return domain.Err.ErrObj.InternalServer, http.StatusInternalServerError
		}
	}
	return nil, http.StatusCreated
}
func GetAudioSummarization(fileHeader []*multipart.FileHeader) (*mlsgrpc.DiarisationResponse, error) {
	if len(fileHeader) == 0 {
		return nil, nil
	}
	file, err := fileHeader[0].Open()
	defer file.Close()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Error(err)
		return nil, err

	}

	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error(err)
		return nil, err

	}
	defer conn.Close()

	client := mlsgrpc.NewDiarisationClient(conn)
	response, err := client.TranscribeAudio(context.Background(), &mlsgrpc.DiarisationRequest{
		Audio: buf.Bytes(),
	})
	if err != nil {
		log.Error(err)
		return nil, err

	}
	return response, nil
}

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
	
	readedAudio, httpCode, err := readMultipartDataAudio(r)
	RecordCreateRequest.Auido = readedAudio
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), httpCode)
		w.WriteHeader(httpCode)
		return
	}
	formdata := r.MultipartForm
	audioResponse, err := GetAudioSummarization(formdata.File["audio"] )
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// diarisation := make([]string, 0)
	// for i := range audioResponse{

	// }
	// RecordCreateRequest.Auido
	text := ""
	sanitizer.SanitizeMedicRecordCreateRequest(RecordCreateRequest)
	if audioResponse != nil {
		text = audioResponse.Text
	}
	es, err := handler.RecordUsecase.CreateMedicRecord(diaryId, userId, *RecordCreateRequest, text)
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
	saveMultipartDataFiles(imageNames, r.MultipartForm.File["images"])
	saveMultipartDataFiles(RecordCreateRequest.Auido, r.MultipartForm.File["audio"])
	
	es.ImageList = make([]domain.RecordImageInfo,0)
	for i := range (imageNames){
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
	
	request := new(domain.MedicRecordUpdateImageRequest)
	images, httpCode, err := readMultipartDataImages(r)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), httpCode)
		w.WriteHeader(httpCode)
		return
	}
	request.Images = images
	sanitizer.SanitizeMedicRecordImages(request)

	medicRecordUpdateResponse, err := handler.RecordUsecase.UpdateMedicRecordImage(medicId, recordId, *request)
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
	saveMultipartDataFiles(imageNames, r.MultipartForm.File["images"])

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

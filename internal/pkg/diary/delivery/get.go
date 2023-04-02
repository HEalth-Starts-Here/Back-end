package diarydelivery

import (
	"hesh/internal/pkg/domain"
	"mime/multipart"

	// "path/filepath"
	// "eventool/internal/pkg/sessions"
	// "hesh/internal/pkg/utils/cast"
	"hesh/internal/pkg/utils/cast"
	"hesh/internal/pkg/utils/config"
	"hesh/internal/pkg/utils/filesaver"
	"hesh/internal/pkg/utils/log"
	"hesh/internal/pkg/utils/sanitizer"

	"path/filepath"
	"strconv"

	// "strings"

	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mailru/easyjson"
	// "encoding/json"
	// "io"
	// "os"
)

func (handler *DiaryHandler) CreateDiary(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// sessionId, err := sessions.CheckSession(r)
	// if err == domain.Err.ErrObj.UserNotLoggedIn {
	// 	http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusForbidden)
	// 	return
	// }

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	queryParameter := r.URL.Query().Get("vk_user_id")
	medicId64, err := strconv.ParseUint(queryParameter, 10, 32)
	medicId := (uint32)(medicId64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	DiaryCreateRequest := new(domain.DiaryCreateRequest)
	DiaryCreateRequest.SetDefault()

	err = easyjson.Unmarshal(b, DiaryCreateRequest)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// if cast.IntToStr(sessionId) != EventCreatingRequest.UserId {
	// 	http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
	// 	w.WriteHeader(http.StatusBadRequest)
	// }

	sanitizer.SanitizeDiaryCreating(DiaryCreateRequest)

	es, err := handler.DiaryUsecase.CreateDiary(*DiaryCreateRequest, medicId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	out, err := easyjson.Marshal(es)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(out)
}

func (handler *DiaryHandler) LinkDiary(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// sessionId, err := sessions.CheckSession(r)
	// if err == domain.Err.ErrObj.UserNotLoggedIn {
	// 	http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusForbidden)
	// 	return
	// }

	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	queryParameter := r.URL.Query().Get("vk_user_id")
	medicId64, err := strconv.ParseUint(queryParameter, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	medicId := (uint32)(medicId64)

	params := mux.Vars(r)
	diaryId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// DiaryCreateRequest := new(domain.DiaryCreateRequest)
	// DiaryCreateRequest.SetDefault()

	// if cast.IntToStr(sessionId) != EventCreatingRequest.UserId {
	// 	http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
	// 	w.WriteHeader(http.StatusBadRequest)
	// }

	es, err := handler.DiaryUsecase.LinkDiary(diaryId, medicId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	out, err := easyjson.Marshal(es)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(out)
}

func (handler *DiaryHandler) DeleteDiary(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// sessionId, err := sessions.CheckSession(r)
	// if err == domain.Err.ErrObj.UserNotLoggedIn {
	// 	http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusForbidden)
	// 	return
	// }

	params := mux.Vars(r)
	diaryId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	err = handler.DiaryUsecase.DeleteDiary(diaryId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *DiaryHandler) GetDiary(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// categoryString := r.URL.Query().Get("category")
	// categories := strings.Split(categoryString, " ")

	diaryList, err := handler.DiaryUsecase.GetDiary()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	out, err := easyjson.Marshal(diaryList)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (handler *DiaryHandler) GetCertainDiary(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// userId, err := sessions.CheckSession(r);
	// if err == domain.Err.ErrObj.UserNotLoggedIn {
	// 	http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusBadRequest)
	// 	return
	// }

	params := mux.Vars(r)
	diaryId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	diary, err := handler.DiaryUsecase.GetCertainDiary(diaryId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	out, err := easyjson.Marshal(diary)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

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

func validExtenstions(files []*multipart.FileHeader) bool {
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

func readMultipartDataImages(r *http.Request) ([]domain.ImageInfoUsecase, error, int) {

	formdata := r.MultipartForm
	//get the *fileheaders
	files := formdata.File["images"] // grab the filenames
	imageInfo := []domain.ImageInfoUsecase{}
	// filePaths := []string{}
	for i, _ := range files { // loop through the files one by one
		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			// TODO: add mapping from error to http code
			return []domain.ImageInfoUsecase{}, domain.Err.ErrObj.InternalServer, http.StatusInternalServerError
		}

		if !validExtenstions(files) {
			return []domain.ImageInfoUsecase{}, domain.Err.ErrObj.BadFileExtension, http.StatusBadRequest
		}
		ar, err := cast.StringToFloat64((r.Form["areas"])[i])
		if err != nil {
			return []domain.ImageInfoUsecase{}, domain.Err.ErrObj.BadInput, http.StatusBadRequest
		}
		imageInfo = append(imageInfo, domain.ImageInfoUsecase{Name: extractName(files[i].Filename), Area: ar})
	}
	return imageInfo, nil, http.StatusCreated
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

func (handler *DiaryHandler) CreateRecord(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// sessionId, err := sessions.CheckSession(r)
	// if err == domain.Err.ErrObj.UserNotLoggedIn {
	// 	http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusForbidden)
	// 	return
	// }

	params := mux.Vars(r)
	diaryId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}

	err = r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		return
	}
	imageInfo, err, httpCode := readMultipartDataImages(r)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), httpCode)
		w.WriteHeader(httpCode)
		return
	}
	// imageInfo, err, httpCode := readAndSaveMultipartDataFiles(r)
	// if err != nil {
	// 	log.Error(err)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	w.WriteHeader(httpCode)
	// 	return
	// }
	RecordCreateRequest := new(domain.RecordCreateRequest)
	RecordCreateRequest.SetDefault()
	RecordCreateRequest.Title = fmt.Sprintf("%v", (r.Form["title"])[0])
	RecordCreateRequest.Description = fmt.Sprintf("%v", (r.Form["description"])[0])
	characteristicsRequest := [](*uint8){&RecordCreateRequest.Characteristics.Itching,
		&RecordCreateRequest.Characteristics.Pain,
		&RecordCreateRequest.Characteristics.Edema,
		&RecordCreateRequest.Characteristics.Redness,
		&RecordCreateRequest.Characteristics.Dryness,
		&RecordCreateRequest.Characteristics.Peeling}
	characteristics := []string{"itching", "pain", "edema", "redness", "dryness", "peeling"}
	for i := range characteristics {
		uint8value, err := strconv.Atoi(r.Form[characteristics[i]][0])
		if err != nil {
			log.Error(err)
			http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
			return
		}
		*characteristicsRequest[i] = uint8(uint8value)

	}

	sanitizer.SanitizeRecordCreating(RecordCreateRequest)

	//TODO: check if file with this name already esist
	es, err := handler.DiaryUsecase.CreateRecord(diaryId, *RecordCreateRequest, imageInfo)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// We check if file with this name alredy exist in usecase
	imageNames := []string{}
	for i := range imageInfo {
		imageNames = append(imageNames, imageInfo[i].Name)
	}
	saveMultipartDataFiles(imageNames, r.MultipartForm.File["images"])

	out, err := easyjson.Marshal(es)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(out)

	// TODO: return files
}

func (handler *DiaryHandler) UpdateDiary(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// sessionId, err := sessions.CheckSession(r)
	// if err == domain.Err.ErrObj.UserNotLoggedIn {
	// 	http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusForbidden)
	// 	return
	// }

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	DiaryUpdateRequest := new(domain.DiaryUpdateRequest)
	// DiaryCreateRequest.SetDefault()

	err = easyjson.Unmarshal(b, DiaryUpdateRequest)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// if cast.IntToStr(sessionId) != EventCreatingRequest.UserId {
	// 	http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
	// 	w.WriteHeader(http.StatusBadRequest)
	// }

	sanitizer.SanitizeDiaryUpdating(DiaryUpdateRequest)

	es, err := handler.DiaryUsecase.UpdateDiary(*DiaryUpdateRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	out, err := easyjson.Marshal(es)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(out)
}

// func (handler *DiaryHandler) UpdateRecord(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()

// 	// sessionId, err := sessions.CheckSession(r)
// 	// if err == domain.Err.ErrObj.UserNotLoggedIn {
// 	// 	http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusForbidden)
// 	// 	return
// 	// }
// 	err := r.ParseMultipartForm(32 << 20) // maxMemory 32MB
// 	if err != nil {
// 		log.Error(err)
// 		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	formdata := r.MultipartForm
//  	//get the *fileheaders
//  	files := formdata.File["images"] // grab the filenames
// 	imageInfo := []domain.ImageInfoUsecase{}
// 	// filePaths := []string{}
//  	for i, _ := range files { // loop through the files one by one
//  		file, err := files[i].Open()
//  		defer file.Close()
//  		if err != nil {
// 			log.Error(err)
// 			http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
// 			return
//  		}

// 		if !validExtenstions(files){
// 			log.Error(domain.Err.ErrObj.BadFileExtension)
// 			http.Error(w, domain.Err.ErrObj.BadFileExtension.Error(), http.StatusBadRequest)
// 			return
// 		}
// 		ar, err := cast.StringToFloat64((r.Form["areas"])[i])
// 		if err != nil {
// 			log.Error(err)
// 			http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
// 			return
//  		}
// 		imageInfo = append(imageInfo, domain.ImageInfoUsecase{Name: extractName(files[i].Filename), Area: ar})
// 		// filePaths = append(filePaths, files[i].Filename)
// 		println(config.DevConfigStore.LoadedFilesPath + files[i].Filename)
//  		out, err := os.Create(config.DevConfigStore.LoadedFilesPath + files[i].Filename)
//  		defer out.Close()
//  		if err != nil {
// 			log.Error(err)
// 			http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
// 			return
//  		}

//  		_, err = io.Copy(out, file) // file not files[i] !

//  		if err != nil {
// 			log.Error(err)
// 			http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
// 			return
//  		}
//  	}

// 	RecordCreateRequest := new(domain.RecordCreateRequest)
// 	RecordCreateRequest.SetDefault()
// 	RecordCreateRequest.Title = fmt.Sprintf("%v", (r.Form["title"])[0])
// 	RecordCreateRequest.Description = fmt.Sprintf("%v", (r.Form["description"])[0])
// 	characteristicsRequest := [](*uint8){   &RecordCreateRequest.Characteristics.Itching,
// 											&RecordCreateRequest.Characteristics.Pain,
// 											&RecordCreateRequest.Characteristics.Edema,
// 											&RecordCreateRequest.Characteristics.Redness,
// 											&RecordCreateRequest.Characteristics.Dryness,
// 											&RecordCreateRequest.Characteristics.Peeling}
// 	characteristics := []string{"itching", "pain", "edema", "redness", "dryness", "peeling"}
// 	for i := range characteristics{
// 		uint8value, err := strconv.Atoi(r.Form[characteristics[i]][0])
// 		if err != nil {
// 			log.Error(err)
// 			http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		*characteristicsRequest[i] = uint8(uint8value)

// 	}

// 	// if cast.IntToStr(sessionId) != EventCreatingRequest.UserId {
// 	// 	http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
// 	// 	w.WriteHeader(http.StatusBadRequest)
// 	// }
// 	sanitizer.SanitizeRecordCreating(RecordCreateRequest)

// 	// imageInfo := []domain.ImageInfoUsecase{}
// 	// for i := range imageNames {
// 	// 	imageInfo = append(imageInfo, domain.ImageInfoUsecase{Name:imageNames[i], Area: 1.1})
// 	// }
// 	//TODO: Соз
// 	es, err := handler.DiaryUsecase.CreateRecord(diaryId, *RecordCreateRequest, imageInfo)
// 	if err != nil {
// 		log.Error(err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	out, err := easyjson.Marshal(es)
// 	if err != nil {
// 		log.Error(err)
// 		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// 	w.Write(out)

// 	// TODO: return files
// }

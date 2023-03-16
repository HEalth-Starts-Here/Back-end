package diarydelivery

import (
	"hesh/internal/pkg/domain"
	// "path/filepath"
	// "eventool/internal/pkg/sessions"
	// "hesh/internal/pkg/utils/cast"
	"hesh/internal/pkg/utils/config"
	"hesh/internal/pkg/utils/log"
	"hesh/internal/pkg/utils/sanitizer"
	"strconv"

	// "strings"

	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mailru/easyjson"
	// "encoding/json"
	"io"
	"os"
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

	DiaryCreatingRequest := new(domain.DiaryCreatingRequest)
	DiaryCreatingRequest.SetDefault()
	
	err = easyjson.Unmarshal(b, DiaryCreatingRequest)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// if cast.IntToStr(sessionId) != EventCreatingRequest.UserId {
	// 	http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
	// 	w.WriteHeader(http.StatusBadRequest)
	// }

	sanitizer.SanitizeDiaryCreating(DiaryCreatingRequest)

	es, err := handler.DiaryUsecase.CreateDiary(*DiaryCreatingRequest)
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

func (handler *DiaryHandler) GetDiary (w http.ResponseWriter, r *http.Request) {
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
	formdata := r.MultipartForm
 	//get the *fileheaders
 	files := formdata.File["images"] // grab the filenames
	filepaths := []string{}
 	for i, _ := range files { // loop through the files one by one
 		file, err := files[i].Open()
 		defer file.Close()
 		if err != nil {
			log.Error(err)
			http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
			return
 		}
		filepaths = append(filepaths, files[i].Filename)
		println(config.DevConfigStore.LoadedFilesPath + files[i].Filename)
 		out, err := os.Create(config.DevConfigStore.LoadedFilesPath + files[i].Filename)
 		defer out.Close()
 		if err != nil {
			log.Error(err)
			http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
			return
 		}

 		_, err = io.Copy(out, file) // file not files[i] !

 		if err != nil {
			log.Error(err)
			http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
			return
 		}
 	}

	RecordCreatingRequest := new(domain.RecordCreatingRequest)
	RecordCreatingRequest.SetDefault()
	RecordCreatingRequest.Title = fmt.Sprintf("%v", (r.Form["title"])[0])
	RecordCreatingRequest.Description = fmt.Sprintf("%v", (r.Form["description"])[0])
	characteristicsRequest := [](*uint8){   &RecordCreatingRequest.Characteristics.Itching, 
											&RecordCreatingRequest.Characteristics.Pain, 
											&RecordCreatingRequest.Characteristics.Edema, 
											&RecordCreatingRequest.Characteristics.Redness, 
											&RecordCreatingRequest.Characteristics.Dryness, 
											&RecordCreatingRequest.Characteristics.Peeling}
	characteristics := []string{"itching", "pain", "edema", "redness", "dryness", "peeling"}
	for i := range characteristics{
		uint8value, err := strconv.Atoi(r.Form[characteristics[i]][0])
		if err != nil {
			log.Error(err)
			http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
			return
		}
		*characteristicsRequest[i] = uint8(uint8value)
		
	}
	RecordCreatingRequest.FilePaths = filepaths

	// if cast.IntToStr(sessionId) != EventCreatingRequest.UserId {
	// 	http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
	// 	w.WriteHeader(http.StatusBadRequest)
	// }

	sanitizer.SanitizeRecordCreating(RecordCreatingRequest)

	es, err := handler.DiaryUsecase.CreateRecord(diaryId, *RecordCreatingRequest)
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


// func (handler *EventHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	categoryList, err := handler.EventUsecase.GetCategory()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
	
// 	out, err := easyjson.Marshal(categoryList)
// 	if err != nil {
// 		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
// 		return
// 	}
	
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(out)
// }

// func (handler *EventHandler) EventSignUp (w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	userId, err := sessions.CheckSession(r)
// 	if err == domain.Err.ErrObj.UserNotLoggedIn {
// 		http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusForbidden)
// 		return
// 	}

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	params := mux.Vars(r)
// 	eventId, err := strconv.ParseUint(params["id"], 10, 64)
// 	if err != nil {
// 		http.Error(w, domain.Err.ErrObj.Uint64Cast.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	err = handler.EventUsecase.EventSignUp(eventId, userId)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	} 

// 	w.WriteHeader(http.StatusCreated)
// }

// func (handler *EventHandler) CancelEventSignUp(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	var err error
// 	var userId uint64
// 	if userId, err = sessions.CheckSession(r); err == domain.Err.ErrObj.UserNotLoggedIn {
// 		http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	params := mux.Vars(r)
// 	eventId, err := strconv.ParseUint(params["id"], 10, 64)
// 	if err != nil {
// 		http.Error(w, domain.Err.ErrObj.Uint64Cast.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	err = handler.EventUsecase.CancelEventSignUp(eventId, userId)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	} 


// 	w.WriteHeader(http.StatusOK)
// }

// func (handler *EventHandler) GetRecomendedEvent(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	userId, err := sessions.CheckSession(r);
// 	if err == domain.Err.ErrObj.UserNotLoggedIn {
// 		http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusBadRequest)
// 		return
// 	}
	
// 	eventList, err := handler.EventUsecase.GetRecomendedEvent(userId)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
	
// 	out, err := easyjson.Marshal(eventList)
// 	if err != nil {
// 		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
// 		return
// 	}
	
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(out)
// }

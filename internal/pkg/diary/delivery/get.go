package diarydelivery

import (
	"hesh/internal/pkg/domain"
	// "eventool/internal/pkg/sessions"
	// "hesh/internal/pkg/utils/cast"
	"hesh/internal/pkg/utils/sanitizer"
	// "strconv"
	// "strings"

	"io/ioutil"
	"net/http"

	// "github.com/gorilla/mux"

	"github.com/mailru/easyjson"
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
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// if cast.IntToStr(sessionId) != EventCreatingRequest.UserId {
	// 	http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
	// 	w.WriteHeader(http.StatusBadRequest)
	// }

	sanitizer.SanitizeDiaryCreating(DiaryCreatingRequest)

	es, err := handler.DiaryUsecase.CreateDiary(*DiaryCreatingRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

// func (handler *EventHandler) GetCertainEvent(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	// userId, err := sessions.CheckSession(r);
// 	// if err == domain.Err.ErrObj.UserNotLoggedIn {
// 	// 	http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusBadRequest)
// 	// 	return
// 	// }

// 	params := mux.Vars(r)
// 	eventId, err := strconv.ParseUint(params["id"], 10, 64)
// 	if err != nil {
// 		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
// 		return
// 	}
	
// 	// event, err := handler.EventUsecase.GetCertainEvent(eventId, userId)
// 	event, err := handler.EventUsecase.GetCertainEvent(eventId)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
	
// 	out, err := easyjson.Marshal(event)
// 	if err != nil {
// 		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
// 		return
// 	}
	
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(out)
// }


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

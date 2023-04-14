package commentdelivery

import (
	"fmt"
	"hesh/internal/pkg/domain"
	"io/ioutil"

	// "path/filepath"
	// "eventool/internal/pkg/sessions"
	// "hesh/internal/pkg/utils/cast"
	// "hesh/internal/pkg/utils/log"
	"hesh/internal/pkg/utils/sanitizer"

	"strconv"

	// "strings"

	"net/http"

	"github.com/gorilla/mux"

	"github.com/mailru/easyjson"
	// "encoding/json"
	// "io"
	// "os"
)

// func (handler *DiaryHandler) CreateDiary(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	// sessionId, err := sessions.CheckSession(r)
// 	// if err == domain.Err.ErrObj.UserNotLoggedIn {
// 	// 	http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusForbidden)
// 	// 	return
// 	// }

// 	b, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	// TODO add check is this user exist
// 	queryParameter := r.URL.Query().Get("vk_user_id")
// 	medicId, err := strconv.ParseUint(queryParameter, 10, 64)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	DiaryCreateRequest := new(domain.DiaryCreateRequest)
// 	DiaryCreateRequest.SetDefault()

// 	err = easyjson.Unmarshal(b, DiaryCreateRequest)
// 	if err != nil {
// 		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// if cast.IntToStr(sessionId) != EventCreatingRequest.UserId {
// 	// 	http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
// 	// 	w.WriteHeader(http.StatusBadRequest)
// 	// }

// 	sanitizer.SanitizeDiaryCreating(DiaryCreateRequest)

// 	es, err := handler.DiaryUsecase.CreateDiary(*DiaryCreateRequest, medicId)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	out, err := easyjson.Marshal(es)
// 	if err != nil {
// 		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// 	w.Write(out)
// }

// func (handler *DiaryHandler) LinkDiary(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	// sessionId, err := sessions.CheckSession(r)
// 	// if err == domain.Err.ErrObj.UserNotLoggedIn {
// 	// 	http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusForbidden)
// 	// 	return
// 	// }

// 	_, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	queryParameter := r.URL.Query().Get("vk_user_id")
// 	medicId, err := strconv.ParseUint(queryParameter, 10, 64)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	params := mux.Vars(r)
// 	diaryId, err := strconv.ParseUint(params["id"], 10, 64)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	// DiaryCreateRequest := new(domain.DiaryCreateRequest)
// 	// DiaryCreateRequest.SetDefault()

// 	// if cast.IntToStr(sessionId) != EventCreatingRequest.UserId {
// 	// 	http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
// 	// 	w.WriteHeader(http.StatusBadRequest)
// 	// }

// 	es, err := handler.DiaryUsecase.LinkDiary(diaryId, medicId)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	out, err := easyjson.Marshal(es)
// 	if err != nil {
// 		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// 	w.Write(out)
// }

// func (handler *DiaryHandler) DeleteDiary(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	// sessionId, err := sessions.CheckSession(r)
// 	// if err == domain.Err.ErrObj.UserNotLoggedIn {
// 	// 	http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusForbidden)
// 	// 	return
// 	// }

// 	params := mux.Vars(r)
// 	diaryId, err := strconv.ParseUint(params["id"], 10, 64)
// 	if err != nil {
// 		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	err = handler.DiaryUsecase.DeleteDiary(diaryId)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

func (handler *CommentHandler) CreateComment (w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// categoryString := r.URL.Query().Get("category")
	// categories := strings.Split(categoryString, " ")
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
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	CommentCreateRequest := new(domain.BasicCommentInfo)
	CommentCreateRequest.SetDefault()

	err = easyjson.Unmarshal(b, CommentCreateRequest)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// if cast.IntToStr(sessionId) != EventCreatingRequest.UserId {
	// 	http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
	// 	w.WriteHeader(http.StatusBadRequest)
	// }

	fmt.Printf("CommentCreateRequest: %v\n", CommentCreateRequest)
	sanitizer.SanitizeBasicCommentInfo(CommentCreateRequest)
	fmt.Printf("CommentCreateRequest: %v\n", CommentCreateRequest)
	response, err := handler.CommentUsecase.CreateComment(diaryId, userId, *CommentCreateRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	out, err := easyjson.Marshal(response)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

// func (handler *DiaryHandler) GetCertainDiary(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	// userId, err := sessions.CheckSession(r);
// 	// if err == domain.Err.ErrObj.UserNotLoggedIn {
// 	// 	http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusBadRequest)
// 	// 	return
// 	// }

// 	params := mux.Vars(r)
// 	diaryId, err := strconv.ParseUint(params["id"], 10, 64)
// 	if err != nil {
// 		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
// 		return
// 	}
	
// 	queryParameter := r.URL.Query().Get("vk_user_id")
// 	userId, err := strconv.ParseUint(queryParameter, 10, 64)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	diary, err := handler.DiaryUsecase.GetCertainDiary(diaryId, userId)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	out, err := easyjson.Marshal(diary)
// 	if err != nil {
// 		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Write(out)
// }

// func (handler *DiaryHandler) UpdateDiary(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	// sessionId, err := sessions.CheckSession(r)
// 	// if err == domain.Err.ErrObj.UserNotLoggedIn {
// 	// 	http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusForbidden)
// 	// 	return
// 	// }

// 	params := mux.Vars(r)
// 	diaryId, err := strconv.ParseUint(params["id"], 10, 64)
// 	if err != nil {
// 		log.Error(err)
// 		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	b, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	DiaryUpdateRequest := new(domain.DiaryUpdateRequest)
// 	// DiaryCreateRequest.SetDefault()

// 	err = easyjson.Unmarshal(b, DiaryUpdateRequest)
// 	if err != nil {
// 		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	// if cast.IntToStr(sessionId) != EventCreatingRequest.UserId {
// 	// 	http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
// 	// 	w.WriteHeader(http.StatusBadRequest)
// 	// }

// 	sanitizer.SanitizeDiaryUpdating(DiaryUpdateRequest)

// 	es, err := handler.DiaryUsecase.UpdateDiary(*DiaryUpdateRequest, diaryId)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	out, err := easyjson.Marshal(es)
// 	if err != nil {
// 		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// 	w.Write(out)
// }


package notedelivery

import (
	"hesh/internal/pkg/domain"

	"strconv"

	"net/http"

	"github.com/gorilla/mux"

	"github.com/mailru/easyjson"
)

// func (handler *NoteHandler) CreateNote (w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	queryParameter := r.URL.Query().Get("vk_user_id")
// 	userId, err := strconv.ParseUint(queryParameter, 10, 64)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	params := mux.Vars(r)
// 	dId, err := strconv.ParseUint(params["id"], 10, 64)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	b, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	CommentCreateRequest := new(domain.BasicCommentInfo)
// 	CommentCreateRequest.SetDefault()

// 	err = easyjson.Unmarshal(b, CommentCreateRequest)
// 	if err != nil {
// 		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	fmt.Printf("CommentCreateRequest: %v\n", CommentCreateRequest)
// 	sanitizer.SanitizeBasicCommentInfo(CommentCreateRequest)
// 	fmt.Printf("CommentCreateRequest: %v\n", CommentCreateRequest)
// 	response, err := handler.CommentUsecase.CreateComment(diaryId, userId, *CommentCreateRequest)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	out, err := easyjson.Marshal(response)
// 	if err != nil {
// 		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Write(out)
// }

func (handler *NoteHandler) GetNote (w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// categoryString := r.URL.Query().Get("category")
	// categories := strings.Split(categoryString, " ")
	queryParameter := r.URL.Query().Get("vk_user_id")
	medicId, err := strconv.ParseUint(queryParameter, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	params := mux.Vars(r)
	isMedicRecord, err := strconv.ParseBool(params["isMedicRecord"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	noteId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	noteList, err := handler.NoteUsecase.GetNote(medicId, isMedicRecord, noteId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	out, err := easyjson.Marshal(noteList)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

// func (handler *CommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
// 	defer r.Body.Close()
// 	// sessionId, err := sessions.CheckSession(r)
// 	// if err == domain.Err.ErrObj.UserNotLoggedIn {
// 	// 	http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusForbidden)
// 	// 	return
// 	// }
// 	queryParameter := r.URL.Query().Get("vk_user_id")
// 	userId, err := strconv.ParseUint(queryParameter, 10, 64)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	params := mux.Vars(r)
// 	commentId, err := strconv.ParseUint(params["id"], 10, 64)
// 	if err != nil {
// 		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	err = handler.CommentUsecase.DeleteComment(userId, commentId)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		w.WriteHeader(http.StatusBadRequest)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }

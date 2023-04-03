package userdelivery

import (
	"hesh/internal/pkg/domain"
	// "mime/multipart"

	// "path/filepath"
	// "eventool/internal/pkg/sessions"
	// "hesh/internal/pkg/utils/cast"
	// "hesh/internal/pkg/utils/cast"
	// "hesh/internal/pkg/utils/config"
	// "hesh/internal/pkg/utils/filesaver"
	// "hesh/internal/pkg/utils/log"
	"hesh/internal/pkg/utils/sanitizer"

	// "path/filepath"
	"strconv"


	// "fmt"
	"io/ioutil"
	"net/http"

	// "github.com/gorilla/mux"

	"github.com/mailru/easyjson"
)

func (handler *UserHandler) UserInit(w http.ResponseWriter, r *http.Request) {
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
	// TODO add check is this user exist
	queryParameter := r.URL.Query().Get("vk_user_id")
	userId, err := strconv.ParseUint(queryParameter, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	UserInitRequest := new(domain.UserInitRequest)
	// DiaryCreateRequest.SetDefault()

	err = easyjson.Unmarshal(b, UserInitRequest)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// if cast.IntToStr(sessionId) != EventCreatingRequest.UserId {
	// 	http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
	// 	w.WriteHeader(http.StatusBadRequest)
	// }

	sanitizer.SanitizeUserInit(UserInitRequest)

	us, err := handler.UserUsecase.UserInit(*UserInitRequest, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	out, err := easyjson.Marshal(us)
	if err != nil {
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(out)
}

package diarydelivery

import (
	"fmt"
	"hesh/internal/pkg/domain"
	"hesh/internal/pkg/utils/log"
	"hesh/internal/pkg/utils/sanitizer"
	"io/ioutil"

	"strconv"

	"net/http"

	"github.com/gorilla/mux"

	"github.com/SevereCloud/vksdk/v2/api"
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
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// TODO add check is this user exist
	queryParameter := r.URL.Query().Get("vk_user_id")
	medicId, err := strconv.ParseUint(queryParameter, 10, 64)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	DiaryCreateRequest := new(domain.DiaryCreateRequest)
	DiaryCreateRequest.SetDefault()

	err = easyjson.Unmarshal(b, DiaryCreateRequest)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Printf("DiaryCreateRequest.DiaryBasicInfo.Reminder.StartDate: %v\n", DiaryCreateRequest.DiaryBasicInfo.Reminder.StartDate)
	
	// if cast.IntToStr(sessionId) != EventCreatingRequest.UserId {
		// 	http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		// 	w.WriteHeader(http.StatusBadRequest)
		// }
		
	sanitizer.SanitizeDiaryCreating(DiaryCreateRequest)
	fmt.Printf("DiaryCreateRequest.DiaryBasicInfo.Reminder.StartDate: %v\n", DiaryCreateRequest.DiaryBasicInfo.Reminder.StartDate)

	es, err := handler.DiaryUsecase.CreateDiary(*DiaryCreateRequest, medicId)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	out, err := easyjson.Marshal(es)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(out)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
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
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	queryParameters := r.URL.Query()
	vkUserId := queryParameters.Get("vk_user_id")
	patientId, err := strconv.ParseUint(vkUserId, 10, 64)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	linkToken := queryParameters.Get("linktoken")
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	diaryId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		log.Error(err)
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

	es, err := handler.DiaryUsecase.LinkDiary(patientId, diaryId, linkToken)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	out, err := easyjson.Marshal(es)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(out)
	if err != nil {
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
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
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.DiaryUsecase.DeleteDiary(diaryId)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (handler *DiaryHandler) GetDiary(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()


	vk := api.NewVK("e80e2119e80e2119e80e21198ceb1d081fee80ee80e21198c168a958ccfd793e077d5da")

	users, err := vk.UsersGet(api.Params{
		"user_ids": 165523569,
		"fields": "photo_50,verified,photo_id,bdate",
		// "fields": "photo_50,verified",
	})
	if err != nil {
		log.Error(err)
	}
	fmt.Printf("users: %v\n", users)

	var userIds []int
	userIds = append(userIds, 165523569)
	// userIds = append(userIds, 165523569)
	notidications, err := vk.NotificationsSendMessage(api.Params{
		"user_ids": userIds,
		"message":  "Вам пришла электронная повестка! Узнать подробности можно в личном кабинете на портале ГосУслуг (gosuslugi.ru)",
		// "sending_mode":  5,
	})
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Printf("notidications: %v\n", notidications)
	println()
	println()



	// categoryString := r.URL.Query().Get("category")
	// categories := strings.Split(categoryString, " ")
	queryParameter := r.URL.Query().Get("vk_user_id")
	userId, err := strconv.ParseUint(queryParameter, 10, 64)
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	diaryList, err := handler.DiaryUsecase.GetDiary(userId)
	if err != nil {
		httpStatus := http.StatusBadRequest
		log.Error(err)
		http.Error(w, err.Error(), httpStatus)
		w.WriteHeader(httpStatus)
		return
	}

	out, err := easyjson.Marshal(diaryList)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		log.Error(err)
		http.Error(w, err.Error(), httpStatus)
		w.WriteHeader(httpStatus)
		return
	}

	_, err = w.Write(out)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), httpStatus)
		w.WriteHeader(httpStatus)
		return
	}
	w.WriteHeader(http.StatusOK)
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
		httpStatus := http.StatusBadRequest
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), httpStatus)
		w.WriteHeader(httpStatus)
		return
	}

	queryParameter := r.URL.Query().Get("vk_user_id")
	userId, err := strconv.ParseUint(queryParameter, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	diary, err := handler.DiaryUsecase.GetCertainDiary(diaryId, userId)
	if err != nil {
		httpStatus := http.StatusBadRequest
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), httpStatus)
		w.WriteHeader(httpStatus)
		return
	}

	out, err := easyjson.Marshal(diary)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), httpStatus)
		w.WriteHeader(httpStatus)
		return
	}

	_, err = w.Write(out)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), httpStatus)
		w.WriteHeader(httpStatus)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *DiaryHandler) UpdateDiary(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// sessionId, err := sessions.CheckSession(r)
	// if err == domain.Err.ErrObj.UserNotLoggedIn {
	// 	http.Error(w, domain.Err.ErrObj.UserNotLoggedIn.Error(), http.StatusForbidden)
	// 	return
	// }

	params := mux.Vars(r)
	diaryId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		httpStatus := http.StatusBadRequest
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), httpStatus)
		w.WriteHeader(httpStatus)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), httpStatus)
		w.WriteHeader(httpStatus)
		return
	}

	DiaryUpdateRequest := new(domain.DiaryUpdateRequest)
	// DiaryCreateRequest.SetDefault()
	fmt.Printf("DiaryUpdateRequest.DiaryBasicInfo.Reminder.StartDate: %v\n", DiaryUpdateRequest.DiaryBasicInfo.Reminder.StartDate)
	err = easyjson.Unmarshal(b, DiaryUpdateRequest)
	if err != nil {
		httpStatus := http.StatusBadRequest
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.BadInput.Error(), httpStatus)
		w.WriteHeader(httpStatus)
		return
	}

	// if cast.IntToStr(sessionId) != EventCreatingRequest.UserId {

	// 	http.Error(w, domain.Err.ErrObj.BadInput.Error(), http.StatusBadRequest)
	// 	w.WriteHeader(http.StatusBadRequest)
	// }

	sanitizer.SanitizeDiaryUpdating(DiaryUpdateRequest)

	es, err := handler.DiaryUsecase.UpdateDiary(*DiaryUpdateRequest, diaryId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	out, err := easyjson.Marshal(es)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), httpStatus)
		w.WriteHeader(httpStatus)
		return
	}

	_, err = w.Write(out)
	if err != nil {
		httpStatus := http.StatusInternalServerError
		log.Error(err)
		http.Error(w, domain.Err.ErrObj.InternalServer.Error(), httpStatus)
		w.WriteHeader(httpStatus)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

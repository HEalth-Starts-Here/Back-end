package diarydelivery

import (
	"hesh/internal/pkg/domain"

	"github.com/gorilla/mux"
)

type DiaryHandler struct {
	DiaryUsecase domain.DiaryUsecase
}

func SetDiaryHandlers(router *mux.Router, pu domain.DiaryUsecase) {
	handler := &DiaryHandler{
		DiaryUsecase: pu,
	}
	router.HandleFunc(CreateDiaryUrl, handler.CreateDiary).Methods("POST", "OPTIONS")
	// router.HandleFunc(LinkDiaryUrl, handler.LinkDiary).Methods("POST", "OPTIONS")
	router.HandleFunc(DeleteDiaryUrl, handler.DeleteDiary).Methods("POST", "OPTIONS")
	router.HandleFunc(PutCertainDiaryUrl, handler.UpdateDiary).Methods("PUT", "OPTIONS") 
	router.HandleFunc(GetDiaryUrl, handler.GetDiary).Methods("GET", "OPTIONS")
	router.HandleFunc(GetCertainDiaryUrl, handler.GetCertainDiary).Methods("GET", "OPTIONS")

	router.HandleFunc(CreateRecordUrl, handler.CreateRecord).Methods("POST", "OPTIONS")
	// router.HandleFunc(PutRecordUrl, handler.UpdateRecord).Methods("PUT", "OPTIONS") 
}

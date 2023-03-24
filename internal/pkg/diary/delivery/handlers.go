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
	router.HandleFunc(GetDiaryUrl, handler.GetDiary).Methods("GET", "OPTIONS")
	router.HandleFunc(GetCertainDiaryUrl, handler.GetCertainDiary).Methods("GET", "OPTIONS")
	router.HandleFunc(CreateRecordUrl, handler.CreateRecord).Methods("POST", "OPTIONS")
	router.HandleFunc(PutCertainDiaryUrl, handler.CreateRecord).Methods("PUT", "OPTIONS")
	// router.HandleFunc(CreateRecordUrl, handler.CreateRecord).Methods("POST", "OPTIONS")

	// router.HandleFunc(GetRecomendedDiary, handler.GetRecomendedDiary).Methods("GET", "OPTIONS")
	// router.HandleFunc(GetCatagoryUrl, handler.GetCategory).Methods("GET", "OPTIONS")
	// router.HandleFunc(EventSignUpUrl, handler.EventSignUp).Methods("POST", "OPTIONS")
	// router.HandleFunc(CancelEventSignUpUrl, handler.CancelEventSignUp).Methods("POST", "OPTIONS")
	
	// router.HandleFunc(DeleteEventUrl, handler.DeleteEvent).Methods("GET", "OPTIONS")
	// router.HandleFunc(AlterEventUrl, handler.AlterEvent).Methods("GET", "OPTIONS")
}

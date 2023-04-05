package recorddelivery

import (
	"hesh/internal/pkg/domain"
	"github.com/gorilla/mux"
)

type RecordHandler struct {
	RecordUsecase domain.RecordUsecase
}

func SetRecordHandlers(router *mux.Router, ru domain.RecordUsecase) {
	handler := &RecordHandler{
		RecordUsecase: ru,
	}
	router.HandleFunc(recordMedicCreateUrl, handler.CreateMedicRecord).Methods("POST", "OPTIONS")
	router.HandleFunc(recordMedicGetUrl, handler.GetMedicRecord).Methods("GET", "OPTIONS")
	router.HandleFunc(recordMedicUpdateTextUrl, handler.UpdateTextMedicRecord).Methods("PUT", "OPTIONS")
}

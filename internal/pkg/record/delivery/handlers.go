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
	router.HandleFunc(recordMedicGet, handler.GetMedicRecord).Methods("GET", "OPTIONS")
}

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
	// MEDIC
	router.HandleFunc(recordMedicCreateUrl, handler.CreateMedicRecord).Methods("POST", "OPTIONS")	
	router.HandleFunc(recordMedicGetUrl, handler.GetMedicRecord).Methods("GET", "OPTIONS")
	router.HandleFunc(recordMedicGetDiarisationsUrl, handler.GetMedicRecordDiarisations).Methods("GET", "OPTIONS")
	router.HandleFunc(recordMedicUpdateTextUrl, handler.UpdateTextMedicRecord).Methods("PUT", "OPTIONS")
	router.HandleFunc(recordMedicUpdateImageUrl, handler.UpdateImageMedicRecord).Methods("PUT", "OPTIONS")
	router.HandleFunc(recordMedicDeleteImageUrl, handler.DeleteMedicRecord).Methods("POST", "OPTIONS")
	
	// PATIENT
	router.HandleFunc(recordPatientCreateUrl, handler.CreatePatientRecord).Methods("POST", "OPTIONS")
	router.HandleFunc(recordPatientGetUrl, handler.GetPatientRecord).Methods("GET", "OPTIONS")
}

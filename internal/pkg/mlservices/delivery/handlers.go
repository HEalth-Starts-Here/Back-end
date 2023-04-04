package mlservicesdelivery

import (
	"hesh/internal/pkg/domain"

	"github.com/gorilla/mux"
)

type MLServicesHandler struct {
	MLServicesUsecase domain.MLServicesUsecase
}

func SetMLServicesHandlers(router *mux.Router, mlsu domain.MLServicesUsecase) {
	handler := &MLServicesHandler{
		MLServicesUsecase: mlsu,
	}
	router.HandleFunc(DetermineAreaUrl, handler.DetermineArea).Methods("POST", "OPTIONS")
	router.HandleFunc(EstimateImageUrl, handler.ImageQualityAssesment).Methods("POST", "OPTIONS")
}

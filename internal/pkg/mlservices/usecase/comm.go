package mlservicesusecase

import (
	"hesh/internal/pkg/domain"
)

type MLServicesUsecase struct {
	mlservicesRepo domain.MLServicesRepository
}

func InitMLServicesUsc(mlsr domain.MLServicesRepository) domain.MLServicesUsecase {
	return &MLServicesUsecase{
		mlservicesRepo: mlsr,
	}
}

package mlservicesusecase

import (
	"hesh/internal/pkg/domain"
	// "hesh/internal/pkg/utils/cast"
	// "hesh/internal/pkg/utils/log"

	// usrusecase "eventool/internal/pkg/user/usecase"
	// usrdelivery "eventool/internal/pkg/user/delivery/rest"
	// usrusecase "eventool/internal/pkg/user/usecase"
	// "usrdelivery"
	// "strings"
)

type MLServicesUsecase struct {
	mlservicesRepo domain.MLServicesRepository
}

func InitMLServicesUsc(mlsr domain.MLServicesRepository) domain.MLServicesUsecase {
	return &MLServicesUsecase{
		mlservicesRepo: mlsr,
	}
}

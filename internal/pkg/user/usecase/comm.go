package userusecase

import (
	"hesh/internal/pkg/domain"
)

type UserUsecase struct {
	userRepo domain.UserRepository
}

// func trimTitle(title *string) {
// 	*title = strings.Trim(*title, " ")
// }

func InitUserUsc(ur domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{
		userRepo: ur,
	}
}

func (eu UserUsecase) UserInit(userId uint64) (bool, domain.UserInfo, error) {
	isEsisted, UserInfo, err := eu.userRepo.UserInit(userId)
	if err != nil {
		return false, domain.UserInfo{}, err
	}
	return isEsisted, UserInfo, nil
}

func (eu UserUsecase) RegisterMedic(userInfoRequest domain.RegisterMedicRequest, medicId uint64) (domain.UserInfo, error) {
	UserInfo, err := eu.userRepo.RegisterMedic(userInfoRequest, medicId)
	if err != nil {
		return domain.UserInfo{}, err
	}
	return UserInfo, nil
}

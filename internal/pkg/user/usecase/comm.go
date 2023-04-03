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

func (eu UserUsecase) UserInit(userInitInfo domain.UserInitRequest, userId uint64) (domain.UserInitResponse, error) {
	// alreadyExist, err := eu.diaryRepo.DiaryAlreadyExist(diaryData)
	// if err != nil {
	// 	return domain.DiaryCreateResponse{}, err
	// }

	// if alreadyExist {
	// 	return domain.DiaryCreateResponse{}, domain.Err.ErrObj.PlaylistExist
	// }

	if !userInitInfo.IsValid() {
		return domain.UserInitResponse{}, domain.Err.ErrObj.InvalidUsername
	}

	UserInitResponse, err := eu.userRepo.UserInit(userInitInfo, userId)
	if err != nil {
		return domain.UserInitResponse{}, err
	}
	return UserInitResponse, nil
}

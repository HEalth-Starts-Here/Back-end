package searchusecase

import (
	"hesh/internal/pkg/domain"
)

type SearchUsecase struct {
	SearchRepo domain.SearchRepository
}

func InitSearchUsc(sr domain.SearchRepository) domain.SearchUsecase {
	return &SearchUsecase{
		SearchRepo: sr,
	}
}


func (su SearchUsecase) CheckUserRole(userId uint64) (bool, bool, error) {
	userExist, isMedic, err := su.SearchRepo.CheckUserRole(userId)
	if err != nil {
		return false, false, err
	}
	return userExist, isMedic, nil
}

func (su SearchUsecase) SearchDiary(userId uint64, searchParams domain.SearchDiaryRequest) (domain.DiaryListResponse, error) {
	if !searchParams.IsValid() {
		return domain.DiaryListResponse{}, domain.Err.ErrObj.InvalidText
	}
	isExist, isMedic, err := su.CheckUserRole(userId)
	if err != nil {
		return domain.DiaryListResponse{}, err
	}
	if !isExist {
		return domain.DiaryListResponse{}, domain.Err.ErrObj.UserNotExist
	}

	// isExist, isMedic, err := nu.CheckUserRole(userId)
	// if err != nil {
	// 	return domain.CommentCreateResponse{}, err
	// }
	// if !isExist {
	// 	return domain.CommentCreateResponse{}, domain.Err.ErrObj.UserNotExist
	// }
	// userExist, err := su.CheckUserExist(userId)
	// if err != nil {
	// 	return domain.DiaryListResponse{}, err
	// }
	// if !userExist {
	// 	return domain.DiaryListResponse{}, domain.Err.ErrObj.UserDoestExist
	// }

	response, err := su.SearchRepo.SearchDiary(userId, isMedic, searchParams)
	if err != nil {
		return domain.DiaryListResponse{}, err
	}
	return response, nil
}

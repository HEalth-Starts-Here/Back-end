package diaryusecase

import (
	"hesh/internal/pkg/domain"
	"hesh/internal/pkg/utils/randomizer"

	// "hesh/internal/pkg/utils/cast"
	// "hesh/internal/pkg/utils/log"

	// usrusecase "eventool/internal/pkg/user/usecase"
	// userrepository "hesh/internal/pkg/user/repository"
	// usrusecase "eventool/internal/pkg/user/usecase"
	// "usrdelivery"
	"strings"
)

const tokenLength = 256

type DiaryUsecase struct {
	diaryRepo domain.DiaryRepository
}

func trimTitle(title *string) {
	*title = strings.Trim(*title, " ")
}

func InitDiaryUsc(pr domain.DiaryRepository) domain.DiaryUsecase {
	return &DiaryUsecase{
		diaryRepo: pr,
	}
}

func (du DiaryUsecase) CreateDiary(diaryData domain.DiaryCreateRequest, medicId uint64) (domain.DiaryCreateResponse, error) {
	// alreadyExist, err := eu.diaryRepo.DiaryAlreadyExist(diaryData)
	// if err != nil {
	// 	return domain.DiaryCreateResponse{}, err
	// }

	// if alreadyExist {
	// 	return domain.DiaryCreateResponse{}, domain.Err.ErrObj.PlaylistExist
	// }

	if !diaryData.IsValid() {
		return domain.DiaryCreateResponse{}, domain.Err.ErrObj.InvalidTitleOrDescription
	}

	DiaryCreateResponse, err := du.diaryRepo.CreateDiary(diaryData, medicId)
	if err != nil {
		return domain.DiaryCreateResponse{}, err
	}

	token, err := randomizer.GenerateRandomString(tokenLength)

	err = du.diaryRepo.CreateLinkToken(DiaryCreateResponse.Id, token)
	if err != nil {
		return domain.DiaryCreateResponse{}, err
	}
	DiaryCreateResponse.LinkToken = token
	return DiaryCreateResponse, nil
}

func (du DiaryUsecase) LinkDiary(patientId uint64, diaryId uint64, linkToken string) (domain.DiaryLinkResponse, error) {
	// alreadyExist, err := eu.diaryRepo.DiaryAlreadyExist(diaryData)
	// if err != nil {
	// 	return domain.DiaryCreateResponse{}, err
	// }

	// if alreadyExist {
	// 	return domain.DiaryCreateResponse{}, domain.Err.ErrObj.PlaylistExist
	// }

	isLinkExist, err := du.diaryRepo.CheckAndDeleteToken(diaryId, linkToken)
	if err != nil {
		return domain.DiaryLinkResponse{},  err
	}
	if !isLinkExist {
		return domain.DiaryLinkResponse{}, domain.Err.ErrObj.InvalidLinkToken
	} 

	DiaryCreateResponse, err := du.diaryRepo.LinkDiary(patientId, diaryId)
	if err != nil {
		return domain.DiaryLinkResponse{}, err
	}
	// DiaryCreateResponse.Categories, err = eu.eventRepo.CreateEventCategory(eventCreatingResponse.Id, eventData.Categories)
	// if err != nil {
	// 	return domain.EventCreatingResponse{}, err
	// }
	return DiaryCreateResponse, nil
}

func (du DiaryUsecase) DeleteDiary(diaryId uint64) error {
	// alreadyExist, err := eu.diaryRepo.DiaryAlreadyExist(diaryData)
	// if err != nil {
	// 	return domain.DiaryCreateResponse{}, err
	// }

	// if alreadyExist {
	// 	return domain.DiaryCreateResponse{}, domain.Err.ErrObj.PlaylistExist
	// }

	err := du.diaryRepo.DeleteDiary(diaryId)
	if err != nil {
		return err
	}

	return nil
}

func (du DiaryUsecase) GetDiary(userId uint64) (domain.DiaryListResponse, error) {

	feed, err := du.diaryRepo.GetDiary(userId)

	if err != nil {
		return domain.DiaryListResponse{}, err
	}

	return feed, nil
}

func (du DiaryUsecase) GetCertainDiary(diaryId uint64, userId uint64) (domain.DiaryResponse, error) {
	diary := domain.DiaryResponse{}
	diary, err := du.diaryRepo.GetCertainDiary(diaryId)
	if err != nil {
		return domain.DiaryResponse{}, err
	}

	isExisted, isMedic, err := du.diaryRepo.GetUserRole(userId) 
	if err != nil {
		return domain.DiaryResponse{}, err
	}

	if (!isExisted) {
		return domain.DiaryResponse{}, domain.Err.ErrObj.UserNotExist
	}

	if (!isMedic) {
		diary.Diary.DiaryBasicInfo.Anamnesis = ""
		diary.Diary.DiaryBasicInfo.Objectively = ""
		diary.Diary.DiaryBasicInfo.Diagnosis = ""
		diary.Diary.DiaryBasicInfo.Complaints = ""
	}

	return diary, nil
}


func (du DiaryUsecase) UpdateDiary(updateDiaryData domain.DiaryUpdateRequest, diaryId uint64) (domain.DiaryUpdateResponse, error) {
	// alreadyExist, err := eu.diaryRepo.DiaryAlreadyExist(diaryData)
	// if err != nil {
	// 	return domain.DiaryCreateResponse{}, err
	// }

	// if alreadyExist {
	// 	return domain.DiaryCreateResponse{}, domain.Err.ErrObj.PlaylistExist
	// }

	if !updateDiaryData.IsValid() {
		return domain.DiaryUpdateResponse{}, domain.Err.ErrObj.InvalidTitleOrDescription
	}
	DiaryUpdateResponse, err := du.diaryRepo.UpdateDiary(updateDiaryData, diaryId)
	if err != nil {
		return domain.DiaryUpdateResponse{}, err
	}

	// DiaryCreateResponse.Categories, err = eu.eventRepo.CreateEventCategory(eventCreatingResponse.Id, eventData.Categories)
	// if err != nil {
	// 	return domain.EventCreatingResponse{}, err
	// }
	return DiaryUpdateResponse, nil
}

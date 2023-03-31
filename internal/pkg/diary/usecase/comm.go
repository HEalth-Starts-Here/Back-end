package plausecase

import (
	"hesh/internal/pkg/domain"
	// "hesh/internal/pkg/utils/cast"
	// "hesh/internal/pkg/utils/log"

	// usrusecase "eventool/internal/pkg/user/usecase"
	// usrdelivery "eventool/internal/pkg/user/delivery/rest"
	// usrusecase "eventool/internal/pkg/user/usecase"
	// "usrdelivery"
	"strings"
)

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


func (eu DiaryUsecase) CreateDiary(diaryData domain.DiaryCreatingRequest) (domain.DiaryCreatingResponse, error) {
	// alreadyExist, err := eu.diaryRepo.DiaryAlreadyExist(diaryData)
	// if err != nil {
	// 	return domain.DiaryCreatingResponse{}, err
	// }

	// if alreadyExist {
	// 	return domain.DiaryCreatingResponse{}, domain.Err.ErrObj.PlaylistExist
	// }

	if !diaryData.IsValid() {
		return domain.DiaryCreatingResponse{}, domain.Err.ErrObj.InvalidTitleOrDescription
	}

	diaryCreatingResponse, err := eu.diaryRepo.CreateDiary(diaryData)
	if err != nil {
		return domain.DiaryCreatingResponse{}, err
	}

	// diaryCreatingResponse.Categories, err = eu.eventRepo.CreateEventCategory(eventCreatingResponse.Id, eventData.Categories)
	// if err != nil {
	// 	return domain.EventCreatingResponse{}, err
	// }
	return diaryCreatingResponse, nil
}

func (eu DiaryUsecase) DeleteDiary(diaryId uint64) (error) {
	// alreadyExist, err := eu.diaryRepo.DiaryAlreadyExist(diaryData)
	// if err != nil {
	// 	return domain.DiaryCreatingResponse{}, err
	// }

	// if alreadyExist {
	// 	return domain.DiaryCreatingResponse{}, domain.Err.ErrObj.PlaylistExist
	// }

	err := eu.diaryRepo.DeleteDiary(diaryId)
	if err != nil {
		return err
	}

	return nil
}

func (eu DiaryUsecase) GetDiary() (domain.DiaryListResponse, error) {
	
	feed, err := eu.diaryRepo.GetDiary()
	
	if err != nil {
		return domain.DiaryListResponse{}, err
	}

	return feed, nil
}

func (eu DiaryUsecase) GetCertainDiary(diaryId uint64) (domain.DiaryResponse, error) {
	diar1y := domain.RecordCreatingResponse{}
	diar1y.SetDefault()
	diary := domain.DiaryResponse{}
	diary, err := eu.diaryRepo.GetCertainDiary(diaryId)
	
	if err != nil {
		return domain.DiaryResponse{}, err
	}

	return diary, nil
}

func (eu DiaryUsecase) CreateRecord(diaryId uint64, recordData domain.RecordCreatingRequest, imageInfo []domain.ImageInfoUsecase) (domain.RecordCreatingResponse, error) {
	// alreadyExist, err := eu.diaryRepo.DiaryAlreadyExist(diaryData)
	// if err != nil {
	// 	return domain.DiaryCreatingResponse{}, err
	// }

	// if alreadyExist {
	// 	return domain.DiaryCreatingResponse{}, domain.Err.ErrObj.PlaylistExist
	// }

	if !recordData.IsValid() {
		return domain.RecordCreatingResponse{}, domain.Err.ErrObj.InvalidTitleOrDescription
	}
	Area := 0.0
	for i := range imageInfo {
		Area += imageInfo[i].Area
	}
	// TODO solve the problem with the same filenames. For example with generating filenames or with creating folders for every record
	diaryCreatingResponse, err := eu.diaryRepo.CreateRecord(diaryId, recordData, imageInfo, Area)
	if err != nil {
		return domain.RecordCreatingResponse{}, err
	}

	// diaryCreatingResponse.Categories, err = eu.eventRepo.CreateEventCategory(eventCreatingResponse.Id, eventData.Categories)
	// if err != nil {
	// 	return domain.EventCreatingResponse{}, err
	// }
	return diaryCreatingResponse, nil
}

func (eu DiaryUsecase) UpdateDiary(updateDiaryData domain.DiaryUpdatingRequest) (domain.DiaryUpdatingResponse, error) {
	// alreadyExist, err := eu.diaryRepo.DiaryAlreadyExist(diaryData)
	// if err != nil {
	// 	return domain.DiaryCreatingResponse{}, err
	// }

	// if alreadyExist {
	// 	return domain.DiaryCreatingResponse{}, domain.Err.ErrObj.PlaylistExist
	// }

	if !updateDiaryData.IsValid() {
		return domain.DiaryUpdatingResponse{}, domain.Err.ErrObj.InvalidTitleOrDescription
	}
	diaryUpdatingResponse, err := eu.diaryRepo.UpdateDiary(updateDiaryData)
	if err != nil {
		return domain.DiaryUpdatingResponse{}, err
	}

	// diaryCreatingResponse.Categories, err = eu.eventRepo.CreateEventCategory(eventCreatingResponse.Id, eventData.Categories)
	// if err != nil {
	// 	return domain.EventCreatingResponse{}, err
	// }
	return diaryUpdatingResponse, nil
}

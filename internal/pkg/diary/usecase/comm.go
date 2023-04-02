package plausecase

import (
	"hesh/internal/pkg/domain"
	"hesh/internal/pkg/utils/filesaver"
	"path/filepath"

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

func (eu DiaryUsecase) CreateDiary(diaryData domain.DiaryCreateRequest, medicId uint32) (domain.DiaryCreateResponse, error) {
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

	DiaryCreateResponse, err := eu.diaryRepo.CreateDiary(diaryData, medicId)
	if err != nil {
		return domain.DiaryCreateResponse{}, err
	}

	// DiaryCreateResponse.Categories, err = eu.eventRepo.CreateEventCategory(eventCreatingResponse.Id, eventData.Categories)
	// if err != nil {
	// 	return domain.EventCreatingResponse{}, err
	// }
	return DiaryCreateResponse, nil
}

func (eu DiaryUsecase) LinkDiary(diaryId uint64, medicId uint32) (domain.DiaryLinkResponse, error) {
	// alreadyExist, err := eu.diaryRepo.DiaryAlreadyExist(diaryData)
	// if err != nil {
	// 	return domain.DiaryCreateResponse{}, err
	// }

	// if alreadyExist {
	// 	return domain.DiaryCreateResponse{}, domain.Err.ErrObj.PlaylistExist
	// }

	DiaryCreateResponse, err := eu.diaryRepo.LinkDiary(diaryId, medicId)
	if err != nil {
		return domain.DiaryLinkResponse{}, err
	}

	// DiaryCreateResponse.Categories, err = eu.eventRepo.CreateEventCategory(eventCreatingResponse.Id, eventData.Categories)
	// if err != nil {
	// 	return domain.EventCreatingResponse{}, err
	// }
	return DiaryCreateResponse, nil
}

func (eu DiaryUsecase) DeleteDiary(diaryId uint64) error {
	// alreadyExist, err := eu.diaryRepo.DiaryAlreadyExist(diaryData)
	// if err != nil {
	// 	return domain.DiaryCreateResponse{}, err
	// }

	// if alreadyExist {
	// 	return domain.DiaryCreateResponse{}, domain.Err.ErrObj.PlaylistExist
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
	diar1y := domain.RecordCreateResponse{}
	diar1y.SetDefault()
	diary := domain.DiaryResponse{}
	diary, err := eu.diaryRepo.GetCertainDiary(diaryId)

	if err != nil {
		return domain.DiaryResponse{}, err
	}

	return diary, nil
}

func (eu DiaryUsecase) CreateRecord(diaryId uint64, recordData domain.RecordCreateRequest, imageInfo []domain.ImageInfoUsecase) (domain.RecordCreateResponse, error) {
	// alreadyExist, err := eu.diaryRepo.DiaryAlreadyExist(diaryData)
	// if err != nil {
	// 	return domain.DiaryCreateResponse{}, err
	// }

	// if alreadyExist {
	// 	return domain.DiaryCreateResponse{}, domain.Err.ErrObj.PlaylistExist
	// }

	if !recordData.IsValid() {
		return domain.RecordCreateResponse{}, domain.Err.ErrObj.InvalidTitleOrDescription
	}
	Area := 0.0
	for i := range imageInfo {
		Area += imageInfo[i].Area
	}
	alreadyUsed, err := eu.diaryRepo.GetImageNames()
	if err != nil {
		return domain.RecordCreateResponse{}, err
	}
	imageNames := filesaver.GetUniqueFileNames(len(imageInfo), alreadyUsed)
	for i := 0; i < len(imageInfo); i++ {
		imageInfo[i].Name = imageNames[i] + filepath.Ext(imageInfo[i].Name)

	}
	// TODO solve the problem with the same filenames. For example with generating filenames or with creating folders for every record
	DiaryCreateResponse, err := eu.diaryRepo.CreateRecord(diaryId, recordData, imageInfo, Area)
	if err != nil {
		return domain.RecordCreateResponse{}, err
	}

	// DiaryCreateResponse.Categories, err = eu.eventRepo.CreateEventCategory(eventCreatingResponse.Id, eventData.Categories)
	// if err != nil {
	// 	return domain.EventCreatingResponse{}, err
	// }
	return DiaryCreateResponse, nil
}

func (eu DiaryUsecase) UpdateDiary(updateDiaryData domain.DiaryUpdateRequest) (domain.DiaryUpdateResponse, error) {
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
	DiaryUpdateResponse, err := eu.diaryRepo.UpdateDiary(updateDiaryData)
	if err != nil {
		return domain.DiaryUpdateResponse{}, err
	}

	// DiaryCreateResponse.Categories, err = eu.eventRepo.CreateEventCategory(eventCreatingResponse.Id, eventData.Categories)
	// if err != nil {
	// 	return domain.EventCreatingResponse{}, err
	// }
	return DiaryUpdateResponse, nil
}

package diaryusecase

import (
	"fmt"
	"hesh/internal/pkg/domain"
	"hesh/internal/pkg/utils/randomizer"
)

const tokenLength = 256

type DiaryUsecase struct {
	diaryRepo domain.DiaryRepository
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
	fmt.Printf("diaryData.DiaryBasicInfo.Reminder.StartDate: %v\n", diaryData.DiaryBasicInfo.Reminder.StartDate)
	if diaryData.DiaryBasicInfo.Reminder.StartDate == "" {
		diaryData.DiaryBasicInfo.Reminder.StartDate = "1970.01.01"
	}
	if !diaryData.IsValid() {
		return domain.DiaryCreateResponse{}, domain.Err.ErrObj.InvalidTitleOrDescription
	}
	fmt.Printf("diaryData.DiaryBasicInfo.Reminder.StartDate: %v\n", diaryData.DiaryBasicInfo.Reminder.StartDate)

	DiaryCreateResponse, err := du.diaryRepo.CreateDiary(diaryData, medicId)
	if err != nil {
		return domain.DiaryCreateResponse{}, err
	}

	token, err := randomizer.GenerateRandomString(tokenLength)
	if err != nil {
		return domain.DiaryCreateResponse{}, err
	}
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
		return domain.DiaryLinkResponse{}, err
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

func (du DiaryUsecase) CompleteDiary(medicId, diaryId uint64) error {
	// alreadyExist, err := eu.diaryRepo.DiaryAlreadyExist(diaryData)
	// if err != nil {
	// 	return domain.DiaryCreateResponse{}, err
	// }

	// if alreadyExist {
	// 	return domain.DiaryCreateResponse{}, domain.Err.ErrObj.PlaylistExist
	// }


	err := du.diaryRepo.CompleteDiary(diaryId)
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

	if !isExisted {
		return domain.DiaryResponse{}, domain.Err.ErrObj.UserNotExist
	}

	if !isMedic {
		diary.Diary.DiaryBasicInfo.Anamnesis = ""
		diary.Diary.DiaryBasicInfo.Objectively = ""
		diary.Diary.DiaryBasicInfo.Diagnosis = ""
		diary.Diary.DiaryBasicInfo.Complaints = ""
	}

	return diary, nil
}

func (du DiaryUsecase) UpdateDiary(DiaryData domain.DiaryUpdateRequest, diaryId uint64) (domain.DiaryUpdateResponse, error) {
	// alreadyExist, err := eu.diaryRepo.DiaryAlreadyExist(diaryData)
	// if err != nil {
	// 	return domain.DiaryCreateResponse{}, err
	// }

	// if alreadyExist {
	// 	return domain.DiaryCreateResponse{}, domain.Err.ErrObj.PlaylistExist
	// }
	if DiaryData.DiaryBasicInfo.Reminder.StartDate == "" {
		DiaryData.DiaryBasicInfo.Reminder.StartDate = "1970.01.01"
	}
	if !DiaryData.IsValid() {
		return domain.DiaryUpdateResponse{}, domain.Err.ErrObj.InvalidTitleOrDescription
	}
	DiaryUpdateResponse, err := du.diaryRepo.UpdateDiary(DiaryData, diaryId)
	if err != nil {
		return domain.DiaryUpdateResponse{}, err
	}

	// DiaryCreateResponse.Categories, err = eu.eventRepo.CreateEventCategory(eventCreatingResponse.Id, eventData.Categories)
	// if err != nil {
	// 	return domain.EventCreatingResponse{}, err
	// }
	return DiaryUpdateResponse, nil
}

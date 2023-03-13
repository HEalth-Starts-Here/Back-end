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
		return domain.DiaryCreatingResponse{}, domain.Err.ErrObj.InvalidTitle
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

func (eu DiaryUsecase) GetDiary() (domain.DiaryListResponse, error) {
	
	feed, err := eu.diaryRepo.GetDiary()
	
	if err != nil {
		return domain.DiaryListResponse{}, err
	}

	return feed, nil
}

func (eu DiaryUsecase) GetCertainDiary(diaryId uint64) (domain.DiaryResponse, error) {

	diary, err := eu.diaryRepo.GetCertainDiary(diaryId)
	
	if err != nil {
		return domain.DiaryResponse{}, err
	}

	return diary, nil
}

func (eu DiaryUsecase) CreateRecord(diaryId uint64, recordData domain.RecordCreatingRequest) (domain.RecordCreatingResponse, error) {
	// alreadyExist, err := eu.diaryRepo.DiaryAlreadyExist(diaryData)
	// if err != nil {
	// 	return domain.DiaryCreatingResponse{}, err
	// }

	// if alreadyExist {
	// 	return domain.DiaryCreatingResponse{}, domain.Err.ErrObj.PlaylistExist
	// }

	if !recordData.IsValid() {
		return domain.RecordCreatingResponse{}, domain.Err.ErrObj.InvalidTitle
	}

	diaryCreatingResponse, err := eu.diaryRepo.CreateRecord(diaryId, recordData)
	if err != nil {
		return domain.RecordCreatingResponse{}, err
	}

	// diaryCreatingResponse.Categories, err = eu.eventRepo.CreateEventCategory(eventCreatingResponse.Id, eventData.Categories)
	// if err != nil {
	// 	return domain.EventCreatingResponse{}, err
	// }
	return diaryCreatingResponse, nil
}

// func (eu EventUsecase) GetCategory() (domain.CategoryListResponse, error) {

// 	categoryList, err := eu.eventRepo.GetCategory()
	
// 	if err != nil {
// 		return domain.CategoryListResponse{}, err
// 	}

// 	return categoryList, nil
// }

// func (eu EventUsecase) EventSignUp(eventId uint64, userId uint64) (error)  {

// 	userAge, err := eu.eventRepo.GetUserAge(userId)
// 	if err != nil {
// 		return err
// 	}
// 	eventMinAge, eventMaxAge, err := eu.eventRepo.GetEventAges(eventId)
// 	if err != nil {
// 		return err
// 	}


// 	isValidUser, err := eu.IsUserValidForEvent(eventMinAge, eventMaxAge, userAge)
// 	if err != nil {
// 		return err
// 	}
// 	if (!isValidUser){
// 		return domain.Err.ErrObj.BadInput
// 	}


// 	err = eu.eventRepo.SignUpUserForEvent(eventId, userId)
// 	if err != nil {
// 		return domain.Err.ErrObj.UserAlreadySignUpForThisEvent
// 	}

// 	return nil
// }

// func (eu EventUsecase) IsUserValidForEvent(minAge uint16, maxAge uint16, age uint64) (bool, error)  {
// 	if (cast.IntToStr(age) < cast.Uint16ToStr(minAge)){
// 		return false, nil
// 	}
// 	if (cast.IntToStr(age) > cast.Uint16ToStr(maxAge) && cast.Uint16ToStr(maxAge) != "0"){
// 		return false, nil
// 	}

// 	return true, nil
// }

// func (eu EventUsecase) CancelEventSignUp(eventId uint64, userId uint64) (error)  {

// 	err := eu.eventRepo.CancelEventSignUp(eventId, userId)
	
// 	if err != nil {
// 		return domain.Err.ErrObj.UserDontSignUpForThisEvent
// 	}

// 	return nil
// }

// func (eu EventUsecase) GetRecomendedEvent(userId uint64) (domain.EventListResponse, error) {

// 	categories, err := eu.eventRepo.GetUserCategory(userId)
	
// 	if err != nil {
// 		log.Error(err)
// 		return domain.EventListResponse{}, err
// 	}
	
// 	eventList, err := eu.GetEvent(categories)
	
// 	if err != nil {
// 		return domain.EventListResponse{}, err
// 	}

// 	return eventList, nil
// }

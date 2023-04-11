package recordusecase

import (
	"hesh/internal/pkg/domain"
	"hesh/internal/pkg/utils/config"
	"hesh/internal/pkg/utils/filesaver"
	"path/filepath"
)

type RecordUsecase struct {
	recordRepo domain.RecordRepository
}

// func trimTitle(title *string) {
// 	*title = strings.Trim(*title, " ")
// }

func InitRecordUsc(rr domain.RecordRepository) domain.RecordUsecase {
	return &RecordUsecase{
		recordRepo: rr,
	}
}

func (ru RecordUsecase) CheckMedicDiaryAccess(medicId uint64, diaryId uint64) (bool, error) {
	ownerId, err := ru.recordRepo.GetMedicIdFromDiary(diaryId)
	if err != nil {
		return false, err
	}
	if medicId != ownerId {
		return false, nil
	}
	return true, nil
}

func (ru RecordUsecase) CheckMedicRecordAccess(medicId uint64, medicRecordId uint64) (bool, error) {
	ownerId, err := ru.recordRepo.GetMedicIdFromDiaryOfRecord(medicRecordId)
	if err != nil {
		return false, err
	}
	if medicId != ownerId {
		return false, nil
	}
	return true, nil
}

func (ru RecordUsecase) CheckDiaryExist(diaryId uint64) (bool, error) {
	diaryExist, err := ru.recordRepo.DiaryExist(diaryId)
	if err != nil {
		return false, err
	}
	return diaryExist, nil
}

func (ru RecordUsecase) CheckRecordExist(recordId uint64) (bool, error) {
	medicRecordExist, err := ru.recordRepo.MedicRecordExist(recordId)
	if err != nil {
		return false, err
	}
	return medicRecordExist, nil
}

func (ru RecordUsecase) CheckMedicExist(medicId uint64) (bool, error) {
	medicExist, err := ru.recordRepo.MedicExist(medicId)
	if err != nil {
		return false, err
	}
	return medicExist, nil
}

func (ru RecordUsecase) CheckMedicAndDiaryExistAndMedicHaveAccess(medicId, diaryId uint64) (error) {
	
	medicExist, err := ru.CheckMedicExist(medicId)
	if err != nil {
		return err
	}
	if !medicExist {
		return domain.Err.ErrObj.MedicDoestExist
	}
	
	diaryExist, err := ru.CheckDiaryExist(diaryId)
	if err != nil {
		return err
	}
	if !diaryExist {
		return domain.Err.ErrObj.DiaryDoestExist
	}

	haveAccess, err := ru.CheckMedicDiaryAccess(medicId, diaryId)
	if err != nil {
		return err
	}
	if !haveAccess {
		return domain.Err.ErrObj.UserHaveNoAccess
	}
	return nil
}

func (ru RecordUsecase) CheckMedicAndDiaryAndRecordExistAndMedicHaveAccess(medicId, recordId uint64) (error) {
	medicExist, err := ru.CheckMedicExist(medicId)
	if err != nil {
		return err
	}
	if !medicExist {
		return domain.Err.ErrObj.MedicDoestExist
	}
	
	recordExist, err := ru.CheckRecordExist(recordId)
	if err != nil {
		return err
	}
	if !recordExist {
		return domain.Err.ErrObj.MedicRecordDoestExist
	}

	haveAccess, err := ru.CheckMedicRecordAccess(medicId, recordId)
	if err != nil {
		return err
	}
	if !haveAccess {
		return domain.Err.ErrObj.UserHaveNoAccess
	}
	return nil
}

func (ru RecordUsecase) CreateMedicRecord(diaryId uint64, medicId uint64, recordData domain.MedicRecordCreateRequest) (domain.MedicRecordCreateResponse, error) {
	err := ru.CheckMedicAndDiaryExistAndMedicHaveAccess(medicId, diaryId)
	if err != nil {
		return domain.MedicRecordCreateResponse{}, err
	}
	// ru.CheckMedicAndDiaryAndRecordExistAndMedicHaveAccess()
	// diaryExist, err := ru.CheckDiaryExist(diaryId)
	// if err != nil {
	// 	return domain.MedicRecordCreateResponse{}, err
	// }
	// if !diaryExist {
	// 	return domain.MedicRecordCreateResponse{}, domain.Err.ErrObj.DiaryDoestExist
	// }
	
	// medicExist, err := ru.CheckMedicExist(medicId)
	// if err != nil {
	// 	return domain.MedicRecordCreateResponse{}, err
	// }
	// if !medicExist {
	// 	return domain.MedicRecordCreateResponse{}, domain.Err.ErrObj.MedicDoestExist
	// }
	
	// haveAccess, err := ru.CheckMedicAccess(medicId, diaryId)
	// if err != nil {
	// 	return domain.MedicRecordCreateResponse{}, err
	// }
	// if !haveAccess {
	// 	return domain.MedicRecordCreateResponse{}, domain.Err.ErrObj.UserHaveNoAccess
	// }

	if !recordData.IsValid() {
		return domain.MedicRecordCreateResponse{}, domain.Err.ErrObj.InvalidTitleOrDescription
	}
	alreadyUsed, err := ru.recordRepo.GetImageNames()
	if err != nil {
		return domain.MedicRecordCreateResponse{}, err
	}
	imageNames := filesaver.GetUniqueFileNames(len(recordData.Images), alreadyUsed)
	for i := 0; i < len(imageNames); i++ {
		recordData.Images[i].ImageName = imageNames[i] + filepath.Ext(recordData.Images[i].ImageName)
		imageNames[i] = recordData.Images[i].ImageName

	}
	RecordCreateResponse, err := ru.recordRepo.CreateMedicRecord(diaryId, recordData)
	if err != nil {
		return domain.MedicRecordCreateResponse{}, err
	}

	// imagenames := []string{}
	// for i := range recordData.Images {
	// 	imagenames = append(imagenames, recordData.Images[i].ImageName)
	// }
	// println(imageNames)
	// TODO check case with 0 images
	_, err = ru.recordRepo.CreateRecordImageLists(true, RecordCreateResponse.Id, imageNames)
	if err != nil {
		return domain.MedicRecordCreateResponse{}, err
	}
	for i := range RecordCreateResponse.ImageList {
		RecordCreateResponse.ImageList[i].ImageName = imageNames[i]
	}

	tags := [][]string{}
	for i := range recordData.Images {
		tags = append(tags, recordData.Images[i].Tags)
	}

	// imageIds, tags, err = ru.recordRepo.CreateImageTags(imageIds, tags)
	// for i := range imageIds {

	// 	tags = append(tags, recordData.Images[i].Tags)
	// }
	return RecordCreateResponse, nil
}

func (ru RecordUsecase) GetMedicRecord (userId, recordId uint64) (domain.MedicRecordCreateResponse, error) {
	// alreadyExist, err := eu.diaryRepo.DiaryAlreadyExist(diaryData)
	// if err != nil {
	// 	return domain.DiaryCreateResponse{}, err
	// }

	// if alreadyExist {
	// 	return domain.DiaryCreateResponse{}, domain.Err.ErrObj.PlaylistExist
	// }
	
	// TODO: Check if this user has access to this record


	diaryId, recordId, creatingDate, BasicRecordInfo, err := ru.recordRepo.GetRecordTextInfo(true, recordId)
	if err != nil {
		return domain.MedicRecordCreateResponse{}, err
	}
	
	response := domain.MedicRecordCreateResponse{
		DiaryId: diaryId,
		Id: recordId,
		CreatingDate: creatingDate,
		BasicInfo: BasicRecordInfo,
		// Diarisation: diarisation,
		ImageList: nil,
	}
	imageNames, err := ru.recordRepo.GetRecordImageNames(true, recordId)
	if err != nil {
		return domain.MedicRecordCreateResponse{}, err
	}
	

	response.ImageList = make([]domain.RecordImageInfo, 0)
	for i := range imageNames {
		response.ImageList = append(response.ImageList, domain.RecordImageInfo{
			ImageName: imageNames[i],
			Tags: nil,
		})
	}
	return response, nil
}

func (ru RecordUsecase) GetMedicRecordDiarisations (userId, medicRecordId uint64) (domain.GetDiarisationsResponse, error) {
	// alreadyExist, err := eu.diaryRepo.DiaryAlreadyExist(diaryData)
	// if err != nil {
	// 	return domain.DiaryCreateResponse{}, err
	// }

	// if alreadyExist {
	// 	return domain.DiaryCreateResponse{}, domain.Err.ErrObj.PlaylistExist
	// }
	
	// TODO: Check if this user has access to this record
	response, err := ru.recordRepo.GetMedicRecordDiarisations(medicRecordId)
	if err != nil {
		return domain.GetDiarisationsResponse{}, err
	}
	return response, nil
}


func (ru RecordUsecase) UpdateMedicRecordText(medicId uint64, recordId uint64, updateTextMedicRecordData domain.MedicRecordBasicInfo) (domain.MedicRecordUpdateTextResponse, error) {
	// alreadyExist, err := eu.diaryRepo.DiaryAlreadyExist(diaryData)
	// if err != nil {
	// 	return domain.DiaryCreateResponse{}, err
	// }

	// if alreadyExist {
	// 	return domain.DiaryCreateResponse{}, domain.Err.ErrObj.PlaylistExist
	// }

	if !updateTextMedicRecordData.IsValid() {
		return domain.MedicRecordUpdateTextResponse{}, domain.Err.ErrObj.InvalidTitleOrDescription
	}
	MedicRecordTextUpdateResponse, err := ru.recordRepo.UpdateMedicRecordText(recordId, updateTextMedicRecordData)
	if err != nil {
		return domain.MedicRecordUpdateTextResponse{}, err
	}

	// DiaryCreateResponse.Categories, err = eu.eventRepo.CreateEventCategory(eventCreatingResponse.Id, eventData.Categories)
	// if err != nil {
	// 	return domain.EventCreatingResponse{}, err
	// }
	return MedicRecordTextUpdateResponse, nil
}

func (ru RecordUsecase) UpdateMedicRecordImage(medicId uint64, recordId uint64, updateImageMedicRecordData domain.MedicRecordUpdateImageRequest) (domain.RecordUpdateImageResponse, error) {
	// TODO: add check if this medic is owner of this record
	
	if !updateImageMedicRecordData.IsValid() {
		return domain.RecordUpdateImageResponse{}, domain.Err.ErrObj.InvalidTitleOrDescription
	}

	alreadyUsed, err := ru.recordRepo.GetImageNames()
	if err != nil {
		return domain.RecordUpdateImageResponse{}, err
	}
	imageNames := filesaver.GetUniqueFileNames(len(updateImageMedicRecordData.Images), alreadyUsed)
	for i := 0; i < len(imageNames); i++ {
		updateImageMedicRecordData.Images[i].ImageName = imageNames[i] + filepath.Ext(updateImageMedicRecordData.Images[i].ImageName)
		imageNames[i] = imageNames[i] + filepath.Ext(updateImageMedicRecordData.Images[i].ImageName)

	}
	updateResponse, err := ru.recordRepo.DeleteRecordImage(true, recordId)
	if err != nil {
		return domain.RecordUpdateImageResponse{}, err
	}
	deletedImages := make([]string, 0)
	for i := range updateResponse.Images{
		deletedImages = append(deletedImages, updateResponse.Images[i].ImageName)
	}
	err = filesaver.DeleteFiles("", config.DevConfigStore.LoadedFilesPath, deletedImages)
	if err != nil {
		return domain.RecordUpdateImageResponse{}, err
	}
	updateResponse.Images = make([]domain.RecordImageInfo, 0)
	_, err = ru.recordRepo.CreateRecordImageLists(true, recordId, imageNames)
	if err != nil {
		return domain.RecordUpdateImageResponse{}, err
	}
	// response := domain.RecordUpdateImageResponse{}
	for i := range updateImageMedicRecordData.Images {
		updateResponse.Images = append(updateResponse.Images, domain.RecordImageInfo{
			ImageName: imageNames[i],
			Tags: []string{},
		})
	}
	//TODO update tags

	// imageIds, tags, err = ru.recordRepo.CreateImageTags(imageIds, tags)
	// for i := range imageIds {

	// 	tags = append(tags, recordData.Images[i].Tags)
	// }
	return updateResponse, nil
}

func (ru RecordUsecase) DeleteMedicRecord(medicId uint64, recordId uint64) (error) {
	// TODO: add check if this medic is owner of this record

	imageList, err := ru.recordRepo.GetRecordImageNames(true, recordId)
	if err != nil {
		return err
	}
	
	err = ru.recordRepo.DeleteRecord(true, recordId)
	if err != nil {
		return err
	}
	// deleteResponse, err := ru.recordRepo.DeleteRecordImage(true, recordId)
	// if err != nil {
	// 	return err
	// }
	// deletedImages := make([]string, 0)
	// for i := range deleteResponse.Images{
	// 	deletedImages = append(deletedImages, deleteResponse.Images[i].ImageName)
	// }
	err = filesaver.DeleteFiles("", config.DevConfigStore.LoadedFilesPath, imageList)
	if err != nil {
		return err
	}
	
	//TODO delete tags

	// imageIds, tags, err = ru.recordRepo.CreateImageTags(imageIds, tags)
	// for i := range imageIds {

	// 	tags = append(tags, recordData.Images[i].Tags)
	// }
	return nil
}
package recordusecase

import (
	"hesh/internal/pkg/domain"
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

func (ru RecordUsecase) CreateMedicRecord(diaryId uint64, medicId uint64, recordData domain.MedicRecordCreateRequest) (domain.MedicRecordCreateResponse, error) {
	// alreadyExist, err := eu.diaryRepo.DiaryAlreadyExist(diaryData)
	// if err != nil {
	// 	return domain.DiaryCreateResponse{}, err
	// }

	// if alreadyExist {
	// 	return domain.DiaryCreateResponse{}, domain.Err.ErrObj.PlaylistExist
	// }
	
	// TODO: Check if this user has access to this diary 
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


	diaryId, recordId, creatingDate, BasicRecordInfo,err := ru.recordRepo.GetRecordTextInfo(true, recordId)
	if err != nil {
		return domain.MedicRecordCreateResponse{}, err
	}
	
	response := domain.MedicRecordCreateResponse{
		DiaryId: diaryId,
		Id: recordId,
		CreatingDate: creatingDate,
		BasicInfo: BasicRecordInfo,
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

package domain

// TODO write valid path
const (

)

type DetermineAreaResponse struct {
	Area 		int32		`json:"area"`
}

type MLServicesRepository interface {
	// CreateDiary(diary DiaryCreatingRequest) (DiaryCreatingResponse, error)
	// GetDiary() (DiaryListResponse, error)  
	// GetCertainDiary(diaryId uint64) (DiaryResponse, error)
	// CreateRecord(diaryId uint64, record RecordCreatingRequest, imageInfo []ImageInfoUsecase, Area float64) (RecordCreatingResponse, error)
}

type MLServicesUsecase interface {
	// CreateDiary(diary DiaryCreatingRequest) (DiaryCreatingResponse, error)
	// GetDiary() (DiaryListResponse, error) 
	// GetCertainDiary(diaryId uint64) (DiaryResponse, error)
	// CreateRecord(diaryId uint64, record RecordCreatingRequest, imageInfo []ImageInfoUsecase) (RecordCreatingResponse, error)
}

package domain

// TODO write valid path
const ()

type DetermineAreaResponse struct {
	Area int32 `json:"area"`
}

type MLServicesRepository interface {
	// CreateDiary(diary DiaryCreateRequest) (DiaryCreateResponse, error)
	// GetDiary() (DiaryListResponse, error)
	// GetCertainDiary(diaryId uint64) (DiaryResponse, error)
	// CreateRecord(diaryId uint64, record RecordCreateRequest, imageInfo []ImageInfoUsecase, Area float64) (RecordCreateResponse, error)
}

type MLServicesUsecase interface {
	// CreateDiary(diary DiaryCreateRequest) (DiaryCreateResponse, error)
	// GetDiary() (DiaryListResponse, error)
	// GetCertainDiary(diaryId uint64) (DiaryResponse, error)
	// CreateRecord(diaryId uint64, record RecordCreateRequest, imageInfo []ImageInfoUsecase) (RecordCreateResponse, error)
}

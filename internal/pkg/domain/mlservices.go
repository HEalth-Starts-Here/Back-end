package domain

// TODO write valid path
const ()

type ImageQualityAssesment struct {
	Assesment bool `json:"assesment"`
}

type DetermineAreaResponse struct {
	Area int32 `json:"area"`
}

type DiarisationInfo struct {
	Diarisation string `json:"diarisation"`
	Filename 	string `json:"filename"`
}

type MedicRecordDiarisationsResponse struct {
	DiarisationsInfo DiarisationInfo `json:"diarisationsinfo"`
}

type DiarisationResponse struct {
	Id				uint64 `json:"id"`
	CreatingDate	string `json:"creatingdate"`
	MedicRecordId	uint64 `json:"medicrecordid"`
	DiarisationInfo DiarisationInfo
}

type MLServicesRepository interface {
	GetAudioNames() (map[string]struct{}, error)
	CreateMedicRecordDiarisation (medicRecordId uint64, DiarisationInfo DiarisationInfo) (DiarisationResponse, error)
	// CreateDiary(diary DiaryCreateRequest) (DiaryCreateResponse, error)
	// GetDiary() (DiaryListResponse, error)
	// GetCertainDiary(diaryId uint64) (DiaryResponse, error)
	// CreateRecord(diaryId uint64, record RecordCreateRequest, imageInfo []ImageInfoUsecase, Area float64) (RecordCreateResponse, error)
}

type MLServicesUsecase interface {
	CreateMedicRecordDiarisations(medicId uint64, recordId uint64, DiarisationInfo DiarisationInfo) (DiarisationResponse, error)

	// CreateDiary(diary DiaryCreateRequest) (DiaryCreateResponse, error)
	// GetDiary() (DiaryListResponse, error)
	// GetCertainDiary(diaryId uint64) (DiaryResponse, error)
	// CreateRecord(diaryId uint64, record RecordCreateRequest, imageInfo []ImageInfoUsecase) (RecordCreateResponse, error)
}

package domain

const ()

type ImageQualityAssesment struct {
	Assesment bool `json:"assesment"`
}

type DetermineAreaResponse struct {
	Area int32 `json:"area"`
}

type DiarisationBeforeCompletingInfo struct {
	// Diarisation string `json:"diarisation"`
	Filename string `json:"filename"`
}

type DiarisationInfo struct {
	Diarisation string `json:"diarisation"`
	Filename    string `json:"filename"`
	IsComplete  bool   `json:"iscomplete"`
}

type DiarisationResponse struct {
	Id              uint64          `json:"id"`
	CreatingDate    string          `json:"creatingdate"`
	MedicRecordId   uint64          `json:"medicrecordid"`
	DiarisationInfo DiarisationInfo `json:"diarisationinfo"`
}

type MLServicesRepository interface {
	GetAudioNames() (map[string]struct{}, error)
	CreateMedicRecordDiarisation(medicRecordId uint64, DiarisationInfo DiarisationBeforeCompletingInfo) (DiarisationResponse, error)
	SetDiarisationText(diarisationId uint64, diarisationText string) error
}

type MLServicesUsecase interface {
	CreateMedicRecordDiarisations(medicId uint64, recordId uint64, DiarisationInfo DiarisationBeforeCompletingInfo) (DiarisationResponse, error)
	SetDiarisationText(diarisationId uint64, diarisationText string) error
	// CreateMedicRecordDiarisations(medicId uint64, recordId uint64, DiarisationInfo DiarisationInfo) (DiarisationResponse, error)
}

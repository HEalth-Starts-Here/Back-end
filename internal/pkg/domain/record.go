package domain

const (
	maxTitleLenght          = 50
	maxTreatmentLenght      = 1000
	maxRecomendationsLenght = 1000
	maxComplaintsLenght     = 1000
	maxDetailsLenght        = 3000
	maxImageNameLenght      = 200
	maxTagsLenght           = 50
)

func (record *MedicRecordBasicInfo) SetDefault() {
	record.Title = ""
	record.Treatment = ""
	record.Recommendations = ""
	record.Details = ""
}

func (record *RecordImageInfo) SetDefault() {
	record.ImageName = ""
	record.Tags = nil
}

func (record *MedicRecordCreateRequest) SetDefault() {
	record.BasicInfo.SetDefault()
	for i := range record.Images {
		record.Images[i].SetDefault()
	}
}

func (record *PatientRecordBasicInfo) SetDefault() {
	record.Title = ""
	record.Complaints = ""
	record.Treatment = ""
	record.Details = ""
}

func (record *PatientRecordCreateRequest) SetDefault() {
	record.BasicInfo.SetDefault()
	for i := range record.Images {
		record.Images[i].SetDefault()
	}
}

// TODO: add returning errors
func (record MedicRecordBasicInfo) IsValid() bool {
	if len(record.Title) > maxTitleLenght ||
		len(record.Treatment) > maxDetailsLenght ||
		len(record.Recommendations) > maxRecomendationsLenght ||
		len(record.Details) > maxDetailsLenght {
		return false
	}
	return true
}

// TODO: add returning errors
func (request RecordUpdateImageRequest) IsValid() bool {
	for i := range request.Images {
		if !(request.Images[i].IsValid()) {
			return false
		}
	}
	return true
}

func (imageInfo RecordImageInfo) IsValid() bool {
	if len(imageInfo.ImageName) > maxImageNameLenght {
		return false
	}
	for i := range imageInfo.Tags {
		if len(imageInfo.Tags[i]) > maxTagsLenght {
			return false
		}
	}
	return true
}

func (medicRecord MedicRecordCreateRequest) IsValid() bool {
	for i := range medicRecord.Images {
		if !medicRecord.Images[i].IsValid() {
			return false
		}
	}
	if !medicRecord.BasicInfo.IsValid() {
		return false
	}
	return true
}

func (record PatientRecordBasicInfo) IsValid() bool {
	if len(record.Title) > maxTitleLenght ||
		len(record.Complaints) > maxComplaintsLenght ||
		len(record.Treatment) > maxDetailsLenght ||
		len(record.Details) > maxDetailsLenght {
		return false
	}
	return true
}

func (patientRecord PatientRecordCreateRequest) IsValid() bool {
	for i := range patientRecord.Images {
		if !patientRecord.Images[i].IsValid() {
			return false
		}
	}
	if !patientRecord.BasicInfo.IsValid() {
		return false
	}
	return true
}

type MedicRecordBasicInfo struct {
	Title           string `json:"title"`
	Treatment       string `json:"treatment"`
	Recommendations string `json:"recommendations"`
	Details         string `json:"details"`
}

type RecordImageInfo struct {
	ImageName string   `json:"imagename"`
	Tags      []string `json:"tags"`
}

type MedicRecordCreateRequest struct {
	BasicInfo MedicRecordBasicInfo `json:"basicinfo"`
	Images    []RecordImageInfo    `json:"images"`
	// Auido			[]string		 	 `json:"audio"`
	// Diarisation		string			 	 `json:"diarisation"`
}

type MedicRecordCreateResponse struct {
	DiaryId      uint64               `json:"diaryid"`
	Id           uint64               `json:"id"`
	CreatingDate string               `json:"creatingdate"`
	BasicInfo    MedicRecordBasicInfo `json:"basicinfo"`
	ImageList    []RecordImageInfo    `json:"imagelist"`
	// Diarisation		string					`json:"diarisation"`
}

type MedicRecordUpdateTextResponse struct {
	DiaryId      uint64               `json:"diaryid"`
	Id           uint64               `json:"id"`
	CreatingDate string               `json:"creatingdate"`
	BasicInfo    MedicRecordBasicInfo `json:"basicinfo"`
}

type RecordUpdateImageRequest struct {
	Images []RecordImageInfo `json:"images"`
}

type RecordUpdateImageResponse struct {
	Id           uint64            `json:"id"`
	DiaryId      uint64            `json:"diaryId"`
	CreatingDate string            `json:"creatingdate"`
	Images       []RecordImageInfo `json:"images"`
}

type DiarisationInListResponse struct {
	Id              uint64 `json:"id"`
	CreatingDate    string `json:"creatingdate"`
	DiarisationInfo DiarisationInfo
}

type GetDiarisationsResponse struct {
	MedicRecordId   uint64                      `json:"medicrecordid"`
	DiarisationList []DiarisationInListResponse `json:"diarisationlist"`
}

type PatientRecordBasicInfo struct {
	Title      string `json:"title"`
	Complaints string `json:"complaints"`
	Treatment  string `json:"treatment"`
	Details    string `json:"details"`
}

type PatientRecordCreateRequest struct {
	BasicInfo PatientRecordBasicInfo `json:"basicinfo"`
	Images    []RecordImageInfo      `json:"images"`
	// Diarisation		string			 	 `json:"diarisation"`
}

type PatientRecordCreateResponse struct {
	DiaryId      uint64                 `json:"diaryid"`
	Id           uint64                 `json:"id"`
	CreatingDate string                 `json:"creatingdate"`
	BasicInfo    PatientRecordBasicInfo `json:"basicinfo"`
	ImageList    []RecordImageInfo      `json:"imagelist"`
	// Diarisation		string					`json:"diarisation"`
}

type PatientRecordUpdateTextResponse struct {
	DiaryId      uint64                 `json:"diaryid"`
	Id           uint64                 `json:"id"`
	CreatingDate string                 `json:"creatingdate"`
	BasicInfo    PatientRecordBasicInfo `json:"basicinfo"`
}

type RecordRepository interface {
	// MEDIC
	CreateMedicRecord(diaryId uint64, record MedicRecordCreateRequest) (MedicRecordCreateResponse, error)
	GetImageNames() (map[string]struct{}, error)
	CreateRecordImageLists(isMedic bool, recordId uint64, imageInfo []string) ([]uint64, error)
	CreateImageTags(imageIds []uint64, tags [][]string) ([]uint64, [][]string, error)
	GetRecordTextInfo(recordId uint64) (uint64, uint64, string, MedicRecordBasicInfo, error)
	GetRecordImageNames(isMedic bool, recordId uint64) ([]string, error)
	UpdateMedicRecordText(recordId uint64, medicRecordBasicInfo MedicRecordBasicInfo) (MedicRecordUpdateTextResponse, error)
	DeleteRecordImage(isMedic bool, recordId uint64) (RecordUpdateImageResponse, error)
	GetMedicIdFromDiary(diaryId uint64) (uint64, error)
	GetMedicIdFromDiaryOfRecord(recordId uint64) (uint64, error)
	DiaryExist(diaryId uint64) (bool, error)
	MedicExist(medicId uint64) (bool, error)
	MedicRecordExist(medicId uint64) (bool, error)
	DeleteRecord(isMedic bool, recordId uint64) error
	GetMedicRecordDiarisations(medicRecordId uint64) (GetDiarisationsResponse, error)

	// PATIENT
	CreatePatientRecord(diaryId uint64, record PatientRecordCreateRequest) (PatientRecordCreateResponse, error)
	GetPatientRecordTextInfo(recordId uint64) (uint64, uint64, string, PatientRecordBasicInfo, error)
	UpdatePatientRecordText(recordId uint64, patientRecordBasicInfo PatientRecordBasicInfo) (PatientRecordUpdateTextResponse, error)
}

type RecordUsecase interface {
	DeleteRecord(isMedic bool, medicId uint64, recordId uint64) error
	// MEDIC
	CreateMedicRecord(diaryId uint64, medicId uint64, record MedicRecordCreateRequest) (MedicRecordCreateResponse, error)
	GetMedicRecord(userId, recordId uint64) (MedicRecordCreateResponse, error)
	UpdateMedicRecordText(medicId uint64, recordId uint64, medicRecordBasicInfo MedicRecordBasicInfo) (MedicRecordUpdateTextResponse, error)
	UpdateRecordImage(isMedic bool, medicId uint64, recordId uint64, updateTextMedicRecordData RecordUpdateImageRequest) (RecordUpdateImageResponse, error)
	CheckMedicDiaryAccess(medicId uint64, diaryId uint64) (bool, error)
	CheckMedicRecordAccess(medicId uint64, diaryId uint64) (bool, error)
	CheckDiaryExist(diaryId uint64) (bool, error)
	CheckRecordExist(recordId uint64) (bool, error)
	CheckMedicExist(medicId uint64) (bool, error)
	CheckMedicAndDiaryExistAndMedicHaveAccess(medicId, diaryId uint64) error
	CheckMedicAndDiaryAndRecordExistAndMedicHaveAccess(medicId, diaryId uint64) error
	GetMedicRecordDiarisations(userId, medicRecordId uint64) (GetDiarisationsResponse, error)

	// PATIENT
	CreatePatientRecord(patientId, diaryId uint64, record PatientRecordCreateRequest) (PatientRecordCreateResponse, error)
	GetPatientRecord(userId, recordId uint64) (PatientRecordCreateResponse, error)
	UpdatePatientRecordText(patientId uint64, recordId uint64, patientRecordBasicInfo PatientRecordBasicInfo) (PatientRecordUpdateTextResponse, error)
}

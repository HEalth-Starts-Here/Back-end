package domain

const (
	maxTitleLenght = 50
	maxTreatmentLenght = 1000
	maxRecomendationsLenght = 1000
	maxDetailsLenght = 3000
	maxImageNameLenght = 200
	maxTagsLenght = 50
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
	for i := range record.Images{
		record.Images[i].SetDefault()
	}
}


// TODO: add returning errors
func (record MedicRecordBasicInfo) IsValid() (bool) {
	if  len(record.Title) > maxTitleLenght  || 
		len(record.Treatment) > maxDetailsLenght  || 
		len(record.Recommendations) > maxRecomendationsLenght ||
		len(record.Details) > maxDetailsLenght{
		return false
	}
	return true
}

func (imageInfo RecordImageInfo) IsValid() (bool) {
	if  len(imageInfo.ImageName) > maxImageNameLenght{
		return false
	}
	for i := range(imageInfo.Tags){
		if len(imageInfo.Tags[i]) > maxTagsLenght{
			return false
		}
	}
	return true
}

func (medicRecord MedicRecordCreateRequest) IsValid() (bool) {
	for i := range medicRecord.Images{
		if !medicRecord.Images[i].IsValid(){
			return false
		}
	}
	if  !medicRecord.BasicInfo.IsValid () {
		return false
	}
	return true
}

type MedicRecordBasicInfo struct {
	Title			string	`json:"title"`
	Treatment		string	`json:"treatment"`
	Recommendations string	`json:"recommendations"`
	Details			string	`json:"details"`
}

type RecordImageInfo struct {
	ImageName		string	 `json:"imagename"`
	Tags			[]string `json:"tags"`
}

type MedicRecordCreateRequest struct {
	BasicInfo	MedicRecordBasicInfo `json:"basicinfo"` 
	Images		[]RecordImageInfo 	 `json:"images"` 
}

type MedicRecordCreateResponse struct {
	DiaryId			uint64					`json:"diaryid"` 
	Id				uint64					`json:"id"` 
	CreatingDate	string					`json:"creatingdate"` 
	BasicInfo		MedicRecordBasicInfo	`json:"basicinfo"`
	ImageList		[]RecordImageInfo		`json:"imagelist"`
}

type RecordRepository interface {
	CreateMedicRecord(diaryId uint64, record MedicRecordCreateRequest) (MedicRecordCreateResponse, error)
	GetImageNames() (map[string]struct{}, error)
	CreateRecordImageLists(isMedic bool,recordId uint64, imageInfo []string) ([]uint64, error) 
	CreateImageTags(imageIds []uint64, tags [][]string) ([]uint64, [][]string, error) 
	GetRecordTextInfo(isMedic bool, recordId uint64,) (uint64, uint64, string, MedicRecordBasicInfo, error) 
	GetRecordImageNames(isMedic bool, recordId uint64) ([]string, error) 
}

type RecordUsecase interface {
	CreateMedicRecord(diaryId uint64, medicId uint64, record MedicRecordCreateRequest) (MedicRecordCreateResponse, error)
	GetMedicRecord(userId, recordId uint64) (MedicRecordCreateResponse, error)
}

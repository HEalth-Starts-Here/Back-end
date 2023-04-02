package domain

// TODO write valid path
const (
	// BaseEventPicture = "/home/ubuntu/lolkek/static/event/event.png"
	maxDiaryTitleLength        = 200
	maxRecordTitleLength       = 200
	maxDiaryObjectivelyLength  = 1000
	maxRecordDescriptionLength = 3000
)

func (er *DiaryCreateRequest) SetDefault() {
	er.DiaryBasicInfo.Title = ""
	er.DiaryBasicInfo.Complaints = ""
	er.DiaryBasicInfo.Anamnesis = ""
	er.DiaryBasicInfo.Objectively = ""
	er.DiaryBasicInfo.Diagnosis = ""
	return
}

func (er *RecordCreateRequest) SetDefault() {
	er.Description = ""
	er.Title = ""
	er.Characteristics = Characteristics{}
	return
}

func (er *RecordCreateResponse) SetDefault() {
	er.Id = 0
	er.DiaryId = 0
	er.CreatingDate = "2022-04-10 15:47:24"
	er.Description = ""
	er.Title = ""
	er.Area = 0
	er.Characteristics = Characteristics{}
	er.ImageList = []ImageInfo{}
	return
}

func (er DiaryCreateRequest) IsValid() (isValid bool) {
	if len(er.DiaryBasicInfo.Title) > maxDiaryTitleLength || len(er.DiaryBasicInfo.Objectively) > maxDiaryObjectivelyLength {
		return false
	}
	return true
}

func (er RecordCreateRequest) IsValid() (isValid bool) {
	if len(er.Title) > maxRecordTitleLength || len(er.Description) > maxRecordDescriptionLength {
		return false
	}
	characteristicsList := [](*uint8){&er.Characteristics.Dryness,
		&er.Characteristics.Edema,
		&er.Characteristics.Itching,
		&er.Characteristics.Pain,
		&er.Characteristics.Peeling,
		&er.Characteristics.Redness,
	}
	for ch := range characteristicsList {
		if ch > 10 {
			return false
		}
	}
	return true
}

func (er DiaryUpdateRequest) IsValid() (isValid bool) {
	if len(er.DiaryBasicInfo.Title) > maxDiaryTitleLength || len(er.DiaryBasicInfo.Objectively) > maxDiaryObjectivelyLength {
		return false
	}
	return true
}

type DiaryUpdateRequest struct {
	DiaryBasicInfo	DiaryBasicInfo `json:"diarybasicinfo"`
}

type DiaryUpdateResponse struct {
	Id				uint64 `json:"id"`
	MedicId			uint32 `json:"medicid"`
	PatientId		uint32 `json:"patientid"`
	CreatingDate	string `json:"creatingdate"`
	DiaryBasicInfo	DiaryBasicInfo `json:"diarybasicinfo"`
}

type DiaryBasicInfo struct {
	Title       string `json:"title"`
	Complaints  string `json:"complaints"`
	Anamnesis   string `json:"anamnesis"`
	Objectively string `json:"objectively"`
	Diagnosis   string `json:"diagnosis"`
}

type DiaryCreateRequest struct {
	// MedicId		   uint32 `json:"medicid"`
	// PatientId	   uint32 `json:"patientid"`
	DiaryBasicInfo DiaryBasicInfo `json:"diarybasicinfo"`
}

type DiaryCreateResponse struct {
	Id             uint64 `json:"id"`
	MedicId        uint32 `json:"medicid"`
	PatientId      uint32 `json:"patientid"`
	CreatingDate   string `json:"creatingdate"`
	DiaryBasicInfo DiaryBasicInfo `json:"diarybasicinfo"`
}

type DiaryLinkResponse struct {
	Id             uint64 `json:"id"`
	MedicId        uint32 `json:"medicid"`
	MedicName      string `json:"medicname"`
	PatientId      uint32 `json:"patientid"`
	CreatingDate   string `json:"creatingdate"`
	DiaryBasicInfo DiaryBasicInfo `json:"diarybasicinfo"`
}

type DiaryInList struct {
	Id				uint64 `json:"id"`
	Title   		string `json:"title"`
	MedicId			uint32 `json:"medicid"`
	MedicName   	string `json:"medicname"`
	PatientId   	uint32 `json:"patientid"`
	PatientName		string `json:"patientname"`
	CreatingDate	string `json:"creatingdate"`
	Objectively		string `json:"objectively"`
}

type DiaryListResponse struct {
	DiaryList []DiaryInList `json:"diarylist"`
}

type RecordBasicInfo struct {
	CreatingDate		string `json:"creatingdate"`
	Title				string `json:"title"`
	Details				string `json:"details"`
}

type Records struct { 
	MedicRecordList   []RecordBasicInfo `json:"medicrecordlist"`
	PatientRecordList []RecordBasicInfo `json:"patientrecordlist"`
}

type DiaryResponse struct {
	PatientName			string `json:"patientname"`
	Diary				DiaryLinkResponse `json:"diary"`
	Records				Records `json:"records"`
}

type RecordCreateRequest struct {
	Title           string          `json:"title"`
	Description     string          `json:"description"`
	Characteristics Characteristics `json:"characteristics"`
}

type ImageInfo struct {
	Id       uint64  `json:"id"`
	RecordId uint64  `json:"recordid"`
	Name     string  `json:"name"`
	Area     float64 `json:"area"`
}

type RecordCreateResponse struct {
	Id              uint64          `json:"id"`
	DiaryId         uint64          `json:"diaryid"`
	CreatingDate    string          `json:"creatingdate"`
	Description     string          `json:"description"`
	Title           string          `json:"title"`
	Area            float64         `json:"area"`
	Characteristics Characteristics `json:"characteristics"`
	ImageList       []ImageInfo     `json:"imagelist"`
}

type RecordUpdateRequest struct {
	Id                  uint64              `json:"id"`
	RecordCreateRequest RecordCreateRequest `json:"RecordCreateRequest"`
}

type Characteristics struct {
	Dryness uint8 `json:"dryness"`
	Edema   uint8 `json:"edema"`
	Itching uint8 `json:"itching"` // defin all not required
	Pain    uint8 `json:"pain"`
	Peeling uint8 `json:"peeling"`
	Redness uint8 `json:"redness"`
}

type ImageInfoUsecase struct {
	Name string  `json:"name"`
	Area float64 `json:"area"`
}

// Area 				   float32 `json:"area"` // cm^2
// // TODO Define palm area as 1%

type DiaryRepository interface {
	CreateDiary(diary DiaryCreateRequest, medicId uint32) (DiaryCreateResponse, error)
	LinkDiary(diaryId uint64, medicId uint32) (DiaryLinkResponse, error)
	DeleteDiary(diaryid uint64) error
	GetDiary(userId uint32) (DiaryListResponse, error)
	GetCertainDiary(diaryId uint64) (DiaryResponse, error)
	CreateRecord(diaryId uint64, record RecordCreateRequest, imageInfo []ImageInfoUsecase, Area float64) (RecordCreateResponse, error)
	UpdateDiary(diary DiaryUpdateRequest, diaryId uint64) (DiaryUpdateResponse, error)
	GetImageNames() (map[string]struct{}, error)

	// DiaryAlreadyExist(diary DiaryCreateRequest) (bool, error)
	// GetCategory() (CategoryListResponse, error)
	// CreateEventCategory(eventId uint64, categories []string) ([]string, error)
	// SignUpUserForEvent(eventId uint64, userId uint64) (error)
	// CancelEventSignUp(eventId uint64, userId uint64) (error)
	// GetUserCategory(id uint64) ([]string, error)
	// GetUserAge(id uint64) (uint64, error)
	// GetEventAges(id uint64) (uint16, uint16, error)
}

type DiaryUsecase interface {
	CreateDiary(diary DiaryCreateRequest, medicId uint32) (DiaryCreateResponse, error)
	LinkDiary(diaryId uint64, medicId uint32) (DiaryLinkResponse, error)
	DeleteDiary(diaryid uint64) error
	GetDiary(userId uint32) (DiaryListResponse, error)
	GetCertainDiary(diaryId uint64) (DiaryResponse, error)
	CreateRecord(diaryId uint64, record RecordCreateRequest, imageInfo []ImageInfoUsecase) (RecordCreateResponse, error)
	UpdateDiary(diary DiaryUpdateRequest, diaryId uint64) (DiaryUpdateResponse, error)

	// GetCategory() (CategoryListResponse, error)
	// EventSignUp(eventId uint64, userId uint64)(error)
	// CancelEventSignUp(eventId uint64, userId uint64) (error)
	// GetRecomendedEvent(userId uint64) (EventListResponse, error)
}

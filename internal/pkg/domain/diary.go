package domain

// TODO write valid path
const (
	// BaseEventPicture = "/home/ubuntu/lolkek/static/event/event.png"
	maxDiaryTitleLength        = 200
	maxDiaryComplaintsLength   = 1000
	maxDiaryAnamnesisLength	   = 1000
	maxDiaryObjectivelyLength  = 1000
	maxDiaryDiagnosisLength	   = 200

	maxRecordTitleLength       = 200
	maxRecordDescriptionLength = 3000

	maxReminderFrequency	   = 50
)

func (r *ReminderInfo) SetDefault() {
	r.Variant = false
	r.Frequency = 0
	r.StartDate = ""
}

func (d *DiaryBasicInfo) SetDefault() {
	d.Title = ""
	d.Complaints = ""
	d.Anamnesis = ""
	d.Objectively = ""
	d.Diagnosis = ""
	d.Reminder.SetDefault()
}


func (er *DiaryCreateRequest) SetDefault() {
	er.DiaryBasicInfo.SetDefault()
}

// func (er *RecordCreateRequest) SetDefault() {
// 	er.Description = ""
// 	er.Title = ""
// 	er.Characteristics = Characteristics{}
// }

func (er *RecordCreateResponse) SetDefault() {
	er.Id = 0
	er.DiaryId = 0
	er.CreatingDate = "2022-04-10 15:47:24"
	er.Description = ""
	er.Title = ""
	er.Area = 0
	er.Characteristics = Characteristics{}
	er.ImageList = []ImageInfo{}
}

func (d DiaryBasicInfo) IsValid() (isValid bool) {
	if len(d.Title) > maxDiaryTitleLength ||
	len(d.Complaints) > maxDiaryComplaintsLength ||
	len(d.Anamnesis) > maxDiaryAnamnesisLength ||
	len(d.Objectively) > maxDiaryObjectivelyLength ||
	len(d.Diagnosis) > maxDiaryDiagnosisLength ||
	!d.Reminder.IsValid() {
		return false
	}	
	return true
}

func (r ReminderInfo) IsValid() (isValid bool) {
	if (r.Frequency > maxReminderFrequency) {
		return false
	}
	return true
}

func (dr DiaryCreateRequest) IsValid() (isValid bool) {
	return dr.DiaryBasicInfo.IsValid()
}

// func (er RecordCreateRequest) IsValid() (isValid bool) {
// 	if len(er.Title) > maxRecordTitleLength || len(er.Description) > maxRecordDescriptionLength {
// 		return false
// 	}
// 	characteristicsList := [](*uint8){&er.Characteristics.Dryness,
// 		&er.Characteristics.Edema,
// 		&er.Characteristics.Itching,
// 		&er.Characteristics.Pain,
// 		&er.Characteristics.Peeling,
// 		&er.Characteristics.Redness,
// 	}
// 	for ch := range characteristicsList {
// 		if ch > 10 {
// 			return false
// 		}
// 	}
// 	return true
// }

func (er DiaryUpdateRequest) IsValid() (isValid bool) {
	return er.DiaryBasicInfo.IsValid()
}

type DiaryUpdateRequest struct {
	DiaryBasicInfo DiaryBasicInfo `json:"diarybasicinfo"`
}

type DiaryUpdateResponse struct {
	Id             uint64         `json:"id"`
	MedicId        uint64         `json:"medicid"`
	PatientId      uint64         `json:"patientid"`
	CreatingDate   string         `json:"creatingdate"`
	DiaryBasicInfo DiaryBasicInfo `json:"diarybasicinfo"`
}

type ReminderInfo struct {
	Variant		bool `json:"variant"`
	Frequency	uint64 `json:"frequency"`
	StartDate	string `json:"startdate"`
}

type DiaryBasicInfo struct {
	Title       string			`json:"title"`
	Complaints  string			`json:"complaints"`
	Anamnesis   string			`json:"anamnesis"`
	Objectively string			`json:"objectively"`
	Diagnosis   string			`json:"diagnosis"`
	Reminder	ReminderInfo	`json:"reminder"`
}

type DiaryCreateRequest struct {
	DiaryBasicInfo DiaryBasicInfo `json:"diarybasicinfo"`
}

type DiaryCreateResponse struct {
	Id             uint64         `json:"id"`
	MedicId        uint64         `json:"medicid"`
	PatientId      uint64         `json:"patientid"`
	CreatingDate   string         `json:"creatingdate"`
	DiaryBasicInfo DiaryBasicInfo `json:"diarybasicinfo"`
	LinkToken      string         `json:"linktoken"`
}

type DiaryLinkResponse struct {
	Id             uint64         `json:"id"`
	MedicId        uint64         `json:"medicid"`
	MedicName      string         `json:"medicname"`
	PatientId      uint64         `json:"patientid"`
	CreatingDate   string         `json:"creatingdate"`
	DiaryBasicInfo DiaryBasicInfo `json:"diarybasicinfo"`
}

type DiaryInList struct {
	Id           uint64 `json:"id"`
	Title        string `json:"title"`
	MedicId      uint64 `json:"medicid"`
	MedicName    string `json:"medicname"`
	PatientId    uint64 `json:"patientid"`
	PatientName  string `json:"patientname"`
	CreatingDate string `json:"creatingdate"`
	Objectively  string `json:"objectively"`
	LinkToken    string `json:"linktoken"`
	IsComplete   bool	`json:"iscomplete"`
}

type DiaryListResponse struct {
	DiaryList []DiaryInList `json:"diarylist"`
}

type RecordInDiaryBasicInfo struct {
	Id           uint64 `json:"id"`
	CreatingDate string `json:"creatingdate"`
	Title        string `json:"title"`
	Details      string `json:"details"`
}

type Records struct {
	MedicRecordList   []RecordInDiaryBasicInfo        `json:"medicrecordlist"`
	PatientRecordList []PatientRecordInDiaryBasicInfo `json:"patientrecordlist"`
}
type PatientRecordInDiaryBasicInfo struct {
	RecordInDiaryBasicInfo RecordInDiaryBasicInfo `json:"recordindiarybasicinfo"`
	Feelings               uint64                 `json:"feelings"`
}

type DiaryResponse struct {
	PatientName string            `json:"patientname"`
	Diary       DiaryLinkResponse `json:"diary"`
	Records     Records           `json:"records"`
}

// type RecordCreateRequest struct {
// 	Title           string          `json:"title"`
// 	Description     string          `json:"description"`
// 	Characteristics Characteristics `json:"characteristics"`
// }

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

// type RecordUpdateRequest struct {
// 	Id                  uint64              `json:"id"`
// 	RecordCreateRequest RecordCreateRequest `json:"RecordCreateRequest"`
// }

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
	CreateDiary(diary DiaryCreateRequest, medicId uint64) (DiaryCreateResponse, error)
	LinkDiary(patientId uint64, diaryId uint64) (DiaryLinkResponse, error)
	DeleteDiary(diaryid uint64) error
	CompleteDiary(diaryid uint64) error
	GetDiary(userId uint64) (DiaryListResponse, error)
	GetUserRole(userId uint64) (bool, bool, error)
	GetCertainDiary(diaryId uint64) (DiaryResponse, error)
	UpdateDiary(diary DiaryUpdateRequest, diaryId uint64) (DiaryUpdateResponse, error)
	CreateLinkToken(diaryId uint64, linkToken string) error
	CheckAndDeleteToken(diaryId uint64, linkToken string) (bool, error)
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
	CreateDiary(diary DiaryCreateRequest, medicId uint64) (DiaryCreateResponse, error)
	LinkDiary(patientId uint64, diaryId uint64, linkToken string) (DiaryLinkResponse, error)
	DeleteDiary(diaryid uint64) error
	CompleteDiary(medicid, diaryid uint64) error
	GetDiary(userId uint64) (DiaryListResponse, error)
	GetCertainDiary(diaryId uint64, userId uint64) (DiaryResponse, error)
	UpdateDiary(diary DiaryUpdateRequest, diaryId uint64) (DiaryUpdateResponse, error)

	// GetCategory() (CategoryListResponse, error)
	// EventSignUp(eventId uint64, userId uint64)(error)
	// CancelEventSignUp(eventId uint64, userId uint64) (error)
	// GetRecomendedEvent(userId uint64) (EventListResponse, error)
}

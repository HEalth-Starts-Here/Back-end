package domain

// TODO write valid path
const (
	// BaseEventPicture = "/home/ubuntu/lolkek/static/event/event.png"
	maxEventTitleLength = 200
	maxRecordDescriptionLength = 3000
)

type DiaryCreatingRequest struct {
	Category               uint32   `json:"category"`
	MedicId                uint32   `json:"medicid"`
	PatientId              uint32   `json:"patientid"`
	Title                  string   `json:"title"`
	Description            string   `json:"description"`
}

func (er *DiaryCreatingRequest) SetDefault() () {
	er.Category = 0
	er.MedicId = 0
	er.PatientId = 0
	er.Title = ""
	er.Description = ""
	return
}

func (er *RecordCreatingRequest) SetDefault() () {
	er.DiaryId = 0
	er.Description = ""
	er.PosterPath = ""
	return
}

func (er DiaryCreatingRequest) IsValid() (isValid bool) {
	if len(er.Title) > maxEventTitleLength {
		return false
	}
	return true
}

func (er RecordCreatingRequest) IsValid() (isValid bool) {
	if len(er.Description) > maxRecordDescriptionLength {
		return false
	}
	return true
}

type DiaryCreatingResponse struct {
	Id                     uint64   `json:"id"`
	Category               uint32   `json:"category"`
	MedicId                uint32   `json:"medicid"`
	PatientId              uint32   `json:"patientid"`
	CreatingDate           string   `json:"creatingdate"`
	Title                  string   `json:"title"`
	Description            string   `json:"description"`
}

type RecordsCreatingResponse struct {
	Id                     uint64   `json:"id"`
	DiaryId                uint64   `json:"medicid"`
	Description            string   `json:"description"`
	PosterPath             string   `json:"posterpath"`
}

type DiaryListResponse struct {
	DiaryList []DiaryCreatingResponse `json:"diarylist"`
}

type DiaryResponse struct {
	Diary DiaryCreatingResponse `json:"diary"`
	RecordsList []RecordsCreatingResponse `json:"records"`
}

type RecordCreatingRequest struct {
	DiaryId                uint64   `json:"diaryid"`
	Description            string   `json:"description"`
	PosterPath             string   `json:"posterpath"`
}

type RecordCreatingResponse struct {
	Id                     uint64   `json:"id"`
	DiaryId                uint64   `json:"diaryid"`
	Description            string   `json:"description"`
	PosterPath             string   `json:"posterpath"`
}

// type CategoryResponse struct {
// 	Name             string   `json:"name"`
// 	ImagePath             string   `json:"ImagePath"`
// }

// type CategoryListResponse struct {
// 	CategoryList []CategoryResponse `json:"categorylist"`
// }

type DiaryRepository interface {
	CreateDiary(diary DiaryCreatingRequest) (DiaryCreatingResponse, error)
	GetDiary() (DiaryListResponse, error)  
	GetCertainDiary(diaryId uint64) (DiaryResponse, error)
	CreateRecord(record RecordCreatingRequest) (RecordCreatingResponse, error)


	// DiaryAlreadyExist(diary DiaryCreatingRequest) (bool, error)
	// GetCategory() (CategoryListResponse, error)
	// CreateEventCategory(eventId uint64, categories []string) ([]string, error)
	// SignUpUserForEvent(eventId uint64, userId uint64) (error)
	// CancelEventSignUp(eventId uint64, userId uint64) (error) 
	// GetUserCategory(id uint64) ([]string, error)
	// GetUserAge(id uint64) (uint64, error)
	// GetEventAges(id uint64) (uint16, uint16, error)
}

type DiaryUsecase interface {
	CreateDiary(diary DiaryCreatingRequest) (DiaryCreatingResponse, error)
	GetDiary() (DiaryListResponse, error) 
	GetCertainDiary(diaryId uint64) (DiaryResponse, error)
	CreateRecord(record RecordCreatingRequest) (RecordCreatingResponse, error)


	// GetCategory() (CategoryListResponse, error)
	// EventSignUp(eventId uint64, userId uint64)(error)
	// CancelEventSignUp(eventId uint64, userId uint64) (error) 
	// GetRecomendedEvent(userId uint64) (EventListResponse, error)
}

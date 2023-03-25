package domain

// TODO write valid path
const (
	// BaseEventPicture = "/home/ubuntu/lolkek/static/event/event.png"
	maxDiaryTitleLength = 200
	maxRecordTitleLength = 200
	maxDiaryDescriptionLength = 3000
	maxRecordDescriptionLength = 3000
)

func (er *DiaryCreatingRequest) SetDefault() () {
	er.MedicId = 0
	er.PatientId = 0
	er.Title = ""
	er.Description = ""
	return
}

func (er *RecordCreatingRequest) SetDefault() () {
	er.Description = ""
	er.Title = ""
	er.Characteristics = Characteristics{}
	return
}

func (er *RecordCreatingResponse) SetDefault() () {
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

func (er DiaryCreatingRequest) IsValid() (isValid bool) {
	if len(er.Title) > maxDiaryTitleLength || len(er.Description) > maxDiaryDescriptionLength{
		return false
	}
	return true
}

func (er RecordCreatingRequest) IsValid() (isValid bool) {
	if len(er.Title) > maxRecordTitleLength || len(er.Description) > maxRecordDescriptionLength{
		return false
	}
	characteristicsList := [](*uint8){  &er.Characteristics.Dryness, 
										&er.Characteristics.Edema, 
										&er.Characteristics.Itching, 
										&er.Characteristics.Pain, 
										&er.Characteristics.Peeling, 
										&er.Characteristics.Redness,
									}
	for ch := range characteristicsList{
		if ch > 10{
			return false
		}
	}
	return true
}

func (er DiaryUpdatingRequest) IsValid() (isValid bool) {
	if len(er.Title) > maxDiaryTitleLength || len(er.Description) > maxDiaryDescriptionLength{
		return false
	}
	return true
}

type DiaryUpdatingRequest struct {
	Id                	   uint64  	`json:"id"`
	Title                  string   `json:"title"`
	Description            string   `json:"description"`
}

type DiaryUpdatingResponse struct {
	Id                     uint64   `json:"id"`
	Title                  string   `json:"title"`
	Description            string   `json:"description"`
}

type DiaryCreatingRequest struct {
	MedicId                uint32   `json:"medicid"`
	PatientId              uint32   `json:"patientid"`
	Title                  string   `json:"title"`
	Description            string   `json:"description"`
}

type DiaryCreatingResponse struct {
	Id                     uint64   `json:"id"`
	MedicId                uint32   `json:"medicid"`
	PatientId              uint32   `json:"patientid"`
	CreatingDate           string   `json:"creatingdate"`
	Title                  string   `json:"title"`
	Description            string   `json:"description"`
}

// type RecordsCreatingResponse struct {
// 	Id                     uint64   `json:"id"`
// 	DiaryId                uint64   `json:"diaryid"`
// 	Description            string   `json:"description"`
// 	PosterPath             string   `json:"posterpath"`
// }

type DiaryListResponse struct {
	DiaryList []DiaryCreatingResponse `json:"diarylist"`
}

type DiaryResponse struct {
	Diary DiaryCreatingResponse `json:"diary"`
	RecordsList []RecordCreatingResponse `json:"records"`
}

type RecordCreatingRequest struct {
	Title           	   string     		 `json:"title"` 
	Description            string			 `json:"description"`
	Characteristics		   Characteristics	 `json:"characteristics"`
}

type ImageInfo struct {
	Id                     uint64   `json:"id"`
	RecordId               uint64   `json:"recordid"`
	Name	               string   `json:"name"`
	Area	               float64  `json:"area"`
}

type RecordCreatingResponse struct {
	Id                     uint64 		     `json:"id"`
	DiaryId                uint64  			 `json:"diaryid"`
	CreatingDate           string   		 `json:"creatingdate"`
	Description            string  			 `json:"description"`
	Title                  string  			 `json:"title"`
	Area			   	   float64			 `json:"area"`
	Characteristics		   Characteristics	 `json:"characteristics"`
	ImageList		   	   []ImageInfo		 `json:"imagelist"`

}

type RecordUpdatingRequest struct {
	Id                	   uint64  	`json:"id"`
	Title                  string   `json:"title"`
	Description            string   `json:"description"`
}

type Characteristics struct {
	Dryness 			   uint8 `json:"dryness"`
	Edema 				   uint8 `json:"edema"`
	Itching 			   uint8 `json:"itching"` // defin all not required
	Pain 				   uint8 `json:"pain"`
	Peeling 			   uint8 `json:"peeling"`
	Redness 			   uint8 `json:"redness"`
}

type ImageInfoUsecase struct {
	Name	               string   `json:"name"`
	Area	               float64  `json:"area"`
}

// Area 				   float32 `json:"area"` // cm^2
// // TODO Define palm area as 1%	

type DiaryRepository interface {
	CreateDiary(diary DiaryCreatingRequest) (DiaryCreatingResponse, error)
	GetDiary() (DiaryListResponse, error)  
	GetCertainDiary(diaryId uint64) (DiaryResponse, error)
	CreateRecord(diaryId uint64, record RecordCreatingRequest, imageInfo []ImageInfoUsecase, Area float64) (RecordCreatingResponse, error)
	UpdateDiary(diary DiaryUpdatingRequest) (DiaryUpdatingResponse, error)

	

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
	CreateRecord(diaryId uint64, record RecordCreatingRequest, imageInfo []ImageInfoUsecase) (RecordCreatingResponse, error)
	UpdateDiary(diary DiaryUpdatingRequest) (DiaryUpdatingResponse, error)


	// GetCategory() (CategoryListResponse, error)
	// EventSignUp(eventId uint64, userId uint64)(error)
	// CancelEventSignUp(eventId uint64, userId uint64) (error) 
	// GetRecomendedEvent(userId uint64) (EventListResponse, error)
}

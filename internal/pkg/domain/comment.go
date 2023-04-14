package domain

const (
	maxCommentTextLength        = 1000
)

func (cr *BasicCommentInfo) SetDefault() {
	cr.Text = ""
	return
}

func (cr BasicCommentInfo) IsValid() (isValid bool) {
	if len(cr.Text) > maxCommentTextLength {
		return false
	}
	return true
}

type BasicCommentInfo struct {
	Text string `json:"text"`
}

type CommentCreateResponse struct {
	Id             		uint64         		`json:"id"`
	BasicCommentInfo    BasicCommentInfo    `json:"basiccommentinfo"`
	AuthorIsMedic  		bool           		`json:"authorismedic"`
	IsReaded	   		bool           		`json:"isreaded"`
	CreatingDate   		string		   		`json:"creatingdate"`
}

// type DiaryBasicInfo struct {
// 	Title       string `json:"title"`
// 	Complaints  string `json:"complaints"`
// 	Anamnesis   string `json:"anamnesis"`
// 	Objectively string `json:"objectively"`
// 	Diagnosis   string `json:"diagnosis"`
// }

// type DiaryCreateRequest struct {
// 	// MedicId		   uint64 `json:"medicid"`
// 	// PatientId	   uint64 `json:"patientid"`
// 	DiaryBasicInfo DiaryBasicInfo `json:"diarybasicinfo"`
// }

// type DiaryCreateResponse struct {
// 	Id             uint64         `json:"id"`
// 	MedicId        uint64         `json:"medicid"`
// 	PatientId      uint64         `json:"patientid"`
// 	CreatingDate   string         `json:"creatingdate"`
// 	DiaryBasicInfo DiaryBasicInfo `json:"diarybasicinfo"`
// }

// type DiaryLinkResponse struct {
// 	Id             uint64         `json:"id"`
// 	MedicId        uint64         `json:"medicid"`
// 	MedicName      string         `json:"medicname"`
// 	PatientId      uint64         `json:"patientid"`
// 	CreatingDate   string         `json:"creatingdate"`
// 	DiaryBasicInfo DiaryBasicInfo `json:"diarybasicinfo"`
// }

// type DiaryInList struct {
// 	Id           uint64 `json:"id"`
// 	Title        string `json:"title"`
// 	MedicId      uint64 `json:"medicid"`
// 	MedicName    string `json:"medicname"`
// 	PatientId    uint64 `json:"patientid"`
// 	PatientName  string `json:"patientname"`
// 	CreatingDate string `json:"creatingdate"`
// 	Objectively  string `json:"objectively"`
// }

// type DiaryListResponse struct {
// 	DiaryList []DiaryInList `json:"diarylist"`
// }

// type RecordBasicInfo struct {
// 	Id 			 uint64 `json:"id"`
// 	CreatingDate string `json:"creatingdate"`
// 	Title        string `json:"title"`
// 	Details      string `json:"details"`
// }

// type Records struct {
// 	MedicRecordList   []RecordBasicInfo `json:"medicrecordlist"`
// 	PatientRecordList []RecordBasicInfo `json:"patientrecordlist"`
// }

// type DiaryResponse struct {
// 	PatientName string            `json:"patientname"`
// 	Diary       DiaryLinkResponse `json:"diary"`
// 	Records     Records           `json:"records"`
// }

// type RecordCreateRequest struct {
// 	Title           string          `json:"title"`
// 	Description     string          `json:"description"`
// 	Characteristics Characteristics `json:"characteristics"`
// }

// type ImageInfo struct {
// 	Id       uint64  `json:"id"`
// 	RecordId uint64  `json:"recordid"`
// 	Name     string  `json:"name"`
// 	Area     float64 `json:"area"`
// }

// type RecordCreateResponse struct {
// 	Id              uint64          `json:"id"`
// 	DiaryId         uint64          `json:"diaryid"`
// 	CreatingDate    string          `json:"creatingdate"`
// 	Description     string          `json:"description"`
// 	Title           string          `json:"title"`
// 	Area            float64         `json:"area"`
// 	Characteristics Characteristics `json:"characteristics"`
// 	ImageList       []ImageInfo     `json:"imagelist"`
// }

// type RecordUpdateRequest struct {
// 	Id                  uint64              `json:"id"`
// 	RecordCreateRequest RecordCreateRequest `json:"RecordCreateRequest"`
// }

// type Characteristics struct {
// 	Dryness uint8 `json:"dryness"`
// 	Edema   uint8 `json:"edema"`
// 	Itching uint8 `json:"itching"` // defin all not required
// 	Pain    uint8 `json:"pain"`
// 	Peeling uint8 `json:"peeling"`
// 	Redness uint8 `json:"redness"`
// }

// type ImageInfoUsecase struct {
// 	Name string  `json:"name"`
// 	Area float64 `json:"area"`
// }

// // Area 				   float32 `json:"area"` // cm^2
// // // TODO Define palm area as 1%

type CommentRepository interface {
	CreateComment(diaryId uint64, authorIsMedic bool, commentInfo BasicCommentInfo) (CommentCreateResponse, error)
	CheckUserRole(userId uint64) (bool, bool, error)
}

type CommentUsecase interface {
	CreateComment(diaryId uint64, userId uint64, commentInfo BasicCommentInfo) (CommentCreateResponse, error)
	CheckUserRole(userId uint64) (bool, bool, error)
}

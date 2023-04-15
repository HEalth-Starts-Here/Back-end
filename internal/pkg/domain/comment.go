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
	DiaryId             uint64         		`json:"diaryid"`
	BasicCommentInfo    BasicCommentInfo    `json:"basiccommentinfo"`
	AuthorIsMedic  		bool           		`json:"authorismedic"`
	IsReaded	   		bool           		`json:"isreaded"`
	CreatingDate   		string		   		`json:"creatingdate"`
}

type CommentRepository interface {
	CreateComment(diaryId uint64, authorIsMedic bool, commentInfo BasicCommentInfo) (CommentCreateResponse, error)
	CheckUserRole(userId uint64) (bool, bool, error)
}

type CommentUsecase interface {
	CreateComment(diaryId uint64, userId uint64, commentInfo BasicCommentInfo) (CommentCreateResponse, error)
	CheckUserRole(userId uint64) (bool, bool, error)
}

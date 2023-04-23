package domain

const (
	maxCommentTextLength = 1000
)

func (cr *BasicCommentInfo) SetDefault() {
	cr.Text = ""
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
	CommentInListInfo CommentInListInfo `json:"commentinlistinfo"`
	DiaryId           uint64            `json:"diaryid"`
}

type CommentInListInfo struct {
	Id               uint64           `json:"id"`
	AuthorIsMedic    bool             `json:"authorismedic"`
	IsReaded         bool             `json:"isreaded"`
	CreatingDate     string           `json:"creatingdate"`
	BasicCommentInfo BasicCommentInfo `json:"basiccommentinfo"`
}

type GetCommentResponse struct {
	CommentList []CommentInListInfo `json:"commentlist"`
	DiaryId     uint64              `json:"diaryid"`
}

type CommentRepository interface {
	CreateComment(diaryId uint64, authorIsMedic bool, commentInfo BasicCommentInfo) (CommentCreateResponse, error)
	CheckUserRole(userId uint64) (bool, bool, error)
	GetComment(diaryId uint64) (GetCommentResponse, error)
	DeleteComment(commentId uint64) error
}

type CommentUsecase interface {
	CreateComment(diaryId uint64, userId uint64, commentInfo BasicCommentInfo) (CommentCreateResponse, error)
	CheckUserRole(userId uint64) (bool, bool, error)
	GetComment(userId, diaryId uint64) (GetCommentResponse, error)
	DeleteComment(userId, commentId uint64) error
}

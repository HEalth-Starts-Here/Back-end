package domain

const (
	maxNoteTextLength = 1000
)

func (basicNoteInfo *BasicNoteInfo) SetDefault() {
	basicNoteInfo.Text = ""
	return
}

func (basicNoteInfo BasicNoteInfo) IsValid() (isValid bool) {
	if len(basicNoteInfo.Text) > maxNoteTextLength {
		return false
	}
	return true
}

type BasicNoteInfo struct {
	Text string `json:"text"`
}

type NoteCreateResponse struct {
	IsMedicRecord  bool           `json:"ismedicrecord"`
	RecordId       uint64         `json:"recordid"`
	NoteInListInfo NoteInListInfo `json:"noteinlistinfo"`
}

type NoteInListInfo struct {
	Id            uint64        `json:"id"`
	CreatingDate  string        `json:"creatingdate"`
	BasicNoteInfo BasicNoteInfo `json:"basicnoteinfo"`
}

type GetNoteResponse struct {
	IsMedicRecord bool             `json:"ismedicrecord"`
	RecordId      uint64           `json:"recordid"`
	NoteList      []NoteInListInfo `json:"notelist"`
}

type NoteRepository interface {
	CreateNote(isMedicRecord bool, recordId uint64, noteCreateRequest BasicNoteInfo) (NoteCreateResponse, error)
	// CheckUserRole(userId uint64) (bool, bool, error)
	GetNote(isMedicRecord bool, recordId uint64) (GetNoteResponse, error)
	// DeleteComment(commentId uint64) (error)
	DeleteNote(noteId uint64) error
}

type NoteUsecase interface {
	// CreateComment(diaryId uint64, userId uint64, commentInfo BasicCommentInfo) (CommentCreateResponse, error)
	// CheckUserRole(userId uint64) (bool, bool, error)
	CreateNote(medicId uint64, isMedicRecord bool, recordId uint64, noteCreateRequest *BasicNoteInfo) (NoteCreateResponse, error)
	GetNote(medicId uint64, isMedicRecord bool, recordId uint64) (GetNoteResponse, error)
	// DeleteComment(userId, commentId uint64) (error)
	DeleteNote(medicId uint64, isMedicRecord bool, noteId uint64) error
}

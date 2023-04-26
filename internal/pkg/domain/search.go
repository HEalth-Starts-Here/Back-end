package domain

const (
	maxSearchTextLength = 1000
)

func (searchDiaryRequest *SearchDiaryRequest) SetDefault() {
	searchDiaryRequest.Text = ""
	return
}

func (searchDiaryRequest *SearchDiaryRequest) IsValid() (isValid bool) {
	if len(searchDiaryRequest.Text) > maxSearchTextLength {
		return false
	}
	return true
}

type SearchDiaryRequest struct {
	Text string `json:"text"`
}

type SearchRepository interface {
	SearchDiary(userId uint64, isMedic bool, searchParams SearchDiaryRequest) (DiaryListResponse, error)
	CheckUserRole(userId uint64) (bool, bool, error)
}

type SearchUsecase interface {
	SearchDiary(userId uint64, searchParams SearchDiaryRequest) (DiaryListResponse, error)	
	CheckUserRole(userId uint64) (bool, bool, error)
}

package sanitizer

import (
	"hesh/internal/pkg/domain"

	"github.com/microcosm-cc/bluemonday"
)

func sanitizeText (text *string) {
	sanitizer := bluemonday.UGCPolicy()
	*text = sanitizer.Sanitize(*text)
}

func SanitizeDiaryCreating(diary *domain.DiaryCreateRequest) {
	sanitizeText(&diary.DiaryBasicInfo.Title)
	sanitizeText(&diary.DiaryBasicInfo.Complaints)
	sanitizeText(&diary.DiaryBasicInfo.Anamnesis)
	sanitizeText(&diary.DiaryBasicInfo.Objectively)
}

func SanitizeDiaryUpdating(diary *domain.DiaryUpdateRequest) {
	sanitizeText(&diary.DiaryBasicInfo.Title)
	sanitizeText(&diary.DiaryBasicInfo.Complaints)
	sanitizeText(&diary.DiaryBasicInfo.Anamnesis)
	sanitizeText(&diary.DiaryBasicInfo.Objectively)
	sanitizeText(&diary.DiaryBasicInfo.Diagnosis)
}

func SanitizeRecordCreating(record *domain.RecordCreateRequest) {
	sanitizeText(&record.Title)
	sanitizeText(&record.Description)
}

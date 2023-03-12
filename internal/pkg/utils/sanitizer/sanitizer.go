package sanitizer

import (
	"hesh/internal/pkg/domain"

	"github.com/microcosm-cc/bluemonday"
)

func SanitizeDiaryCreating(diary *domain.DiaryCreatingRequest) {
	sanitizer := bluemonday.UGCPolicy()
	diary.Title = sanitizer.Sanitize(diary.Title)
	diary.Description = sanitizer.Sanitize(diary.Description)
}

func SanitizeRecordCreating(record *domain.RecordCreatingRequest) {
	sanitizer := bluemonday.UGCPolicy()
	record.Description = sanitizer.Sanitize(record.Description)
	record.PosterPath = sanitizer.Sanitize(record.PosterPath)
}

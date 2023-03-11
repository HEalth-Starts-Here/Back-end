package sanitizer

import (
	"hesh/internal/pkg/domain"

	"github.com/microcosm-cc/bluemonday"
)

func SanitizeDiaryCreating(event *domain.DiaryCreatingRequest) {
	sanitizer := bluemonday.UGCPolicy()
	event.Title = sanitizer.Sanitize(event.Title)
	event.Description = sanitizer.Sanitize(event.Description)
}

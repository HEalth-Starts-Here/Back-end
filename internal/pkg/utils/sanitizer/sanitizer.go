package sanitizer

import (
	"hesh/internal/pkg/domain"

	"github.com/microcosm-cc/bluemonday"
)

func sanitizeText (text *string) {
	sanitizer := bluemonday.UGCPolicy()
	*text = sanitizer.Sanitize(*text)
}

// func SanitizeUserInit(userInitInfo *domain.UserInitRequest) {
// 	sanitizeText(&userInitInfo.InitBasicInfo.Name)
// }

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

func SanitizeMedicRecordBasicInfo (record *domain.MedicRecordBasicInfo){
	sanitizeText(&record.Title)
	sanitizeText(&record.Treatment)
	sanitizeText(&record.Recommendations)
	sanitizeText(&record.Details)
}

func SanitizeImageInfo (imageInfo *domain.RecordImageInfo){
	sanitizeText(&imageInfo.ImageName)
	for i := 0; i< (len(imageInfo.Tags)); i++ {
		sanitizeText(&imageInfo.Tags[i])
	}
}

func SanitizeMedicRecordCreateRequest(record *domain.MedicRecordCreateRequest) {
	SanitizeMedicRecordBasicInfo(&record.BasicInfo)
	for i := range (record.Images){
		SanitizeImageInfo(&record.Images[i])
	}
}

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

func SanitizeMedicRecordImages (request (*domain.MedicRecordUpdateImageRequest)){
	for i := range request.Images {
		sanitizeText(&((request.Images[i]).ImageName))
		for j := range request.Images[i].Tags {
			sanitizeText(&((request.Images[i]).Tags[j]))
		}
	}
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
	// for i := range (record.Auido){
	// 	sanitizeText(&record.Auido[i])
	// }
}

func SanitizePatientRecordBasicInfo (record *domain.PatientRecordBasicInfo){
	sanitizeText(&record.Title)
	sanitizeText(&record.Treatment)
	sanitizeText(&record.Complaints)
	sanitizeText(&record.Details)
}

func SanitizePatientRecordCreateRequest(record *domain.PatientRecordCreateRequest) {
	SanitizePatientRecordBasicInfo(&record.BasicInfo)
	for i := range (record.Images){
		SanitizeImageInfo(&record.Images[i])
	}
}



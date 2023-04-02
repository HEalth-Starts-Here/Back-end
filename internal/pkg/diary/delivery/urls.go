package diarydelivery

const (
	diaryUrl      		  = "/diary"
	GetDiaryUrl   		  = diaryUrl + "/get"
	CreateDiaryUrl		  = diaryUrl + "/create"
	LinkDiaryUrl		  = diaryUrl + "/link/{id:[0-9]+}"
	DeleteDiaryUrl		  = diaryUrl + "/delete/{id:[0-9]+}"
	GetCertainDiaryUrl    = GetDiaryUrl + "/{id:[0-9]+}"
	CreateRecordUrl       = CreateDiaryUrl + "/{id:[0-9]+}"
	PutCertainDiaryUrl	  = diaryUrl + "/update"
	PutRecordUrl		  = diaryUrl + "/recordUpdate"
	DetermineAreaUrl      = "/determinearea"
)

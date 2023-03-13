package diarydelivery

const (
	diaryUrl      		  = "/diary"
	GetDiaryUrl   		  = diaryUrl + "/get"
	CreateDiaryUrl		  = diaryUrl + "/create"
	GetCertainDiaryUrl    = GetDiaryUrl + "/{id:[0-9]+}"
	CreateRecordUrl       = CreateDiaryUrl + "/{id:[0-9]+}"
	// GetRecomendedEvent    = GetEventUrl + "/recomended"
	// GetCatagoryUrl        = eventUrl + "/category"
	// EventSignUpUrl		  = eventUrl + "/signup/{id:[0-9]+}"
	// CancelEventSignUpUrl  = eventUrl + "/signup/cancel/{id:[0-9]+}"
	// DeleteEventUrl = eventUrl + "/delete"
	// AlterEventUrl  = eventUrl + "/alter"
)

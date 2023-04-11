package recorddelivery

const (
	recordUrl						= "/record"
	recordMedicUrl					= recordUrl + "/medic"
	recordMedicCreateUrl			= recordMedicUrl + "/create/{id:[0-9]+}"
	// recordMedicAddDiarisationUr		= recordMedicUrl + "/diarisation/{id:[0-9]+}"
	recordMedicGetUrl				= recordMedicUrl + "/get/{id:[0-9]+}"
	recordMedicGetDiarisationsUrl	= recordMedicUrl + "/get/diarisations/{id:[0-9]+}"
	recordMedicUpdateUrl			= recordMedicUrl + "/update"
	recordMedicUpdateTextUrl		= recordMedicUpdateUrl +"/text/{id:[0-9]+}" 
	recordMedicUpdateImageUrl		= recordMedicUpdateUrl +"/image/{id:[0-9]+}" 
	recordMedicDeleteImageUrl		= recordMedicUrl + "/delete/{id:[0-9]+}"
)

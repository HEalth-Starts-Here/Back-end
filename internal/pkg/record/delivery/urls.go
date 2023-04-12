package recorddelivery

const (
	recordUrl						= "/record"
	
	// MEDIC
	recordMedicUrl					= recordUrl + "/medic"
	recordMedicCreateUrl			= recordMedicUrl + "/create/{id:[0-9]+}"
	// recordMedicAddDiarisationUr		= recordMedicUrl + "/diarisation/{id:[0-9]+}"
	recordMedicGetUrl				= recordMedicUrl + "/get/{id:[0-9]+}"
	recordMedicGetDiarisationsUrl	= recordMedicUrl + "/get/diarisations/{id:[0-9]+}"
	recordMedicUpdateUrl			= recordMedicUrl + "/update"
	recordMedicUpdateTextUrl		= recordMedicUpdateUrl +"/text/{id:[0-9]+}" 
	recordMedicUpdateImageUrl		= recordMedicUpdateUrl +"/image/{id:[0-9]+}" 
	recordMedicDeleteImageUrl		= recordMedicUrl + "/delete/{id:[0-9]+}"

	// PATIENT
	recordPatientUrl				= recordUrl + "/patient"
	recordPatientCreateUrl			= recordPatientUrl + "/create/{id:[0-9]+}"
	recordPatientGetUrl				= recordPatientUrl + "/get/{id:[0-9]+}"
	recordPatientUpdateUrl			= recordPatientUrl + "/update"
	recordPatientUpdateTextUrl		= recordPatientUpdateUrl + "/text/{id:[0-9]+}"
	recordPatientUpdateImageUrl		= recordPatientUpdateUrl + "/image/{id:[0-9]+}"
	recordPatientDeleteImageUrl		= recordPatientUrl + "/delete/{id:[0-9]+}"
)

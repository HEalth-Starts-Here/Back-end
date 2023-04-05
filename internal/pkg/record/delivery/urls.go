package recorddelivery

const (
	recordUrl      		  = "/record"
	recordMedicUrl	  	  = recordUrl + "/medic"
	recordMedicCreateUrl  = recordMedicUrl + "/create/{id:[0-9]+}"
	recordMedicGet		  = recordMedicUrl + "/get/{id:[0-9]+}"
)

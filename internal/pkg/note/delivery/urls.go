package notedelivery

const (
	isMedicRecordUrl		= "/{isMedicRecord:[a-z]+}"
	idUrl					= "/{id:[0-9]+}"
	noteUrl      			= "/note"
	GetNoteUrl  	 		= noteUrl + "/get" + isMedicRecordUrl + idUrl
	// CreateNoteUrl   		= commentUrl + "/create" + idUrl
	// DeleteCommentUrl   		= commentUrl + "/delete" + idUrl
)

package commentdelivery

const (
	idUrl					= "/{id:[0-9]+}"
	commentUrl      		= "/comment"
	GetCommentUrl   		= commentUrl + "/get" + idUrl
	CreateCommentUrl   		= commentUrl + "/create" + idUrl
)

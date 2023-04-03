package domain

const (
)

// func (ur *UserInitRequest) SetDefault() {
// 	ur.InitBasicInfo.name = ""
// 	return
// }

// TODO: add returning errors
func (ur UserInitRequest) IsValid() (bool) {
	if len(ur.InitBasicInfo.Name) == 0 {
		return false
	}
	return true
}

type InitBasicInfo struct {
	Name string		`json:"name"`
	IsMedic bool	`json:"ismedic"`
}

type UserInitRequest struct {
	InitBasicInfo InitBasicInfo `json:"initbasicinfo"`
}

type UserInitResponse struct {
	Id            uint64        `json:"id"`
	InitBasicInfo InitBasicInfo `json:"initbasicinfo"`
}

type UserRepository interface {
	UserInit(initBasicInfo UserInitRequest, userId uint64) (UserInitResponse, error)
}

type UserUsecase interface {
	UserInit(initBasicInfo UserInitRequest, userId uint64) (UserInitResponse, error)
}

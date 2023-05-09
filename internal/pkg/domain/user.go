package domain

const (
	maxNameLength = 200
)

type UserInfo struct {
	Id      uint64 `json:"id"`
	Name    string `json:"name"`
	IsMedic bool   `json:"ismedic"`
}

func (registerMedicRequest *RegisterMedicRequest) SetDefault() {
	registerMedicRequest.Name = ""
	return
}

func (registerMedicRequest *RegisterMedicRequest) IsValid() (isValid bool) {
	if len(registerMedicRequest.Name) > maxNameLength {
		return false
	}
	return true
}

func (registerPatientRequest *RegisterPatientRequest) SetDefault() {
	registerPatientRequest.Name = ""
	return
}

func (registerPatientRequest *RegisterPatientRequest) IsValid() (isValid bool) {
	if len(registerPatientRequest.Name) > maxNameLength {
		return false
	}
	return true
}


type RegisterMedicRequest struct {
	Name string `json:"name"`
}

type RegisterPatientRequest struct {
	Name    string `json:"name"`
	DiaryId uint64 `json:"diaryid"`
}

type RegisterPatientResponse struct {
	UserInfo UserInfo `json:"userinfo"`
	DiaryId  uint64   `json:"diaryid"`
}

type UserRepository interface {
	UserInit(userId uint64) (bool, UserInfo, error)
	RegisterMedic(registerMedicRequest RegisterMedicRequest, userId uint64) (UserInfo, error)
	RegisterPatient(registerPatientRequest RegisterPatientRequest, userId uint64) (UserInfo, error)
	LinkPatientToDiary(patientId, diaryId uint64) (uint64, uint64, error)
	CheckAndDeleteToken(diaryId uint64, linkToken string) (bool, error)
}

type UserUsecase interface {
	UserInit(userId uint64) (bool, UserInfo, error)
	RegisterMedic(registerMedicRequest RegisterMedicRequest, userId uint64) (UserInfo, error)
	RegisterPatient(registerPatientRequest RegisterPatientRequest, userId uint64, linkToken string) (RegisterPatientResponse, error)
}

package userrepository

import (
	"hesh/internal/pkg/database"
	"hesh/internal/pkg/domain"

	"hesh/internal/pkg/utils/cast"
	"hesh/internal/pkg/utils/log"
)

type dbuserrepository struct {
	dbm *database.DBManager
}

func InitUserRep(manager *database.DBManager) domain.UserRepository {
	return &dbuserrepository{
		dbm: manager,
	}
}


func (er *dbuserrepository) UserInit(userId uint64) (bool, domain.UserInfo, error) {
	query := queryGetUserInfo
	// if userInitInfo.InitBasicInfo.IsMedic{
	// 	query = queryMedicInit
	// } else {
	// 	query = queryPatientInit
	// }
	resp, err := er.dbm.Query(query,
		userId)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return false, domain.UserInfo{}, err
	}
	if len(resp) == 0 {
		return false, domain.UserInfo{}, nil
	}
	return true, 
	domain.UserInfo{
		Id:           		cast.ToUint64(resp[0][0]),
		Name:           	cast.ToString(resp[0][1]),
		IsMedic:            cast.ToBool(resp[0][2]),},
		nil
}

func (er *dbuserrepository) RegisterMedic (userInfoRequest domain.RegisterMedicRequest, medicId uint64) (domain.UserInfo, error) {
	query := queryRegisterMedic
	// if userInitInfo.InitBasicInfo.IsMedic{
	// 	query = queryMedicInit
	// } else {
	// 	query = queryPatientInit
	// }
	resp, err := er.dbm.Query(query,
		medicId, userInfoRequest.Name)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return domain.UserInfo{}, err
	}
	return domain.UserInfo{
		Id:           		cast.ToUint64(resp[0][0]),
		Name:           	cast.ToString(resp[0][1]),
		IsMedic:            true}, 
		nil
}


func (er *dbuserrepository) RegisterPatient (patientInfoRequest domain.RegisterPatientRequest, patientId uint64) (domain.UserInfo, error) {
	query := queryRegisterPatient
	// if userInitInfo.InitBasicInfo.IsMedic{
	// 	query = queryMedicInit
	// } else {
	// 	query = queryPatientInit
	// }
	resp, err := er.dbm.Query(query,patientId, patientInfoRequest.Name)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return domain.UserInfo{}, err
	}
	return domain.UserInfo{
		Id:           		cast.ToUint64(resp[0][0]),
		Name:           	cast.ToString(resp[0][1]),
		IsMedic:            false}, 
		nil
}

func (er *dbuserrepository) CheckAndDeleteToken (diaryId uint64, linkToken string) (bool, error) {
	query := queryDeleteLinkToken
	// if userInitInfo.InitBasicInfo.IsMedic{
	// 	query = queryMedicInit
	// } else {
	// 	query = queryPatientInit
	// }
	resp, err := er.dbm.Query(query, diaryId, linkToken)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return false, err
	}
	if len(resp) == 0 {
		return false, nil
	}
	return true, nil
}

func (er *dbuserrepository) LinkPatientToDiary (patientId, diaryId uint64) (uint64, uint64, error) {
	query := queryLinkPatientToDiary
	resp, err := er.dbm.Query(query,
		patientId, diaryId)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return 0, 0, err
	}
	// return 0, 0, nil
	patientId = cast.ToUint64(resp[0][0]) 
	diaryId = cast.ToUint64(resp[0][1])
	return patientId, diaryId, nil
}


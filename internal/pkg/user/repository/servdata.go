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

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


func (er *dbuserrepository) UserInit(userInitInfo domain.UserInitRequest, userId uint64) (domain.UserInitResponse, error) {
	var query string
	if userInitInfo.InitBasicInfo.IsMedic{
		query = queryMedicInit
	} else {
		query = queryPatientInit
	}
	resp, err := er.dbm.Query(query,
		userId,
		userInitInfo.InitBasicInfo.Name)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		println(userId)
		println(userInitInfo.InitBasicInfo.Name)
		return domain.UserInitResponse{}, err
	}

	return domain.UserInitResponse{
		Id:           cast.ToUint64(resp[0][0]),
		InitBasicInfo:     domain.InitBasicInfo{
			Name: cast.ToString(resp[0][1]),
			IsMedic: userInitInfo.InitBasicInfo.IsMedic,
		},
	}, nil
}

package mlservicesrepository

import (
	"hesh/internal/pkg/database"
	"hesh/internal/pkg/domain"
)

type dbmlservicesrepository struct {
	dbm *database.DBManager
}

func InitMLServicesRep(manager *database.DBManager) domain.MLServicesRepository {
	return &dbmlservicesrepository{
		dbm: manager,
	}
}

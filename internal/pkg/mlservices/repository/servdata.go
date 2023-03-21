package mlservicesrepository

import (
	"hesh/internal/pkg/database"
	"hesh/internal/pkg/domain"
	// "strings"

	// diaryrepository "hesh/internal/pkg/diary/repository"
	// "hesh/internal/pkg/utils/cast"
	// "hesh/internal/pkg/utils/log"

	// "strings"
	// "time"
)

type dbmlservicesrepository struct {
	dbm *database.DBManager
}

func InitMLServicesRep(manager *database.DBManager) domain.MLServicesRepository {
	return &dbmlservicesrepository{
		dbm: manager,
	}
}

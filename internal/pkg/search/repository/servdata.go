package searchrepository

import (
	"fmt"
	"hesh/internal/pkg/database"
	"hesh/internal/pkg/domain"
	"hesh/internal/pkg/utils/cast"
	"hesh/internal/pkg/utils/log"
)

type dbsearchrepository struct {
	dbm *database.DBManager
}

func InitSearchRep(manager *database.DBManager) domain.SearchRepository {
	return &dbsearchrepository{
		dbm: manager,
	}
}

func (er *dbsearchrepository) CheckUserRole(userId uint64) (bool, bool, error) {
	query := fmt.Sprintf(queryCheckUserRole, er.dbm.EncryptionKey, er.dbm.EncryptionKey)
	resp, err := er.dbm.Query(query,
		userId)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return false, false, err
	}
	if len(resp) == 0 {
		return false, false, nil
	}
	return true, cast.ToBool(resp[0][0]), nil
}

func (cr *dbsearchrepository) SearchDiary (userId uint64, isMedic bool, searchParams domain.SearchDiaryRequest) (domain.DiaryListResponse, error) {
	var resp []database.DBbyterow
	var err error

	// query := queryGetNote3
	query := querySearchDiary
	visavis := "medics"
	if isMedic {
		visavis = "patients"
	}
	textInDiary := "%" + searchParams.Text +"%"
	query = fmt.Sprintf(query,cr.dbm.EncryptionKey, cr.dbm.EncryptionKey, cr.dbm.EncryptionKey, cr.dbm.EncryptionKey, textInDiary, textInDiary, visavis, textInDiary)
	resp, err = cr.dbm.Query(query, userId)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return domain.DiaryListResponse{}, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		return domain.DiaryListResponse{}, nil
	}
	diaries := make([]domain.DiaryInList, 0)
	for i := range resp {
		patientId := uint64(0)
		if (resp[i][3] != nil) {
			patientId = cast.ToUint64(resp[i][3])
		}
		diaries = append(diaries, domain.DiaryInList{
			Id:           cast.ToUint64(resp[i][0]),
			MedicId:      cast.ToUint64(resp[i][1]),
			MedicName:    cast.ToString(resp[i][2]),
			PatientId:    patientId,
			PatientName:  cast.ToString(resp[i][4]),
			CreatingDate: cast.TimeToStr(cast.ToTime(resp[i][5]), true),
			Title:        cast.ToString(resp[i][6]),
			Objectively:  cast.ToString(resp[i][7]),
			LinkToken:    cast.ToString(resp[i][8]),
			IsComplete:   cast.ToBool(resp[i][9]),
		})
	}

	return domain.DiaryListResponse{DiaryList: diaries}, nil
}

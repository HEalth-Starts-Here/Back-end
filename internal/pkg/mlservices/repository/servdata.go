package mlservicesrepository

import (
	"hesh/internal/pkg/database"
	"hesh/internal/pkg/domain"
	"hesh/internal/pkg/utils/cast"
	"hesh/internal/pkg/utils/log"
	"time"
)

type dbmlservicesrepository struct {
	dbm *database.DBManager
}

func InitMLServicesRep(manager *database.DBManager) domain.MLServicesRepository {
	return &dbmlservicesrepository{
		dbm: manager,
	}
}


func (msr *dbmlservicesrepository) GetAudioNames() (map[string]struct{}, error) {
	var resp []database.DBbyterow
	var err error
	query := queryGetAudioList
	resp, err = msr.dbm.Query(query)

	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return nil, domain.Err.ErrObj.InternalServer
	}

	imageNames := make(map[string]struct{}, 0)
	for i := range resp {
		imageNames[cast.ToString(resp[i][0])] = struct{}{}
	}

	return imageNames, nil
}

func (rr *dbmlservicesrepository) CreateMedicRecordDiarisation(recordId uint64, record domain.DiarisationInfo) (domain.DiarisationResponse, error) {
	query := queryCreateMedicRecordDiarisation
	resp, err := rr.dbm.Query(query,
		recordId,
		time.Now().Format("2006.01.02 15:04:05"),
		record.Diarisation,
		record.Filename)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return domain.DiarisationResponse{}, err
	}
	response := domain.DiarisationResponse{
		Id:           cast.ToUint64(resp[0][0]),
		CreatingDate: cast.TimeToStr(cast.ToTime(resp[0][1]), true),
		MedicRecordId: cast.ToUint64(resp[0][2]),
		DiarisationInfo: domain.DiarisationInfo{
			Diarisation: cast.ToString(resp[0][3]),
			Filename: cast.ToString(resp[0][4]),
		},
	}
	
	if err != nil {
		log.Error(err)
		return domain.DiarisationResponse{}, err
	}
	return response, nil
}
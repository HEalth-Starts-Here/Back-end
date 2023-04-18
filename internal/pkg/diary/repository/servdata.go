package diaryrepository

import (
	"hesh/internal/pkg/database"
	"hesh/internal/pkg/domain"

	"hesh/internal/pkg/utils/cast"
	"hesh/internal/pkg/utils/log"

	"time"
)

type dbdiaryrepository struct {
	dbm *database.DBManager
}

func InitDiaryRep(manager *database.DBManager) domain.DiaryRepository {
	return &dbdiaryrepository{
		dbm: manager,
	}
}

func (er *dbdiaryrepository) CreateDiary(diary domain.DiaryCreateRequest, medicId uint64) (domain.DiaryCreateResponse, error) {
	query := queryCreateDiary
	resp, err := er.dbm.Query(query,
		medicId,
		// nil,
		time.Now().Format("2006.01.02 15:04:05"),
		diary.DiaryBasicInfo.Title,
		diary.DiaryBasicInfo.Complaints,
		diary.DiaryBasicInfo.Anamnesis,
		diary.DiaryBasicInfo.Objectively,
		diary.DiaryBasicInfo.Diagnosis)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return domain.DiaryCreateResponse{}, err
	}

	return domain.DiaryCreateResponse{
		Id:           cast.ToUint64(resp[0][0]),
		MedicId:      cast.ToUint64(resp[0][1]),
		PatientId:    0,
		CreatingDate: cast.TimeToStr(cast.ToTime(resp[0][2]), true),
		DiaryBasicInfo: domain.DiaryBasicInfo{
			Title:       cast.ToString(resp[0][3]),
			Complaints:  cast.ToString(resp[0][4]),
			Anamnesis:   cast.ToString(resp[0][5]),
			Objectively: cast.ToString(resp[0][6]),
			Diagnosis:   cast.ToString(resp[0][7]),
		},
	}, nil
}

func (er *dbdiaryrepository) LinkDiary(diaryId uint64, medicId uint64) (domain.DiaryLinkResponse, error) {
	query := queryLinkDiary
	resp, err := er.dbm.Query(query,
		diaryId,
		medicId)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return domain.DiaryLinkResponse{}, err
	}

	return domain.DiaryLinkResponse{
		Id:           cast.ToUint64(resp[0][0]),
		MedicId:      cast.ToUint64(resp[0][1]),
		MedicName:    cast.ToString(resp[0][2]),
		PatientId:    cast.ToUint64(resp[0][3]),
		CreatingDate: cast.TimeToStr(cast.ToTime(resp[0][4]), true),
		DiaryBasicInfo: domain.DiaryBasicInfo{
			Title:       cast.ToString(resp[0][5]),
			Complaints:  cast.ToString(resp[0][6]),
			Anamnesis:   cast.ToString(resp[0][7]),
			Objectively: cast.ToString(resp[0][8]),
			Diagnosis:   cast.ToString(resp[0][9]),
		},
	}, nil
}

func (er *dbdiaryrepository) DeleteDiary(diaryId uint64) error {
	query := queryDeleteDiary
	_, err := er.dbm.Query(query, diaryId)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
	}
	return err
}

func (cr *dbdiaryrepository) GetDiary(userId uint64) (domain.DiaryListResponse, error) {
	var resp []database.DBbyterow
	var err error

	query := queryDiaryList
	resp, err = cr.dbm.Query(query, userId)

	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return domain.DiaryListResponse{}, domain.Err.ErrObj.InternalServer
	}

	diaries := make([]domain.DiaryInList, 0)
	for i := range resp {
		diaries = append(diaries, domain.DiaryInList{
			Id:           cast.ToUint64(resp[i][0]),
			MedicId:      cast.ToUint64(resp[i][1]),
			MedicName:    cast.ToString(resp[i][2]),
			PatientId:    cast.ToUint64(resp[i][3]),
			PatientName:  cast.ToString(resp[i][4]),
			CreatingDate: cast.TimeToStr(cast.ToTime(resp[0][5]), true),
			Title:        cast.ToString(resp[i][6]),
			Objectively:  cast.ToString(resp[i][7]),
		})
	}

	out := domain.DiaryListResponse{
		DiaryList: diaries,
	}

	return out, nil
}

func (er *dbdiaryrepository) GetUserRole(userId uint64) (bool, bool, error) {
	query := queryGetUserRole
	resp, err := er.dbm.Query(query, userId)
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

func (dr *dbdiaryrepository) GetCertainDiary(diaryId uint64) (domain.DiaryResponse, error) {
	var resp []database.DBbyterow
	var err error
	query := queryGetCertainDiaryMainInfo
	resp, err = dr.dbm.Query(query, diaryId)

	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return domain.DiaryResponse{}, domain.Err.ErrObj.InternalServer
	}

	if len(resp) == 0 {
		log.Warn("{GetCertainDiary}")
		log.Error(domain.Err.ErrObj.SmallDb)
		return domain.DiaryResponse{}, domain.Err.ErrObj.SmallDb
	}

	diary := domain.DiaryResponse{
		PatientName: cast.ToString(resp[0][0]),
		Diary: domain.DiaryLinkResponse{
			Id:           cast.ToUint64(resp[0][1]),
			MedicId:      cast.ToUint64(resp[0][2]),
			MedicName:    cast.ToString(resp[0][3]),
			PatientId:    cast.ToUint64(resp[0][4]),
			CreatingDate: cast.TimeToStr(cast.ToTime(resp[0][5]), true),
			DiaryBasicInfo: domain.DiaryBasicInfo{
				Title:       cast.ToString(resp[0][6]),
				Complaints:  cast.ToString(resp[0][7]),
				Anamnesis:   cast.ToString(resp[0][8]),
				Objectively: cast.ToString(resp[0][9]),
				Diagnosis:   cast.ToString(resp[0][10]),
			},
		},
	}

	var resp2 []database.DBbyterow
	var err2 error
	query2 := queryGetCertainDiaryMedicRecords
	resp2, err2 = dr.dbm.Query(query2, diaryId)

	if err2 != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err2)
		return domain.DiaryResponse{}, domain.Err.ErrObj.InternalServer
	}

	medicRecords := make([]domain.RecordInDiaryBasicInfo, 0)
	for i := range resp2 {
		RecordCreateResponse := domain.RecordInDiaryBasicInfo{
			Id:           cast.ToUint64(resp2[i][0]),
			CreatingDate: cast.TimeToStr(cast.ToTime(resp2[i][1]), true),
			Title:        cast.ToString(resp2[i][2]),
			Details:      cast.ToString(resp2[i][3]),
		}
		medicRecords = append(medicRecords, RecordCreateResponse)
	}

	var resp3 []database.DBbyterow
	var err3 error
	//TODO replace this query with queryGetCertainDiaryRecords
	query3 := queryGetCertainDiaryPatientRecords
	resp3, err3 = dr.dbm.Query(query3, diaryId)

	if err3 != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query3)
		log.Error(err3)
		return domain.DiaryResponse{}, domain.Err.ErrObj.InternalServer
	}
	patientRecords := make([]domain.PatientRecordInDiaryBasicInfo, 0)
	for i := range resp3 {
		RecordCreateResponse := domain.PatientRecordInDiaryBasicInfo{
			RecordInDiaryBasicInfo: domain.RecordInDiaryBasicInfo{
				Id:           cast.ToUint64(resp3[i][0]),
				CreatingDate: cast.TimeToStr(cast.ToTime(resp3[i][1]), true),
				Title:        cast.ToString(resp3[i][2]),
				Details:      cast.ToString(resp3[i][3]),
			},
			Feelings: cast.ToUint64(resp3[i][4]),
		}
		patientRecords = append(patientRecords, RecordCreateResponse)
	}

	out := domain.DiaryResponse{
		PatientName: diary.PatientName,
		Diary:       diary.Diary,
		Records: domain.Records{
			MedicRecordList:   medicRecords,
			PatientRecordList: patientRecords,
		},
	}

	return out, nil
}

func (er *dbdiaryrepository) UpdateDiary(diary domain.DiaryUpdateRequest, diaryId uint64) (domain.DiaryUpdateResponse, error) {
	query := queryUpdateDiary
	resp, err := er.dbm.Query(query,
		diary.DiaryBasicInfo.Title,
		diary.DiaryBasicInfo.Complaints,
		diary.DiaryBasicInfo.Anamnesis,
		diary.DiaryBasicInfo.Objectively,
		diary.DiaryBasicInfo.Diagnosis,
		diaryId)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return domain.DiaryUpdateResponse{}, err
	}

	return domain.DiaryUpdateResponse{
		Id:           cast.ToUint64(resp[0][0]),
		MedicId:      cast.ToUint64(resp[0][1]),
		PatientId:    cast.ToUint64(resp[0][2]),
		CreatingDate: cast.TimeToStr(cast.ToTime(resp[0][3]), true),
		DiaryBasicInfo: domain.DiaryBasicInfo{
			Title:       cast.ToString(resp[0][4]),
			Complaints:  cast.ToString(resp[0][5]),
			Anamnesis:   cast.ToString(resp[0][6]),
			Objectively: cast.ToString(resp[0][7]),
			Diagnosis:   cast.ToString(resp[0][8]),
		},
	}, nil
}

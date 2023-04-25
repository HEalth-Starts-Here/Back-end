package diaryrepository

import (
	"fmt"
	"hesh/internal/pkg/database"
	"hesh/internal/pkg/domain"

	// null "github.com/volatiletech/null/v9"

	"hesh/internal/pkg/utils/cast"
	"hesh/internal/pkg/utils/log"

	"time"
	// "database/sql"
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
	fmt.Printf("diary.DiaryBasicInfo.Reminder.StartDate: %v\n", diary.DiaryBasicInfo.Reminder.StartDate)
	startDate, err := cast.StringToDate(diary.DiaryBasicInfo.Reminder.StartDate)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return domain.DiaryCreateResponse{}, err
	}
	fmt.Printf("startDate: %v\n", startDate)
	resp, err := er.dbm.Query(query,
		medicId,
		// nil,
		time.Now().Format("2006.01.02 15:04:05"),
		diary.DiaryBasicInfo.Title,
		diary.DiaryBasicInfo.Complaints,
		diary.DiaryBasicInfo.Anamnesis,
		diary.DiaryBasicInfo.Objectively,
		diary.DiaryBasicInfo.Diagnosis,
		diary.DiaryBasicInfo.Reminder.Variant,
		diary.DiaryBasicInfo.Reminder.Frequency,
		startDate)
		// diary.DiaryBasicInfo.Reminder.StartDate)
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
			Reminder : domain.ReminderInfo {
				Variant: cast.ToBool(resp[0][8]),
				Frequency: cast.ToUint64(resp[0][9]),
				StartDate: cast.TimeToStr(cast.ToDate(resp[0][10]), false),
			},
		},
	}, nil
}

func (er *dbdiaryrepository) CreateLinkToken(diaryId uint64, token string) error {
	query := queryCreateDiaryLinkToken
	_, err := er.dbm.Query(query,
		diaryId,
		token)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return err
	}

	return nil
}

func (er *dbdiaryrepository) CheckAndDeleteToken(diaryId uint64, linkToken string) (bool, error) {
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

func (er *dbdiaryrepository) LinkDiary(patientId uint64, diaryId uint64) (domain.DiaryLinkResponse, error) {
	query := queryLinkDiary
	resp, err := er.dbm.Query(query,
		diaryId,
		patientId)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return domain.DiaryLinkResponse{}, err
	}

	return domain.DiaryLinkResponse{
		Id:        cast.ToUint64(resp[0][0]),
		MedicId:   cast.ToUint64(resp[0][1]),
		MedicName: cast.ToString(resp[0][2]),
		// TODO get new value after update
		PatientId:    patientId,
		CreatingDate: cast.TimeToStr(cast.ToTime(resp[0][4]), true),
		DiaryBasicInfo: domain.DiaryBasicInfo{
			Title:       cast.ToString(resp[0][5]),
			Complaints:  cast.ToString(resp[0][6]),
			Anamnesis:   cast.ToString(resp[0][7]),
			Objectively: cast.ToString(resp[0][8]),
			Diagnosis:   cast.ToString(resp[0][9]),
			Reminder: domain.ReminderInfo{
				Variant: cast.ToBool(resp[0][10]),
				Frequency: cast.ToUint64(resp[0][11]),
				StartDate: cast.TimeToStr(cast.ToDate(resp[0][12]), false),
				
			},
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

func (er *dbdiaryrepository) CompleteDiary(diaryId uint64) error {
	query := queryCompleteDiary
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
	// if len(resp) == 0 {
	// 	return domain.DiaryListResponse{}, domain.Err.ErrObj.SmallDb
	// }
	diaries := make([]domain.DiaryInList, 0)
	for i := range resp {
		// a = (resp[i][3])
		patientId := uint64(0)
		if len(resp[i][3]) != 0 {
			patientId = cast.ToUint64(resp[i][3])
		}
		diaries = append(diaries, domain.DiaryInList{
			Id:           cast.ToUint64(resp[i][0]),
			MedicId:      cast.ToUint64(resp[i][1]),
			MedicName:    cast.ToString(resp[i][2]),
			PatientId:    patientId,
			PatientName:  cast.ToString(resp[i][4]),
			CreatingDate: cast.TimeToStr(cast.ToTime(resp[0][5]), true),
			Title:        cast.ToString(resp[i][6]),
			Objectively:  cast.ToString(resp[i][7]),
			LinkToken:    cast.ToString(resp[i][8]),
			IsComplete:   cast.ToBool(resp[i][9]),
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
	// var a sql.NullInt64
	// var a null.Bytes{}
	// var b []byte
	// b = nil

	patientId := uint64(0)
	if (resp[0][4]) != nil {
		patientId = cast.ToUint64(resp[0][4])
	}
	// a = sql.NullInt64(resp[0][3])
	// if len(resp[0][3]) != 0 {
	// 	patientId = cast.ToUint64(resp[0][3])
	// }
	diary := domain.DiaryResponse{
		PatientName: cast.ToString(resp[0][0]),
		Diary: domain.DiaryLinkResponse{
			Id:        cast.ToUint64(resp[0][1]),
			MedicId:   cast.ToUint64(resp[0][2]),
			MedicName: cast.ToString(resp[0][3]),
			// PatientId:    cast.ToString(resp[0][4]),
			PatientId:    patientId,
			CreatingDate: cast.TimeToStr(cast.ToTime(resp[0][5]), true),
			DiaryBasicInfo: domain.DiaryBasicInfo{
				Title:       cast.ToString(resp[0][6]),
				Complaints:  cast.ToString(resp[0][7]),
				Anamnesis:   cast.ToString(resp[0][8]),
				Objectively: cast.ToString(resp[0][9]),
				Diagnosis:   cast.ToString(resp[0][10]),
				Reminder: domain.ReminderInfo{
					Variant: cast.ToBool(resp[0][11]),
					Frequency: cast.ToUint64(resp[0][12]),
					StartDate: cast.TimeToStr(cast.ToDate(resp[0][13]), false),
				},			
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
	startDate, err := cast.StringToDate(diary.DiaryBasicInfo.Reminder.StartDate)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return domain.DiaryUpdateResponse{}, err
	}
	fmt.Printf("startDate: %v\n", startDate)
	resp, err := er.dbm.Query(query,
		diary.DiaryBasicInfo.Title,
		diary.DiaryBasicInfo.Complaints,
		diary.DiaryBasicInfo.Anamnesis,
		diary.DiaryBasicInfo.Objectively,
		diary.DiaryBasicInfo.Diagnosis,
		diary.DiaryBasicInfo.Reminder.Variant,
		diary.DiaryBasicInfo.Reminder.Frequency,
		startDate,
		diaryId)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return domain.DiaryUpdateResponse{}, err
	}
	patientId := uint64(0)
	if resp[0][2] != nil {
		patientId = cast.ToUint64(resp[0][2])
	}

	return domain.DiaryUpdateResponse{
		Id:      cast.ToUint64(resp[0][0]),
		MedicId: cast.ToUint64(resp[0][1]),
		// PatientId:    cast.ToUint64(resp[0][2]),
		PatientId:    patientId,
		CreatingDate: cast.TimeToStr(cast.ToTime(resp[0][3]), true),
		DiaryBasicInfo: domain.DiaryBasicInfo{
			Title:       cast.ToString(resp[0][4]),
			Complaints:  cast.ToString(resp[0][5]),
			Anamnesis:   cast.ToString(resp[0][6]),
			Objectively: cast.ToString(resp[0][7]),
			Diagnosis:   cast.ToString(resp[0][8]),
			Reminder: domain.ReminderInfo{
				Variant: cast.ToBool(resp[0][9]),
				Frequency: cast.ToUint64(resp[0][10]),
				StartDate: cast.TimeToStr(cast.ToDate(resp[0][11]), false),
			},
		},
	}, nil
}

package diaryrepository

import (
	"hesh/internal/pkg/database"
	"hesh/internal/pkg/domain"
	"strings"

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

func (cr *dbdiaryrepository) GetImageNames() (map[string]struct{}, error) {
	var resp []database.DBbyterow
	var err error
	query := ""

	query = queryGetImageList
	resp, err = cr.dbm.Query(query)

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

func (er *dbdiaryrepository) CreateDiary(diary domain.DiaryCreateRequest, medicId uint64) (domain.DiaryCreateResponse, error) {
	query := queryCreateDiary
	resp, err := er.dbm.Query(query,
		medicId,
		0,
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
	_, err := er.dbm.Query(queryDeleteDiary, diaryId)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + queryCreateDiary)
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

func (er *dbdiaryrepository) GetRecordImageLists(recordId uint64) ([]domain.ImageInfo, error) {
	query := queryGetImageList
	resp, err := er.dbm.Query(query, recordId)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return []domain.ImageInfo{}, err
	}
	response := []domain.ImageInfo{}
	for i := range resp {
		response = append(response, domain.ImageInfo{
			Id:       cast.ToUint64(resp[i][0]),
			RecordId: cast.ToUint64(resp[i][1]),
			Name:     cast.ToString(resp[i][2]),
			Area:     cast.ToFloat64(resp[i][3]),
		})
	}
	return response, nil
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
	println(cast.ToBool(resp[0][0]))
	println(cast.ToBool(resp[0][0]))
	println(cast.ToBool(resp[0][0]))
	println(cast.ToBool(resp[0][0]))
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

	records1 := make([]domain.RecordBasicInfo, 0)
	for i := range resp2 {
		RecordCreateResponse := domain.RecordBasicInfo{
			Id:			  cast.ToUint64(resp2[i][0]),
			CreatingDate: cast.TimeToStr(cast.ToTime(resp2[i][1]), true),
			Title:        cast.ToString(resp2[i][2]),
			Details:      cast.ToString(resp2[i][3]),
		}
		records1 = append(records1, RecordCreateResponse)
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
	records2 := make([]domain.RecordBasicInfo, 0)
	for i := range resp3 {
		RecordCreateResponse := domain.RecordBasicInfo{
			Id:			  cast.ToUint64(resp2[i][0]),
			CreatingDate: cast.TimeToStr(cast.ToTime(resp2[i][1]), true),
			Title:        cast.ToString(resp2[i][2]),
			Details:      cast.ToString(resp2[i][3]),
		}
		records2 = append(records2, RecordCreateResponse)
	}

	out := domain.DiaryResponse{
		PatientName: diary.PatientName,
		Diary:       diary.Diary,
		Records: domain.Records{
			MedicRecordList:   records1,
			PatientRecordList: records2,
		},
	}

	return out, nil
}

func (er *dbdiaryrepository) CreateRecordImageLists(recordId uint64, imageInfo []domain.ImageInfoUsecase) ([]domain.ImageInfo, error) {
	if len(imageInfo) == 0 {
		return []domain.ImageInfo{}, nil
	}
	queryBuilder := strings.Builder{}
	// arrayForQuery := ""
	queryBuilder.Write([]byte(queryCreateRecordImageListFirstPart))
	for i := range imageInfo {
		queryBuilder.Write([]byte("("))
		queryBuilder.Write([]byte(cast.IntToStr(recordId)))
		queryBuilder.Write([]byte(","))
		queryBuilder.Write([]byte("'" + imageInfo[i].Name + "'"))
		queryBuilder.Write([]byte(","))
		queryBuilder.Write([]byte(cast.FlToStr(imageInfo[i].Area)))
		queryBuilder.Write([]byte(")"))
		if i != len(imageInfo)-1 {
			queryBuilder.Write([]byte(","))
		}
	}
	queryBuilder.Write([]byte(queryCreateRecordImageListSecondPart))
	query := queryBuilder.String()
	resp, err := er.dbm.Query(queryBuilder.String())
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return []domain.ImageInfo{}, err
	}
	response := []domain.ImageInfo{}
	for i := range resp {
		response = append(response, domain.ImageInfo{
			Id:       cast.ToUint64(resp[i][0]),
			RecordId: cast.ToUint64(resp[i][1]),
			Name:     cast.ToString(resp[i][2]),
			Area:     cast.ToFloat64(resp[i][3]),
		})
	}
	return response, nil
}

func (er *dbdiaryrepository) CreateRecord(diaryId uint64, record domain.RecordCreateRequest, imageInfo []domain.ImageInfoUsecase, Area float64) (domain.RecordCreateResponse, error) {
	query := queryCreateRecord
	resp, err := er.dbm.Query(queryCreateRecord,
		diaryId,
		time.Now().Format("2006.01.02 15:04:05"),
		record.Title,
		record.Description,
		Area,
		record.Characteristics.Dryness,
		record.Characteristics.Edema,
		record.Characteristics.Itching,
		record.Characteristics.Pain,
		record.Characteristics.Peeling,
		record.Characteristics.Redness)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return domain.RecordCreateResponse{}, err
	}

	response := domain.RecordCreateResponse{
		Id:           cast.ToUint64(resp[0][0]),
		DiaryId:      cast.ToUint64(resp[0][1]),
		CreatingDate: cast.TimeToStr(cast.ToTime(resp[0][2]), true),
		Title:        cast.ToString(resp[0][3]),
		Description:  cast.ToString(resp[0][4]),
		Area:         cast.ToFloat64(resp[0][5]),

		Characteristics: domain.Characteristics{
			Dryness: cast.ToUint8(resp[0][6]),
			Edema:   cast.ToUint8(resp[0][7]),
			Itching: cast.ToUint8(resp[0][8]),
			Pain:    cast.ToUint8(resp[0][9]),
			Peeling: cast.ToUint8(resp[0][10]),
			Redness: cast.ToUint8(resp[0][11]),
		},
		ImageList: []domain.ImageInfo{},
	}

	imagenames := []string{}
	for i := range imageInfo {
		imagenames = append(imagenames, imageInfo[i].Name)
	}
	// TODO check case with 0 images
	response.ImageList, err = er.CreateRecordImageLists(response.Id, imageInfo)
	if err != nil {
		log.Error(err)
		return domain.RecordCreateResponse{}, err
	}
	return response, nil
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

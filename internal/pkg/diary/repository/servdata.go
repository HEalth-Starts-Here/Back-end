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

func (er *dbdiaryrepository) CreateDiary(diary domain.DiaryCreateRequest, medicId uint32) (domain.DiaryCreateResponse, error) {
	query := queryCreateDiary
	resp, err := er.dbm.Query(	query, 
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
		MedicId:      cast.ToUint32(resp[0][1]),
		PatientId:    cast.ToUint32(resp[0][2]),
		CreatingDate: cast.TimeToStr(cast.ToTime(resp[0][3]), true),
		DiaryBasicInfo: domain.DiaryBasicInfo{
			Title:        cast.ToString(resp[0][4]),
			Complaints:   cast.ToString(resp[0][5]),
			Anamnesis:    cast.ToString(resp[0][6]),
			Objectively:  cast.ToString(resp[0][7]),
			Diagnosis:    cast.ToString(resp[0][8]),
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

// func (er *dbdiaryrepository) CreateEventCategory(eventId uint64, categories []string) ([]string, error) {
// 	// var sb strings.Builder
// 	// sb.WriteString(queryCreateEventCategoryFirstPart)
// 	// for i, el := range(categories) {
// 	// 	sb.WriteString(queryCreateEventCategorySecondPart)
// 	// 	sb.WriteString(string(i * 2))
// 	// 	sb.WriteString(queryCreateEventCategoryThirdPart)
// 	// 	sb.WriteString(string(i * 2 + 1))
// 	// 	sb.WriteString(queryCreateEventCategoryForthPart)
// 	// 	sb.WriteString(",")
// 	// }
// 	// sb.WriteString(queryCreateEventCategoryFifthPart)
// 	// var resp []database.DBbyterow
// 	var err error
// 	for i, _ := range categories {
// 		_, err = er.dbm.Query(queryCreateEventCategory, eventId, categories[i])
// 		if err != nil {
// 			log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
// 			log.Error(err)
// 			return nil, err
// 		}
// 	}

// 	return categories, nil
// }

// func (er *dbeventrepository) EventAlreadyExist(event domain.EventCreatingRequest) (bool, error) {
// 	resp, err := er.dbm.Query(queryCheckEvent, event.Title, event.Longitude, event.Latitude)
// 	if err != nil {
// 		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
// 		log.Error(err)
// 		return false, err
// 	}

// 	if cast.ToUint64(resp[0][0]) != 0 {
// 		return true, nil
// 	}
// 	return false, nil
// }

func (cr *dbdiaryrepository) GetDiary() (domain.DiaryListResponse, error) {
	var resp []database.DBbyterow
	var err error

	query := queryDiaryList
	resp, err = cr.dbm.Query(query)

	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return domain.DiaryListResponse{}, domain.Err.ErrObj.InternalServer
	}

	diaries := make([]domain.DiaryInList, 0)
	for i := range resp {
		diaries = append(diaries, domain.DiaryInList{
			Id: 			cast.ToUint64(resp[i][0]),
			MedicId:		cast.ToUint32(resp[i][1]),
			MedicName:		cast.ToString(resp[i][2]),
			PatientId:		cast.ToUint32(resp[i][3]),
			PatientName:	cast.ToString(resp[i][4]),
			CreatingDate:	cast.TimeToStr(cast.ToTime(resp[0][5]), true),
			Title: 			cast.ToString(resp[i][6]),
			Objectively:	cast.ToString(resp[i][7]),
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

	diary := domain.DiaryCreateResponse{
		Id:           cast.ToUint64(resp[0][0]),
		MedicId:      cast.ToUint32(resp[0][1]),
		PatientId:    cast.ToUint32(resp[0][2]),
		CreatingDate: cast.TimeToStr(cast.ToTime(resp[0][3]), true),
		DiaryBasicInfo: domain.DiaryBasicInfo{
			Title:        cast.ToString(resp[0][4]),
			Complaints:   cast.ToString(resp[0][5]),
			Anamnesis:    cast.ToString(resp[0][6]),
			Objectively:  cast.ToString(resp[0][7]),
			Diagnosis:    cast.ToString(resp[0][8]),
		},
	}

	var resp2 []database.DBbyterow
	var err2 error
	query2 := queryGetCertainDiaryRecords
	resp2, err2 = dr.dbm.Query(query2, diaryId)

	if err2 != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err2)
		return domain.DiaryResponse{}, domain.Err.ErrObj.InternalServer
	}

	records := make([]domain.RecordCreateResponse, 0)
	for i := range resp2 {
		RecordCreateResponse := domain.RecordCreateResponse{
			Id:           cast.ToUint64(resp2[i][0]),
			DiaryId:      cast.ToUint64(resp2[i][1]),
			CreatingDate: cast.TimeToStr(cast.ToTime(resp2[i][2]), true),
			Description:  cast.ToString(resp2[i][3]),
			Title:        cast.ToString(resp2[i][4]),
			Area:         cast.ToFloat64(resp2[i][5]),
			Characteristics: domain.Characteristics{
				Dryness: cast.ToUint8(resp2[i][6]),
				Edema:   cast.ToUint8(resp2[i][7]),
				Itching: cast.ToUint8(resp2[i][8]),
				Pain:    cast.ToUint8(resp2[i][9]),
				Peeling: cast.ToUint8(resp2[i][10]),
				Redness: cast.ToUint8(resp2[i][11]),
			},
		}
		RecordCreateResponse.ImageList, err = dr.GetRecordImageLists(RecordCreateResponse.Id)
		if err != nil {
			log.Error(err)
			return domain.DiaryResponse{}, domain.Err.ErrObj.InternalServer
		}
		records = append(records, RecordCreateResponse)
	}

	out := domain.DiaryResponse{
		Diary:       diary,
		RecordsList: records,
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

func (er *dbdiaryrepository) UpdateDiary(diary domain.DiaryUpdateRequest) (domain.DiaryUpdateResponse, error) {
	query := queryUpdateDiary
	resp, err := er.dbm.Query(query,
		diary.Title,
		diary.Description,
		diary.Id)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return domain.DiaryUpdateResponse{}, err
	}

	return domain.DiaryUpdateResponse{
		Id:          cast.ToUint64(resp[0][0]),
		Title:       cast.ToString(resp[0][1]),
		Description: cast.ToString(resp[0][2]),
	}, nil
}

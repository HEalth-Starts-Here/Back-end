package diaryrepository

import (
	"hesh/internal/pkg/database"
	"hesh/internal/pkg/domain"
	"strings"

	// diaryrepository "hesh/internal/pkg/diary/repository"
	"hesh/internal/pkg/utils/cast"
	"hesh/internal/pkg/utils/log"

	// "strings"
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

func (er *dbdiaryrepository) CreateDiary(diary domain.DiaryCreatingRequest) (domain.DiaryCreatingResponse, error) {
	resp, err := er.dbm.Query(queryCreateDiary, diary.MedicId,
		diary.PatientId, time.Now().Format("2006.01.02 15:04:05"), diary.Title, diary.Description)
	if err != nil {
		log.Warn("{CreateDiary} in query: " + queryCreateDiary)
		log.Error(err)
		return domain.DiaryCreatingResponse{}, err
	}

	return domain.DiaryCreatingResponse{
		Id:                     cast.ToUint64(resp[0][0]),
		MedicId:                cast.ToUint32(resp[0][1]),
		PatientId:              cast.ToUint32(resp[0][2]),
		CreatingDate:           cast.TimeToStr(cast.ToTime(resp[0][3]), true),
		Title:                  cast.ToString(resp[0][4]),
		Description:            cast.ToString(resp[0][5]),
	}, nil
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
// 			log.Warn("{CreateEventCategory} in query: " + queryCreateEventCategory)
// 			log.Error(err)
// 			return nil, err
// 		}
// 	}

// 	return categories, nil
// }

// func (er *dbeventrepository) EventAlreadyExist(event domain.EventCreatingRequest) (bool, error) {
// 	resp, err := er.dbm.Query(queryCheckEvent, event.Title, event.Longitude, event.Latitude)
// 	if err != nil {
// 		log.Warn("{EventCreating} in query: " + queryCheckEvent)
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
	query := ""

	query = queryDiaryList
	resp, err = cr.dbm.Query(query)


	if err != nil {
		log.Warn("{GetDiary} in query: " + query)
		log.Error(err)
		return domain.DiaryListResponse{}, domain.Err.ErrObj.InternalServer
	}

	diaries := make([]domain.DiaryCreatingResponse, 0)
	for i := range resp {
		diaries = append(diaries, domain.DiaryCreatingResponse{
			Id:                     cast.ToUint64(resp[i][0]),
			MedicId:                cast.ToUint32(resp[i][1]),
			PatientId:              cast.ToUint32(resp[i][2]),
			CreatingDate:           cast.TimeToStr(cast.ToTime(resp[i][3]), true),
			Title:                  cast.ToString(resp[i][4]),
			Description:            cast.ToString(resp[i][5]),
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
		log.Warn("{GetRecordImageLists} in query: " + query)
		log.Error(err)
		return []domain.ImageInfo{}, err
	}
	response := []domain.ImageInfo{}
	for i := range resp {
		response = append(response, domain.ImageInfo{
			Id:       cast.ToUint64(resp[i][0]),
			RecordId: cast.ToUint64(resp[i][1]),
			Name:	  cast.ToString(resp[i][2]),
			Area:	  cast.ToFloat64(resp[i][3]),
		} )
	}
	return response, nil
}

func (dr *dbdiaryrepository) GetCertainDiary(diaryId uint64) (domain.DiaryResponse, error) {
	var resp []database.DBbyterow
	var err error
	query := queryGetCertainDiaryMainInfo
	resp, err = dr.dbm.Query(query, diaryId)

	if err != nil {
		log.Warn("{GetCertainDiary} in query: " + query)
		log.Error(err)
		return domain.DiaryResponse{}, domain.Err.ErrObj.InternalServer
	}

	if len(resp) == 0 {
		log.Warn("{GetCertainDiary}")
		log.Error(domain.Err.ErrObj.SmallDb)
		return domain.DiaryResponse{}, domain.Err.ErrObj.SmallDb
	}

	diary := domain.DiaryCreatingResponse{
		Id:                     cast.ToUint64(resp[0][0]),
		MedicId:                cast.ToUint32(resp[0][1]),
		PatientId:              cast.ToUint32(resp[0][2]),
		CreatingDate:           cast.TimeToStr(cast.ToTime(resp[0][3]), true),
		Title:                  cast.ToString(resp[0][4]),
		Description:            cast.ToString(resp[0][5]),
	}

	var resp2 []database.DBbyterow
	var err2 error
	query2 := queryGetCertainDiaryRecords
	resp2, err2 = dr.dbm.Query(query2, diaryId)


	if err2 != nil {
		log.Warn("{GetCertainDiaryRecords} in query: " + query2)
		log.Error(err2)
		return domain.DiaryResponse{}, domain.Err.ErrObj.InternalServer
	}

	records := make([]domain.RecordCreatingResponse, 0)
	for i := range resp2 {
		recordCreatingResponse := domain.RecordCreatingResponse{
			Id:                     cast.ToUint64(resp2[i][0]),
			DiaryId:                cast.ToUint64(resp2[i][1]),
			CreatingDate:           cast.TimeToStr(cast.ToTime(resp2[i][2]), true),
			Description:            cast.ToString(resp2[i][3]),
			Title:            	    cast.ToString(resp2[i][4]),
			Area:            	    cast.ToFloat64(resp2[i][5]),
			// TODO add characteristics and image info
		}
		recordCreatingResponse.ImageList, err = dr.GetRecordImageLists(recordCreatingResponse.Id)
		if err != nil {
			log.Error(err)
			return domain.DiaryResponse{}, domain.Err.ErrObj.InternalServer
		}
		records = append(records, recordCreatingResponse)
	}


	out := domain.DiaryResponse{
		Diary: diary,
		RecordsList: records,
	}

	return out, nil
}

func (er *dbdiaryrepository) CreateRecordImageLists(recordId uint64, imageInfo []domain.ImageInfoUsecase) ([]domain.ImageInfo, error) {
	query := strings.Builder{}
	// arrayForQuery := ""
	query.Write([]byte(queryCreateRecordImageListFirstPart))
	for i := range imageInfo {
		query.Write([]byte("("))
		query.Write([]byte(cast.IntToStr(recordId)))
		query.Write([]byte(","))
		query.Write([]byte("'" + imageInfo[i].Name + "'"))
		query.Write([]byte(","))
		query.Write([]byte(cast.FlToStr(imageInfo[i].Area)))
		query.Write([]byte(")"))
		if i != len(imageInfo) - 1 {
			query.Write([]byte(","))
		}
	}
	query.Write([]byte(queryCreateRecordImageListSecondPart))
	resp, err := er.dbm.Query(query.String())
	if err != nil {
		log.Warn("{CreateRecordImageLists} in query: " + query.String())
		log.Error(err)
		return []domain.ImageInfo{}, err
	}
	response := []domain.ImageInfo{}
	for i := range resp {
		response = append(response, domain.ImageInfo{
			Id:       cast.ToUint64(resp[i][0]),
			RecordId: cast.ToUint64(resp[i][1]),
			Name:	  cast.ToString(resp[i][2]),
			Area:	  cast.ToFloat64(resp[i][3]),
		} )
	}
	return response, nil
}

func (er *dbdiaryrepository) CreateRecord(diaryId uint64, record domain.RecordCreatingRequest, imageInfo []domain.ImageInfoUsecase, Area float64) (domain.RecordCreatingResponse, error) {
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
		log.Warn("{CreateRecord} in query: " + queryCreateRecord)
		log.Error(err)
		return domain.RecordCreatingResponse{}, err
	}

	response := domain.RecordCreatingResponse{
		Id:                     cast.ToUint64(resp[0][0]),
		DiaryId:                cast.ToUint64(resp[0][1]),
		CreatingDate:           cast.TimeToStr(cast.ToTime(resp[0][2]), true),
		Title:            	    cast.ToString(resp[0][3]),
		Description:            cast.ToString(resp[0][4]),
		Area:          		    cast.ToFloat64(resp[0][5]),

		Characteristics: domain.Characteristics{
			Dryness:		    cast.ToUint8(resp[0][6]),
			Edema: 				cast.ToUint8(resp[0][7]),
			Itching: 			cast.ToUint8(resp[0][8]),
			Pain:				cast.ToUint8(resp[0][9]),
			Peeling: 			cast.ToUint8(resp[0][10]),
			Redness:			cast.ToUint8(resp[0][11]),
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
		return domain.RecordCreatingResponse{}, err
	}	// resp, err = er.dbm.Query(queryGetImageList, response.Id)
	// if err != nil {
	// 	log.Warn("{CreateRecord} in query: " + queryGetImageList)
	// 	log.Error(err)
	// 	return domain.RecordCreatingResponse{}, err
	// }

	// for i := range resp {
	// 	response.ImageList = append(response.ImageList, cast.ToString(resp[i][0]))
	// }

	return response, nil
}


// func (cr *dbeventrepository) GetCategory() (domain.CategoryListResponse, error) {
// 	var resp []database.DBbyterow
// 	var err error
// 	resp, err = cr.dbm.Query(queryGetCategoryList)

// 	if err != nil {
// 		log.Warn("{GetCategory} in query: " + queryGetCategoryList)
// 		log.Error(err)
// 		return domain.CategoryListResponse{}, domain.Err.ErrObj.InternalServer
// 	}

// 	if len(resp) == 0 {
// 		log.Warn("{GetCategory}")
// 		log.Error(domain.Err.ErrObj.SmallDb)
// 		return domain.CategoryListResponse{}, domain.Err.ErrObj.SmallDb
// 	}

// 	categoryList := make([]domain.CategoryResponse, 0)
// 	for i := range resp {
// 		categoryList = append(categoryList, domain.CategoryResponse{
// 			Name:      cast.ToString(resp[i][0]),
// 			ImagePath: cast.ToString(resp[i][1]),
// 		})
// 	}

// 	out := domain.CategoryListResponse{
// 		CategoryList: categoryList,
// 	}

// 	return out, nil
// }

// func (er *dbeventrepository) SignUpUserForEvent(eventId uint64, userId uint64) error {
// 	_, err := er.dbm.Query(querySignUpUserForEvent, cast.IntToStr(eventId), cast.IntToStr(userId))
// 	if err != nil {
// 		log.Warn("{SignUpUserForEvent} in query: " + querySignUpUserForEvent)
// 		log.Error(err)
// 		return err
// 	}

// 	return nil
// }

// func (er *dbeventrepository) CancelEventSignUp(eventId uint64, userId uint64) error {
// 	query := queryCancelSignUpUserForEvent
// 	_, err := er.dbm.Query(query, cast.IntToStr(eventId), cast.IntToStr(userId))
// 	if err != nil {
// 		log.Warn("{CancelEventSignUp} in query: " + query)
// 		log.Error(err)
// 		return err
// 	}

// 	return nil
// }

// func (ur *dbeventrepository) GetUserCategory(id uint64) ([]string, error) {
// 	query := usrrepository.QueryGetCategory
// 	resp, err := ur.dbm.Query(query, id)

// 	if len(resp) == 0 {
// 		log.Warn("{GetCategory}")
// 		log.Error(domain.Err.ErrObj.NoCategory)
// 		return nil, domain.Err.ErrObj.NoCategory
// 	}
// 	if err != nil {
// 		log.Warn("{GetCategory} in query: " + query)
// 		log.Error(err)
// 		return nil, domain.Err.ErrObj.InternalServer
// 	}

// 	category := make([]string, 0)
// 	for i := range resp {
// 		category = append(category, cast.ToString(resp[i][0]))
// 	}

// 	return category, nil
// }

// func (ur *dbeventrepository) GetUserAge(id uint64) (uint64, error) {
// 	query := queryGetUserAge
// 	resp, err := ur.dbm.Query(query, id)

// 	if len(resp) == 0 {
// 		log.Warn("{GetCategory}")
// 		er := domain.Err.ErrObj.NoUser
// 		log.Error(er)
// 		return 0, er
// 	}
// 	if err != nil {
// 		log.Warn("{GetCategory} in query: " + query)
// 		log.Error(err)
// 		return 0, domain.Err.ErrObj.InternalServer
// 	}

// 	age := cast.ToUint64(resp[0][0])

// 	return age, nil
// }

// func (ur *dbeventrepository) GetEventAges(id uint64) (uint16, uint16, error) {
// 	query := queryGetEventAges
// 	resp, err := ur.dbm.Query(query, id)

// 	if len(resp) == 0 {
// 		log.Warn("{GetEventAges}")
// 		er := domain.Err.ErrObj.NoUser
// 		log.Error(er)
// 		return 0, 0, er
// 	}
// 	if err != nil {
// 		log.Warn("{GetEventAges} in query: " + query)
// 		log.Error(err)
// 		return 0, 0, domain.Err.ErrObj.InternalServer
// 	}

// 	minAge := cast.ToUint16(resp[0][0])
// 	maxAge := cast.ToUint16(resp[0][1])

// 	return minAge, maxAge, nil
// }

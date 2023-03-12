package diaryrepository

import (
	"hesh/internal/pkg/database"
	"hesh/internal/pkg/domain"
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
	resp, err := er.dbm.Query(queryCreateDiary, diary.Category, diary.MedicId,
		diary.PatientId, time.Now().Format("2006.01.02 15:04:05"), diary.Title, diary.Description)
	if err != nil {
		log.Warn("{CreateEvent} in query: " + queryCreateDiary)
		log.Error(err)
		return domain.DiaryCreatingResponse{}, err
	}

	return domain.DiaryCreatingResponse{
		Id:                     cast.ToUint64(resp[0][0]),
		Category:               cast.ToUint32(resp[0][1]),
		MedicId:                cast.ToUint32(resp[0][2]),
		PatientId:              cast.ToUint32(resp[0][3]),
		CreatingDate:           cast.ToString(resp[0][4]),
		Title:                  cast.ToString(resp[0][5]),
		Description:            cast.ToString(resp[0][6]),
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
			Category:               cast.ToUint32(resp[i][1]),
			MedicId:                cast.ToUint32(resp[i][2]),
			PatientId:              cast.ToUint32(resp[i][3]),
			CreatingDate:           cast.ToString(resp[i][4]),
			Title:                  cast.ToString(resp[i][5]),
			Description:            cast.ToString(resp[i][6]),
			// Id:                     cast.ToUint64(resp[i][0]),
			// PosterPath:             cast.ToString(resp[i][1]),
			// Title:                  cast.ToString(resp[i][2]),
			// Rating:                 cast.FlToStr((cast.ToFloat64(resp[i][3]))),
			// VotesNum:               cast.ToUint64(resp[i][4]),
			// Description:            cast.ToString(resp[i][5]),
			// UserId:                 cast.ToString(resp[i][6]),
			// Longitude:              cast.FlToStr((cast.ToFloat64(resp[i][7]))),
			// Latitude:               cast.FlToStr((cast.ToFloat64(resp[i][8]))),
			// CurrentMembersQuantity: cast.ToUint64(resp[i][9]),
			// MaxMembersQuantity:     cast.ToUint64(resp[i][10]),
			// MinMembersQuantity:     cast.ToUint64(resp[i][11]),
			// CreatingDate:           cast.ToString(resp[i][12]),
			// StartDate:              cast.ToString(resp[i][13]),
			// EndDate:                cast.ToString(resp[i][14]),
			// MinAge:                 cast.ToString(resp[i][15]),
			// MaxAge:                 cast.ToString(resp[i][16]),
			// Price:                  cast.ToString(resp[i][17]),
		})
	}

	out := domain.DiaryListResponse{
		DiaryList: diaries,
	}

	return out, nil
}

func (cr *dbdiaryrepository) GetCertainDiary(diaryId uint64) (domain.DiaryResponse, error) {
	var resp []database.DBbyterow
	var err error
	query := queryGetCertainDiaryMainInfo
	resp, err = cr.dbm.Query(query, diaryId)

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
		Category:               cast.ToUint32(resp[0][1]),
		MedicId:                cast.ToUint32(resp[0][2]),
		PatientId:              cast.ToUint32(resp[0][3]),
		CreatingDate:           cast.ToString(resp[0][4]),
		Title:                  cast.ToString(resp[0][5]),
		Description:            cast.ToString(resp[0][6]),
	}

	var resp2 []database.DBbyterow
	var err2 error
	query2 := queryGetCertainDiaryRecords
	resp2, err2 = cr.dbm.Query(query2, diaryId)


	if err2 != nil {
		log.Warn("{GetCertainDiaryRecords} in query: " + query2)
		log.Error(err2)
		return domain.DiaryResponse{}, domain.Err.ErrObj.InternalServer
	}

	records := make([]domain.RecordsCreatingResponse, 0)
	for i := range resp2 {
		records = append(records, domain.RecordsCreatingResponse{
			Id:                     cast.ToUint64(resp2[i][0]),
			DiaryId:                cast.ToUint64(resp2[i][1]),
			Description:            cast.ToString(resp2[i][2]),
			PosterPath:             cast.ToString(resp2[i][3]),
		})
	}


	out := domain.DiaryResponse{
		Diary: diary,
		RecordsList: records,
	}

	return out, nil
}

func (er *dbdiaryrepository) CreateRecord(record domain.RecordCreatingRequest) (domain.RecordCreatingResponse, error) {
	resp, err := er.dbm.Query(queryCreateRecord, record.DiaryId, record.PosterPath, record.Description)
	if err != nil {
		log.Warn("{CreateRecord} in query: " + queryCreateDiary)
		log.Error(err)
		return domain.RecordCreatingResponse{}, err
	}

	return domain.RecordCreatingResponse{
		Id:                     cast.ToUint64(resp[0][0]),
		DiaryId:                cast.ToUint64(resp[0][1]),
		PosterPath:             cast.ToString(resp[0][2]),
		Description:            cast.ToString(resp[0][3]),
	}, nil
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

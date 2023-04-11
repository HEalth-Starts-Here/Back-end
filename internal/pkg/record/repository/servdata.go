package recordrepository

import (
	"hesh/internal/pkg/database"
	"hesh/internal/pkg/domain"
	"strconv"
	"strings"
	"time"

	"hesh/internal/pkg/utils/cast"
	"hesh/internal/pkg/utils/log"
)

type dbrecordrepository struct {
	dbm *database.DBManager
}

func InitRecordRep(manager *database.DBManager) domain.RecordRepository {
	return &dbrecordrepository{
		dbm: manager,
	}
}


func (cr *dbrecordrepository) MedicExist(medicId uint64) (bool, error) {
	var resp []database.DBbyterow
	var err error
	query := queryMedicExist
	resp, err = cr.dbm.Query(query, medicId)

	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return false, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		return false, nil
	}
	return true, nil
}

func (cr *dbrecordrepository) DiaryExist(diaryId uint64) (bool, error) {
	var resp []database.DBbyterow
	var err error
	query := queryDiaryExist
	resp, err = cr.dbm.Query(query, diaryId)

	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return false, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		return false, nil
	}
	return true, nil
}

func (cr *dbrecordrepository) MedicRecordExist(diaryId uint64) (bool, error) {
	var resp []database.DBbyterow
	var err error
	query := queryMedicRecordExist
	resp, err = cr.dbm.Query(query, diaryId)

	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return false, domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		return false, nil
	}
	return true, nil
}

// func (cr *dbrecordrepository) MedicExist(diaryId uint64) (bool, error) {
// 	var resp []database.DBbyterow
// 	var err error
// 	query := queryDiaryExist
// 	resp, err = cr.dbm.Query(query, diaryId)

// 	if err != nil {
// 		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
// 		log.Error(err)
// 		return false, domain.Err.ErrObj.InternalServer
// 	}
// 	if len(resp) == 0 {
// 		return false, nil
// 	}
// 	return true, nil
// }

func (cr *dbrecordrepository) GetImageNames() (map[string]struct{}, error) {
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

func (er *dbrecordrepository) CreateRecordImageLists(isMedic bool,recordId uint64, imageInfo []string) ([]uint64, error) {
	if len(imageInfo) == 0 {
		return []uint64{}, nil
	}
	queryBuilder := strings.Builder{}
	// arrayForQuery := ""
	queryBuilder.Write([]byte(queryCreateRecordImageListFirstPart))
	for i := range imageInfo {
		isMedicString := strconv.FormatBool(isMedic)
		queryBuilder.Write([]byte("("))
		queryBuilder.Write([]byte(isMedicString))
		queryBuilder.Write([]byte(","))
		queryBuilder.Write([]byte(cast.IntToStr(recordId)))
		queryBuilder.Write([]byte(","))
		queryBuilder.Write([]byte("'" + imageInfo[i] + "'"))
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
		return []uint64{}, err
	}
	response := []uint64{}
	for i := range resp {
		response = append(response,	cast.ToUint64(resp[i][0]))
	}
	
	return response, nil
}

func (rr *dbrecordrepository) CreateMedicRecord(diaryId uint64, record domain.MedicRecordCreateRequest) (domain.MedicRecordCreateResponse, error) {
	query := queryCreateMedicRecord
	resp, err := rr.dbm.Query(query,
		diaryId,
		time.Now().Format("2006.01.02 15:04:05"),
		record.BasicInfo.Title,
		record.BasicInfo.Treatment,
		record.BasicInfo.Recommendations,
		record.BasicInfo.Details,
		// record.Diarisation
	)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return domain.MedicRecordCreateResponse{}, err
	}
	response := domain.MedicRecordCreateResponse{
		Id:           cast.ToUint64(resp[0][0]),
		DiaryId:      cast.ToUint64(resp[0][1]),
		CreatingDate: cast.TimeToStr(cast.ToTime(resp[0][2]), true),
		BasicInfo: domain.MedicRecordBasicInfo{
			Title: cast.ToString(resp[0][3]),
			Treatment: cast.ToString(resp[0][4]),
			Recommendations: cast.ToString(resp[0][5]),
			Details: cast.ToString(resp[0][6]),
		},
		// Diarisation: cast.ToString(resp[0][7]),
		ImageList: nil,
	}
	
	if err != nil {
		log.Error(err)
		return domain.MedicRecordCreateResponse{}, err
	}
	return response, nil
}

func (rr *dbrecordrepository) CreateImageTags(imageIds []uint64, tags [][]string)  ([]uint64, [][]string, error) {
	if len(imageIds) == 0 {
		return []uint64{}, [][]string{}, nil
	}
	queryBuilder := strings.Builder{}
	// arrayForQuery := ""
	queryBuilder.Write([]byte(queryCreateImageTagListFirstPart))
	for i := range imageIds {
		for j := range(tags[i]) {
			queryBuilder.Write([]byte("("))
			queryBuilder.Write([]byte(cast.IntToStr(imageIds[i])))
			queryBuilder.Write([]byte(","))
			queryBuilder.Write([]byte(tags[i][j]))
			// if(j != len(tags[i]) - 1){
			// 	queryBuilder.Write([]byte(","))
			// }
			queryBuilder.Write([]byte(")"))
			if i != len(imageIds) || j != len(tags[i]) - 1 {
				queryBuilder.Write([]byte(","))
			}
		}
	}
	queryBuilder.Write([]byte(queryCreateImageTagListSecondPart))
	query := queryBuilder.String()
	_, err := rr.dbm.Query(queryBuilder.String())
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return []uint64{}, [][]string{},err
	}

	// response := []uint64{}
	// for i := range resp {
	// 	response = append(response,	cast.ToUint64(resp[i][0]))
	// }

	// TODO parse response from query
	return imageIds, tags, nil
}


func (dr *dbrecordrepository) GetRecordTextInfo(isMedic bool, recordId uint64,) (uint64, uint64, string, domain.MedicRecordBasicInfo, error) {
	var resp []database.DBbyterow
	var err error
	query := queryGetMedicRecordInfo
	resp, err = dr.dbm.Query(query, recordId)

	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return 0, 0, "", domain.MedicRecordBasicInfo{}, domain.Err.ErrObj.InternalServer
	}

	if len(resp) == 0 {
		log.Warn(cast.GetCurrentFuncName())
		log.Error(domain.Err.ErrObj.SmallDb)
		return 0, 0, "", domain.MedicRecordBasicInfo{}, domain.Err.ErrObj.SmallDb
	}
	return cast.ToUint64(resp[0][0]), cast.ToUint64(resp[0][1]), cast.TimeToStr(cast.ToTime(resp[0][2]), true), domain.MedicRecordBasicInfo{
		Title: cast.ToString(resp[0][3]),
		Treatment: cast.ToString(resp[0][4]),
		Recommendations: cast.ToString(resp[0][5]),
		Details: cast.ToString(resp[0][6]),
		},  nil
}


func (cr *dbrecordrepository) GetRecordImageNames(isMedic bool, recordId uint64) ([]string, error) {
	var resp []database.DBbyterow
	var err error
	query := ""

	query = queryGetRecordImageList
	resp, err = cr.dbm.Query(query, isMedic, recordId)

	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return nil, domain.Err.ErrObj.InternalServer
	}

	imageNames := make([]string, 0)
	for i := range resp {
		imageNames = append(imageNames, cast.ToString(resp[i][0]))
	}

	return imageNames, nil
}


func (er *dbrecordrepository) UpdateMedicRecordText(recordId uint64, medicRecordBasicInfo domain.MedicRecordBasicInfo) (domain.MedicRecordUpdateTextResponse, error) {
	query := queryUpdateTextMedicRecord
	resp, err := er.dbm.Query(query,
		medicRecordBasicInfo.Title,
		medicRecordBasicInfo.Treatment,
		medicRecordBasicInfo.Recommendations,
		medicRecordBasicInfo.Details,
		recordId)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return domain.MedicRecordUpdateTextResponse{}, err
	}

	return domain.MedicRecordUpdateTextResponse{
		Id:           cast.ToUint64(resp[0][0]),
		DiaryId:      cast.ToUint64(resp[0][1]),
		CreatingDate: cast.TimeToStr(cast.ToTime(resp[0][2]), true),
		BasicInfo: domain.MedicRecordBasicInfo{
			Title:       		cast.ToString(resp[0][3]),
			Treatment:       	cast.ToString(resp[0][4]),
			Recommendations:    cast.ToString(resp[0][5]),
			Details:			cast.ToString(resp[0][6]),
		},
	}, nil
}

// func (er *dbdiaryrepository) GetRecordImageLists(recordId uint64) ([]domain.ImageInfo, error) {
// 	query := queryGetImageList
// 	resp, err := er.dbm.Query(query, recordId)
// 	if err != nil {
// 		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
// 		log.Error(err)
// 		return []domain.ImageInfo{}, err
// 	}
// 	response := []domain.ImageInfo{}
// 	for i := range resp {
// 		response = append(response, domain.ImageInfo{
// 			Id:       cast.ToUint64(resp[i][0]),
// 			RecordId: cast.ToUint64(resp[i][1]),
// 			Name:     cast.ToString(resp[i][2]),
// 			Area:     cast.ToFloat64(resp[i][3]),
// 		})
// 	}
// 	return response, nil
// }

func (dr *dbrecordrepository) DeleteRecord(isMedic bool, recordId uint64,) (error) {
	var resp []database.DBbyterow
	var err error
	query := queryDeleteMedicRecord
	resp, err = dr.dbm.Query(query, recordId)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return domain.Err.ErrObj.InternalServer
	}
	if len(resp) == 0 {
		log.Warn(cast.GetCurrentFuncName())
		log.Error(domain.Err.ErrObj.SmallDb)
		return domain.Err.ErrObj.SmallDb
	}
	return nil
}

func (er *dbrecordrepository) DeleteRecordImage(isMedic bool, recordId uint64) (domain.RecordUpdateImageResponse, error) {
	query := queryDeleteImageMedicRecord
	resp, err := er.dbm.Query(query,
		isMedic,
		recordId)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return domain.RecordUpdateImageResponse{}, err
	}
	getRecordQuery := queryGetBasicUpdateImageMedicRecord
	getResp, getErr := er.dbm.Query(getRecordQuery,
		recordId)
	if getErr != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + getRecordQuery)
		log.Error(getErr)
		return domain.RecordUpdateImageResponse{}, getErr
	}
	recordUpdateImageResponse := domain.RecordUpdateImageResponse{
		DiaryId: cast.ToUint64(getResp[0][0]),
		Id: cast.ToUint64(getResp[0][1]),
		CreatingDate: cast.TimeToStr(cast.ToTime(getResp[0][2]), true),
	}
	// deletedImages := make([]string, 0)
	for i := range resp {
		recordUpdateImageResponse.Images = append(recordUpdateImageResponse.Images, domain.RecordImageInfo{	
			ImageName: cast.ToString(resp[i][0]),
			Tags:  nil,
		})
	}
	return recordUpdateImageResponse, nil
}

func (dr *dbrecordrepository) GetMedicIdFromDiary(diaryId uint64,) (uint64, error) {
	var resp []database.DBbyterow
	var err error
	query := queryGetMedicIdFromDiary
	resp, err = dr.dbm.Query(query, diaryId)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return 0, domain.Err.ErrObj.InternalServer
	}

	if len(resp) == 0 {
		log.Warn(cast.GetCurrentFuncName())
		log.Error(domain.Err.ErrObj.SmallDb)
		return 0, domain.Err.ErrObj.SmallDb
	}

	return cast.ToUint64(resp[0][0]), nil
}


func (dr *dbrecordrepository) GetMedicIdFromDiaryOfRecord(recordId uint64,) (uint64, error) {
	var resp []database.DBbyterow
	var err error
	query := queryGetMedicIdFromDiaryOfRecord
	resp, err = dr.dbm.Query(query, recordId)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return 0, domain.Err.ErrObj.InternalServer
	}

	if len(resp) == 0 {
		log.Warn(cast.GetCurrentFuncName())
		log.Error(domain.Err.ErrObj.SmallDb)
		return 0, domain.Err.ErrObj.SmallDb
	}

	return cast.ToUint64(resp[0][0]), nil
}

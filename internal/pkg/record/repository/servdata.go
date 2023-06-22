package recordrepository

import (
	"fmt"
	"hesh/internal/pkg/database"
	"hesh/internal/pkg/domain"

	// "strconv"
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
	query := fmt.Sprintf(queryMedicExist, cr.dbm.EncryptionKey)
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

func (cr *dbrecordrepository) UserExist(userId uint64) (bool, error) {
	var resp []database.DBbyterow
	var err error
	query := fmt.Sprintf(queryUserExist, cr.dbm.EncryptionKey, cr.dbm.EncryptionKey)
	resp, err = cr.dbm.Query(query, userId)

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

// MEDIC
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

func (er *dbrecordrepository) CreateRecordImageLists(isMedic bool, recordId uint64, imageInfo []string) ([]uint64, error) {
	if len(imageInfo) == 0 {
		return []uint64{}, nil
	}
	queryBuilder := strings.Builder{}
	// arrayForQuery := ""
	if isMedic {
		queryBuilder.Write([]byte(queryCreateMedicRecordImageListFirstPart))
	} else {
		queryBuilder.Write([]byte(queryCreatePatientRecordImageListFirstPart))
	}
	for i := range imageInfo {
		// isMedicString := strconv.FormatBool(isMedic)
		queryBuilder.Write([]byte("("))
		// queryBuilder.Write([]byte(isMedicString))
		// queryBuilder.Write([]byte(","))
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
		response = append(response, cast.ToUint64(resp[i][0]))
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
			Title:           cast.ToString(resp[0][3]),
			Treatment:       cast.ToString(resp[0][4]),
			Recommendations: cast.ToString(resp[0][5]),
			Details:         cast.ToString(resp[0][6]),
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

func (rr *dbrecordrepository) CreateImageTags(imageIds []uint64, tags [][]string) ([]uint64, [][]string, error) {
	if len(imageIds) == 0 {
		return []uint64{}, [][]string{}, nil
	}
	queryBuilder := strings.Builder{}
	// arrayForQuery := ""
	queryBuilder.Write([]byte(queryCreateImageTagListFirstPart))
	for i := range imageIds {
		for j := range tags[i] {
			queryBuilder.Write([]byte("("))
			queryBuilder.Write([]byte(cast.IntToStr(imageIds[i])))
			queryBuilder.Write([]byte(","))
			queryBuilder.Write([]byte(tags[i][j]))
			// if(j != len(tags[i]) - 1){
			// 	queryBuilder.Write([]byte(","))
			// }
			queryBuilder.Write([]byte(")"))
			if i != len(imageIds) || j != len(tags[i])-1 {
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
		return []uint64{}, [][]string{}, err
	}

	// response := []uint64{}
	// for i := range resp {
	// 	response = append(response,	cast.ToUint64(resp[i][0]))
	// }

	// TODO parse response from query
	return imageIds, tags, nil
}

func (dr *dbrecordrepository) GetRecordTextInfo(recordId uint64) (uint64, uint64, string, domain.MedicRecordBasicInfo, error) {
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
		Title:           cast.ToString(resp[0][3]),
		Treatment:       cast.ToString(resp[0][4]),
		Recommendations: cast.ToString(resp[0][5]),
		Details:         cast.ToString(resp[0][6]),
	}, nil
}

func (dr *dbrecordrepository) GetMedicRecordDiarisations(medicRecordId uint64) (domain.GetDiarisationsResponse, error) {
	var resp []database.DBbyterow
	query := queryGetMedicRecordDiarisationList
	resp, err := dr.dbm.Query(query, medicRecordId)

	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return domain.GetDiarisationsResponse{}, domain.Err.ErrObj.InternalServer
	}

	if len(resp) == 0 {
		return domain.GetDiarisationsResponse{MedicRecordId: medicRecordId}, nil
	}
	response := domain.GetDiarisationsResponse{}
	response.MedicRecordId = cast.ToUint64(resp[0][1])
	response.DiarisationList = make([]domain.DiarisationInListResponse, 0)
	for i := range resp {
		response.DiarisationList = append(response.DiarisationList, domain.DiarisationInListResponse{
			Id: cast.ToUint64(resp[i][0]),
			// [i] [1] used before
			CreatingDate: cast.TimeToStr(cast.ToTime(resp[i][2]), true),
			DiarisationInfo: domain.DiarisationInfo{
				Diarisation: cast.ToString(resp[i][3]),
				Filename:    cast.ToString(resp[i][4]),
				IsComplete:  cast.ToBool(resp[i][5]),
			},
		})
	}
	return response, nil
}

func (cr *dbrecordrepository) GetRecordImageNames(isMedic bool, recordId uint64) ([]string, error) {
	var resp []database.DBbyterow
	var err error
	query := ""

	if isMedic {
		query = queryGetMedicRecordImageList
	} else {
		query = queryGetPatientRecordImageList
	}
	resp, err = cr.dbm.Query(query, recordId)

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
			Title:           cast.ToString(resp[0][3]),
			Treatment:       cast.ToString(resp[0][4]),
			Recommendations: cast.ToString(resp[0][5]),
			Details:         cast.ToString(resp[0][6]),
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

func (dr *dbrecordrepository) DeleteRecord(isMedic bool, recordId uint64) error {
	var resp []database.DBbyterow
	var err error
	var query string
	if isMedic {
		query = queryDeleteMedicRecord
	} else {
		query = queryDeletePatientRecord
	}
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
	var query string
	if isMedic {
		query = queryDeleteImageMedicRecord
	} else {
		query = queryDeleteImagePatientRecord
	}
	resp, err := er.dbm.Query(query,
		recordId)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return domain.RecordUpdateImageResponse{}, err
	}
	var getRecordQuery string
	if isMedic {
		getRecordQuery = queryGetBasicUpdateImageMedicRecord
	} else {
		getRecordQuery = queryGetBasicUpdateImagePatientRecord
	}
	getResp, getErr := er.dbm.Query(getRecordQuery,
		recordId)
	if getErr != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + getRecordQuery)
		log.Error(getErr)
		return domain.RecordUpdateImageResponse{}, getErr
	}
	recordUpdateImageResponse := domain.RecordUpdateImageResponse{
		DiaryId:      cast.ToUint64(getResp[0][0]),
		Id:           cast.ToUint64(getResp[0][1]),
		CreatingDate: cast.TimeToStr(cast.ToTime(getResp[0][2]), true),
	}
	// deletedImages := make([]string, 0)
	for i := range resp {
		recordUpdateImageResponse.Images = append(recordUpdateImageResponse.Images, domain.RecordImageInfo{
			ImageName: cast.ToString(resp[i][0]),
			Tags:      nil,
		})
	}
	return recordUpdateImageResponse, nil
}

func (dr *dbrecordrepository) GetMedicIdFromDiary(diaryId uint64) (uint64, error) {
	var resp []database.DBbyterow
	var err error
	query := fmt.Sprintf(queryGetMedicIdFromDiary, dr.dbm.EncryptionKey)
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

func (dr *dbrecordrepository) GetMedicAndPatientIdsFromDiaryOfRecord(recordId uint64) (uint64, uint64, error) {
	var resp []database.DBbyterow
	var err error
	query := fmt.Sprintf(queryGetMedicAndPatientIdsFromDiaryOfRecord, dr.dbm.EncryptionKey, dr.dbm.EncryptionKey)
	resp, err = dr.dbm.Query(query, recordId)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return 0, 0, domain.Err.ErrObj.InternalServer
	}

	if len(resp) == 0 {
		log.Warn(cast.GetCurrentFuncName())
		log.Error(domain.Err.ErrObj.SmallDb)
		return 0, 0, domain.Err.ErrObj.SmallDb
	}
	patientId := uint64(0)
	if resp[0][1] != nil {
		patientId = cast.ToUint64(resp[0][1])
	}
	return cast.ToUint64(resp[0][0]), patientId, nil
}

// PATIENT
func (rr *dbrecordrepository) CreatePatientRecord(diaryId uint64, record domain.PatientRecordCreateRequest) (domain.PatientRecordCreateResponse, error) {
	query := queryCreatePatientRecord
	resp, err := rr.dbm.Query(query,
		diaryId,
		time.Now().Format("2006.01.02 15:04:05"),
		record.BasicInfo.Title,
		record.BasicInfo.Complaints,
		record.BasicInfo.Treatment,
		record.BasicInfo.Details,
		record.BasicInfo.Feelings,
	)
	fmt.Printf("diaryId: %v\n", diaryId)
	fmt.Printf("record.BasicInfo.Title: %v\n", record.BasicInfo.Title)
	fmt.Printf("record.BasicInfo.Complaints: %v\n", record.BasicInfo.Complaints)
	fmt.Printf("record.BasicInfo.Treatment: %v\n", record.BasicInfo.Treatment)
	fmt.Printf("record.BasicInfo.Details: %v\n", record.BasicInfo.Details)
	fmt.Printf("record.BasicInfo.Feelings: %v\n", record.BasicInfo.Feelings)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return domain.PatientRecordCreateResponse{}, err
	}
	response := domain.PatientRecordCreateResponse{
		Id:           cast.ToUint64(resp[0][0]),
		DiaryId:      cast.ToUint64(resp[0][1]),
		CreatingDate: cast.TimeToStr(cast.ToTime(resp[0][2]), true),
		BasicInfo: domain.PatientRecordBasicInfo{
			Title:      cast.ToString(resp[0][3]),
			Complaints: cast.ToString(resp[0][4]),
			Treatment:  cast.ToString(resp[0][5]),
			Details:    cast.ToString(resp[0][6]),
			Feelings:   cast.ToUint64(resp[0][7]),
		},
		ImageList: nil,
	}

	if err != nil {
		log.Error(err)
		return domain.PatientRecordCreateResponse{}, err
	}
	return response, nil
}

func (dr *dbrecordrepository) GetPatientRecordTextInfo(recordId uint64) (uint64, uint64, string, domain.PatientRecordBasicInfo, error) {
	var resp []database.DBbyterow
	var err error
	query := queryGetPatientRecordInfo
	resp, err = dr.dbm.Query(query, recordId)

	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return 0, 0, "", domain.PatientRecordBasicInfo{}, domain.Err.ErrObj.InternalServer
	}

	if len(resp) == 0 {
		log.Warn(cast.GetCurrentFuncName())
		log.Error(domain.Err.ErrObj.SmallDb)
		return 0, 0, "", domain.PatientRecordBasicInfo{}, domain.Err.ErrObj.SmallDb
	}

	return cast.ToUint64(resp[0][0]), cast.ToUint64(resp[0][1]), cast.TimeToStr(cast.ToTime(resp[0][2]), true), domain.PatientRecordBasicInfo{
		Title:      cast.ToString(resp[0][3]),
		Complaints: cast.ToString(resp[0][4]),
		Treatment:  cast.ToString(resp[0][5]),
		Details:    cast.ToString(resp[0][6]),
		Feelings:   cast.ToUint64(resp[0][7]),
	}, nil
}

func (er *dbrecordrepository) UpdatePatientRecordText(recordId uint64, patientRecordBasicInfo domain.PatientRecordBasicInfo) (domain.PatientRecordUpdateTextResponse, error) {
	query := queryUpdateTextPatientRecord
	resp, err := er.dbm.Query(query,
		patientRecordBasicInfo.Title,
		patientRecordBasicInfo.Complaints,
		patientRecordBasicInfo.Treatment,
		patientRecordBasicInfo.Details,
		patientRecordBasicInfo.Feelings,
		recordId)
	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return domain.PatientRecordUpdateTextResponse{}, err
	}

	return domain.PatientRecordUpdateTextResponse{
		Id:           cast.ToUint64(resp[0][0]),
		DiaryId:      cast.ToUint64(resp[0][1]),
		CreatingDate: cast.TimeToStr(cast.ToTime(resp[0][2]), true),
		BasicInfo: domain.PatientRecordBasicInfo{
			Title:      cast.ToString(resp[0][3]),
			Complaints: cast.ToString(resp[0][4]),
			Treatment:  cast.ToString(resp[0][5]),
			Details:    cast.ToString(resp[0][6]),
			Feelings:   cast.ToUint64(resp[0][7]),
		},
	}, nil
}

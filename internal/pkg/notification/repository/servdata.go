package notificationrepository

import (
	"fmt"
	"hesh/internal/pkg/database"
	"hesh/internal/pkg/domain"
	"hesh/internal/pkg/utils/cast"
	"hesh/internal/pkg/utils/log"
	// "hesh/internal/pkg/utils/cast"
	// "hesh/internal/pkg/utils/log"
)

type dbNotificationRepository struct {
	dbm *database.DBManager
}

func InitNotificationRep(manager *database.DBManager) domain.NotificationRepository {
	return &dbNotificationRepository{
		dbm: manager,
	}
}

func (r *dbNotificationRepository) PatientRecordRemind() ([]domain.PatientNotificationInfo, error) {
	var resp []database.DBbyterow
	var err error
	query := fmt.Sprintf(getPatientsReminders, r.dbm.EncryptionKey,"%", r.dbm.EncryptionKey)

	resp, err = r.dbm.Query(query)

	if err != nil {
		log.Warn("{" + cast.GetCurrentFuncName() + "} in query: " + query)
		log.Error(err)
		return nil, domain.Err.ErrObj.InternalServer
	}
	// if len(resp) == 0 {
	// 	return domain.DiaryListResponse{}, domain.Err.ErrObj.SmallDb
	// }
	notifications := make([]domain.PatientNotificationInfo, 0)
	var patientId uint64
	if resp != nil {
		patientId = cast.ToUint64(resp[0][0])
	}
	for i := range resp {
		// currentPatientId := cast.ToUint64(resp[i][0])
		// if currentPatientId != patientId {
		// 	patientId = currentPatientId
		// }
		var patientNotificationInfo domain.PatientNotificationInfo
		if (cast.ToUint64(resp[i][0]) != patientId || i == 0) {
			if i != 0 {
				fmt.Printf("patientNotificationInfo.PatientId: %v\n", patientNotificationInfo.PatientId)
				fmt.Printf("patientNotificationInfo: %v\n", patientNotificationInfo)

				notifications = append(notifications, patientNotificationInfo)
			}
			patientId = cast.ToUint64(resp[i][0])
			fmt.Printf("patientId: %v\n", patientId)
			patientNotificationInfo = domain.PatientNotificationInfo{
				PatientId: patientId,
			}
			fmt.Printf("After init patientNotificationInfo.PatientId: %v\n", patientNotificationInfo.PatientId)
		}
		fmt.Printf("In main loop over resp patientNotificationInfo.PatientId: %v\n", patientNotificationInfo.PatientId)
		patientNotificationInfo.NotificationInfo = append(patientNotificationInfo.NotificationInfo, domain.DiaryNotificationsInfo{
			DiaryName: cast.ToString(resp[i][1]),
			Frequency: cast.ToUint64(resp[i][3]),
		})
		
		if i == len(resp) - 1 {
			fmt.Printf("patientNotificationInfo.PatientId: %v\n", patientNotificationInfo.PatientId)
			fmt.Printf("patientNotificationInfo: %v\n", patientNotificationInfo)
			notifications = append(notifications, patientNotificationInfo)
		}
		println()
		println()
		println()
	}
	return notifications, nil
}
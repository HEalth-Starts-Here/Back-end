package domain

const (
	MaxVkNotificationLength = 254
)

type DiaryNotificationsInfo struct {
	DiaryName	 string
	Frequency	 uint64
}

type PatientNotificationInfo struct {
	PatientId				uint64
	NotificationInfo 	[]DiaryNotificationsInfo
}

type NotificationRepository interface {
	PatientRecordRemind() ([]PatientNotificationInfo, error)
}

type NotificationUsecase interface {
	PatientRecordRemind() ()
}

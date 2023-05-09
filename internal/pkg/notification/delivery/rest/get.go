package notificationdelivery

import (

)

func (handler *NotificationHandler) PatientRecordRemind() {

	handler.NotificationUsecase.PatientRecordRemind()
}

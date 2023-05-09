package notificationdelivery

import (
	// "github.com/gorilla/mux"
	"hesh/internal/pkg/domain"
)

type NotificationHandler struct {
	NotificationUsecase domain.NotificationUsecase

}

func SetNotificationHandlers(u domain.NotificationUsecase) {
	// handler := &NotificationHandler{
	// 	NotificationUsecase: u,
	// }

	u.PatientRecordRemind()
	// router.HandleFunc(postRatingUrl, handler.PostRating).Methods("POST", "OPTIONS")
}
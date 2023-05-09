package setter

import (
	"hesh/internal/pkg/database"

	// "hesh/internal/pkg/utils/config"
	// "hesh/internal/pkg/utils/log"

	// autmcs "hesh/internal/pkg/authorization/delivery/grpc"
	// autdelivery "hesh/internal/pkg/authorization/delivery/rest"

	diarydelivery "hesh/internal/pkg/diary/delivery"
	diaryrepository "hesh/internal/pkg/diary/repository"
	diaryusecase "hesh/internal/pkg/diary/usecase"

	mlservicesdelivery "hesh/internal/pkg/mlservices/delivery"
	mlservicesrepository "hesh/internal/pkg/mlservices/repository"
	mlservicesusecase "hesh/internal/pkg/mlservices/usecase"

	userdelivery "hesh/internal/pkg/user/delivery"
	userrepository "hesh/internal/pkg/user/repository"
	userusecase "hesh/internal/pkg/user/usecase"

	recorddelivery "hesh/internal/pkg/record/delivery"
	recordrepository "hesh/internal/pkg/record/repository"
	recordusecase "hesh/internal/pkg/record/usecase"

	commentdelivery "hesh/internal/pkg/comment/delivery"
	commentrepository "hesh/internal/pkg/comment/repository"
	commentusecase "hesh/internal/pkg/comment/usecase"

	notedelivery "hesh/internal/pkg/note/delivery"
	noterepository "hesh/internal/pkg/note/repository"
	noteusecase "hesh/internal/pkg/note/usecase"

	searchdelivery "hesh/internal/pkg/search/delivery"
	searchrepository "hesh/internal/pkg/search/repository"
	searchusecase "hesh/internal/pkg/search/usecase"

	// notificationdelivery "hesh/internal/pkg/notification/delivery/rest"
	// notificationrepository "hesh/internal/pkg/notification/repository"
	// notificationusecase "hesh/internal/pkg/notification/usecase"

	// notificationmcs "hesh/internal/pkg/notification/delivery/grpc"

	// "hesh/internal/pkg/notification/delivery/rest"

	"github.com/gorilla/mux"

	// "github.com/rs/zerolog/log"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/credentials/insecure"
)

type Data struct {
	Db  *database.DBManager
	Api *mux.Router
}

type Services struct {
	Diary      		Data
	MLServices 		Data
	User       		Data
	Record     		Data
	Comment    		Data
	Note       		Data
	Search     		Data

	Notification	Data
}

// func setNotificationMcs()  {
// 	// notificationconn, err := grpc.Dial(":"+config.DevConfigStore.Mcs.Auth.Port, grpc.WithInsecure())
// 	// if err != nil {
// 	// 	log.Warn("{setAutMcs} mcs Dial")
// 	// }

// 	// return notificationmcs.NewNotificationClient(notificationconn)
// }

func SetHandlers(svs Services) {
	diaryRep := diaryrepository.InitDiaryRep(svs.Diary.Db)
	mlservicesRep := mlservicesrepository.InitMLServicesRep(svs.MLServices.Db)
	userRep := userrepository.InitUserRep(svs.User.Db)
	recordRep := recordrepository.InitRecordRep(svs.Record.Db)
	commentRep := commentrepository.InitCommentRep(svs.Comment.Db)
	noteRep := noterepository.InitNoteRep(svs.Note.Db)
	searchRep := searchrepository.InitSearchRep(svs.Search.Db)
	// notificationRep := notificationrepository.InitNotificationRep(svs.Notification.Db)

	diaryUsc := diaryusecase.InitDiaryUsc(diaryRep)
	mlservicesUsc := mlservicesusecase.InitMLServicesUsc(mlservicesRep)
	userUsc := userusecase.InitUserUsc(userRep)
	recordUsc := recordusecase.InitRecordUsc(recordRep)
	commentUsc := commentusecase.InitCommentUsc(commentRep)
	noteUsc := noteusecase.InitNoteUsc(noteRep)
	searchUsc := searchusecase.InitSearchUsc(searchRep)
	// notificationUsc := notificationusecase.InitNotificationUsc(notificationRep)

	diarydelivery.SetDiaryHandlers(svs.Diary.Api, diaryUsc)
	mlservicesdelivery.SetMLServicesHandlers(svs.MLServices.Api, mlservicesUsc)
	userdelivery.SetUserHandlers(svs.User.Api, userUsc)
	recorddelivery.SetRecordHandlers(svs.Record.Api, recordUsc)
	commentdelivery.SetCommentHandlers(svs.Comment.Api, commentUsc)
	notedelivery.SetNoteHandlers(svs.Note.Api, noteUsc)
	searchdelivery.SetSearchHandlers(svs.Search.Api, searchUsc)
	
	// notificationdelivery.SetNotificationHandlers(notificationUsc)

}

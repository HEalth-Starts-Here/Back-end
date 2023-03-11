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

	// eventdelivery "hesh/internal/pkg/event/delivery"
	// eventrepository "hesh/internal/pkg/event/repository"
	// eventusecase "hesh/internal/pkg/event/usecase"

	"github.com/gorilla/mux"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/credentials/insecure"
)

type Data struct {
	Db  *database.DBManager
	Api *mux.Router
}

type Services struct {
	// Act Data
	// Mov Data
	// User Data
	// Col Data
	// Gen Data
	// Ann Data
	// Ser Data
	Diary Data

	// Rat Data
	// Aut Data
	// Com Data
}

// func setAutMcs() autmcs.AutherClient {
// 	autconn, err := grpc.Dial(":"+config.DevConfigStore.Mcs.Auth.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Warn("{setAutMcs} mcs Dial")
// 	}

// 	return autmcs.NewAutherClient(autconn)
// }

func SetHandlers(svs Services) {
	// userRep := usrrepository.InitUsrRep(svs.User.Db)
	diaryRep := diaryrepository.InitDiaryRep(svs.Diary.Db)

	// userUsc := usrusecase.InitUsrUsc(userRep)
	diaryUsc := diaryusecase.InitDiaryUsc(diaryRep)

	// usrdelivery.SetUsrHandlers(svs.User.Api, userUsc)
	diarydelivery.SetDiaryHandlers(svs.Diary.Api, diaryUsc)

	// autdelivery.SetAutHandlers(svs.Aut.Api, setAutMcs())
}

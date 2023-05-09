package mcsnotification

// import (
// 	"hesh/internal/pkg/notification/repository"
// 	"hesh/internal/pkg/notification/usecase"
// 	"hesh/internal/pkg/database"
// 	"hesh/internal/pkg/utils/config"
// 	"hesh/internal/pkg/utils/log"

// 	"net"
// )

// func RunServer() {
// 	db := database.InitDatabase()
// 	db.Connect()
// 	defer db.Disconnect()

// 	autRep := repository.InitRatRep(db)
// 	autUsc := usecase.InitRatUsc(autRep)

// 	// s := grpc.NewServer()

// 	// proto.RegisterPosterServer(s, autUsc)

// 	// l, err := net.Listen(config.DevConfigStore.Mcs.Rating.ConnType, ":"+config.DevConfigStore.Mcs.Rating.Port)
// 	// if err != nil {
// 	// 	log.Warn("{RunServer} mcs rtng")
// 	// 	log.Error(err)
// 	// }

// 	// s.Serve(l)
// }

package main

import (
	"github.com/gorilla/mux"
	"go_practice/9_clean_arch_db/config"
	"go_practice/9_clean_arch_db/internal/mwares"
	productHandler "go_practice/9_clean_arch_db/internal/product/delivery"
	grpcProduct "go_practice/9_clean_arch_db/internal/product/delivery/grpc"
	productUsecase "go_practice/9_clean_arch_db/internal/product/usecases"
	sessHandler "go_practice/9_clean_arch_db/internal/session/delivery"
	grpcSess "go_practice/9_clean_arch_db/internal/session/delivery/grpc"
	sessUsecase "go_practice/9_clean_arch_db/internal/session/usecases"
	userHandler "go_practice/9_clean_arch_db/internal/user/delivery"
	grpcUser "go_practice/9_clean_arch_db/internal/user/delivery/grpc"
	userUsecase "go_practice/9_clean_arch_db/internal/user/usecases"
	"go_practice/9_clean_arch_db/tools/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	confg, err := config.LoadConfig("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	err = logger.InitLogger(confg.GetLoggerDir(), confg.GetLogLevel())
	if err != nil {
		log.Fatal(err)
	}

	postgresConnection, err := confg.Postgres.GetPostgresDbConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer postgresConnection.Close()

	redisConnection, err := confg.Redis.GetRedisDbConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer redisConnection.Close()

	sessionGrpcConn, err := grpc.Dial(confg.GetAuthConnectionClientString(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer sessionGrpcConn.Close()
	sessServiceClient := grpcSess.NewSessionServiceClient(sessionGrpcConn)

	userGrpcConn, err := grpc.Dial(confg.GetUserConnectionClientString(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer userGrpcConn.Close()
	userServiceClient := grpcUser.NewUserServiceClient(userGrpcConn)

	productGrpcConn, err := grpc.Dial(confg.GetProductConnectionClientString(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer userGrpcConn.Close()
	productServiceClient := grpcProduct.NewProductServiceClient(productGrpcConn)

	productUse := productUsecase.NewProductUsecase(productServiceClient)
	userUse := userUsecase.NewUserUsecase(userServiceClient)
	sessUse := sessUsecase.NewSessionUsecase(sessServiceClient)
	sessHnd := sessHandler.NewSessionHandler(sessUse, userUse)
	userHnd := userHandler.NewUserHandler(userUse, sessUse)
	productHnd := productHandler.NewProductHandler(productUse)

	mux := mux.NewRouter()

	mwManager := mwares.NewMiddlewareManager(sessUse)
	mux.Use(mwManager.PanicCoverMiddleware)
	mux.Use(mwManager.AccessLogMiddleware)
	mux.Use(mwManager.CORS)

	userHnd.Configure(mux, mwManager)
	productHnd.Configure(mux, mwManager)
	sessHnd.Configure(mux, mwManager)

	srv := config.NewServer("8080", mux)
	log.Fatal(srv.Run())
}

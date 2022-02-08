package main

import (
	"github.com/gorilla/mux"
	"go_practice/9_clean_arch_db/config"
	"go_practice/9_clean_arch_db/internal/mwares"
	productHandler "go_practice/9_clean_arch_db/internal/product/delivery"
	productRepository "go_practice/9_clean_arch_db/internal/product/repository"
	productUsecase "go_practice/9_clean_arch_db/internal/product/usecases"
	sessHandler "go_practice/9_clean_arch_db/internal/session/delivery"
	sessRepository "go_practice/9_clean_arch_db/internal/session/repository"
	sessUsecase "go_practice/9_clean_arch_db/internal/session/usecases"
	userHandler "go_practice/9_clean_arch_db/internal/user/delivery"
	userRepository "go_practice/9_clean_arch_db/internal/user/repository"
	userUsecase "go_practice/9_clean_arch_db/internal/user/usecases"
	"go_practice/9_clean_arch_db/tools/logger"
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

	userRep := userRepository.NewUserPgRepository(postgresConnection)
	userUse := userUsecase.NewUserUsecase(userRep)

	productRep := productRepository.NewProductPgRepository(postgresConnection)
	productUse := productUsecase.NewProductUsecase(productRep)
	productHnd := productHandler.NewProductHandler(productUse)

	sessRep := sessRepository.NewSessionRdRepository(redisConnection)
	sessUse := sessUsecase.NewSessionUsecase(sessRep)
	sessHnd := sessHandler.NewSessionHandler(sessUse, userUse)
	userHnd := userHandler.NewUserHandler(userUse, sessUse)

	mux := mux.NewRouter()

	mwManager := mwares.NewMiddlewareManager(sessUse)
	mux.Use(mwManager.PanicCoverMiddleware)
	mux.Use(mwManager.AccessLogMiddleware)

	userHnd.Configure(mux, mwManager)
	productHnd.Configure(mux, mwManager)
	sessHnd.Configure(mux, mwManager)


	srv := config.NewServer("8080", mux)
	log.Fatal(srv.Run())
}


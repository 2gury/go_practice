package main

import (
	"github.com/asaskevich/govalidator"
	mux2 "github.com/gorilla/mux"
	"go_practice/9_clean_arch_db/config"
	"go_practice/9_clean_arch_db/internal/product/delivery"
	"go_practice/9_clean_arch_db/internal/product/repository"
	"go_practice/9_clean_arch_db/internal/product/usecases"
	"log"
)

func main() {
	confg, err := config.LoadConfig("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	dbConnection, err := confg.Database.GetPostgresDbConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConnection.Close()

	productRepository := repository.NewProductPgRepository(dbConnection)
	productUsecase := usecases.NewProductUsecase(productRepository)
	productHandler := delivery.NewProductHandler(productUsecase)

	mux := mux2.NewRouter()
	productHandler.Configure(mux)

	srv := config.NewServer("8080", mux)
	log.Fatal(srv.Run())
}

func init() {
	govalidator.CustomTypeTagMap.Set(
		"title",
		govalidator.CustomTypeValidator(func(i interface{}, o interface{}) bool {
			subject, ok := i.(string)
			if !ok {
				return false
			}
			if len(subject) > 64 {
				return false
			}
			return true
		}),
	)
}

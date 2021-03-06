package main

import (
	"context"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go_practice/11_clean_arch_mongo_db/config"
	"go_practice/11_clean_arch_mongo_db/internal/delivery"
	"go_practice/11_clean_arch_mongo_db/internal/repository"
	"go_practice/11_clean_arch_mongo_db/internal/usecase"
	"log"
	"time"
)

func getMongoCollection() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client())
	if err != nil {
		return nil, err
	}
	ctx, cancel = context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func main() {
	mux := mux.NewRouter()

	client, err := getMongoCollection()
	if err != nil {
		log.Fatal(err.Error())
	}

	coll := client.Database("golang").Collection("products")

	rep := repository.NewProductMongoRepository(coll)
	use := usecase.NewProductUsecase(rep)
	hnd := delivery.NewProductHandler(use)
	hnd.Configure(mux)

	srv := config.NewHttpServer("8080", mux)
	srv.Start()
}



package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/ruziba3vich/to-do-with-mongo/handler"
	"github.com/ruziba3vich/to-do-with-mongo/internal/service"
	"github.com/ruziba3vich/to-do-with-mongo/internal/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	handler := handler.New(
		service.New(
			redis.NewClient(&redis.Options{
				Addr: "localhost:6379",
			}),
			storage.New(client),
		),
	)

	router := gin.Default()

	router.POST("/create", handler.CreateTaskHandler)

	log.Fatal(router.Run(":7777"))
}

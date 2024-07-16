package main

import (
	"context"
	"log"
	"os"
	"time"

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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	handler := handler.New(
		service.New(
			storage.NewTasksStorage(
				client,
				log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
				redis.NewClient(&redis.Options{
					Addr: "localhost:6379",
				},
				),
				client.Database("todo").Collection("tasks"),
				nil,
			),
		),
		ctx,
	)

	router := gin.Default()

	router.POST("/create", handler.CreateTaskHandler)
	// router.GET("/get/all/tasks/by/user", handler.GetIncompleteSubTasksHandler)
	log.Fatal(router.Run(":7777"))
}

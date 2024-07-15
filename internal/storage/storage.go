package storage

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/ruziba3vich/to-do-with-mongo/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Storage struct {
	db      *mongo.Client
	logger  *log.Logger
	redisDb *redis.Client
}

func New(db *mongo.Client, logger *log.Logger, redisDb *redis.Client) *Storage {
	return &Storage{
		db:      db,
		logger:  logger,
		redisDb: redisDb,
	}
}

func (s *Storage) CreateTask(ctx context.Context, task *models.Task) (*models.Task, error) {
	task.Id = primitive.NewObjectID()
	task.CreatedAt = time.Now()
	collection := s.db.Database(task.GetDbName()).Collection(task.GetCollectionName())

	_, err := collection.InsertOne(ctx, task)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}
	stat := s.redisDb.Set(ctx, task.Id.String(), task, time.Hour*24)
	if err := stat.Err(); err != nil {
		s.logger.Println(err)
		return nil, err
	}
	return task, nil
}

func (s *Storage) GetTaskById(ctx context.Context, taskId string) (*models.Task, error) {
	var response models.Task
	result, err := s.redisDb.Get(ctx, taskId).Result()
	if err != nil {
		if err == redis.Nil {
			docId, err := primitive.ObjectIDFromHex(taskId)
			if err != nil {
				s.logger.Printf("could not convert the given ID %s\n", taskId)
				return nil, err
			}

			collection := s.db.Database(response.GetDbName()).Collection(response.GetCollectionName())

			filter := bson.M{"_id": docId}

			if err := collection.FindOne(ctx, filter).Decode(&response); err != nil {
				if err == mongo.ErrNoDocuments {
					s.logger.Println("no data found in the database :", err.Error())
				} else {
					s.logger.Println("error while decoding response :", err.Error())
				}
				return nil, err
			}
			s.redisDb.Set(ctx, taskId, response, time.Hour*24)
		} else {
			s.logger.Println("error while fetching object from redis :", err.Error())
			return nil, err
		}
	}
	if err := json.Unmarshal([]byte(result), &response); err != nil {
		s.logger.Println("error while marshaling result :", err.Error())
		return nil, err
	}
	return &response, err
}

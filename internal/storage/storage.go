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

type TasksStorage struct {
	db              *mongo.Client
	logger          *log.Logger
	redisDb         *redis.Client
	tasksCollection *mongo.Collection
	usersCollection *mongo.Collection
}

func NewTasksStorage(
	db *mongo.Client,
	logger *log.Logger,
	redisDb *redis.Client,
	tasksCollection *mongo.Collection,
	usersCollection *mongo.Collection,
) *TasksStorage {
	return &TasksStorage{
		db:              db,
		logger:          logger,
		redisDb:         redisDb,
		tasksCollection: tasksCollection,
		usersCollection: usersCollection,
	}
}

func (s *TasksStorage) CreateTask(ctx context.Context, task *models.Task) (*models.Task, error) {
	task.Id = primitive.NewObjectID()

	_, err := s.tasksCollection.InsertOne(ctx, task)
	if err != nil {
		s.logger.Println(err)
		return nil, err
	}

	hashedTask, err := json.Marshal(task)
	if err != nil {
		s.logger.Println("error while marshaling task :", err.Error())
		return nil, err
	}

	stat := s.redisDb.Set(ctx, task.Id.String(), hashedTask, time.Hour*24)
	if err := stat.Err(); err != nil {
		s.logger.Println(err)
		return nil, err
	}
	return task, nil
}

/// PRODONiK ---------------------------------- 😎😎😎 ---------------------------------- ///

func (s *TasksStorage) GetTaskById(ctx context.Context, taskId string) (*models.Task, error) {
	var response models.Task
	result, err := s.redisDb.Get(ctx, taskId).Result()
	if err != nil {
		if err == redis.Nil {
			docId, err := primitive.ObjectIDFromHex(taskId)
			if err != nil {
				s.logger.Printf("could not convert the given ID %s\n", taskId)
				return nil, err
			}

			filter := bson.M{"_id": docId}

			if err := s.tasksCollection.FindOne(ctx, filter).Decode(&response); err != nil {
				if err == mongo.ErrNoDocuments {
					s.logger.Println("no data found in the database :", err.Error())
				} else {
					s.logger.Println("error while decoding response :", err.Error())
				}
				return nil, err
			}

			hashedTask, err := json.Marshal(response)
			if err != nil {
				s.logger.Println("could not marshal the response data :", err.Error())
				return nil, err
			}

			if err := s.redisDb.Set(ctx, taskId, hashedTask, time.Hour*24).Err(); err != nil {
				s.logger.Println("could not cache the data :", err.Error())
				return nil, err
			}
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

func (s *TasksStorage) GetTaskByUser(ctx context.Context, user *models.User) (*models.RepeatedModelsResponse, error) {
	filter := bson.M{
		"assigned_to": bson.M{
			"$all": []string{user.Fullname, user.Email},
		},
	}

	cursor, err := s.tasksCollection.Find(ctx, filter)
	if err != nil {
		s.logger.Println("error while fetching data from db :", err.Error())
		return nil, err
	}
	defer cursor.Close(ctx)

	var response models.RepeatedModelsResponse

	for cursor.Next(ctx) {
		var responsecha models.Task
		if err := cursor.Decode(&responsecha); err != nil {
			s.logger.Println("error while decoding data")
			return nil, err
		}
		response.Tasks = append(response.Tasks, &responsecha)
	}

	if err := cursor.Err(); err != nil {
		s.logger.Println("error fetched from the cursor :", err.Error())
		return nil, err
	}
	return &response, nil
}

func (s *TasksStorage) UpdateTaskStatus(ctx context.Context, req *models.UpdateTaskStatusRequest) (*models.Task, error) {
	task, err := s.GetTaskById(ctx, req.TaskId)
	if err != nil {
		s.logger.Println("task not found :", err.Error())
		return nil, err
	}
	task.Status = models.Status(req.TaskId)
	hashedTask, err := json.Marshal(task)
	if err != nil {
		s.logger.Println("error while marshaling data :", err.Error())
		return nil, err
	}
	stat := s.redisDb.Set(ctx, req.TaskId, hashedTask, time.Hour*24)
	if err := stat.Err(); err != nil {
		s.logger.Println("error while caching the data in redis :", err.Error())
		return nil, err
	}
	return task, nil
}

func (s *TasksStorage) GetIncompleteTasksUntillDate(ctx context.Context, req *models.GetIncompleteTasksUntillDateRequest) (*models.RepeatedModelsResponse, error) {
	filter := bson.M{
		"sub_tasks": bson.M{
			"$elemMatch": bson.M{"status": models.Pending},
		},
	}

	cursor, err := s.tasksCollection.Find(ctx, filter)
	if err != nil {
		s.logger.Println("error while fetching data from collection :", err.Error())
		return nil, err
	}

	defer cursor.Close(ctx)

	var response models.RepeatedModelsResponse
	for cursor.Next(ctx) {
		var responsecha models.Task
		if err := cursor.Decode(&responsecha); err != nil {
			s.logger.Println("error while decoding data :", err.Error())
			return nil, err
		}
		response.Tasks = append(response.Tasks, &responsecha)
	}

	if err := cursor.Err(); err != nil {
		s.logger.Println("error fetched from the cursor :", err.Error())
		return nil, err
	}
	return &response, nil
}

package service

import (
	"context"

	"github.com/ruziba3vich/to-do-with-mongo/internal/models"
	"github.com/ruziba3vich/to-do-with-mongo/internal/storage"
)

type Service struct {
	storage *storage.TasksStorage
}

func New(storage *storage.TasksStorage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) CreateTask(ctx context.Context, task *models.Task) (*models.Task, error) {
	return s.storage.CreateTask(ctx, task)
}

func (s *Service) GetTaskById(ctx context.Context, taskId string) (*models.Task, error) {
	return s.storage.GetTaskById(ctx, taskId)
}

func (s *Service) UpdateTaskStatus(ctx context.Context, req *models.UpdateTaskStatusRequest) (*models.Task, error) {
	return s.storage.UpdateTaskStatus(ctx, req)
}

func (s *Service) GetIncompleteSubTasks(ctx context.Context, req *models.User) (*models.RepeatedModelsResponse, error) {
	return s.storage.GetIncompleteSubTasks(ctx, req)
}

func (s *Service) GetTasksUntilDate(ctx context.Context, req *models.GetTasksUntilDateRequest) (*models.RepeatedModelsResponse, error) {
	return s.storage.GetTasksUntilDate(ctx, req)
}

func (s *Service) UpdateSubTaskStatus(ctx context.Context, req *models.UpdateSubTaskStatusRequest) (*models.Task, error) {
	return s.storage.UpdateSubTaskStatus(ctx, req)
}

func (s *Service) AddNewSubTaskIntoTask(ctx context.Context, req *models.AddNewSubTaskIntoTaskRequest) (*models.Task, error) {
	return s.storage.AddNewSubTaskIntoTask(ctx, req)
}

func (s *Service) ChangeTaskUser(ctx context.Context, req *models.ChangeTaskUserRequest) (*models.Task, error) {
	return s.storage.ChangeTaskUser(ctx, req)
}

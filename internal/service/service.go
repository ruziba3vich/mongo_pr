package service

import (
	"github.com/ruziba3vich/to-do-with-mongo/internal/models"
	"github.com/ruziba3vich/to-do-with-mongo/internal/storage"
)

type Service struct {
	storage *storage.Storage
}

func New(storage *storage.Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) CreateTask(task *models.Task) (*models.Task, error) {
	return s.storage.CreateTask(task)
}

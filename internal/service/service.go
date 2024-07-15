package service

import (
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

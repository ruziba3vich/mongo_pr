package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Task struct {
		Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
		Title       string             `json:"title,omitempty"`
		Description string             `json:"description,omitempty"`
		Content     string             `json:"content,omitempty"`
		Completed   bool               `json:"completed,omitempty"`
		CreatedAt   time.Time          `json:"created_at,omitempty"`
	}
)

func (t *Task) GetDbName() string {
	return "todo"
}

func (t *Task) GetCollectionName() string {
	return "tasks"
}

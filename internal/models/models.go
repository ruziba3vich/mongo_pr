package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Status string

	SubTask struct {
		Title  string `json:"title,omitempty"`
		Status Status `json:"status,omitempty"`
	}

	Task struct {
		Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
		Title       *string            `json:"title,omitempty"`
		Description *string            `json:"description,omitempty"`
		Status      Status             `json:"status,omitempty"`
		AssignedTo  []string           `json:"assigned_to,omitempty"`
		DueDate     primitive.DateTime `json:"due_date,omitempty"`
		SubTasks    []*SubTask         `json:"sub_tasks,omitempty"`
	}

	User struct {
		Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
		Fullname string             `json:"full_name,omitempty"`
		Email    string             `json:"email,omitempty"`
	}

	UpdateTaskStatusRequest struct {
		TaskId     string `json:"update_task_status_request"`
		TaskStatus string `json:"task_status"`
	}

	// GetIncompleteSubTasksRequest struct {
	// 	Date primitive.DateTime
	// }

	RepeatedModelsResponse struct {
		Tasks []*Task
	}

	GetTasksUntilDateRequest struct {
		Date primitive.DateTime
	}

	UpdateSubTaskStatusRequest struct {
		TaskId       string `json:"task_id"`
		SubTaskIndex int    `json:"subtask_index"`
		NewStatus    Status `json:"new_status"`
	}

	AddNewSubTaskIntoTaskRequest struct {
		TaskId  string  `json:"task_id"`
		SubTask SubTask `json:"sub_task"`
	}

	ChangeTaskUserRequest struct {
		TaskId           string   `json:"task_id"`
		NewTaskOwnerData []string `json:"new_task_owner_data"`
	}
)

const (
	Completed  = Status("completed")
	Pending    = Status("pending")
	InProgress = Status("in-progress")
)

func (t *Task) GetDbName() string {
	return "todo"
}

func (t *Task) GetCollectionName() string {
	return "tasks"
}

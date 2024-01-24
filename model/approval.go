package model

import "time"

type Approval struct {
	Id        string    `json:"id"`
	TaskId    string    `json:"task_id"`
	Approval  bool      `json:"approval"`
	Feedback  string    `json:"feedback"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

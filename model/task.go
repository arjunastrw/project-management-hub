package model

import "time"

type Task struct {
	Id             string    `json:"id"`
	Name           string    `json:"name"`
	Status         string    `json:"status"`
	Approval       bool      `json:"approval"`
	ApprovalDate   time.Time `json:"approval_date"`
	Feedback       string    `json:"feedback"`
	PersonInCharge string    `json:"person_in_charge"`
	ProjectId      string    `json:"project_id"`
	Deadline       time.Time `json:"deadline"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at"`
}

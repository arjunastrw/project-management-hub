package model

import "time"

type Task struct {
	Id             string     `json:"id"`
	Name           string     `json:"name"`
	Status         string     `json:"status"`
	Approval       bool       `json:"approval"`
	ApprovalDate   *time.Time `json:"approval_date"`
	Feedback       string     `json:"feedback"`
	PersonInCharge string     `json:"person_in_charge"`
	ProjectId      string     `json:"project_id"`
	Deadline       string     `json:"deadline"`
	CreatedAt      time.Time  `json:"-"`
	UpdatedAt      time.Time  `json:"-"`
	DeletedAt      *time.Time `json:"-"`
}

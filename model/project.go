package model

import "time"

type Project struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	ManagerId string    `json:"manager_id"`
	Deadline  time.Time `json:"deadline"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Members   []User    `json:"members"`
	Tasks     []Task    `json:"tasks"`
}

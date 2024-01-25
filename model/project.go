package model

import "time"

type Project struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	ManagerId string     `json:"manager_id"`
	Deadline  string     `json:"deadline"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
	Members   []User     `json:"members"`
	Tasks     []Task     `json:"tasks"`
}

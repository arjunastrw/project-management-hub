package model

import "time"

type Report struct {
	Id         string     `json:"id"`
	User_id    string     `json:"user_id"`
	Report     string     `json:"report"`
	Task_id    string     `json:"task_id"`
	Created_at time.Time  `json:"-"`
	Updated_at time.Time  `json:"-"`
	DeletedAt  *time.Time `json:"-"`
}

// Struktur untuk laporan yang ditampilkan
type ShowReport struct {
	Date    time.Time
	Content Report
}

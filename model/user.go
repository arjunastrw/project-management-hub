<<<<<<< HEAD
package model

import "time"

type User struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Role      string     `json:"role"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Project   []Project  `json:"project"`
	Task      []Task     `json:"task"`
}
=======
package model

import "time"

type User struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Project   []Project `json:"project"`
	Task      []Task    `json:"task"`
}
>>>>>>> 4e0c22da47a4350c832b379f60379ec7e899024d

package repository

import (
	"database/sql"
	"log"
	"math"

	"enigma.com/projectmanagementhub/config"
	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/shared/shared_model"
)

type UserRepository interface {
	GetAll(page int, size int) ([]model.User, shared_model.Paging, error)
	GetById(id string) (model.User, error)
	GetByEmail(email string) (model.User, error)
	CreateUser(payload model.User) (model.User, error)
	Update(payload model.User) (model.User, error)
	Delete(id string) error
}

type userRepository struct {
	db *sql.DB
}

// Delete implements User.
func (u *userRepository) Delete(id string) error {
	//to do make query to delete user
	_, err := u.db.Query(config.DeleteUserById, id)
	if err != nil {
		log.Println("user_repository.Query", err.Error())
		return err
	}

	return nil
}

// GetAll implements User.
func (u *userRepository) GetAll(page int, size int) ([]model.User, shared_model.Paging, error) {
	var users []model.User
	offset := (page - 1) * size
	row, err := u.db.Query(config.GetAllUser, size, offset)
	if err != nil {
		log.Println("user_repository.QueryRow", err.Error())
		return nil, shared_model.Paging{}, err
	}

	for row.Next() {
		user := model.User{}
		//updated_at cannot be nil
		err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			log.Println("userRepository.Rows.Next", err.Error())
			return nil, shared_model.Paging{}, err
		}

		users = append(users, user)
	}

	totalRows := 0

	if err := u.db.QueryRow(config.CountAllUser).Scan(&totalRows); err != nil {
		return nil, shared_model.Paging{}, err
	}

	paging := shared_model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}

	return users, paging, nil
}

// GetByEmail implements User.
func (u *userRepository) GetByEmail(email string) (model.User, error) {
	var user model.User

	err := u.db.QueryRow(config.GetUserByEmail, email).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Println("user not found", err.Error())
		return model.User{}, err
	}

	return user, nil
}

// GetById implements User.
func (u *userRepository) GetById(id string) (model.User, error) {
	var user model.User

	err := u.db.QueryRow(config.GetUserByID, id).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Println("user not found", err.Error())
		return model.User{}, err
	}

	//isi task
	if user.Role == "MANAGER" {
		var projects []model.Project
		row, err := u.db.Query(config.GetProjectByManagerID, user.Id)
		if err != nil {
			log.Println("Currently no project assigned")
		}

		for row.Next() {
			project := model.Project{}
			//updated_at cannot be nil
			err := row.Scan(&project.Id, &project.Name, &project.ManagerId, &project.Deadline, &project.CreatedAt, &project.UpdatedAt)
			if err != nil {
				log.Println("projectRepository.Rows.Next", err.Error())
			}

			projects = append(projects, project)
		}

		user.Project = projects
	} else if user.Role == "TEAM MEMBER" {
		var tasks []model.Task

		row, err := u.db.Query(config.GetTaskByPersonInCharge, user.Id)
		if err != nil {
			log.Println("Currently no task available")
		}
		for row.Next() {
			task := model.Task{}
			err := row.Scan(&task.Id, &task.Name, &task.Status, &task.Approval, &task.PersonInCharge, &task.Deadline, &task.ProjectId, &task.ApprovalDate, &task.Feedback, &task.CreatedAt, &task.UpdatedAt)
			if err != nil {
				log.Println("taskRepository.Rows.Next", err.Error())
			}
			tasks = append(tasks, task)
		}
		user.Task = tasks

		var projects []model.Project
		row, err = u.db.Query(config.GetAllProjectMember, user.Id)
		if err != nil {
			log.Println("Currently no project assigned")
		}

		for row.Next() {
			project := model.Project{}
			//updated_at cannot be nil
			err := row.Scan(&project.Id, &project.Name, &project.ManagerId, &project.Deadline, &project.CreatedAt, &project.UpdatedAt)
			if err != nil {
				log.Println("projectRepository.Rows.Next", err.Error())
			}

			projects = append(projects, project)
		}

		user.Project = projects
	}

	return user, nil
}

// CreateUser implements User.
func (u *userRepository) CreateUser(payload model.User) (model.User, error) {
	var user model.User

	err := u.db.QueryRow(config.CreateUser, payload.Name, payload.Email, payload.Password, payload.Role).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Println("user_repository.QueryRow", err.Error())
		return model.User{}, err
	}
	return user, nil
}

// Update implements User.
func (u *userRepository) Update(payload model.User) (model.User, error) {
	var user model.User
	err := u.db.QueryRow(config.UpdateUser, payload.Id, payload.Name, payload.Email, payload.Password, payload.Role).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Println("user_repository.QueryRow", err.Error())
		return model.User{}, err
	}
	return user, nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

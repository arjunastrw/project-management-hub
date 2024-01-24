package config

const (
	//users
	DeleteUserById = "UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at = null"
	GetAllUser     = "SELECT id, name, email, role, created_at, updated_at FROM users WHERE deleted_at = null ORDER BY created_at DESC LIMIT $1 OFFSET $2"
	GetUserByID    = "SELECT id, name, email, password, role, created_at, updated_at FROM users WHERE id = $1 AND deleted_at = null"
	GetUserByEmail = "SELECT id, name, email, password, role, created_at, updated_at FROM users WHERE email = $1 AND deleted_at = null"
	CreateUser     = "INSERT INTO users(name, email, password, role, deleted_at) VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, null) RETURNING id, name, email, password, role, created_at, updated_at"
	UpdateUser     = "UPDATE users SET name = $2, email = $3, password = $4, role = $5, updated_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at = null RETURNING id, name, email, password, role, created_at, updated_at"
	CountAllUser   = "SELECT COUNT(*) FROM users WHERE deleted_at = null"

	//projects
	GetAll(page int, size int) ([]model.Project, shared_model.Paging, error)
	GetById(id string) (model.Project, error)
	GetByDeadline(date time.Time) ([]model.Project, shared_model.Paging, error)
	GetByManagerId(id string) ([]model.Project, shared_model.Paging, error)
	GetByMemberId(id string) ([]model.Project, shared_model.Paging, error)
	CreateProject(payload model.Project) (model.Project, error)
	EditProjectMember(id string, members []string) ([]model.User, error)
	GetAllProjectMember(id string) ([]model.User, error)
	Update(payload model.Project) (model.Project, error)
	Delete(string string) error

	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    manager_id UUID NOT NULL,
    deadline DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,

	GetAllProject = "SELECT id, name, manager_id, deadline, created_at, updated_at FROM projects WHERE deleted_at = null ORDER BY deadline DESC LIMIT $1 OFFSET $2"

	SelectTasksByAuthor = "select * from tasks where author_id = $1"
	CreateTask          = "Insert Into tasks (title, content, author_id, updated_at) VALUES ($1, $2, $3, CURRENT_TIMESTAMP) RETURNING id, created_at"
	SelectAllTasks      = "SELECT * FROM tasks ORDER BY created_at DESC LIMIT $1 OFFSET $2"
	CountAllTasks       = "SELECT COUNT(*) FROM tasks"
	GetTaskById         = "SELECT * FROM tasks WHERE id = $1"
	UpdateTaskById      = "UPDATE tasks SET title = $2, content = $3, updated_at = CURRENT_TIMESTAMP WHERE ID = $1 RETURNING id, author_id, created_at, updated_at"
	DeleteTask          = "DELETE FROM tasks WHERE id = $1"
)

package config

const (
	//users
	DeleteUserById = "UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at IS NULL"
	GetAllUser     = "SELECT id, name, email, role, created_at, updated_at FROM users WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT $1 OFFSET $2"
	GetUserByID    = "SELECT id, name, email, password, role, created_at, updated_at FROM users WHERE id = $1 AND deleted_at IS NULL"
	GetUserByEmail = "SELECT id, name, email, password, role, created_at, updated_at FROM users WHERE email = $1 AND deleted_at IS NULL"
	CreateUser     = "INSERT INTO users(name, email, password, role) VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP) RETURNING id, name, email, password, role, created_at, updated_at"
	UpdateUser     = "UPDATE users SET name = $2, email = $3, password = $4, role = $5, updated_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at IS NULL RETURNING id, name, email, password, role, created_at, updated_at"
	CountAllUser   = "SELECT COUNT(*) FROM users WHERE deleted_at IS NULL"

	//projects
	GetAllProject         = "SELECT id, name, manager_id, deadline, created_at, updated_at FROM projects WHERE deleted_at IS NULL ORDER BY deadline DESC LIMIT $1 OFFSET $2"
	GetProjectByID        = "SELECT id, name, manager_id, deadline, created_at, updated_at FROM projects WHERE deleted_at IS NULL AND id = $1"
	GetProjectByManagerID = "SELECT id, name, manager_id, deadline, created_at, updated_at FROM projects WHERE deleted_at IS NULL AND manager_id = $1 ORDER BY deadline DESC LIMIT $2 OFFSET $3"
	GetProjectByDeadline  = "SELECT id, name, manager_id, deadline, created_at, updated_at FROM projects WHERE deleted_at IS NULL AND deadline = $1 ORDER BY deadline DESC LIMIT $2 OFFSET $3"

	CreateProject = "INSERT INTO projects(name, manager_id, deadline, updated_at) VALUES ($1, $2, $3, CURRENT_TIMESTAMP) RETURNING id, name, manager_id, deadline, created_at, updated_at"
	UpdateProject = "UPDATE projects SET name = $2, manager_id = $3, deadline = $4, updated_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at IS NULL RETURNING id, name, manager_id, deadline, created_at, updated_at"
	DeleteProject = "UPDATE projects SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at IS NULL"

	AddProjectMember        = "INSERT INTO project_members(member_id, project_id) VALUES ($1, $2)"
	GetAllProjectMember     = "SELECT member_id FROM project_members WHERE project_id = $1 AND deleted_at IS NULL"
	GetAllProjectByMemberID = "SELECT project_id FROM project_members WHERE member_id = $1 AND deleted_at IS NULL"
	DeleteProjectMember     = "UPDATE project_members SET deleted_at = CURRENT_TIMESTAMP WHERE member_id = $1 AND project_id = $2"

	SelectTasksByAuthor = "select * from tasks where author_id = $1"
	CreateTask          = "Insert Into tasks (title, content, author_id, updated_at) VALUES ($1, $2, $3, CURRENT_TIMESTAMP) RETURNING id, created_at"
	SelectAllTasks      = "SELECT * FROM tasks ORDER BY created_at DESC LIMIT $1 OFFSET $2"
	CountAllTasks       = "SELECT COUNT(*) FROM tasks"
	GetTaskById         = "SELECT * FROM tasks WHERE id = $1"
	UpdateTaskById      = "UPDATE tasks SET title = $2, content = $3, updated_at = CURRENT_TIMESTAMP WHERE ID = $1 RETURNING id, author_id, created_at, updated_at"
	DeleteTask          = "DELETE FROM tasks WHERE id = $1"
)

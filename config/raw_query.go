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

	//tasks
	GetAllTask              = "SELECT id, name, status, approval, person_in_charge, deadline, project_id, approval_date, CASE WHEN feedback IS NULL THEN '-' ELSE feedback END, created_at, updated_at FROM tasks WHERE deleted_at IS NULL ORDER BY deadline DESC LIMIT $1 OFFSET $2"
	CountAllTask            = "SELECT COUNT(*) FROM tasks WHERE deleted_at IS NULL"
	GetTaskById             = "SELECT id, name, status, approval, person_in_charge, deadline, project_id, approval_date, CASE WHEN feedback IS NULL THEN '-' ELSE feedback END, created_at, updated_at FROM tasks WHERE id = $1 AND deleted_at IS NULL"
	GetTaskByPersonInCharge = "SELECT id, name, status, approval, person_in_charge, deadline, project_id, approval_date, CASE WHEN feedback IS NULL THEN '-' ELSE feedback END, created_at, updated_at FROM tasks WHERE person_in_charge=$1 AND deleted_at IS NULL"
	GetTaskByProjectId      = "SELECT id, name, status, approval, person_in_charge, deadline, project_id, approval_date, CASE WHEN feedback IS NULL THEN '-' ELSE feedback END, created_at, updated_at FROM tasks WHERE project_id=$1 AND deleted_at IS NULL"
	CreateTask              = "INSERT INTO tasks(name, status, approval, person_in_charge, deadline, project_id, updated_at) VALUES ($1, 'In Progress', false, $2, $3, $4, CURRENT_TIMESTAMP) RETURNING id, name, person_in_charge, deadline, project_id, created_at"
	UpdateTaskByManager     = "UPDATE tasks SET name = $2, status = $3, approval = $4, person_in_charge = $5, deadline = $6, approval_date = CURRENT_TIMESTAMP, feedback = $7, updated_at = CURRENT_TIMESTAMP WHERE id = $1 AND AND deleted_at IS NULL id, name, status, approval, person_in_charge, deadline, project_id, approval_date, CASE WHEN feedback IS NULL THEN '-' ELSE feedback END, created_at, updated_at"
	UpdateTaskByMember      = "UPDATE tasks status = $3, updated_at = CURRENT_TIMESTAMP WHERE id = $1 AND person_in_charge = $2 AND deleted_at IS NULL id, name, status, approval, person_in_charge, deadline, project_id, approval_date, CASE WHEN feedback IS NULL THEN '-' ELSE feedback END, created_at, updated_at"
	DeleteTask              = "UPDATE tasks SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at IS NULL"
)

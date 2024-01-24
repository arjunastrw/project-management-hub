package repository

import (
	"database/sql"
	"log"

	"enigma.com/projectmanagementhub/model"
)

type Report interface {
	CreateReport(payload model.Report) (model.Report, error)
	UpdateReport(payload model.Report) (model.Report, error)
	DeleteReportById(id string) error
	GetReportByTaskId(taskId string) ([]model.Report, error)
	GetReportByProjectId(projectId string) ([]model.Report, error)
}

type report struct {
	db *sql.DB
}

// CreateReport implements Report.
func (r *report) CreateReport(payload model.Report) (model.Report, error) {
	var report model.Report
	query := "INSERT INTO reports(user_id, report, task_id, project_id, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, null, nul) RETURNING id, user_id, report, task_id, project_id, created_at, updated_at, deleted_at"

	err := r.db.QueryRow(query, payload).Scan(&report.Id, &report.User_id, &report.Report, &report.Task_id, &report.Project_id, &report.Created_at, &report.Updated_at, &report.Deleted_at)
	if err != nil {
		log.Println("report_repository.QueryRow", err.Error())
		return model.Report{}, err
	}
	return report, nil

}

// DeleteReportById implements Report.
func (r *report) DeleteReportById(id string) error {
	query := "UPDATE reports SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at = null"
	_, err := r.db.Query(query, id)
	if err != nil {
		log.Println("report_repository.Query", err.Error())
		return err
	}

	return nil
}

// GetReportByProjectId implements Report.
func (r *report) GetReportByProjectId(projectId string) ([]model.Report, error) {
	var reports []model.Report
	query := "SELECT id, user_id, report, task_id, created_at, updated_at, deleted_at FROM reports WHERE id = $1 AND deleted_at = null"

	rows, err := r.db.Query(query, projectId)
	if err != nil {
		log.Println("report_repository.QueryRow", err.Error())
		return nil, err
	}

	for rows.Next() {
		report := model.Report{}
		//updated_at cannot be nil
		err := rows.Scan(&report.Id, &report.User_id, &report.Report, &report.Task_id, &report.Project_id, &report.Created_at, &report.Updated_at, &report.Deleted_at)
		if err != nil {
			log.Println("userRepository.Rows.Next", err.Error())
			return nil, err
		}

		reports = append(reports, report)
	}

	return reports, nil
}

// GetReportByTaskId implements Report.
func (r *report) GetReportByTaskId(taskId string) ([]model.Report, error) {
	var reports []model.Report
	query := "SELECT id, user_id, report, task_id, created_at, updated_at, deleted_at FROM reports WHERE id = $1 AND deleted_at = null"

	rows, err := r.db.Query(query, taskId)
	if err != nil {
		log.Println("report_repository.QueryRow", err.Error())
		return nil, err
	}

	for rows.Next() {
		report := model.Report{}
		//updated_at cannot be nil
		err := rows.Scan(&report.Id, &report.User_id, &report.Report, &report.Task_id, &report.Project_id, &report.Created_at, &report.Updated_at, &report.Deleted_at)
		if err != nil {
			log.Println("reportRepository.Rows.Next", err.Error())
			return nil, err
		}

		reports = append(reports, report)
	}

	return reports, nil
}

// UpdateReport implements Report.
func (r *report) UpdateReport(payload model.Report) (model.Report, error) {
	var report model.Report
	query := "UPDATE reports SET user_id = $2, report = $3, task_id = $4, project_id = $5, updated_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at = null RETURNING  id, user_id, report, task_id, project_id, created_at, updated_at, deleted_at"
	err := r.db.QueryRow(query, payload.Id, payload.User_id, payload.Report, payload.Task_id, payload.Project_id).Scan(&report.Id, &report.User_id, &report.Report, &report.Task_id, &report.Project_id, &report.Created_at, &report.Updated_at, &report.Deleted_at)
	if err != nil {
		log.Println("report_repository.QueryRow", err.Error())
		return model.Report{}, err
	}
	return report, nil
}

func NewReportRepository(db *sql.DB) Report {
	return &report{
		db: db,
	}
}

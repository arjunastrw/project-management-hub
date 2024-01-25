package repository

import (
	"database/sql"
	"log"

	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/report"
)

type ReportRepository interface {
	CreateReport(payload model.Report) (model.Report, error)
	UpdateReport(payload model.Report) (model.Report, error)
	DeleteReportById(id string) error
	GetReportByTaskId(taskId string) ([]model.Report, error)
	GetReportByUserId(userId string) ([]model.Report, error)
}

type reportRepository struct {
	db     *sql.DB
	report report.ReportToTXT
}

// CreateReport implements Report.
func (r *reportRepository) CreateReport(payload model.Report) (model.Report, error) {
	var report model.Report
	query := "INSERT INTO reports(user_id, report, task_id, updated_at) VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP) RETURNING id, user_id, report, task_id, created_at, updated_at"

	err := r.db.QueryRow(query, payload).Scan(&report.Id, &report.User_id, &report.Report, &report.Task_id, &report.Created_at, &report.Updated_at)
	if err != nil {
		log.Println("report_repository.QueryRow", err.Error())
		return model.Report{}, err
	}

	// report to txt
	reportToTXT := model.ShowReport{Content: report}
	err = r.report.WriteReport(reportToTXT)
	if err != nil {
		return model.Report{}, err
	}
	return report, nil

}

// DeleteReportById implements Report.
func (r *reportRepository) DeleteReportById(id string) error {
	query := "UPDATE reports SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1 AND deleted_at IS null"
	_, err := r.db.Query(query, id)
	if err != nil {
		log.Println("report_repository.Query", err.Error())
		return err
	}

	return nil
}

// GetReportByProjectId implements Report.
func (r *reportRepository) GetReportByUserId(userId string) ([]model.Report, error) {
	var reports []model.Report
	query := "SELECT id, user_id, report, task_id, created_at, updated_at FROM reports WHERE user_id = $1 AND deleted_at IS null"

	rows, err := r.db.Query(query, userId)
	if err != nil {
		log.Println("report_repository.QueryRow", err.Error())
		return nil, err
	}

	for rows.Next() {
		report := model.Report{}
		//updated_at cannot be nil
		err := rows.Scan(&report.Id, &report.User_id, &report.Report, &report.Task_id, &report.Created_at, &report.Updated_at)
		if err != nil {
			log.Println("report_Repository.Rows.Next", err.Error())
			return nil, err
		}

		reports = append(reports, report)
	}

	return reports, nil
}

// GetReportByTaskId implements Report.
func (r *reportRepository) GetReportByTaskId(taskId string) ([]model.Report, error) {
	var reports []model.Report
	query := "SELECT id, user_id, report, task_id, created_at, updated_at FROM reports WHERE task_id = $1 AND deleted_at IS null"

	rows, err := r.db.Query(query, taskId)
	if err != nil {
		log.Println("report_repository.QueryRow", err.Error())
		return nil, err
	}

	for rows.Next() {
		report := model.Report{}
		//updated_at cannot be nil
		err := rows.Scan(&report.Id, &report.User_id, &report.Report, &report.Task_id, &report.Created_at, &report.Updated_at)
		if err != nil {
			log.Println("reportRepository.Rows.Next", err.Error())
			return nil, err
		}

		reports = append(reports, report)
	}

	return reports, nil
}

// UpdateReport implements Report.
func (r *reportRepository) UpdateReport(payload model.Report) (model.Report, error) {
	var report model.Report
	query := "UPDATE reports SET report = $3, task_id = $4, updated_at = CURRENT_TIMESTAMP WHERE id = $1 AND user_id = $2 AND deleted_at IS null RETURNING  id, user_id, report, task_id, created_at, updated_at"
	err := r.db.QueryRow(query, payload.Id, payload.User_id, payload.Report, payload.Task_id).Scan(&report.Id, &report.User_id, &report.Report, &report.Task_id, &report.Created_at, &report.Updated_at)
	if err != nil {
		log.Println("report_repository.QueryRow", err.Error())
		return model.Report{}, err
	}

	// report to txt
	reportToTXT := model.ShowReport{Content: report}
	err = r.report.WriteReport(reportToTXT)
	if err != nil {
		return model.Report{}, err
	}
	return report, nil
}

func NewReportRepository(db *sql.DB) ReportRepository {
	return &reportRepository{
		db: db,
	}
}

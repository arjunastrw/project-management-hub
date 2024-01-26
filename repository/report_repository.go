package repository

import (
	"database/sql"
	"fmt"
	"log"

	"enigma.com/projectmanagementhub/config"
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

	err := r.db.QueryRow(config.CreateReport, payload.User_id, payload.Report, payload.Task_id).Scan(&report.Id, &report.User_id, &report.Report, &report.Task_id, &report.Created_at, &report.Updated_at)
	if err != nil {
		log.Println("report_repository.QueryRow", err.Error())
		return model.Report{}, err
	}

	//report to txt
	reportToTXT := model.ShowReport{Content: report}
	err = r.report.WriteReport(reportToTXT)
	if err != nil {
		return model.Report{}, err
	}
	return report, nil

}

// DeleteReportById implements Report.
func (r *reportRepository) DeleteReportById(id string) error {

	_, err := r.db.Query(config.DeleteReportById, id)
	if err != nil {
		log.Println("report_repository.Query", err.Error())
		return err
	}

	return nil
}

// GetReportByProjectId implements Report.
func (r *reportRepository) GetReportByUserId(userId string) ([]model.Report, error) {
	var reports []model.Report

	rows, err := r.db.Query(config.GetReportByUserId, userId)
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

	rows, err := r.db.Query(config.GetReportByTaskId, taskId)
	if err != nil {
		log.Println("report_repository.QueryRow", err.Error())
		return nil, err
	}

	for rows.Next() {
		report := model.Report{}
		//updated_at cannot be nil
		err := rows.Scan(&report.Id, &report.User_id, &report.Report, &report.Task_id, &report.Created_at, &report.Updated_at)
		fmt.Println("ini report :", report)
		if err != nil {
			log.Println("reportRepository.Rows.Next", err.Error())
			return nil, err
		}

		reports = append(reports, report)
	}
	fmt.Println("ini reports :", reports)

	return reports, nil
}

// UpdateReport implements Report.
func (r *reportRepository) UpdateReport(payload model.Report) (model.Report, error) {
	var report model.Report
	err := r.db.QueryRow(config.UpdateReport, payload.Id, payload.User_id, payload.Report, payload.Task_id).Scan(&report.Id, &report.User_id, &report.Report, &report.Task_id, &report.Created_at, &report.Updated_at)
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
		db:     db,
		report: &report.AreportToTXT{},
	}
}

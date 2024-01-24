package usecase

import (
	"fmt"
	"log"

	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/repository"
)

type ReportUsecase interface {
	CreateReport(payload model.Report) (model.Report, error)
	UpdateReport(payload model.Report) (model.Report, error)
	DeleteReportById(id string) error
	GetReportByTaskId(taskId string) ([]model.Report, error)
	GetReportUserId(userId string) (model.Report, error)
}

type reportUsecase struct {
	reportRepository repository.Report
}

// GetReportUserId implements ReportUsecase.
func (*reportUsecase) GetReportUserId(userId string) ([]model.Report, error) {
	if taskId == "" {
		return nil, fmt.Errorf("task id cannot be empty")
	}

	if r.reportRepository == nil {
		return nil, fmt.Errorf("report repository cannot be nil: %v", r.reportRepository)
	}

	reports, err := r.reportRepository.GetReportByTaskId(taskId)
	if err != nil {
		return nil, err
	}
	return reports, nil
}

// CreateReport implements ReportUsecase.
func (r *reportUsecase) CreateReport(payload model.Report) (model.Report, error) {
	_, err := r.reportRepository.GetReportByTaskId(payload.Task_id)
	if err != nil {
		return model.Report{}, err
	}

	if payload.Report == "" || payload.Task_id == "" || payload.User_id == "" {
		log.Fatal("data cannot be empty: %v", payload)
	}

	report, err := r.reportRepository.CreateReport(payload)
	if err != nil {
		return model.Report{}, err
	}

	return report, nil
}

// DeleteReportById implements ReportUsecase.
func (r *reportUsecase) DeleteReportById(id string) error {
	if id == "" {
		return fmt.Errorf("report id cannot be empty")
	}
	_, err := r.reportRepository.GetReportByTaskId(id)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// GetReportByTaskId implements ReportUsecase.
func (r *reportUsecase) GetReportByTaskId(taskId string) ([]model.Report, error) {
	if taskId == "" {
		return nil, fmt.Errorf("task id cannot be empty")
	}

	if r.reportRepository == nil {
		return nil, fmt.Errorf("report repository cannot be nil: %v", r.reportRepository)
	}

	reports, err := r.reportRepository.GetReportByTaskId(taskId)
	if err != nil {
		return nil, err
	}
	return reports, nil
}

// UpdateReport implements ReportUsecase.
func (r *reportUsecase) UpdateReport(payload model.Report) (model.Report, error) {
	_, err := r.reportRepository.GetReportByTaskId(payload.Task_id)
	if err != nil {
		return model.Report{}, err
	}

	if payload.Id == "" || payload.User_id == "" || payload.Report == "" || payload.Task_id == "" {
		return model.Report{}, fmt.Errorf("report cannot be empty")
	}

	reports, err := r.reportRepository.UpdateReport(payload)
	if err != nil {
		return model.Report{}, fmt.Errorf("failed to update report : %v", err)
	}
	return reports, nil
}

func NewReportUsecase(reportRepository repository.Report) ReportUsecase {
	return &reportUsecase{
		reportRepository: reportRepository,
	}
}

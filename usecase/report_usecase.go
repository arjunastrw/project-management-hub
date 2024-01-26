package usecase

import (
	"errors"
	"fmt"

	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/repository"
)

type ReportUsecase interface {
	CreateReport(payload model.Report) (model.Report, error)
	UpdateReport(payload model.Report) (model.Report, error)
	DeleteReportById(id string) error
	GetReportByTaskId(taskId string) ([]model.Report, error)
	GetReportByUserId(userId string) ([]model.Report, error)
}

type reportUsecase struct {
	reportRepository repository.ReportRepository
	taskRepo         repository.TaskRepository
}

// GetReportUserId implements ReportUsecase.
func (r *reportUsecase) GetReportByUserId(userId string) ([]model.Report, error) {
	if userId == "" {
		return []model.Report{}, nil
	}

	if r.reportRepository == nil {
		return nil, fmt.Errorf("report repository cannot be nil: %v", r.reportRepository)
	}

	reports, err := r.reportRepository.GetReportByUserId(userId)
	if err != nil {
		return nil, err
	}
	return reports, nil
}

// CreateReport implements ReportUsecase.
func (r *reportUsecase) CreateReport(payload model.Report) (model.Report, error) {

	_, err := r.taskRepo.GetByPersonInCharge(payload.User_id)
	if err != nil {
		return model.Report{}, err
	}

	if payload.Report == "" || payload.Task_id == "" || payload.User_id == "" {
		return model.Report{}, errors.New("report cannot be empty")
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
	return r.reportRepository.DeleteReportById(id)
}

// GetReportByTaskId implements ReportUsecase.
func (r *reportUsecase) GetReportByTaskId(taskId string) ([]model.Report, error) {
	_, err := r.taskRepo.GetById(taskId)
	if err != nil {
		return []model.Report{}, errors.New("task not found")
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
	_, err := r.taskRepo.GetById(payload.Task_id)
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

func NewReportUsecase(reportRepository repository.ReportRepository, taskRepo repository.TaskRepository) ReportUsecase {
	return &reportUsecase{
		reportRepository: reportRepository,
		taskRepo:         taskRepo,
	}
}

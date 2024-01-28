package usecase_mock

import (
	"enigma.com/projectmanagementhub/model"
	"github.com/stretchr/testify/mock"
)

type ReportUsecaseMock struct {
	mock.Mock
}

// GetReportByTaskId implements usecase.ReportUsecase.

func (r *ReportUsecaseMock) CreateReport(payload model.Report) (model.Report, error) {
	args := r.Called(payload)
	return args.Get(0).(model.Report), args.Error(1)
}

func (r *ReportUsecaseMock) UpdateReport(payload model.Report) (model.Report, error) {
	args := r.Called(payload)
	return args.Get(0).(model.Report), args.Error(1)
}

func (r *ReportUsecaseMock) DeleteReportById(id string) error {
	args := r.Called(id)
	return args.Error(0)
}

func (r *ReportUsecaseMock) GetReportByUserId(userId string) ([]model.Report, error) {
	args := r.Called(userId)
	return args.Get(0).([]model.Report), args.Error(1)
}

func (r *ReportUsecaseMock) GetReportByTaskId(taskId string) ([]model.Report, error) {
	args := r.Called(taskId)
	return args.Get(0).([]model.Report), args.Error(1)
}

func NewReportsUsecaseMock() ReportUsecaseMock {
	return ReportUsecaseMock{}
}

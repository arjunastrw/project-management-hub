package repository_mock

import (
	"enigma.com/projectmanagementhub/model"
	"github.com/stretchr/testify/mock"
)

type ReportRepositoryMock struct {
	mock.Mock
}

func (m *ReportRepositoryMock) CreateReport(payload model.Report) (model.Report, error) {
	// args is a struct that captures the arguments passed to the method
	args := m.Called(payload)

	// Extract the first return value, which should be a model.Report
	result, ok := args.Get(0).(model.Report)
	if !ok {
		// If the type assertion fails, return a zero-initialized model.Report
		return model.Report{}, args.Error(1)
	}

	return result, args.Error(1)
}

func (m *ReportRepositoryMock) UpdateReport(payload model.Report) (model.Report, error) {
	args := m.Called(payload)

	result, ok := args.Get(0).(model.Report)
	if !ok {
		return model.Report{}, args.Error(1)
	}

	return result, args.Error(1)
}

func (m *ReportRepositoryMock) DeleteReportById(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *ReportRepositoryMock) GetReportByTaskId(taskId string) ([]model.Report, error) {
	args := m.Called(taskId)

	result, ok := args.Get(0).([]model.Report)
	if !ok {
		return nil, args.Error(1)
	}

	return result, args.Error(1)
}

func (m *ReportRepositoryMock) GetReportByUserId(userId string) ([]model.Report, error) {
	args := m.Called(userId)

	result, ok := args.Get(0).([]model.Report)
	if !ok {
		return nil, args.Error(1)
	}

	return result, args.Error(1)
}

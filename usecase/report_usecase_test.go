package usecase

import (
	"testing"
	"time"

	"enigma.com/projectmanagementhub/mock/repository_mock"
	"enigma.com/projectmanagementhub/model"
	"github.com/stretchr/testify/suite"
)

type ReportUsecaseSuite struct {
	suite.Suite
	reportRepo *repository_mock.ReportRepositoryMock
	taskRepo   *repository_mock.TaskRepositoryMock
	ReportUc   ReportUsecase
}

var ExpectedReport = model.Report{
	Id:         "ed09d2f3-1878-4e11-adaf-a14326c81657",
	User_id:    "09effbb3-34fe-4719-a1f6-33619f926577",
	Report:     "report",
	Task_id:    "task_id",
	Created_at: time.Now(),
	Updated_at: time.Time{},
	DeletedAt:  &time.Time{},
}

var ExpectedTask = model.Task{
	Id:             "task_id",
	Name:           "task_name",
	Status:         "task_status",
	Approval:       false,
	ApprovalDate:   &time.Time{},
	Feedback:       "task_feedback",
	PersonInCharge: "09effbb3-34fe-4719-a1f6-33619f926577",
	ProjectId:      "project_id",
	Deadline:       "2023-01-01",
}

func (t *ReportUsecaseSuite) SetupTest() {
	t.reportRepo = &repository_mock.ReportRepositoryMock{}
	t.taskRepo = &repository_mock.TaskRepositoryMock{}
	t.ReportUc = NewReportUsecase(t.reportRepo, t.taskRepo)
}

// func unit test to get report by user id
func (t *ReportUsecaseSuite) TestGetReportByUserId_Success() {
	t.reportRepo.On("GetReportByUserId", ExpectedReport.User_id).Return([]model.Report{ExpectedReport}, nil)
	actual, err := t.ReportUc.GetReportByUserId(ExpectedReport.User_id)
	t.NoError(err)

	t.Equal(ExpectedReport.Report, actual[0].Report)

}

func (t *ReportUsecaseSuite) TestGetReportByUserId_Failed() {
	t.reportRepo.On("GetReportByUserId", ExpectedReport.User_id).Return([]model.Report{}, nil)
	_, err := t.ReportUc.GetReportByUserId(ExpectedReport.User_id)
	t.NoError(err)
	t.Nil(err)
}

// func unit test to get report by task id
func (t *ReportUsecaseSuite) TestGetReportByTaskId_Success() {
	t.taskRepo.On("GetById", ExpectedTask.Id).Return(ExpectedTask, nil)
	t.reportRepo.On("GetReportByTaskId", ExpectedReport.Task_id).Return([]model.Report{ExpectedReport}, nil)
	actual, err := t.ReportUc.GetReportByTaskId(ExpectedReport.Task_id)
	t.NoError(err)
	t.Equal(ExpectedReport.Report, actual[0].Report)
}

func (t *ReportUsecaseSuite) TestGetReportByTaskId_Failed() {
	t.taskRepo.On("GetById", ExpectedTask.Id).Return(ExpectedTask, nil)
	t.reportRepo.On("GetReportByTaskId", ExpectedReport.Task_id).Return([]model.Report{}, nil)
	_, err := t.ReportUc.GetReportByTaskId(ExpectedReport.Task_id)
	t.NoError(err)
	t.Nil(err)
}

// func unit test to create report
func (t *ReportUsecaseSuite) TestCreateReport_Success() {
	// t.reportRepo.On("GetById", ExpectedTask.Id).Return(ExpectedTask, nil)
	t.taskRepo.On("GetByPersonInCharge", ExpectedTask.PersonInCharge).Return([]model.Task{ExpectedTask}, nil)
	t.reportRepo.On("CreateReport", ExpectedReport).Return(ExpectedReport, nil)
	actual, err := t.ReportUc.CreateReport(ExpectedReport)
	t.NoError(err)
	t.Nil(err)
	t.Equal(ExpectedReport.Report, actual.Report)
}

func (t *ReportUsecaseSuite) TestCreateReport_Failed() {
	t.taskRepo.On("GetByPersonInCharge", ExpectedTask.PersonInCharge).Return([]model.Task{ExpectedTask}, nil)
	t.reportRepo.On("CreateReport", ExpectedReport).Return(model.Report{}, nil)
	_, err := t.ReportUc.CreateReport(ExpectedReport)
	t.NoError(err)
	t.Nil(err)
}

// func unit test to update report
func (t *ReportUsecaseSuite) TestUpdateReport_Success() {
	t.taskRepo.On("GetById", ExpectedTask.Id).Return(ExpectedTask, nil)
	t.reportRepo.On("UpdateReport", ExpectedReport).Return(ExpectedReport, nil)
	actual, err := t.ReportUc.UpdateReport(ExpectedReport)
	t.NoError(err)
	t.Nil(err)
	t.Equal(ExpectedReport.Report, actual.Report)
}

func (t *ReportUsecaseSuite) TestUpdateReport_Failed() {
	t.taskRepo.On("GetById", ExpectedTask.Id).Return(ExpectedTask, nil)
	t.reportRepo.On("UpdateReport", ExpectedReport).Return(model.Report{}, nil)
	_, err := t.ReportUc.UpdateReport(ExpectedReport)
	t.NoError(err)
	t.Nil(err)
}

// func unit test to delete report
func (t *ReportUsecaseSuite) TestDeleteReportById_Success() {
	t.reportRepo.On("DeleteReportById", ExpectedReport.Id).Return(nil)
	err := t.ReportUc.DeleteReportById(ExpectedReport.Id)
	t.NoError(err)
	t.Nil(err)
}

func (t *ReportUsecaseSuite) TestDeleteReportById_Failed() {
	t.reportRepo.On("DeleteReportById", ExpectedReport.Id).Return(nil)
	err := t.ReportUc.DeleteReportById("")
	t.Error(err)
	t.NotNil(err)
}

func TestReportUscaseSuite(t *testing.T) {
	suite.Run(t, new(ReportUsecaseSuite))
}

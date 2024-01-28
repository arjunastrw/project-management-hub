package repository

import (
	"database/sql"
	"testing"
	"time"

	"enigma.com/projectmanagementhub/config"
	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/report"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"
)

type ReportRepositoryTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    ReportRepository
	reports report.ReportToTXT
}

func (r *ReportRepositoryTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	r.mockDB = db
	r.mockSql = mock
	r.reports = report.NewReportToTXT(config.PathConfig{StaticPath: "file dir path"})
	r.repo = NewReportRepository(r.mockDB, r.reports)
}

var expectedReport = model.Report{
	Id:         "1",
	User_id:    "1",
	Report:     "This is report",
	Task_id:    "1",
	Created_at: time.Now(),
	Updated_at: time.Now(),
	DeletedAt:  nil,
}

func TestReportRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ReportRepositoryTestSuite))
}

func (r *ReportRepositoryTestSuite) TearDownTest() {
	r.mockDB.Close()
}

func (r *ReportRepositoryTestSuite) TestCreateReport_Success() {
	// Ekspektasi bahwa panggilan QueryRowContext akan terjadi
	r.mockSql.ExpectQuery("INSERT INTO reports").
		WithArgs(expectedReport.User_id, expectedReport.Report, expectedReport.Task_id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "report", "task_id", "created_at", "updated_at"}).
			AddRow(expectedReport.Id, expectedReport.User_id, expectedReport.Report, expectedReport.Task_id, expectedReport.Created_at, expectedReport.Updated_at))

	// Melakukan pemanggilan metode yang diuji
	reportCreated, err := r.repo.CreateReport(expectedReport)

	// Mengecek bahwa tidak ada error yang diharapkan
	r.NoError(err, "CreateReport should not return an error")

	// Mengecek bahwa nilai yang dikembalikan sesuai dengan ekspektasi
	r.Equal(expectedReport, reportCreated, "Created report should match the expected report")

	// Memastikan bahwa ekspektasi panggilan database terpenuhi
	err = r.mockSql.ExpectationsWereMet()
	r.NoError(err, "Database expectations were not met")
}

func (r *ReportRepositoryTestSuite) TestCreateReport_Failure() {
	// Ekspektasi bahwa panggilan QueryRowContext akan terjadi
	r.mockSql.ExpectQuery("INSERT INTO reports").
		WithArgs(expectedReport.User_id, expectedReport.Report, expectedReport.Task_id).
		WillReturnError(sql.ErrNoRows)

	// Melakukan pemanggilan metode yang diuji
	_, err := r.repo.CreateReport(expectedReport)

	// Mengecek bahwa error yang diharapkan terjadi
	r.Error(err, "CreateReport should return an error")

	// Memastikan bahwa ekspektasi panggilan database terpenuhi
	err = r.mockSql.ExpectationsWereMet()
	r.NoError(err, "Database expectations were not met")
}

func (r *ReportRepositoryTestSuite) TestUpdateReport_Success() {

	r.mockSql.ExpectQuery("UPDATE reports").
		WithArgs(expectedReport.Id, expectedReport.User_id, expectedReport.Report, expectedReport.Task_id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "report", "task_id", "created_at", "updated_at"}).
			AddRow(expectedReport.Id, expectedReport.User_id, expectedReport.Report, expectedReport.Task_id, expectedReport.Created_at, expectedReport.Updated_at))

	reportUpdated, err := r.repo.UpdateReport(expectedReport)

	r.NoError(err, "UpdateReport should not return an error")
	r.Equal(expectedReport, reportUpdated, "Updated report should match the expected report")

	err = r.mockSql.ExpectationsWereMet()
	r.NoError(err, "Database expectations were not met")
}

func (r *ReportRepositoryTestSuite) TestUpdateReport_Failure() {

	r.mockSql.ExpectQuery("UPDATE reports").
		WithArgs(expectedReport.Id, expectedReport.User_id, expectedReport.Report, expectedReport.Task_id).
		WillReturnError(sql.ErrNoRows)

	_, err := r.repo.UpdateReport(expectedReport)

	r.Error(err, "UpdateReport should return an error")

	err = r.mockSql.ExpectationsWereMet()
	r.NoError(err, "Database expectations were not met")
}

func (r *ReportRepositoryTestSuite) TestDeleteReportById_Success() {

	r.mockSql.ExpectQuery("UPDATE reports SET deleted_at").
		WithArgs(expectedReport.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "report", "task_id", "created_at", "updated_at"}).
			AddRow(expectedReport.Id, expectedReport.User_id, expectedReport.Report, expectedReport.Task_id, expectedReport.Created_at, expectedReport.Updated_at))

	err := r.repo.DeleteReportById(expectedReport.Id)

	r.NoError(err, "DeleteReportById should not return an error")

	err = r.mockSql.ExpectationsWereMet()
	r.NoError(err, "Database expectations were not met")
}

func (r *ReportRepositoryTestSuite) TestDeleteReportById_Failure() {

	r.mockSql.ExpectQuery("UPDATE reports SET deleted_at").
		WithArgs(expectedReport.Id).
		WillReturnError(sql.ErrNoRows)

	err := r.repo.DeleteReportById(expectedReport.Id)

	r.Error(err, "DeleteReportById should return an error")

	err = r.mockSql.ExpectationsWereMet()
	r.NoError(err, "Database expectations were not met")
}

func (r *ReportRepositoryTestSuite) TestGetReportByTaskId_Success() {

	r.mockSql.ExpectQuery("SELECT id, user_id, report, task_id, created_at, updated_at FROM reports").
		WithArgs(expectedReport.Task_id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "report", "task_id", "created_at", "updated_at"}).
			AddRow(expectedReport.Id, expectedReport.User_id, expectedReport.Report, expectedReport.Task_id, expectedReport.Created_at, expectedReport.Updated_at))

	reports, err := r.repo.GetReportByTaskId(expectedReport.Task_id)

	r.NoError(err, "GetReportByTaskId should not return an error")
	r.NotNil(reports, "GetReportByTaskId should return non-nil reports")

	err = r.mockSql.ExpectationsWereMet()
	r.NoError(err, "Database expectations were not met")
}

func (r *ReportRepositoryTestSuite) TestGetReportByTaskId_Failure() {

	r.mockSql.ExpectQuery("SELECT id, user_id, report, task_id, created_at, updated_at FROM reports").
		WithArgs(expectedReport.Task_id).
		WillReturnError(sql.ErrNoRows)

	// Calling the tested method
	reports, err := r.repo.GetReportByTaskId(expectedReport.Task_id)

	r.Error(err, "GetReportByTaskId should return an error")
	r.Nil(reports, "GetReportByTaskId should return nil reports")

	err = r.mockSql.ExpectationsWereMet()
	r.NoError(err, "Database expectations were not met")
}

func (r *ReportRepositoryTestSuite) TestGetReportByUserId_Success() {

	r.mockSql.ExpectQuery("SELECT id, user_id, report, task_id, created_at, updated_at FROM reports").
		WithArgs(expectedReport.User_id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "report", "task_id", "created_at", "updated_at"}).
			AddRow(expectedReport.Id, expectedReport.User_id, expectedReport.Report, expectedReport.Task_id, expectedReport.Created_at, expectedReport.Updated_at))

	reports, err := r.repo.GetReportByUserId(expectedReport.User_id)

	r.NoError(err, "GetReportByUserId should not return an error")
	r.NotNil(reports, "GetReportByUserId should return non-nil reports")

	err = r.mockSql.ExpectationsWereMet()
	r.NoError(err, "Database expectations were not met")
}

func (r *ReportRepositoryTestSuite) TestGetReportByUserId_Failure() {

	r.mockSql.ExpectQuery("SELECT id, user_id, report, task_id, created_at, updated_at FROM reports").
		WithArgs(expectedReport.User_id).
		WillReturnError(sql.ErrNoRows)

	reports, err := r.repo.GetReportByUserId(expectedReport.User_id)

	r.Error(err, "GetReportByUserId should return an error")
	r.Nil(reports, "GetReportByUserId should return nil reports")

	err = r.mockSql.ExpectationsWereMet()
	r.NoError(err, "Database expectations were not met")
}

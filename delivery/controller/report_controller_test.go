package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"enigma.com/projectmanagementhub/mock/middleware_mock"
	"enigma.com/projectmanagementhub/mock/usecase_mock"
	"enigma.com/projectmanagementhub/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type ReportControllerTestSuite struct {
	suite.Suite
	ReportUc       *usecase_mock.ReportUsecaseMock
	rg             *gin.RouterGroup
	authMiddleware *middleware_mock.AuthMiddlewareMock
}

var ExpectedReport = model.Report{
	Id:         "ed09d2f3-1878-4e11-adaf-a14326c81657",
	User_id:    "09effbb3-34fe-4719-a1f6-33619f926577",
	Report:     "report",
	Task_id:    "35fc1b48-a4d1-4bf2-9d34-c35271fc282f",
	Created_at: time.Now(),
	Updated_at: time.Time{},
	DeletedAt:  &time.Time{},
}

func (t *ReportControllerTestSuite) SetupTest() {
	t.ReportUc = new(usecase_mock.ReportUsecaseMock)
	t.authMiddleware = new(middleware_mock.AuthMiddlewareMock)
	r := gin.Default()
	rg := r.Group("/pmh-api/v1")
	rg.Use(t.authMiddleware.RequireToken("ADMIN", "TEAM MEMBER", "MANAGER"))
	t.rg = rg
}

func (t *ReportControllerTestSuite) TestCreateNewReportController() {
	t.ReportUc.On("CreateReport", model.Report{}).Return(ExpectedReport, nil)
	reportController := NewReportController(t.ReportUc, t.authMiddleware, t.rg)
	reportController.Route()
	requestBody := `{"user_id":"09effbb3-34fe-4719-a1f6-33619f926577","report":"report","task_id":"35fc1b48-a4d1-4bf2-9d34-c35271fc282f"}`
	request, err := http.NewRequest("POST", "/pmh-api/v1/createreport", strings.NewReader(requestBody))
	t.Nil(err)

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request
	ctx.Set("TEAM MEMBER", true)
	reportController.CreateNewReportController(ctx)
	t.Equal(http.StatusOK, record.Code)
}

func (t *ReportControllerTestSuite) TestCreateNewReportController_Failed() {
	t.ReportUc.On("CreateReport", model.Report{}).Return(ExpectedReport, nil)
	reportController := NewReportController(t.ReportUc, t.authMiddleware, t.rg)
	reportController.Route()
	requestBody := `{"user_id":"09effbb3-34fe-4719-a1f6-33619f926577","report":"report","task_id":"35fc1b48-a4d1-4bf2-9d34-c35271fc282f"}`
	request, err := http.NewRequest("POST", "/pmh-api/v1/createreport", strings.NewReader(requestBody))
	t.Nil(err)

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request
	ctx.Set("TEAM MEMBER", false) //i change this akses to TEAM MEMBER
	reportController.CreateNewReportController(ctx)
	t.NotEqual(http.StatusUnauthorized, record.Code)
}

func (t *ReportControllerTestSuite) TestUpdateReportController() {
	t.ReportUc.On("UpdateReport", model.Report{}).Return(ExpectedReport, nil)
	reportController := NewReportController(t.ReportUc, t.authMiddleware, t.rg)
	reportController.Route()
	requestBody := `{user_id:09effbb3-34fe-4719-a1f6-33619f926577,report:report,task_id:35fc1b48-a4d1-4bf2-9d34-c35271fc282f}`
	request, err := http.NewRequest("PUT", "/pmh-api/v1/updatereport", strings.NewReader(requestBody))
	t.Nil(err)

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request
	ctx.Set("TEAM MEMBER", true)
	reportController.UpdateReportController(ctx)
	t.Equal(http.StatusOK, record.Code)
}

func (t *ReportControllerTestSuite) TestUpdateReportController_Failed() {
	t.ReportUc.On("UpdateReport", model.Report{}).Return(ExpectedReport, nil)
	reportController := NewReportController(t.ReportUc, t.authMiddleware, t.rg)
	reportController.Route()
	requestBody := `{user_id:09effbb3-34fe-4719-a1f6-33619f926577,report:report,task_id:35fc1b48-a4d1-4bf2-9d34-c35271fc282f}`
	request, err := http.NewRequest("PUT", "/pmh-api/v1/updatereport", strings.NewReader(requestBody))
	t.Nil(err)

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request
	ctx.Set("ADMIN", false)
	reportController.UpdateReportController(ctx)
	t.NotEqual(http.StatusUnauthorized, record.Code)
}

func (t *ReportControllerTestSuite) TestGetReportByTaskIdController() {
	t.ReportUc.On("GetReportByTaskId", ExpectedReport.Task_id).Return([]model.Report{ExpectedReport}, nil)
	reportController := NewReportController(t.ReportUc, t.authMiddleware, t.rg)
	reportController.Route()
	request, err := http.NewRequest("GET", "/pmh-api/v1/get/reporttaskid?taskId=35fc1b48-a4d1-4bf2-9d34-c35271fc282f", nil)
	t.Nil(err)

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request
	ctx.Set("ADMIN", "MANAGER")
	reportController.GetReportByTaskIdController(ctx)
	t.Equal(http.StatusOK, record.Code)
}

func (t *ReportControllerTestSuite) TestGetReportByTaskIdController_failed() {
	t.ReportUc.On("GetReportByTaskId", ExpectedReport.Task_id).Return([]model.Report{ExpectedReport}, nil)
	reportController := NewReportController(t.ReportUc, t.authMiddleware, t.rg)
	reportController.Route()

	request, err := http.NewRequest("GET", "/pmh-api/v1/get/reporttaskid?taskId=35fc1b48-a4d1-4bf2-9d34-c35271fc282f", nil)
	t.Nil(err)

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request
	ctx.Set("TEAM MEMBER", false) // i change this akses to TEAM MEMBER
	reportController.GetReportByTaskIdController(ctx)
	fmt.Println("Response Code:", record.Code)
	t.NotEqual(http.StatusUnauthorized, record.Code)
}

func (t *ReportControllerTestSuite) TestGetReportByUserIdController() {
	t.ReportUc.On("GetReportByUserId", ExpectedReport.User_id).Return([]model.Report{ExpectedReport}, nil)
	reportController := NewReportController(t.ReportUc, t.authMiddleware, t.rg)
	reportController.Route()
	request, err := http.NewRequest("GET", "/pmh-api/v1/get/reportuserid?userId=09effbb3-34fe-4719-a1f6-33619f926577", nil)
	t.NoError(err)

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request
	ctx.Set("ADMIN", "MANAGER")
	reportController.GetReportByUserIdController(ctx)
	t.Equal(http.StatusOK, record.Code)
}

func (t *ReportControllerTestSuite) TestGetReportByUserIdController_failed() {
	t.ReportUc.On("GetReportByUserId", ExpectedReport.User_id).Return([]model.Report{ExpectedReport}, nil)
	reportController := NewReportController(t.ReportUc, t.authMiddleware, t.rg)
	reportController.Route()
	request, err := http.NewRequest("GET", "/pmh-api/v1/get/reportuserid?userId=09effbb3-34fe-4719-a1f6-33619f926577", nil)
	t.NoError(err)

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request
	ctx.Set("TEAM MEMBER", false) // i change this akses to TEAM MEMBER
	reportController.GetReportByUserIdController(ctx)
	t.NotEqual(http.StatusUnauthorized, record.Code)
}

func (t *ReportControllerTestSuite) TestDeleteReportByIdController() {
	t.ReportUc.On("DeleteReportById", ExpectedReport.Id).Return(nil)
	reportController := NewReportController(t.ReportUc, t.authMiddleware, t.rg)
	reportController.Route()
	request, err := http.NewRequest("DELETE", "/pmh-api/v1/delete/report?id=ed09d2f3-1878-4e11-adaf-a14326c81657", nil)
	t.NoError(err)

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request
	ctx.Set("ADMIN", true)
	reportController.DeleteReportByIdController(ctx)
	t.Equal(http.StatusOK, record.Code)
}

func (t *ReportControllerTestSuite) TestDeleteReportByIdController_failed() {
	t.ReportUc.On("DeleteReportById", ExpectedReport.Id).Return(nil)
	reportController := NewReportController(t.ReportUc, t.authMiddleware, t.rg)
	reportController.Route()
	request, err := http.NewRequest("DELETE", "/pmh-api/v1/delete/report?id=ed09d2f3-1878-4e11-adaf-a14326c81657", nil)
	t.NoError(err)

	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request
	ctx.Set("TEAM MEMBER", false) // i change this akses to TEAM MEMBER
	reportController.DeleteReportByIdController(ctx)
	t.NotEqual(http.StatusUnauthorized, record.Code)
}

func TestReportControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ReportControllerTestSuite))
}

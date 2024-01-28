package controller

import (
	"net/http"

	"enigma.com/projectmanagementhub/delivery/middleware"
	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/shared/common"
	"enigma.com/projectmanagementhub/usecase"
	"github.com/gin-gonic/gin"
)

type ReportController struct {
	reportUC       usecase.ReportUsecase
	authMiddleware middleware.AuthMiddleware
	rg             *gin.RouterGroup
}

func NewReportController(reportUC usecase.ReportUsecase, authMiddleware middleware.AuthMiddleware, rg *gin.RouterGroup) *ReportController {
	return &ReportController{
		reportUC:       reportUC,
		authMiddleware: authMiddleware,
		rg:             rg,
	}
}

func (h *ReportController) CreateNewReportController(c *gin.Context) {
	var newReport model.Report
	err := c.ShouldBind(&newReport)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newReport, err = h.reportUC.CreateReport(newReport)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	common.SendSingleResponse(c, newReport, "Success")
}

func (h *ReportController) UpdateReportController(c *gin.Context) {
	var updatedReport model.Report
	err := c.ShouldBind(&updatedReport)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	updatedReport, err = h.reportUC.UpdateReport(updatedReport)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "failed to update report "+err.Error())
		return
	}

	common.SendSingleResponse(c, updatedReport, "succesfully updated report")
}

func (h *ReportController) DeleteReportByIdController(c *gin.Context) {
	id := c.Query("id")
	err := h.reportUC.DeleteReportById(id)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "failed to delete report "+err.Error())
		return
	}
	common.SendSingleResponse(c, "is null", "succesfully deleted report")
}

func (h *ReportController) GetReportByTaskIdController(c *gin.Context) {
	taskId := c.Query("taskId")

	reports, err := h.reportUC.GetReportByTaskId(taskId)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, "failed to get report by task id "+err.Error())
		return
	}
	var response []interface{}
	for _, v := range reports {
		response = append(response, v)
	}
	common.SendSingleResponse(c, response, "Succes to get report by task id")

}

func (h *ReportController) GetReportByUserIdController(c *gin.Context) {
	userId := c.Query("userId")

	reports, err := h.reportUC.GetReportByUserId(userId)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, "failed to get report by user id "+err.Error())
		return
	}
	var response []interface{}
	for _, v := range reports {
		response = append(response, v)
	}
	common.SendSingleResponse(c, response, "Succes to get report by user id")

}

// rg meng group end-point2
func (h *ReportController) Route() {
	h.rg.GET("/get/reporttaskid", h.authMiddleware.RequireToken("ADMIN", "MANAGER"), h.GetReportByTaskIdController)
	h.rg.GET("/get/reportuserid", h.authMiddleware.RequireToken("ADMIN", "MANAGER"), h.GetReportByUserIdController)
	h.rg.POST("/createreport", h.authMiddleware.RequireToken("TEAM MEMBER"), h.CreateNewReportController)
	h.rg.PUT("/updatereport", h.authMiddleware.RequireToken("TEAM MEMBER"), h.UpdateReportController)
	h.rg.DELETE("/deletedreport", h.authMiddleware.RequireToken("ADMIN"), h.DeleteReportByIdController)
}

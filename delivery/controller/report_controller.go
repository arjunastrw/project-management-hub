package controller

import (
	"net/http"

	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/usecase"
	"github.com/gin-gonic/gin"
)

type ReportController struct {
	reportUC usecase.ReportUsecase
	rg       *gin.RouterGroup
}

func NewReportController(reportUC usecase.ReportUsecase, rg *gin.RouterGroup) *ReportController {
	return &ReportController{
		reportUC: reportUC,
		rg:       rg,
	}
}

func (h *ReportController) CreateNewReportController(c *gin.Context) {
	var newReport model.Report
	err := c.ShouldBind(&newReport)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newReport, err = h.reportUC.CreateReport(newReport)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    http.StatusCreated,
		"message": "created report succesfully",
		"Data":    newReport,
	})
}

func (h *ReportController) UpdateReportController(c *gin.Context) {
	var updatedReport model.Report
	err := c.ShouldBind(&updatedReport)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedReport, err = h.reportUC.UpdateReport(updatedReport)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "failed to update report",
			"error":   err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "succesfully updated report",
		"data":    updatedReport,
	})
}

func (h *ReportController) DeleteReportByIdController(c *gin.Context) {
	id := c.Query("id")
	err := h.reportUC.DeleteReportById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "failed to delete task",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "succesfully deleted tsak",
	})
}

func (h *ReportController) GetReportByTaskIdController(c *gin.Context) {
	taskId := c.Query("taskId")

	reports, err := h.reportUC.GetReportByTaskId(taskId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	var response []interface{}
	for _, v := range reports {
		response = append(response, v)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"code":    http.StatusOK,
			"message": "OK",
		},
		"data": response,
	})

}

func (h *ReportController) GetReportByUserIdController(c *gin.Context) {
	userId := c.Query("userId")

	reports, err := h.reportUC.GetReportByTaskId(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	var response []interface{}
	for _, v := range reports {
		response = append(response, v)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"code":    http.StatusOK,
			"message": "OK",
		},
		"data": response,
	})

}

// rg meng group end-point2
func (h *ReportController) Route() {
	h.rg.GET("/get/")
	h.rg.GET("/get/reportuserid", h.GetReportByUserIdController)
	h.rg.POST("/createreport", h.CreateNewReportController)
	h.rg.PUT("/updatereport", h.UpdateReportController)
	h.rg.DELETE("/delettask", h.DeleteReportByIdController)
}

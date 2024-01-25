package controller

import (
	"net/http"

	"enigma.com/projectmanagementhub/model"
	"enigma.com/projectmanagementhub/usecase"
	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	reportUC usecase.ReportUsecase
	rg       *gin.RouterGroup
}

func NewReportHandler(reportUC usecase.ReportUsecase, rg *gin.RouterGroup) *ReportHandler {
	return &ReportHandler{
		reportUC: reportUC,
		rg:       rg,
	}
}

func (h *ReportHandler) CreateNewReportHandler(c *gin.Context) {
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

func (h *ReportHandler) UpdateReportHandler(c *gin.Context) {
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

func (h *ReportHandler) DeleteReportByIdHandler(c *gin.Context) {
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

func (h *ReportHandler) GetReportByTaskIdHandler(c *gin.Context) {
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

func (h *ReportHandler) GetReportByUserIdHandler(c *gin.Context) {
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
func (h *ReportHandler) Route() {
	h.rg.GET("/get/")
	h.rg.GET("/get/reportuserid", h.GetReportByUserIdHandler)
	h.rg.POST("/createreport", h.CreateNewReportHandler)
	h.rg.PUT("/updatereport", h.UpdateReportHandler)
	h.rg.DELETE("/delettask", h.DeleteReportByIdHandler)
}

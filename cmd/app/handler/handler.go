package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CrawlerService interface {
	GetData(ctx context.Context, datType string) ([]map[string]interface{}, error)
	GetDataByYear(ctx context.Context, datType string, year int) ([]map[string]interface{}, error)
	GetDataByYearMonth(ctx context.Context, datType string, year, month int) ([]map[string]interface{}, error)
}

type Handler struct {
	service CrawlerService
}

func NewHandler(crawler CrawlerService) *Handler {
	return &Handler{
		service: crawler,
	}
}

func (h *Handler) GetAllData(context *gin.Context) {
	dataType := context.Param("dataType")
	result, err := h.service.GetData(context, dataType)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, result)
}

func (h *Handler) GetDataByYear(context *gin.Context) {
	dataType := context.Param("dataType")
	year, _ := strconv.Atoi(context.Param("year"))

	result, err := h.service.GetDataByYear(context, dataType, year)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, result)
}

func (h *Handler) GetDataByYearMonth(context *gin.Context) {
	dataType := context.Param("dataType")
	year, _ := strconv.Atoi(context.Param("year"))
	month, _ := strconv.Atoi(context.Param("month"))

	result, err := h.service.GetDataByYearMonth(context, dataType, year, month)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, result)
}

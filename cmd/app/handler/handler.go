package handler

import (
	"net/http"

	"github.com/danielcesario/sspsp-crawler/internal/crawler"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Crawler crawler.Crawler
}

func NewHandler(crawler crawler.Crawler) *Handler {
	return &Handler{
		Crawler: crawler,
	}
}

func (h *Handler) GetAllData(context *gin.Context) {
	dataType := context.Param("dataType")
	result, err := h.Crawler.GetData(context, dataType)
	if err != nil {
		context.JSON(http.StatusBadRequest, nil)
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, result)
}

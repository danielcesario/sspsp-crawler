package main

import (
	"github.com/danielcesario/sspsp-crawler/cmd/app/handler"
	"github.com/danielcesario/sspsp-crawler/internal/crawler"
	"github.com/gin-gonic/gin"
)

func main() {
	service := crawler.NewService()
	handler := handler.NewHandler(service)

	r := gin.Default()
	r.GET("/api/ssp/sp/:dataType", handler.GetAllData)
	r.GET("/api/ssp/sp/:dataType/:year", handler.GetDataByYear)
	r.GET("/api/ssp/sp/:dataType/:year/:month", handler.GetDataByYearMonth)

	r.Run(":8080")
}

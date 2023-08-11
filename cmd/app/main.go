package main

import (
	"github.com/danielcesario/sspsp-crawler/cmd/app/handler"
	"github.com/danielcesario/sspsp-crawler/internal/crawler"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly/v2"
)

func main() {
	collector := colly.NewCollector()
	crawler := crawler.NewCrawler(collector)
	handler := handler.NewHandler(crawler)

	r := gin.Default()
	r.GET("/api/ssp/sp/:dataType", handler.GetAllData)

	r.Run(":8080")

}

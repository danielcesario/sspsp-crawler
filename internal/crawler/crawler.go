package crawler

import (
	"context"
	"fmt"

	crawler "github.com/danielcesario/sspsp-crawler/internal/crawler/datasource"
	"github.com/gocolly/colly/v2"
)

type CrawlerDatasource interface {
	GetAllData(ctx context.Context) ([]map[string]interface{}, error)
}

type Crawler struct {
	Collector *colly.Collector
}

func NewCrawler(collector *colly.Collector) Crawler {
	return Crawler{
		Collector: collector,
	}
}

func (c *Crawler) GetData(ctx context.Context, datType string) ([]map[string]interface{}, error) {
	switch datType {
	case "violencia-contra-mulher":
		return crawler.NewCollectorViolenceAgainstWomen(c.Collector).GetAllData(ctx)
	default:
		return nil, fmt.Errorf("unknown data type")
	}
}

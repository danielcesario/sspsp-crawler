package crawler

import (
	"context"
	"fmt"

	"github.com/danielcesario/sspsp-crawler/internal/crawler/datasource"
)

type CrawlerDatasource interface {
	GetAllData(ctx context.Context) ([]map[string]interface{}, error)
	GetDataByYear(ctx context.Context, year int) ([]map[string]interface{}, error)
}

type Crawler struct {
}

func NewService() *Crawler {
	return &Crawler{}
}

func (c *Crawler) GetData(ctx context.Context, datType string) ([]map[string]interface{}, error) {
	switch datType {
	case "violencia-contra-mulher":
		return datasource.NewCollectorViolenceAgainstWomen().GetAllData(ctx)
	default:
		return nil, fmt.Errorf("unknown data type")
	}
}

func (c *Crawler) GetDataByYear(ctx context.Context, datType string, year int) ([]map[string]interface{}, error) {
	switch datType {
	case "violencia-contra-mulher":
		return datasource.NewCollectorViolenceAgainstWomen().GetDataByYear(ctx, year)
	default:
		return nil, fmt.Errorf("unknown data type")
	}
}

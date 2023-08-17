package datasource

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

type ViolenceAgainstWomen struct {
}

func NewCollectorViolenceAgainstWomen() *ViolenceAgainstWomen {
	return &ViolenceAgainstWomen{}
}

type Event struct {
	EventType     string `json:"event"`
	TotalCapital  int    `json:"total_capital"`
	TotalDemacro  int    `json:"total_demacro"`
	TotalInterior int    `json:"total_interior"`
	Total         int    `json:"total"`
}

type Report struct {
	Month  int     `json:"month"`
	Year   int     `json:"year"`
	Events []Event `json:"events"`
}

func (v *ViolenceAgainstWomen) GetAllData(ctx context.Context) ([]map[string]interface{}, error) {
	collector := colly.NewCollector(
		colly.MaxDepth(5),
		colly.Async(true),
	)

	collector.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 5})

	var AllReports []Report
	collector.OnHTML("div[id^=conteudo_repPeriodo_divPeriodo]", func(div *colly.HTMLElement) {

		divId := div.Attr("id")
		period, _ := getPeriod(divId)
		month, year, err := getMonthYear(*period, div)
		if err != nil {
			fmt.Println("Error getting month year", err)
		}

		MonthReport := Report{
			Month: *month,
			Year:  *year,
		}

		var events []Event
		div.ForEach("table[id^=conteudo_repPeriodo_grdOcorrencias]", func(index int, table *colly.HTMLElement) {

			table.ForEach("tr", func(rowIndex int, row *colly.HTMLElement) {
				if rowIndex == 0 {
					return
				}

				var event Event
				row.ForEach("td", func(cellIndex int, cell *colly.HTMLElement) {
					switch cellIndex {
					case 0:
						event.EventType = strings.TrimSpace(cell.Text)
					case 1:
						event.TotalCapital, _ = strconv.Atoi(strings.TrimSpace(cell.Text))
					case 2:
						event.TotalDemacro, _ = strconv.Atoi(strings.TrimSpace(cell.Text))
					case 3:
						event.TotalInterior, _ = strconv.Atoi(strings.TrimSpace(cell.Text))
					case 4:
						event.Total, _ = strconv.Atoi(strings.TrimSpace(cell.Text))
					}
				})

				events = append(events, event)
			})
		})

		MonthReport.Events = events
		AllReports = append(AllReports, MonthReport)
	})

	collector.Visit("http://www.ssp.sp.gov.br/Estatistica/ViolenciaMulher.aspx")
	collector.Wait()

	var err error
	collector.OnError(func(r *colly.Response, errorReceived error) {
		err = errorReceived
	})

	if err != nil {
		fmt.Println("Visit", err.Error())
		return nil, err
	}

	inrec, err := json.Marshal(AllReports)
	if err != nil {
		fmt.Println("Marshal", err.Error())
		return nil, err
	}

	var inInterface []map[string]interface{}
	err = json.Unmarshal(inrec, &inInterface)
	if err != nil {
		fmt.Println("Unmarshal", err.Error())
		return nil, err
	}

	return inInterface, nil
}

func (v *ViolenceAgainstWomen) GetDataByYear(ctx context.Context, year int) ([]map[string]interface{}, error) {
	collector := colly.NewCollector(
		colly.MaxDepth(5),
		colly.Async(true),
	)

	collector.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 5})

	var AllReports []Report
	collector.OnHTML("div[id^=conteudo_repPeriodo_divPeriodo]", func(div *colly.HTMLElement) {

		divId := div.Attr("id")
		period, _ := getPeriod(divId)
		month, currentYear, err := getMonthYear(*period, div)
		if err != nil {
			fmt.Println("Error getting month year", err)
		}

		if *currentYear != year {
			return
		}

		MonthReport := Report{
			Month: *month,
			Year:  *currentYear,
		}

		var events []Event
		div.ForEach("table[id^=conteudo_repPeriodo_grdOcorrencias]", func(index int, table *colly.HTMLElement) {

			table.ForEach("tr", func(rowIndex int, row *colly.HTMLElement) {
				if rowIndex == 0 {
					return
				}

				var event Event
				row.ForEach("td", func(cellIndex int, cell *colly.HTMLElement) {
					switch cellIndex {
					case 0:
						event.EventType = strings.TrimSpace(cell.Text)
					case 1:
						event.TotalCapital, _ = strconv.Atoi(strings.TrimSpace(cell.Text))
					case 2:
						event.TotalDemacro, _ = strconv.Atoi(strings.TrimSpace(cell.Text))
					case 3:
						event.TotalInterior, _ = strconv.Atoi(strings.TrimSpace(cell.Text))
					case 4:
						event.Total, _ = strconv.Atoi(strings.TrimSpace(cell.Text))
					}
				})

				events = append(events, event)
			})
		})

		MonthReport.Events = events
		AllReports = append(AllReports, MonthReport)

	})

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("OnRequest", r)
	})

	collector.Visit("http://www.ssp.sp.gov.br/Estatistica/ViolenciaMulher.aspx")
	collector.Wait()

	var err error
	collector.OnError(func(r *colly.Response, errorReceived error) {
		err = errorReceived
	})

	if err != nil {
		return nil, err
	}

	inrec, err := json.Marshal(AllReports)
	if err != nil {
		return nil, err
	}

	var inInterface []map[string]interface{}
	err = json.Unmarshal(inrec, &inInterface)
	if err != nil {
		return nil, err
	}

	return inInterface, nil
}

func (v *ViolenceAgainstWomen) GetDataByYearMonth(ctx context.Context, year, month int) ([]map[string]interface{}, error) {
	collector := colly.NewCollector(
		colly.MaxDepth(5),
		colly.Async(true),
	)

	collector.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 5})

	var AllReports []Report
	collector.OnHTML("div[id^=conteudo_repPeriodo_divPeriodo]", func(div *colly.HTMLElement) {

		divId := div.Attr("id")
		period, _ := getPeriod(divId)
		currentMonth, currentYear, err := getMonthYear(*period, div)
		if err != nil {
			fmt.Println("Error getting month year", err)
		}

		if *currentYear != year {
			return
		}

		if *currentMonth != month {
			return
		}

		MonthReport := Report{
			Month: *currentMonth,
			Year:  *currentYear,
		}

		var events []Event
		div.ForEach("table[id^=conteudo_repPeriodo_grdOcorrencias]", func(index int, table *colly.HTMLElement) {

			table.ForEach("tr", func(rowIndex int, row *colly.HTMLElement) {
				if rowIndex == 0 {
					return
				}

				var event Event
				row.ForEach("td", func(cellIndex int, cell *colly.HTMLElement) {
					switch cellIndex {
					case 0:
						event.EventType = strings.TrimSpace(cell.Text)
					case 1:
						event.TotalCapital, _ = strconv.Atoi(strings.TrimSpace(cell.Text))
					case 2:
						event.TotalDemacro, _ = strconv.Atoi(strings.TrimSpace(cell.Text))
					case 3:
						event.TotalInterior, _ = strconv.Atoi(strings.TrimSpace(cell.Text))
					case 4:
						event.Total, _ = strconv.Atoi(strings.TrimSpace(cell.Text))
					}
				})

				events = append(events, event)
			})
		})

		MonthReport.Events = events
		AllReports = append(AllReports, MonthReport)

	})

	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("OnRequest", r)
	})

	collector.Visit("http://www.ssp.sp.gov.br/Estatistica/ViolenciaMulher.aspx")
	collector.Wait()

	var err error
	collector.OnError(func(r *colly.Response, errorReceived error) {
		err = errorReceived
	})

	if err != nil {
		return nil, err
	}

	inrec, err := json.Marshal(AllReports)
	if err != nil {
		return nil, err
	}

	var inInterface []map[string]interface{}
	err = json.Unmarshal(inrec, &inInterface)
	if err != nil {
		return nil, err
	}

	return inInterface, nil
}

func getPeriod(divId string) (*string, error) {
	regex := regexp.MustCompile(`\d{1,3}$`)
	match := regex.FindStringSubmatch(divId)

	if len(match) >= 1 {
		result := match[0]
		return &result, nil
	} else {
		return nil, fmt.Errorf("não foi possível extrair o número")
	}
}

func getMonthYear(period string, div *colly.HTMLElement) (*int, *int, error) {
	dataStr := div.ChildText(fmt.Sprintf("span[id=conteudo_repPeriodo_lblPeriodo_%v]", period))
	regex := regexp.MustCompile(`mês: (\p{L}+) de (\d{4})`)
	match := regex.FindStringSubmatch(dataStr)

	if len(match) >= 3 {
		month := match[1]

		monthNumber, err := getMonthNumber(month)
		if err != nil {
			fmt.Println("Error getMonthNumber", match[1])
			return nil, nil, err
		}

		year, err := strconv.Atoi(match[2])
		if err != nil {
			fmt.Println("Error convert", match[2])
			return nil, nil, err
		}

		return &monthNumber, &year, nil
	} else {
		return nil, nil, fmt.Errorf("não foi possível extrair os valores")
	}
}

func getMonthNumber(monthName string) (int, error) {
	monthMapping := map[string]int{
		"Janeiro":   1,
		"Fevereiro": 2,
		"Março":     3,
		"Abril":     4,
		"Maio":      5,
		"Junho":     6,
		"Julho":     7,
		"Agosto":    8,
		"Setembro":  9,
		"Outubro":   10,
		"Novembro":  11,
		"Dezembro":  12,
	}

	number, found := monthMapping[monthName]
	if !found {
		return 0, fmt.Errorf("mês não encontrado: %s", monthName)
	}

	return number, nil
}

package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

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

func main() {

	c := colly.NewCollector()

	//c.OnHTML("div[id=conteudo_repPeriodo_divPeriodo_5]", func(div *colly.HTMLElement) {
	c.OnHTML("div[id^=conteudo_repPeriodo_divPeriodo]", func(div *colly.HTMLElement) {

		divId := div.Attr("id")
		period, _ := getPeriod(divId)
		month, year, err := getMonthYear(*period, div)
		if err != nil {
			fmt.Println("Error getting month year", err)
		}

		var events []Event

		div.ForEach("table[id^=conteudo_repPeriodo_grdOcorrencias]",
			func(index int, table *colly.HTMLElement) {

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

		var MonthReport Report = Report{
			Month:  *month,
			Year:   *year,
			Events: events,
		}

		MonthReportJson, err := json.Marshal(MonthReport)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(MonthReportJson))
	})

	c.Visit("http://www.ssp.sp.gov.br/Estatistica/ViolenciaMulher.aspx")

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

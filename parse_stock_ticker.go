package parse_stock_ticker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// TickerData is the data returned from the API
type TickerData struct {
	Chart Chart `json:"chart"`
}

// Chart is the chart data from the API
type Chart struct {
	Result []Result `json:"result"`
}

// Result is the result data from the API
type Result struct {
	Meta Meta `json:"meta"`
	Timestamp []int `json:"timestamp"`
	Indicators Indicators `json:"indicators"`
}

// Meta is the meta data from the API
type Meta struct {
	Currency string `json:"currency"`
	Symbol string `json:"symbol"`
	ExchangeName string `json:"exchangeName"`
	InstrumentType string `json:"instrumentType"`
	FirstTradeDate int `json:"firstTradeDate"`
	RegularMarketTime int `json:"regularMarketTime"`
	Gmtoffset int `json:"gmtoffset"`
	Timezone string `json:"timezone"`
	ExchangeTimezoneName string `json:"exchangeTimezoneName"`
	RegularMarketPrice float64 `json:"regularMarketPrice"`
	ChartPreviousClose float64 `json:"chartPreviousClose"`
	PreviousClose float64 `json:"previousClose"`
	Scale float64 `json:"scale"`
	PriceHint float64 `json:"priceHint"`
}

// Indicators is the indicators data from the API
type Indicators struct {
	Quote []Quote `json:"quote"`
}

// Quote is the quote data from the API
type Quote struct {
	Close []float64 `json:"close"`
}

// ParseStockTicker takes a stock ticker and returns the difference in price
func ParseStockTicker(ticker string) (float64, error) {
	url := fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s?region=US&lang=en-US&includePrePost=false&interval=2m&range=1d&corsDomain=finance.yahoo.com&.tsrc=finance", ticker)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var data TickerData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0, err
	}

	return data.Chart.Result[0].Meta.RegularMarketPrice - data.Chart.Result[0].Meta.ChartPreviousClose, nil
}

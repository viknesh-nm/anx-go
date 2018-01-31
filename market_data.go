package anx

import (
	"fmt"

	request "github.com/viknesh-nm/http-request"
)

// Ticker holds general ticker data
type Ticker struct {
	Result string     `json:"result"`
	Data   TickerData `json:"data"`
}

// TickerData holds specific Ticker data
type TickerData struct {
	High   TickerStatData `json:"high"`
	Low    TickerStatData `json:"low"`
	Avg    TickerStatData `json:"avg"`
	Last   TickerStatData `json:"last"`
	Buy    TickerStatData `json:"buy"`
	Sell   TickerStatData `json:"sell"`
	Volume TickerStatData `json:"vol"`
	Vwap   TickerStatData `json:"vwap"`
	Now    string         `json:"now"`
	Time   string         `json:"dataUpdateTime"`
}

// TickerStatData holds Ticker statistics data
type TickerStatData struct {
	Currency     string `json:"currency"`
	Display      string `json:"display"`
	DisplayShort string `json:"display_short"`
	Value        string `json:"value"`
	ValueInt     string `json:"value_int"`
}

// OrderBook holds general orderbook data
type OrderBook struct {
	Result string        `json:"result"`
	Data   OrderBookData `json:"data"`
}

// OrderBookData holds specific orderbook data
type OrderBookData struct {
	Now  string              `json:"now"`
	Time string              `json:"dataUpdateTime"`
	Asks []OrderBookStatData `json:"asks"`
	Bids []OrderBookStatData `json:"bids"`
}

// OrderBookStatData holds price and amount of either Asks or Bids
type OrderBookStatData struct {
	Price  string `json:"price"`
	Amount string `json:"amount"`
}

// GetTicker fetches and returns ticker
func (a *Anx) GetTicker(symbol string) (Ticker, error) {
	ticker := Ticker{}
	a.BaseURL = fmt.Sprintf("%s/%s/%s", v2APIURL, symbol, tickerEndpoint)
	err := request.DoGetHTTPRequest(a.BaseURL, &ticker)
	if err != nil {
		return ticker, err
	}
	return ticker, nil
}

// GetOrderBook fetches and returns orderbook
func (a *Anx) GetOrderBook(symbol string) (OrderBook, error) {
	orderbook := OrderBook{}
	a.BaseURL = fmt.Sprintf("%s/%s/%s", v2APIURL, symbol, orderbookEndpoint)
	err := request.DoGetHTTPRequest(a.BaseURL, &orderbook)
	if err != nil {
		return orderbook, err
	}
	return orderbook, nil
}

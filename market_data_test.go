package anx

import (
	"log"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var err error

// TestGetTicker tests the functionality of GetTicker
func TestGetTicker(t *testing.T) {
	log.Println("TestGetTicker - anx.GetTicker")
	var (
		ticker     Ticker
		tickerStat TickerStatData
		field      string
		// fetchTicker fetches market data ticker
		fetchTicker = func(symbol string) {
			client = New("", "")
			ticker, err = client.GetTicker(symbol)
		}
		fields = []string{"High", "Low", "Avg", "Buy", "Last", "Volume", "Vwap", "Sell"}
	)

	Convey("GetTicker should return Ticker and error", t, func() {
		Convey("First value should be of type Ticker", func() {
			fetchTicker("BTCUSD")
			So(err, ShouldBeNil)
			So(ticker.Result, ShouldEndWith, "success")
		})
		Convey("Second value should be of type error. Blank symbols are invalid and throw an error", func() {
			fetchTicker("")
			So(err, ShouldNotBeNil)
			So(err, ShouldImplement, (*error)(nil))
		})
	})
	Convey("Print successfully fetched ticker data", t, func() {
		log.Println("Anx Ticker Data")
		currency := []string{"BTCUSD", "LTCBTC", "LTCUSD"}
		for _, value := range currency {
			fetchTicker(value)
			if err != nil {
				t.Error(err)
				return
			}
			log.Println(value, ":")
			for _, field = range fields {
				tickerStat = (client.getField(ticker.Data, field)).(TickerStatData)
				log.Println(field, ":", tickerStat.Value)
			}
			log.Println("Time:", ticker.Data.Time)
		}
	})
}

// TestGetOrderBook tests the functionality of GetOrderBook
func TestGetOrderBook(t *testing.T) {
	log.Println("TestGetOrderBook - anx.GetOrderBook")
	var (
		ob, testOB    OrderBook
		fieldStat     []OrderBookStatData
		orderbookStat OrderBookStatData
		field         string
		// fetchOB fetches market data orderbook
		fetchOB = func(symbol string) {
			client = New("", "")
			ob, err = client.GetOrderBook(symbol)
		}
	)

	fields := []string{"Asks", "Bids"}
	Convey("GetOrderBook should return OrderBook and error", t, func() {
		Convey("First value should be of type OrderBook", func() {
			fetchOB("BTCUSD")
			So(ob, ShouldNotResemble, testOB)
			So(err, ShouldBeNil)
			So(ob.Result, ShouldNotBeEmpty)
		})
		Convey("Second value should be of type error. Blank symbols are invalid and throw an error", func() {
			fetchOB("")
			So(err, ShouldNotBeNil)
			So(err, ShouldImplement, (*error)(nil))
			So(ob, ShouldResemble, testOB)
		})
	})
	Convey("Print successfully fetched orderbook data", t, func() {
		log.Println("Anx OrderBook Data")
		currency := []string{"BTCUSD", "LTCUSD", "LTCBTC"}
		for _, value := range currency {
			fetchOB(value)
			if err != nil {
				return
			}
			log.Println(value, ":")
			for _, field = range fields {
				log.Println(field, ":")
				fieldStat = (client.getField(ob.Data, field)).([]OrderBookStatData)
				for _, orderbookStat = range fieldStat {
					log.Println("Price:", orderbookStat.Price)
					log.Println("Amount:", orderbookStat.Amount)
				}
			}
		}
	})
}

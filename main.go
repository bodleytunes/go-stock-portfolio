package main

import (
	"fmt"

	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

type PortfolioItem struct {
	Name       string
	Allocation int
}

type Basket struct {
	weightslicevalue float32
	sumofweights     float32
	percentage       int
	totalexpenditure float32

	stocks []Stock
}

type Stock struct {
	symbol         string
	weight         int
	purchaseAmount float32
	percentage     float32
	closePrice     float32
	purchasePrice  float32
}

func main() {

	spy, _ := quote.NewQuoteFromYahoo("spy", "2020-01-01", "2021-04-10", quote.Daily, true)
	fmt.Print(spy.CSV())
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)

	portfolioItems := map[int][]PortfolioItem{
		1: []PortfolioItem{PortfolioItem{"AAPL", 35}},
		2: []PortfolioItem{PortfolioItem{"AMZN", 95}},
		3: []PortfolioItem{PortfolioItem{"MSFT", 83}},
		4: []PortfolioItem{PortfolioItem{"MSFT", 83}},
	}

	_ = portfolioItems

	var b = Basket{}
	_ = b

}

func New(stocks []Stock) *Basket {

	// return instance of basket
	return &Basket{
		stocks: stocks,
	}

}

func Empty() *Basket {
	return New([]Stock{})
}

func (b *Basket) Add(stock Stock) {
	b.stocks = append(b.stocks, stock)
}

func (b *Basket) WeightSliceValue(weightslicevalue float32) {
	b.weightslicevalue = weightslicevalue
}

func (b *Basket) SumOfWeights(sumofweights float32) {
	b.sumofweights = sumofweights
}

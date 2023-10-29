package service

import (
	"math"
	"sort"

	"github.com/gleison/peculi/internal/model"
)

func AveragePrice(trades []model.Trade) float32 {
	sort.Sort(ByDate(trades))
	var totalQuantity uint16 = 0
	var totalCost float32 = 0.0
	for _, trade := range trades {
		switch trade.Operation {
		case model.BUY:
			totalCost += trade.Cost
			totalQuantity += trade.Quantity
		case model.SELL:
			partialAveragePrice := totalCost / float32(totalQuantity)
			totalCost -= float32(trade.Quantity) * partialAveragePrice
			totalQuantity -= trade.Quantity
		}
	}
	averagePrice := totalCost / float32(totalQuantity)
	return roundTwoDecimals(averagePrice)
}

func roundTwoDecimals(f float32) float32 {
	return float32(math.Round(float64(f)*100) / 100)
}

// ByDate implements sort.Interface based on the Date field
type ByDate []model.Trade

func (trades ByDate) Len() int {
	return len(trades)
}

func (trades ByDate) Less(i, j int) bool {
	comparation := trades[i].Date.Compare(trades[j].Date)
	return comparation <= 0
}

func (trades ByDate) Swap(i, j int) {
	trades[i], trades[j] = trades[j], trades[i]
}

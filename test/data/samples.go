package test

import (
	"time"

	"github.com/gleison/peculi/internal/model"
)

var Trades = []model.Trade{
	{Id: 1, Date: time.Date(2020, 10, 8, 0, 0, 0, 0, time.Local), Ticker: "O", Operation: model.BUY, Quantity: 1, Cost: 63.23},
	{Id: 2, Date: time.Date(2020, 6, 10, 0, 0, 0, 0, time.Local), Ticker: "O", Operation: model.BUY, Quantity: 5, Cost: 316.45},
	{Id: 3, Date: time.Date(2020, 8, 6, 0, 0, 0, 0, time.Local), Ticker: "O", Operation: model.BUY, Quantity: 1, Cost: 61.97},
	{Id: 4, Date: time.Date(2020, 5, 21, 0, 0, 0, 0, time.Local), Ticker: "O", Operation: model.BUY, Quantity: 10, Cost: 523.60},
	{Id: 5, Date: time.Date(2020, 7, 6, 0, 0, 0, 0, time.Local), Ticker: "O", Operation: model.BUY, Quantity: 1, Cost: 62.95},
	{Id: 6, Date: time.Date(2021, 1, 7, 0, 0, 0, 0, time.Local), Ticker: "O", Operation: model.SELL, Quantity: 8, Cost: 470.96},
	{Id: 7, Date: time.Date(2021, 2, 8, 0, 0, 0, 0, time.Local), Ticker: "O", Operation: model.BUY, Quantity: 10, Cost: 575.40},
}

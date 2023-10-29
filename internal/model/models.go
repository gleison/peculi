package model

import (
	"time"
)

type Trade struct {
	Id        int64
	Date      time.Time
	Ticker    string
	Operation Operation
	Quantity  uint16
	Cost      float32
}

// Equivalent to a ENUM
type Operation int

const (
	BUY Operation = iota
	SELL
)

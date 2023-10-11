package domain

import (
	"time"
)

type Transaction struct {
	ID     string
	Title  string
	Price  float64
	UserID string
	Rate   float64
	Date   time.Time
}

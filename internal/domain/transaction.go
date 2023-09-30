package domain

import (
	"time"
)

type Transaction struct {
	ID     string
	Title  string
	Price  float64
	UserID string
	Date   time.Time
}

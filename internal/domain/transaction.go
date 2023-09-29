package domain

import (
	"time"
)

type Transaction struct {
	ID     string
	Title  string
	Price  int
	UserID string
	Date   time.Time
}

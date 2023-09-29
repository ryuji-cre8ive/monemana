package domain

import (
	"time"
)

type Transaction struct {
	ID     string
	Name   string
	Price  int
	UserID string
	Date   time.Time
}

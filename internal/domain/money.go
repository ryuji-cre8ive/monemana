package domain

import (
	"time"
)

type Money struct {
	ID     string
	Name   string
	Amount int
	Date   time.Time
}

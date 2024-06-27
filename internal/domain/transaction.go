package domain

import (
	"time"
)

type Transaction struct {
	ID           string
	Title        string
	Price        uint64
	UserID       string
	TargetUserID string
	RoomID       string
	MessageID    string
	CreatedAt    time.Time
	DeletedAt    *time.Time
}

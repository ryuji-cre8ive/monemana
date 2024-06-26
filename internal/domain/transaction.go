package domain

import (
	"time"
)

type Transaction struct {
	ID           string
	Title        string
	Price        uint64
	UserID       string
	TargetUserID string     // 変更: []string から string へ
	RoomID       string     // 追加
	CreatedAt    time.Time  // 追加
	DeletedAt    *time.Time // 追加
}

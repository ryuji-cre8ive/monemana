package domain

import "time"

type User struct {
	ID          string
	DisplayName string
	CreatedAt   time.Time
	DeletedAt   *time.Time
}

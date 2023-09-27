package stores

import (
	"github.com/ryuji-cre8ive/monemana/internal/domain"
	"gorm.io/gorm"
	"time"
)

type (
	WebhookStore interface {
		CreateTransaction(tx *gorm.Tx) error
	}

	webhookStore struct {
		*gorm.DB
	}
)

func (s *webhookStore) CreateTransaction(tx *gorm.Tx) error {
	s.DB.Create(&domain.Transaction{
		ID:    "test",
		Name:  "test",
		Price: 100,
		Date:  time.Now(),
	})
	return nil
}

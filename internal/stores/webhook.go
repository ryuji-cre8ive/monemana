package stores

import (
	"database/sql"
)

type (
	WebhookStore interface {
		CreateTransaction(tx *sql.Tx) error
	}

	webhookStore struct {
		*sql.DB
	}
)

func (s *webhookStore) CreateTransaction(tx *sql.Tx) error {
	return nil
}

package stores

import (
	"database/sql"
)

type Stores struct {
	DB      *sql.DB
	Webhook WebhookStore
}

func New(db *sql.DB) *Stores {
	return &Stores{
		DB:      db,
		Webhook: &webhookStore{db},
	}
}

func (s *Stores) Begin() (*sql.Tx, error) {
	return s.DB.Begin()
}

func (s *Stores) Commit(tx *sql.Tx) error {
	return tx.Commit()
}

func (s *Stores) RollBack(tx *sql.Tx) error {
	return tx.Rollback()
}

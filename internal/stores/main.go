package stores

import (
	"gorm.io/gorm"
)

type Stores struct {
	DB       *gorm.DB
	Webhook  WebhookStore
	Exchange ExchangeStore
}

func New(db *gorm.DB) *Stores {
	return &Stores{
		DB:       db,
		Webhook:  &webhookStore{db},
		Exchange: &exchangeStore{},
	}
}

// func (s *Stores) Begin() (*gorm.Tx, error) {
// 	return s.DB.Begin()
// }

// func (s *Stores) Commit(tx *gorm.Tx) error {
// 	return tx.Commit()
// }

// func (s *Stores) RollBack(tx *gorm.Tx) error {
// 	return tx.Rollback()
// }

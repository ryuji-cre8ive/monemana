package stores

import (
	"github.com/ryuji-cre8ive/monemana/internal/domain"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
	"time"
)

type (
	WebhookStore interface {
		CreateTransaction(tx *gorm.Tx) error
		GetUser(userID string) (*domain.User, error)
		GetAllUsers() ([]*domain.User, error)
		CreateUser(tx *gorm.Tx, userID string, userName string) error
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

func (s *webhookStore) GetUser(userID string) (*domain.User, error) {
	user := &domain.User{}
	result := s.DB.First(&user, "id = ?", userID)
	if result.Error != nil {
		return nil, xerrors.Errorf("get user err%w", result.Error)
	}
	return user, nil
}

func (s *webhookStore) GetAllUsers() ([]*domain.User, error) {
	users := make([]*domain.User, 0, 0)

	result := s.DB.Find(&users)
	if result.Error != nil {
		return nil, xerrors.Errorf("get all users err%w", result.Error)
	}
	// エイリアス "u" を指定して結合
	// s.DB.Preload("Transactions").Joins("INNER JOIN transactions ON users.id = transactions.user_id").First(&user)
	// fmt.Println("user no transaction", user.Transactions)
	return users, nil
}

func (s *webhookStore) CreateUser(tx *gorm.Tx, userID string, userName string) error {
	result := s.DB.Create(&domain.User{
		ID:           userID,
		Name:         userName,
		Transactions: nil,
	})
	if result.Error != nil {
		return xerrors.Errorf("create user err%w", result.Error)
	}
	return nil
}

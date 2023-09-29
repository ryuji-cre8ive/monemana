package stores

import (
	"github.com/google/uuid"
	"github.com/ryuji-cre8ive/monemana/internal/domain"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
	"time"
)

type (
	WebhookStore interface {
		CreateTransaction(tx *gorm.Tx, name string, price int, userID string) error
		GetUser(userID string) (*domain.User, error)
		GetAllUsers() ([]*domain.User, error)
		CreateUser(tx *gorm.Tx, userID string, userName string) (*domain.User, error)
	}

	webhookStore struct {
		*gorm.DB
	}
)

func (s *webhookStore) CreateTransaction(tx *gorm.Tx, title string, price int, userID string) error {
	uuid := uuid.Must(uuid.NewRandom())
	s.DB.Create(&domain.Transaction{
		ID:     uuid.String(),
		Title:  title,
		Price:  price,
		UserID: userID,
		Date:   time.Now(),
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

func (s *webhookStore) CreateUser(tx *gorm.Tx, userID string, userName string) (*domain.User, error) {
	user := &domain.User{
		ID:           userID,
		Name:         userName,
		Transactions: nil,
	}
	result := s.DB.Create(&user)
	if result.Error != nil {
		return nil, xerrors.Errorf("create user err%w", result.Error)
	}
	return user, nil
}

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
		CreateTransaction(tx *gorm.Tx, name string, price float64, userID string, rate float64) error
		GetUser(userID string) (*domain.User, error)
		GetAllUsers() ([]*domain.User, error)
		CreateUser(tx *gorm.Tx, userID string, userName string) (*domain.User, error)
		ChangeUserName(tx *gorm.Tx, userID string, userName string) (*domain.User, error)
		AggregateTransaction() ([]*domain.User, error)
	}

	webhookStore struct {
		*gorm.DB
	}
)

func (s *webhookStore) CreateTransaction(tx *gorm.Tx, title string, price float64, userID string, rate float64) error {
	uuid := uuid.Must(uuid.NewRandom())
	s.DB.Create(&domain.Transaction{
		ID:     uuid.String(),
		Title:  title,
		Price:  price,
		UserID: userID,
		Rate:   rate,
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

func (s *webhookStore) ChangeUserName(tx *gorm.Tx, userID string, userName string) (*domain.User, error) {
	user := &domain.User{}
	result := s.DB.First(&user, "id = ?", userID)
	if result.Error != nil {
		return nil, xerrors.Errorf("get user err%w", result.Error)
	}
	user.Name = userName
	result = s.DB.Save(&user)
	if result.Error != nil {
		return nil, xerrors.Errorf("save user err%w", result.Error)
	}
	return user, nil
}

func (s *webhookStore) AggregateTransaction() ([]*domain.User, error) {
	users := make([]*domain.User, 0, 0)
	result := s.DB.Debug().Preload("Transactions").Joins("LEFT JOIN transactions ON users.id = transactions.user_id").Find(&users)

	uniqueUsers := make(map[string]*domain.User)

	for _, user := range users {
		// ユーザーIDをキーとしてマップに追加
		uniqueUsers[user.ID] = user
	}

	// ユニークなユーザーを格納するスライスを作成
	uniqueUserSlice := make([]*domain.User, 0, len(uniqueUsers))
	for _, user := range uniqueUsers {
		uniqueUserSlice = append(uniqueUserSlice, user)
	}

	if result.Error != nil {
		return nil, xerrors.Errorf("get transactions err%w", result.Error)
	}
	return uniqueUserSlice, nil
}

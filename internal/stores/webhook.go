package stores

import (
	"fmt"
	"time"

	"github.com/ryuji-cre8ive/monemana/internal/domain"
	"golang.org/x/xerrors"
	"gorm.io/gorm"
)

type (
	WebhookStore interface {
		CreateTransaction(tx *gorm.Tx, id string, title string, price uint64, userID string, targetUserID string, roomID string) error
		GetUser(userID string) (*domain.User, error)
		GetAllUsers() ([]*domain.User, error)
		CreateUser(tx *gorm.Tx, userID string, userName string) error
		AggregateTransaction(roomId string) ([]*domain.Transaction, error)
		CreateRoom(tx *gorm.Tx, roomID string) error
		GetRoomById(roomID string) (*domain.Room, error)
		CheckUserExists(userID string) (bool, error)
		CheckRoomExists(roomID string) (bool, error)
		UpdateUserName(tx *gorm.Tx, userID string, userName string) error
	}

	webhookStore struct {
		*gorm.DB
	}
)

func (s *webhookStore) CreateTransaction(tx *gorm.Tx, id string, title string, price uint64, userID string, targetUserID string, roomID string) error {
	s.DB.Create(&domain.Transaction{
		ID:           id,
		Title:        title,
		Price:        price,
		UserID:       userID,
		TargetUserID: targetUserID,
		RoomID:       roomID,
		CreatedAt:    time.Now(),
		DeletedAt:    nil,
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
	user := &domain.User{
		ID:          userID,
		DisplayName: userName,
		CreatedAt:   time.Now(),
		DeletedAt:   nil,
	}
	result := s.DB.Create(&user)
	if result.Error != nil {
		return xerrors.Errorf("create user err%w", result.Error)
	}
	return nil
}

func (s *webhookStore) AggregateTransaction(roomId string) ([]*domain.Transaction, error) {
	transactions := make([]*domain.Transaction, 0, 0)
	result := s.DB.Where("room_id = ?", roomId).Find(&transactions)

	if result.Error != nil {
		return nil, xerrors.Errorf("get transactions err%w", result.Error)
	}
	return transactions, nil
}

func (s *webhookStore) CreateRoom(tx *gorm.Tx, roomID string) error {
	room := &domain.Room{
		ID:          roomID,
		DisplayName: "",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}
	result := s.DB.Create(&room)
	if result.Error != nil {
		return xerrors.Errorf("create room err%w", result.Error)
	}
	return nil
}

func (s *webhookStore) GetRoomById(roomID string) (*domain.Room, error) {
	room := &domain.Room{}
	result := s.DB.First(&room, "id = ?", roomID)
	if result.Error != nil {
		return nil, result.Error
	}
	return room, nil
}

func (s *webhookStore) CheckUserExists(userID string) (bool, error) {
	user := &domain.User{}
	result := s.DB.First(&user, "id = ?", userID)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (s *webhookStore) CheckRoomExists(roomID string) (bool, error) {
	room := &domain.Room{}
	result := s.DB.First(&room, "id = ?", roomID)
	fmt.Printf("result: %+v\n", result)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (s *webhookStore) UpdateUserName(tx *gorm.Tx, userID string, userName string) error {
	user := &domain.User{}
	result := s.DB.First(&user, "id = ?", userID)
	if result.Error != nil {
		return result.Error
	}
	user.DisplayName = userName
	result = s.DB.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

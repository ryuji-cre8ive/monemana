package usecase

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/ryuji-cre8ive/monemana/internal/domain"
	"github.com/ryuji-cre8ive/monemana/internal/stores"
	"golang.org/x/xerrors"
)

type (
	WebhookUsecase interface {
		CreateTransaction(c echo.Context, title string, price uint64, userId string, targetUserId []string, roomId string, messageId string) error
		AggregateTransaction(c echo.Context, roomId string) (string, error)
		GetUser(userId string) (*domain.User, error)
		CreateUser(c echo.Context, userId string, name string) error
		CheckUserExists(c echo.Context, userId string) (bool, error)
		UpdateUserName(c echo.Context, userId string, name string) error
		DeleteTransaction(c echo.Context, roomId string, messageId string) (bool, error)
	}

	webhookUsecase struct {
		stores *stores.Stores
	}
)

func (u *webhookUsecase) AggregateTransaction(c echo.Context, roomId string) (string, error) {
	transactions, err := u.stores.Webhook.AggregateTransaction(roomId)
	aggregateMessage := ""
	if err != nil {
		return "", xerrors.Errorf("aggregate transaction err: %w", err)
	}

	// ユーザーごとの支払いを集計
	userTransactions := make(map[string]map[string]int64)
	for _, transaction := range transactions {
		if userTransactions[transaction.UserID] == nil {
			userTransactions[transaction.UserID] = make(map[string]int64)
		}
		userTransactions[transaction.UserID][transaction.TargetUserID] += int64(transaction.Price)
	}

	// 相互の支払いを差し引き
	finalBalances := make(map[string]map[string]int64)
	for user, targets := range userTransactions {
		for target, amount := range targets {
			if finalBalances[user] == nil {
				finalBalances[user] = make(map[string]int64)
			}
			finalBalances[user][target] = amount - userTransactions[target][user]
		}
	}

	// 結果を出力
	for user, targets := range finalBalances {
		user, err := u.GetUser(user)
		if err != nil {
			return "", err
		}
		for target, balance := range targets {
			if balance > 0 {
				targetUser, err := u.GetUser(target)
				if err != nil {
					return "", err
				}
				aggregateMessage += fmt.Sprintf("%s→%s: %d円\n", targetUser.DisplayName, user.DisplayName, balance)
			}
		}
	}

	// 最後の改行文字を削除
	aggregateMessage = strings.TrimSuffix(aggregateMessage, "\n")

	return aggregateMessage, nil
}

func (u *webhookUsecase) CreateTransaction(c echo.Context, title string, price uint64, userId string, targetUserIds []string, roomId string, messageId string) error {

	isExistRoom, err := u.stores.Webhook.CheckRoomExists(roomId)

	if !isExistRoom || err != nil {
		if err := u.stores.Webhook.CreateRoom(nil, roomId); err != nil {
			return xerrors.Errorf("create room err: %w", err)
		}
	}

	for _, targetUserId := range targetUserIds {
		id := uuid.Must(uuid.NewRandom()).String()
		if err := u.stores.Webhook.CreateTransaction(nil, id, title, price, userId, targetUserId, roomId, messageId); err != nil {
			return xerrors.Errorf("create transaction err: %w", err)
		}
	}
	return nil
}

func (u *webhookUsecase) GetUser(userId string) (*domain.User, error) {
	user, err := u.stores.Webhook.GetUser(userId)
	if err != nil {
		return nil, xerrors.Errorf("get user err: %w", err)
	}
	return user, nil
}

func (u *webhookUsecase) CreateUser(c echo.Context, userId string, name string) error {
	if err := u.stores.Webhook.CreateUser(nil, userId, name); err != nil {
		return xerrors.Errorf("create user err: %w", err)
	}
	return nil
}

func (u *webhookUsecase) CheckUserExists(c echo.Context, userId string) (bool, error) {
	return u.stores.Webhook.CheckUserExists(userId)
}
func (u *webhookUsecase) UpdateUserName(c echo.Context, userId string, name string) error {
	if err := u.stores.Webhook.UpdateUserName(nil, userId, name); err != nil {
		return xerrors.Errorf("update user name err: %w", err)
	}
	return nil
}
func (u *webhookUsecase) DeleteTransaction(c echo.Context, roomId string, messageId string) (bool, error) {
	isDeleted, err := u.stores.Webhook.DeleteTransaction(nil, roomId, messageId)
	if err != nil {
		return false, xerrors.Errorf("delete transaction err: %w", err)
	}
	return isDeleted, nil
}

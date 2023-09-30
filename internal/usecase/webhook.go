package usecase

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/ryuji-cre8ive/monemana/internal/stores"
	"golang.org/x/xerrors"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type (
	WebhookUsecase interface {
		PostWebhook(c echo.Context) error
	}

	webhookUsecase struct {
		stores *stores.Stores
	}
)

func (u *webhookUsecase) PostWebhook(c echo.Context) error {
	pattern := regexp.MustCompile("[\\s　]+")
	Secret := os.Getenv("LINE_BOT_CHANNEL_SECRET")
	Token := os.Getenv("LINE_BOT_CHANNEL_TOKEN")

	bot, botErr := linebot.New(Secret, Token)
	if botErr != nil {
		return xerrors.Errorf("linebot.New error: %w", botErr)
	}
	events, parseErr := bot.ParseRequest(c.Request())
	if parseErr != nil {
		return xerrors.Errorf("bot.ParseRequest error: %w", parseErr)
	}
	for _, event := range events {
		user, err := u.stores.Webhook.GetUser(event.Source.UserID)
		fmt.Println("user", user)
		if err != nil {
			fmt.Println("user not found")
			//モック的にユーザーネームをIDと同一で登録
			if user, err = u.stores.Webhook.CreateUser(nil, event.Source.UserID, event.Source.UserID); err != nil {
				fmt.Println("create user err", err)
			}
		}

		if event.Type == linebot.EventTypeMessage {
			switch event.Message.(type) {
			case *linebot.TextMessage:
				text := event.Message.(*linebot.TextMessage).Text
				if strings.Contains(text, "名前変更") {
					fmt.Println("名前変更desuyooooooooo")
					userName := pattern.Split(text, -1)[1]
					fmt.Println("userName", userName)
					if _, err := u.stores.Webhook.ChangeUserName(nil, user.ID, userName); err != nil {
						return xerrors.Errorf("change user name err: %w", err)
					}
					if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("名前変更したよ！")).Do(); err != nil {
						return xerrors.Errorf("reply message err: %w", err)
					}
					return nil
				}
				if len(pattern.Split(text, -1)) == 2 {
					title, priceStr := pattern.Split(text, -1)[0], pattern.Split(text, -1)[1]
					price, parseIntErr := strconv.ParseFloat(priceStr, 64)
					if parseIntErr != nil {
						return xerrors.Errorf("price parse err: %w", parseIntErr)
					}
					if err := u.stores.Webhook.CreateTransaction(nil, title, price, user.ID); err != nil {
						return xerrors.Errorf("create transaction err: %w", err)
					}
					if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("登録が完了したよ！")).Do(); err != nil {
						fmt.Print(err)
					}
				}

				if strings.Contains(text, "集計") {
					users, err := u.stores.Webhook.AggregateTransaction()
					if err != nil {
						return xerrors.Errorf("aggregate transaction err: %w", err)
					}
					var message string
					aggregate := map[string]float64{}
					for _, user := range users {
						aggregate[string(user.Name)] = 0
						for _, transaction := range *user.Transactions {
							aggregate[string(user.Name)] += transaction.Price
						}
					}

					for name, price := range aggregate {
						fmt.Println("price", price)
						priceStr := fmt.Sprintf("%.2f", price)
						message += name + ": " + priceStr + "RM\n"
					}
					// 最後の改行文字を削除
					if len(message) > 0 {
						message = message[:len(message)-1]
					}
					if _, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message)).Do(); err != nil {
						return xerrors.Errorf("reply message err: %w", err)
					}
				}
			}
		}
	}

	return nil
}

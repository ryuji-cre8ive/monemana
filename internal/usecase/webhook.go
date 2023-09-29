package usecase

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/ryuji-cre8ive/monemana/internal/stores"
	"golang.org/x/xerrors"
	"os"
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
				if len(strings.Split(text, " ")) == 2 {
					title, priceStr := strings.Split(text, " ")[0], strings.Split(text, " ")[1]
					price, parseIntErr := strconv.Atoi(priceStr)
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
				// if strings.Contains(text, "名前変更") {

				// }
			}
		}
	}

	return nil
}

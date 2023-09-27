package usecase

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/ryuji-cre8ive/monemana/internal/stores"
	"os"
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

	bot, err := linebot.New(Secret, Token)
	fmt.Print("in webhook")
	events, err := bot.ParseRequest(c.Request())
	fmt.Print("events: ", events)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			fmt.Print(err)
		}
	}
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch event.Message.(type) {
			case *linebot.TextMessage:
				if event.Message.(*linebot.TextMessage).Text == "お問い合わせ" {
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("お問い合わせいただきありがとうございます、")).Do(); err != nil {
						fmt.Print(err)
					}
				}
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ご連絡いただきありがとうございます。")).Do(); err != nil {
					fmt.Print(err)
				}
			}
		}
	}
	return nil
}

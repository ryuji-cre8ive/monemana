package main

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/line/line-bot-sdk-go/linebot"
	"golang.org/x/xerrors"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	Secret := os.Getenv("LINE_BOT_CHANNEL_SECRET")
	Token := os.Getenv("LINE_BOT_CHANNEL_TOKEN")

	e := echo.New()
	e.Use(middleware.Logger())

	bot, err := linebot.New(Secret, Token)
	if err != nil {
		log.Fatal(xerrors.Errorf("failed to create new bot: %w", err))
	}

	e.GET("/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.POST("/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.GET("/webhook", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.POST("/webhook", func(c echo.Context) error {
		log.Print("in webhook")
		events, err := bot.ParseRequest(c.Request())
		log.Print("events: ", events)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				log.Print(err)
			}
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
		return nil
	})

	e.Logger.Fatal(e.Start(":" + port))
}

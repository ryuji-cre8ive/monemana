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
		port = ":8080"
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
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/webhook", func(c echo.Context) error {
		if err := bot.ReplyMessage("Hello, World!"); err != nil {
			return xerrors.Errorf("failed to reply message: %w", err)
		}
		return nil
	})

	e.Logger.Fatal(e.Start(port))
}

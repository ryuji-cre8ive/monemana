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
	port := ":8080"
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

	e.GET("/webhook", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	message := linebot.NewTextMessage("Hello World!")

	if _, err := bot.BroadcastMessage(message).Do(); err != nil {
		log.Fatal(err)
	}

	e.Logger.Fatal(e.Start(port))
}

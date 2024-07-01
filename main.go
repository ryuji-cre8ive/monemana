package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/ryuji-cre8ive/monemana/internal/database"
	"github.com/ryuji-cre8ive/monemana/internal/stores"
	"github.com/ryuji-cre8ive/monemana/internal/ui"
	"github.com/ryuji-cre8ive/monemana/internal/usecase"
	"golang.org/x/xerrors"
)

func main() {
	err := godotenv.Load()

	db, err := database.New()
	if err != nil {
		log.Fatal(xerrors.Errorf("failed to connect to database: %w", err))
	}
	postgres, err := db.DB()
	defer postgres.Close()

	e := ui.Echo()

	s := stores.New(db)
	ss := usecase.New(s)
	h := ui.New(ss)

	ui.SetApi(e, h)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err != nil {
		log.Fatal(xerrors.Errorf("failed to create new bot: %w", err))
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	go func() {
		if err := e.Start(":" + port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

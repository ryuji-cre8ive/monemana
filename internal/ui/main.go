package ui

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ryuji-cre8ive/monemana/internal/usecase"
)

type Handler struct {
	WebhookHandler
}

func New(u *usecase.Usecase) *Handler {
	return &Handler{
		WebhookHandler: &webhookHandler{u.Webhook},
	}
}

func SetApi(e *echo.Echo, h *Handler) {
	g := e.Group("/api/v1")
	g.POST("/webhook", h.WebhookHandler.PostWebhook)
	g.GET("/healthcheck", HealthCheckHandler)
}

func Echo() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	return e
}

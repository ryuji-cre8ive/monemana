package ui

import (
	"net/http"

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
	authMiddleware := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			signature := c.Request().Header.Get("X-Line-Signature")
			userAgent := c.Request().Header.Get("User-Agent")
			if signature != "" && userAgent == "LineBotWebhook/2.0" {
				return next(c)
			} else {
				return c.String(http.StatusUnauthorized, "Invalid signature or User-Agent")
			}
		}
	}

	g.POST("/webhook", h.WebhookHandler.PostWebhook, authMiddleware)
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

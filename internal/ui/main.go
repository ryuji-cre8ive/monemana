package ui

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ryuji-cre8ive/monemana/internal/usecase"
	"github.com/ryuji-cre8ive/monemana/internal/utils"
	"net/http"
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
	g.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		signature := c.Request().Header.Get("X-Line-Signature")
		if utils.IsValidSignature(signature, reqBody) {
			// 署名が正しい場合、リクエストを処理
			fmt.Println("signature is valid")
			// ここでリクエストの処理を行う
		} else {
			fmt.Println("signature is invalid")
			// 署名が正しくない場合、エラーを返すなどの処理を行う
			c.String(http.StatusUnauthorized, "Invalid signature")
		}
	}))
	g.GET("/healthcheck", HealthCheckHandler)
	g.POST("/webhook", h.WebhookHandler.PostWebhook)

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

package ui

import (
	"github.com/labstack/echo/v4"
	"github.com/ryuji-cre8ive/monemana/internal/usecase"

	"golang.org/x/xerrors"
)

type (
	WebhookHandler interface {
		PostWebhook(c echo.Context) error
	}

	webhookHandler struct {
		usecase.WebhookUsecase
	}
)

func (h *webhookHandler) PostWebhook(c echo.Context) error {
	err := h.WebhookUsecase.PostWebhook(c)
	if err != nil {
		return xerrors.Errorf("failed to post webhook: %w", err)
	}
	return c.NoContent(200)
}

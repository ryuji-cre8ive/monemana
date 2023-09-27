package usecase

import "github.com/ryuji-cre8ive/monemana/internal/stores"

type Usecase struct {
	Webhook WebhookUsecase
}

func New(s *stores.Stores) *Usecase {
	return &Usecase{
		Webhook: &webhookUsecase{stores: s},
	}
}

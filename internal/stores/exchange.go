package stores

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/ryuji-cre8ive/monemana/internal/domain"
	"os"

	"golang.org/x/xerrors"
)

type (
	ExchangeStore interface {
		GetExchangeRate() (*domain.Exchange, error)
	}
	exchangeStore struct{}
)

func (s *exchangeStore) GetExchangeRate() (*domain.Exchange, error) {
	exchangeAccessKey := os.Getenv("EXCHANGE_ACCESS_KEY")
	var exchange *domain.Exchange
	client := resty.New()
	_, err := client.R().SetResult(&exchange).Get("http://api.exchangeratesapi.io/v1/latest?access_key=" + exchangeAccessKey)
	fmt.Println("usd", exchange.Rates["USD"])
	if err != nil {
		return nil, xerrors.Errorf("get exchange rate err: %w", err)
	}
	return exchange, nil
}

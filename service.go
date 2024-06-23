package main

import (
	"context"
	"fmt"
	"time"
)

type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, item string) (float64, error) {
	return MockPriceFetcher(ctx, item)

}

var priceMocks = map[string]float64{
	"BTC": 10_000.0,
	"ETH": 500.0,
	"WO":  0.0,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	//mimic the HTTP roundtrip
	time.Sleep(100 * time.Millisecond)
	price, ok := priceMocks[ticker]
	if !ok {
		return price, fmt.Errorf("the given ticker (%s) is not supported", ticker)
	}

	return price, nil
}

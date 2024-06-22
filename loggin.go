package main

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type logginService struct {
	next PriceFetcher
}

func NewLoggingService(svc PriceFetcher) PriceFetcher {
	return &logginService{next: svc}
}

func (s *logginService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		logrus.WithFields((logrus.Fields{
			"took":  time.Since(begin),
			"error": err,
			"price": price,
		})).Info("fetchPrice")
	}(time.Now())

	return s.next.FetchPrice(ctx, ticker)
}

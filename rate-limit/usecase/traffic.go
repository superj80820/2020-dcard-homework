package usecase

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/superj80820/2020-dcard-homework/domain"
)

// RateLimitUsecase ...
type rateLimitUsecase struct {
	RateLimitRepo domain.RateLimitRepository
}

// NewRateLimitUsecase ...
func NewRateLimitUsecase(rateLimitRepo domain.RateLimitRepository) domain.RateLimitUsecase {
	return &rateLimitUsecase{
		RateLimitRepo: rateLimitRepo,
	}
}

func (ru *rateLimitUsecase) IsTooManyRequests(ctx context.Context, IP string) (bool, int, error) {
	rateLimit, err := ru.RateLimitRepo.GetByIP(ctx, IP)
	if err != nil {
		logrus.WithFields(logrus.Fields{"logID": "99a5f41a-e232-4872-99b9-6a7cef6eaee0", "Error": err}).Error("Get repository error")
		return false, 0, errors.Wrap(err, "Get repository error")
	}

	rateLimit.Count++
	if err := ru.RateLimitRepo.Store(ctx, rateLimit); err != nil {
		logrus.WithFields(logrus.Fields{"logID": "8125830d-9fa3-45a9-b1d2-453ea3f3a858", "Error": err}).Error("Store repository error")
		return false, 0, errors.Wrap(err, "Store repository error")
	}

	if rateLimit.Count > 60 {
		return true, rateLimit.Count, nil
	}

	return false, rateLimit.Count, nil
}

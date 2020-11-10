package domain

import "context"

// RateLimit ...
type RateLimit struct {
	Count      int
	IP         string
	ExpireTime int
}

// RateLimitRepository ...
type RateLimitRepository interface {
	GetByIP(ctx context.Context, IP string) (*RateLimit, error)
	Store(ctx context.Context, rateLimit *RateLimit) error
}

// RateLimitUsecase ..
type RateLimitUsecase interface {
	IsTooManyRequests(ctx context.Context, IP string) (bool, int, error)
}

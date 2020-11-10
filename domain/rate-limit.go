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
	Store(ctx context.Context, r *RateLimit) error
}

// RateLimitUsecase ..
type RateLimitUsecase interface {
	Store(ctx context.Context, r *RateLimit) error
	isTooManyRequests(ctx context.Context, IP string) (bool, error)
}

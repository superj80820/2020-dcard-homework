package domain

import "context"

// RateLimit ...
type RateLimit struct {
	count      uint32
	IP         string
	ExpireTime uint32
}

// RateLimitRepository ...
type RateLimitRepository interface {
	GetByIP(ctx context.Context, IP string) (*RateLimit, error)
	Store(ctx context.Context, u *RateLimit) error
}

// RateLimitUsecase ..
type RateLimitUsecase interface {
	Store(ctx context.Context, u *RateLimit) error
	isTooManyRequests(ctx context.Context, IP string) (bool, error)
}

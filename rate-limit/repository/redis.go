package repository

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/superj80820/2020-dcard-homework/domain"
)

type RedisRedisRateLimitRepository struct {
	Client *redis.Client
}

// NewRedisRateLimitRepository ...
func NewRedisRateLimitRepository(Client *redis.Client) domain.RateLimitRepository {
	return &RedisRedisRateLimitRepository{Client}
}

func (c *RedisRedisRateLimitRepository) GetByIP(ctx context.Context, IP string) (*domain.RateLimit, error) {
	count, err := c.Client.Get(ctx, IP).Int()
	if err != nil {
		return nil, errors.Wrap(err, "Get count error")
	}
	expireTime, err := c.Client.TTL(ctx, IP).Result()
	if err != nil {
		return nil, errors.Wrap(err, "Get TTL error")
	}

	return &domain.RateLimit{Count: count, IP: IP, ExpireTime: int(expireTime)}, nil
}

func (c *RedisRedisRateLimitRepository) Store(ctx context.Context, rateLimit *domain.RateLimit) error {
	expireTime := time.Duration(rateLimit.ExpireTime) * time.Second
	if err := c.Client.Set(ctx, rateLimit.IP, rateLimit.Count, expireTime).Err(); err != nil {
		return errors.Wrap(err, "Store to redis error")
	}
	return nil
}

package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/superj80820/2020-dcard-homework/domain"
)

type redisRedisRateLimitRepository struct {
	Client *redis.Client
}

// NewRedisRateLimitRepository ...
func NewRedisRateLimitRepository(Client *redis.Client) domain.RateLimitRepository {
	return &redisRedisRateLimitRepository{Client}
}

func (c *redisRedisRateLimitRepository) GetByIP(ctx context.Context, IP string) (*domain.RateLimit, error) {
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

func (c *redisRedisRateLimitRepository) Store(ctx context.Context, r *domain.RateLimit) error {
	expireTime := time.Duration(r.ExpireTime) * time.Second
	if err := c.Client.Set(ctx, r.IP, r.Count, expireTime).Err(); err != nil {
		return errors.Wrap(err, "Store to redis error")
	}
	return nil
}

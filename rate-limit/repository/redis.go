package repository

import (
	"context"
	"time"

	goredislib "github.com/go-redis/redis/v8"
	redsync "github.com/go-redsync/redsync/v4"

	"github.com/pkg/errors"
	"github.com/superj80820/2020-dcard-homework/domain"
)

type redisRedisRateLimitRepository struct {
	Client *goredislib.Client
	Mutex  *redsync.Mutex
}

// NewRedisRateLimitRepository ...
func NewRedisRateLimitRepository(client *goredislib.Client, mutex *redsync.Mutex) domain.RateLimitRepository {
	return &redisRedisRateLimitRepository{
		Client: client,
		Mutex:  mutex,
	}
}

func (r *redisRedisRateLimitRepository) Lock() (err error) {
	if err := r.Mutex.Lock(); err != nil {
		return errors.Wrap(err, "Lock error")
	}
	return nil
}

func (r *redisRedisRateLimitRepository) Unlock() (err error) {
	if ok, err := r.Mutex.Unlock(); !ok || err != nil {
		return errors.Wrap(err, "Unlock error")
	}
	return nil
}

func (r *redisRedisRateLimitRepository) GetByIP(ctx context.Context, IP string) (*domain.RateLimit, bool, error) {
	count, err := r.Client.Get(ctx, IP).Int()
	if err == goredislib.Nil {
		return &domain.RateLimit{}, false, nil
	} else if err != nil {
		return nil, true, errors.Wrap(err, "Get count error")
	}
	expireTime, err := r.Client.TTL(ctx, IP).Result()
	if err != nil {
		return nil, true, errors.Wrap(err, "Get TTL error")
	}

	return &domain.RateLimit{Count: count, IP: IP, ExpireTime: int(expireTime / time.Second)}, true, nil
}

func (r *redisRedisRateLimitRepository) Store(ctx context.Context, rateLimit *domain.RateLimit) error {
	expireTime := time.Duration(rateLimit.ExpireTime) * time.Second

	if err := r.Client.Set(ctx, rateLimit.IP, rateLimit.Count, expireTime).Err(); err != nil {
		return errors.Wrap(err, "Store to redis error")
	}
	return nil
}

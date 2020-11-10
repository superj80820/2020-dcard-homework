package repository_test

import (
	"context"
	"testing"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"github.com/stretchr/testify/assert"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"

	"github.com/superj80820/2020-dcard-homework/domain"
	_rateLimitRepo "github.com/superj80820/2020-dcard-homework/rate-limit/repository"
)

func newTestRedis() *redis.Client {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	return client
}

func TestRedisSimple(t *testing.T) {
	redisClient := newTestRedis()
	redisPool := goredis.NewPool(redisClient)
	redisMutex := redsync.New(redisPool).NewMutex("dcard-service")

	rateLimitRepo := _rateLimitRepo.NewRedisRateLimitRepository(redisClient, redisMutex)
	rateLimitRepo.Store(context.TODO(), &domain.RateLimit{
		Count:      1,
		IP:         "127.0.1.1",
		ExpireTime: 10,
	})

	count, _ := redisClient.Get(context.TODO(), "127.0.1.1").Int()
	assert.EqualValues(t, 1, count)
}

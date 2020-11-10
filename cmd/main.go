package main

import (
	"github.com/gin-gonic/gin"
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_rateLimitHttpDelivery "github.com/superj80820/2020-dcard-homework/rate-limit/delivery/http"
	_rateLimitRepo "github.com/superj80820/2020-dcard-homework/rate-limit/repository"
	_rateLimitUsecase "github.com/superj80820/2020-dcard-homework/rate-limit/usecase"
)

func init() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	restfulHost := viper.GetString("RESTFULHOST")
	restfulPort := viper.GetString("RESTFULPORT")
	redisAddress := viper.GetString("REDISADDRESS")

	redisClient := goredislib.NewClient(&goredislib.Options{
		Addr:     redisAddress,
		Password: "",
		DB:       0,
	})
	redisPool := goredis.NewPool(redisClient)
	redisMutex := redsync.New(redisPool).NewMutex("dcard-service")

	rateLimitRepo := _rateLimitRepo.NewRedisRateLimitRepository(redisClient, redisMutex)

	rateLimitUsecase := _rateLimitUsecase.NewRateLimitUsecase(rateLimitRepo)

	logrus.Info("HTTP server started")
	engine := gin.Default()

	_rateLimitHttpDelivery.NewRateLimitHandler(engine, rateLimitUsecase)

	logrus.Fatal(engine.Run(restfulHost + ":" + restfulPort))
}

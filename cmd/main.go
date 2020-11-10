package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
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

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	rateLimitRepo := _rateLimitRepo.RedisRedisRateLimitRepository(rdb)

	rateLimitUsecase := _rateLimitUsecase.NewRateLimitUsecase(rateLimitRepo)

	log.Print("HTTP server started")
	engine := gin.Default()

	_rateLimitHttpDelivery.NewRateLimitHandler(engine, rateLimitUsecase)

	log.Fatal(engine.Run(restfulHost + ":" + restfulPort))
}

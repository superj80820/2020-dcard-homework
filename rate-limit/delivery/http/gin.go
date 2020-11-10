package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/superj80820/2020-dcard-homework/domain"
)

// rateLimitHandler ...
type rateLimitHandler struct {
	RateLimitUsecase domain.RateLimitUsecase
}

// NewRateLimitHandler ...
func NewRateLimitHandler(e *gin.Engine, rateLimitUsecase domain.RateLimitUsecase) {
	handler := &rateLimitHandler{
		RateLimitUsecase: rateLimitUsecase,
	}

	e.GET("/", handler.getIndexWithRateLimit)
}

func (rh *rateLimitHandler) getIndexWithRateLimit(c *gin.Context) {
	isTooManyRequests, requestCount, err := rh.RateLimitUsecase.IsTooManyRequests(c, c.ClientIP())
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Internal error",
		})
		return
	}
	if isTooManyRequests {
		c.String(http.StatusOK, "Error")
		return
	}
	c.String(http.StatusOK, strconv.Itoa(requestCount))
}

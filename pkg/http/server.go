package http

import (
	"github.com/brhamidi/fizzbuzz/pkg/logger"
	"github.com/brhamidi/fizzbuzz/pkg/storage"
	"github.com/gin-gonic/gin"
)

type handler struct {
	store storage.Storage
	log   logger.Logger
}

const (
	healthRoute   = "/health"
	FizzbuzzRoute = "/fizzbuzz"
	StatsRoute    = "/stats"
)

func NewServer(env string, store storage.Storage, log logger.Logger) *gin.Engine {
	h := handler{store, log}

	gin.SetMode(env)

	router := gin.New()
	router.Use(gin.Recovery())

	// useful for monitoring our service and CI/CD tools
	router.GET(healthRoute, h.Health)

	router.GET(FizzbuzzRoute, h.Fizzbuzz)

	router.GET(StatsRoute, h.GetStats)
	router.DELETE(StatsRoute, h.DeleteStats)

	return router
}

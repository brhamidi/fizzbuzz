package http

import (
	"github.com/brhamidi/fizzbuzz/pkg/logger"
	"github.com/brhamidi/fizzbuzz/pkg/storage"
	"github.com/gin-gonic/gin"
)

type handler struct {
	s   storage.Storage
	log logger.Logger
}

const (
	healthRoute = "/health"
)

func NewServer(env string, s storage.Storage, log logger.Logger) *gin.Engine {
	h := handler{s, log}

	gin.SetMode(env)

	router := gin.New()
	router.Use(gin.Recovery())

	router.GET(healthRoute, h.Health)

	return router
}

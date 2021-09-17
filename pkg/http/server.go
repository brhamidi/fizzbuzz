package http

import (
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	"github.com/brhamidi/fizzbuzz/pkg/logger"
	"github.com/brhamidi/fizzbuzz/pkg/storage"
	"github.com/gin-gonic/gin"

	_ "github.com/brhamidi/fizzbuzz/docs"
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

// @title Fizzbuzz Rest Server
// @version 1.0
// @description This is a fizzbuz server with a statistic endpoint

// @contact.name Brahim Hamidi
// @contact.email brahim.hamidi75@gmail.com

// @license.name GPL 3
// @license.url https://www.gnu.org/licenses/quick-guide-gplv3.html

// @host localhost:3000
func NewServer(env string, store storage.Storage, log logger.Logger) *gin.Engine {
	h := handler{store, log}

	gin.SetMode(env)

	router := gin.New()
	router.Use(gin.Recovery())

	// swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// useful for monitoring our service and CI/CD tools
	router.GET(healthRoute, h.Health)

	// we dont want access log for the route above
	router.Use(AccessLog(log))

	router.GET(FizzbuzzRoute, h.Fizzbuzz)

	router.GET(StatsRoute, h.GetStats)
	router.DELETE(StatsRoute, h.DeleteStats)

	return router
}

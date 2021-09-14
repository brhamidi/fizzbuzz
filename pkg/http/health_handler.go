package http

import "github.com/gin-gonic/gin"

type HealthResp struct {
	Status bool `json:"status"`
}

func (h handler) Health(c *gin.Context) {
	h.log.Info("requested")
	c.JSON(200, HealthResp{true})
}

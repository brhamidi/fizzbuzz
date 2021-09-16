package http

import "github.com/gin-gonic/gin"

type HealthResp struct {
	Status bool `json:"status"`
} //@name Health Reponse

// Health method http GET
// @Summary Health check
// @Description Healthcheck endpoint, to ensure that the service is running.
// @Tags Health
// @Accept  json
// @Produce  json
// @Success 200 {object} HealthResp
// @Router /health [get]
func (h handler) Health(c *gin.Context) {
	c.JSON(200, &ResponseSuccess{&HealthResp{true}})
}

package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const templateStatsResponse = `The most popular fizzbuzz requests with %d hits are: {%s}`

// @Summary Retreive statistic regarding fizzbuzz request
// @Description Returns the query that was performed the most and also the number of times
// @Tags Statistic
// @Produce json
// @Success 200 {object} ResponseSuccess
// @Success 204
// @Failure 500,400 {object} ResponseError
// @Router /stats [get]
func (h handler) GetStats(c *gin.Context) {
	key, hits, err := h.store.Max()
	if err != nil {
		h.log.Error(fmt.Errorf("%w: %s", errUserInternal, err))
		c.JSON(http.StatusInternalServerError, newResponseError(errUserInternal))
		return
	}

	if hits == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	formattedResp := fmt.Sprintf(templateStatsResponse, hits, key)
	c.JSON(http.StatusOK, &ResponseSuccess{formattedResp})
}

// @Summary Reset statistics
// @Description Reset all statistics regarding the queries made in the past
// @Tags Statistic
// @Success 200
// @Failure 500 {object} ResponseError
// @Router /stats [delete]
func (h handler) DeleteStats(c *gin.Context) {
	if err := h.store.Reset(); err != nil {
		h.log.Error(fmt.Errorf("%w: %s", errUserInternal, err))
		c.JSON(http.StatusInternalServerError, newResponseError(errUserInternal))
		return
	}

	c.Status(http.StatusOK)
}

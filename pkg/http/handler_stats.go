package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const templateStatsResponse = `The most popular fizzbuzz requests with %d hits are: {%s}`

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

func (h handler) DeleteStats(c *gin.Context) {
	if err := h.store.Reset(); err != nil {
		h.log.Error(fmt.Errorf("%w: %s", errUserInternal, err))
		c.JSON(http.StatusInternalServerError, newResponseError(errUserInternal))
		return
	}

	c.Status(http.StatusOK)
}

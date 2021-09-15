package http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const templateStatsResponse = `The most popular fizzbuzz requests with %d hits are: {int1:%s int2:%s limit:%s str1:%s str2:%s}`

func (h handler) GetStats(c *gin.Context) {
	key, hits, err := h.s.Max()
	if err != nil {
		h.log.Error(fmt.Errorf("%w: %s", errUserInternal, err))
		c.JSON(http.StatusInternalServerError, NewResponseError(errUserInternal))
		return
	}

	if hits == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	q := strings.Split(key, ",")
	h.log.Info(q)
	formattedResp := fmt.Sprintf(templateStatsResponse, hits, q[0], q[1], q[2], q[3], q[4])
	c.JSON(http.StatusOK, &ResponseSuccess{formattedResp})
}

func (h handler) DeleteStats(c *gin.Context) {
	if err := h.s.Reset(); err != nil {
		h.log.Error(fmt.Errorf("%w: %s", errUserInternal, err))
		c.JSON(http.StatusInternalServerError, NewResponseError(errUserInternal))
		return
	}

	c.Status(http.StatusOK)
}

package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Todo add description
// and swagger tag for auto generating yml file
func (h handler) Fizzbuzz(c *gin.Context) {
	var input queries

	if err := c.ShouldBindQuery(&input); err != nil {
		errWrapped := fmt.Errorf("%w: %s", errInvalidQueries, err)
		c.JSON(http.StatusBadRequest, NewResponseError(errWrapped))
		return
	}

	if err := input.isValid(); err != nil {
		errWrapped := fmt.Errorf("%w: %s", errInvalidQueries, err)
		c.JSON(http.StatusBadRequest, NewResponseError(errWrapped))
		return
	}

	result := input.Compute()

	if err := h.s.Increment(input.String()); err != nil {
		errWrapped := fmt.Errorf("%w: %s", errStorage, err)
		h.log.Error(errWrapped)
	}

	c.JSON(200, &ResponseSuccess{result})
}

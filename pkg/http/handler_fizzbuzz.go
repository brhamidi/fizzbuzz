package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Todo add description
// and swagger tag for auto generating yml file
func (h handler) Fizzbuzz(c *gin.Context) {
	var input Fizzbuzz

	if err := c.ShouldBindQuery(&input); err != nil {
		errWrapped := fmt.Errorf("%w: %s", errInvalidQuery, err)
		c.JSON(http.StatusBadRequest, newResponseError(errWrapped))
		return
	}

	if err := input.IsValid(); err != nil {
		errWrapped := fmt.Errorf("%w: %s", errInvalidQuery, err)
		c.JSON(http.StatusBadRequest, newResponseError(errWrapped))
		return
	}

	result := input.Compute()

	if err := h.store.Increment(input.Key()); err != nil {
		h.log.Error(fmt.Errorf("%w: %s", errStorage, err))
	}

	c.JSON(200, &ResponseSuccess{result})
}

package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	errInvalidQueries = errors.New("failed to validate input queries")
	errStorage        = errors.New("failed to perform storage operation")
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

	// increment stats DB

	c.JSON(200, &ResponseSuccess{result})
}

package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Fizzbuzz
// @Summary Return fizzbuzz result
// @Description Fizzbuzz operation
// @Tags Fizzbuzz Operation
// @Produce json
// @Param int1 query int true "int1 query parameter"
// @Param int2 query int true "int2 query parameter"
// @Param limit query int true "limit query parameter"
// @Param str1 query string true "str1 query parameter"
// @Param str2 query string true "str2 query parameter"
// @Success 200 {object} ResponseSuccess{data=[]string}
// @Failure 400 {object} ResponseError
// @Router /fizzbuzz [get]
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

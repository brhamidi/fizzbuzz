package http

import (
	"errors"
	"fmt"
	"strconv"
)

// Todo add description
// TODO maybe remove validator, too overkill and verbose for this project
type Fizzbuzz struct {
	V1    int    `form:"int1"`
	V2    int    `form:"int2"`
	Limit int    `form:"limit"`
	S1    string `form:"str1"`
	S2    string `form:"str2"`
}

func (q Fizzbuzz) Key() string {
	return fmt.Sprintf("%d,%d,%d,%s,%s", q.V1, q.V2, q.Limit, q.S1, q.S2)
}

func (q Fizzbuzz) IsValid() error {
	if q.V1 <= 0 || q.V2 <= 0 || q.Limit <= 0 {
		return errors.New("integer parameter must be greater than zero")
	}
	if q.S1 == "" || q.S2 == "" {
		return errors.New("string parameter must be not empty")
	}
	return nil
}

func (q Fizzbuzz) Compute() []string {
	r := make([]string, 0, q.Limit)
	for i := 1; i <= q.Limit; i++ {
		if i%q.V1 == 0 && i%q.V2 == 0 {
			r = append(r, q.S1+q.S2)
		} else if i%q.V1 == 0 {
			r = append(r, q.S1)
		} else if i%q.V2 == 0 {
			r = append(r, q.S2)
		} else {
			r = append(r, strconv.Itoa(i))
		}
	}
	return r
}

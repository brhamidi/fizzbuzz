package http

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Todo add description
// TODO maybe remove validator, too overkill and verbose for this project
type queries struct {
	V1    int    `validate:"required,gt=0" form:"int1"`
	V2    int    `validate:"required,gt=0" form:"int2"`
	Limit int    `validate:"required,gt=0"  form:"limit"`
	S1    string `validate:"required" form:"str1"`
	S2    string `validate:"required" form:"str2"`
}

func (q queries) String() string {
	return fmt.Sprintf("%d,%d,%d,%s,%s", q.V1, q.V2, q.Limit, q.S1, q.S2)
}

func (q queries) isValid() error {
	var err error
	if err = validate.Struct(q); err == nil {
		return nil
	}

	verrs, ok := err.(validator.ValidationErrors)
	if !ok {
		return fmt.Errorf("%w. Another error occured: failed to assert type <validator.ValidationErrors>", err)
	}

	errs := make([]string, 0, len(verrs))
	for _, err := range verrs {
		errs = append(errs, err.Error())
	}
	return errors.New(strings.Join(errs, ". "))
}

func (q queries) Compute() []string {
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

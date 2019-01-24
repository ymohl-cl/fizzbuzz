package algo

import (
	"errors"
	"strconv"
)

// errors list
const (
	ErrModuloZero    = "specification use an impossible modulo operation by zero"
	ErrBadRange      = "start must be smaller than limit"
	ErrBadValues     = "values can't be empty"
	ErrValuesTooLong = "values are limited to 254 characters"
)

// constraints values
const (
	limitValue = 254
)

// Algo define the specifics configuration to run the FizzBuzz
type Algo struct {
	Start  int
	Limit  int
	N1     int
	N2     int
	Value1 string
	Value2 string
}

// Run start the fizz buzz test wich count from Start to Limit
// if Start is biggest than Limit or equal an error will be returned
// if Start, Limit, N1, or N2 == 0 an error will be returned
// if i is multiple of n1, so add v1 text
// is i is multiple of n2, so add v2 text
// else just add number as text
func (a Algo) Run() ([]string, error) {
	var result []string
	var err error

	if err = a.check(); err != nil {
		return nil, err
	}
	for i := a.Start; i <= a.Limit; i++ {
		var value string
		var flag bool
		if i%a.N1 == 0 {
			value += a.Value1
			flag = true
		}
		if i%a.N2 == 0 {
			value += a.Value2
			flag = true
		}
		if !flag {
			value = strconv.Itoa(int(i))
		}
		result = append(result, value)
	}
	return result, nil
}

func (a Algo) check() error {
	if a.Start >= a.Limit {
		return errors.New(ErrBadRange)
	}
	if a.Start < 0 && a.Limit > 0 {
		return errors.New(ErrModuloZero)
	}
	if a.Start == 0 || a.Limit == 0 || a.N1 == 0 || a.N2 == 0 {
		return errors.New(ErrModuloZero)
	}
	if a.Value1 == "" || a.Value2 == "" {
		return errors.New(ErrBadValues)
	}
	if len(a.Value1) > limitValue || len(a.Value2) > limitValue {
		return errors.New(ErrValuesTooLong)
	}
	return nil
}

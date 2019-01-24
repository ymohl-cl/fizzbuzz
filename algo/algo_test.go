package algo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func dataTest() Algo {
	return Algo{
		Start:  1,
		Limit:  100,
		N1:     3,
		N2:     5,
		Value1: "fizz",
		Value2: "buzz",
	}
}

func TestRun(t *testing.T) {
	// should return an error because start is bigger than limit
	func() {
		// init
		a := dataTest()
		a.Start = a.Limit + 42
		ret, err := a.Run()

		// assert
		if assert.Error(t, err) {
			assert.Equal(t, ErrBadRange, err.Error())
			assert.Nil(t, ret)
		}
	}()
	// should return an error because start equal zero value
	func() {
		// init
		a := dataTest()
		a.Start = 0
		ret, err := a.Run()

		// assert
		if assert.Error(t, err) {
			assert.Equal(t, ErrModuloZero, err.Error())
			assert.Nil(t, ret)
		}
	}()
	// should return an error because limit equal zero value
	func() {
		// init
		a := dataTest()
		a.Start = -3
		a.Limit = 0
		ret, err := a.Run()

		// assert
		if assert.Error(t, err) {
			assert.Equal(t, ErrModuloZero, err.Error())
			assert.Nil(t, ret)
		}
	}()
	// should return an error because N1 equal zero value
	func() {
		// init
		a := dataTest()
		a.N1 = 0
		ret, err := a.Run()

		// assert
		if assert.Error(t, err) {
			assert.Equal(t, ErrModuloZero, err.Error())
			assert.Nil(t, ret)
		}
	}()
	// should return an error because N2 equal zero value
	func() {
		// init
		a := dataTest()
		a.N2 = 0
		ret, err := a.Run()

		// assert
		if assert.Error(t, err) {
			assert.Equal(t, ErrModuloZero, err.Error())
			assert.Nil(t, ret)
		}
	}()
	// should return an error because range specificion check a zero value
	func() {
		// init
		a := dataTest()
		a.Start = -10
		ret, err := a.Run()

		// assert
		if assert.Error(t, err) {
			assert.Equal(t, ErrModuloZero, err.Error())
			assert.Nil(t, ret)
		}
	}()
	// should return an error because value1 is not define
	func() {
		// init
		a := dataTest()
		a.Value1 = ""
		ret, err := a.Run()

		// assert
		if assert.Error(t, err) {
			assert.Equal(t, ErrBadValues, err.Error())
			assert.Nil(t, ret)
		}
	}()
	// should return an error because value2 is not define
	func() {
		// init
		a := dataTest()
		a.Value2 = ""
		ret, err := a.Run()

		// assert
		if assert.Error(t, err) {
			assert.Equal(t, ErrBadValues, err.Error())
			assert.Nil(t, ret)
		}
	}()
	// should return an error because value1 is too long
	func() {
		// init
		a := dataTest()
		a.Value1 = strings.Repeat("a", 255)
		ret, err := a.Run()

		// assert
		if assert.Error(t, err) {
			assert.Equal(t, ErrValuesTooLong, err.Error())
			assert.Nil(t, ret)
		}
	}()
	// should return an error because value2 is too long
	func() {
		// init
		a := dataTest()
		a.Value2 = strings.Repeat("a", 255)
		ret, err := a.Run()

		// assert
		if assert.Error(t, err) {
			assert.Equal(t, ErrValuesTooLong, err.Error())
			assert.Nil(t, ret)
		}
	}()

	// default: should be ok with positives values
	func() {
		// init
		expected := []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16"}
		a := dataTest()
		a.Limit = 16
		ret, err := a.Run()

		// assert
		if assert.NoError(t, err) {
			assert.Equal(t, expected, ret)
		}
	}()

	// default: should be ok with negatives values
	func() {
		// init
		expected := []string{"-16", "fizzbuzz", "-14", "-13", "fizz", "-11", "buzz", "fizz", "-8", "-7", "fizz", "buzz", "-4", "fizz", "-2", "-1"}
		a := dataTest()
		a.Start = -16
		a.Limit = -1
		ret, err := a.Run()

		// assert
		if assert.NoError(t, err) {
			assert.Equal(t, expected, ret)
		}
	}()
}

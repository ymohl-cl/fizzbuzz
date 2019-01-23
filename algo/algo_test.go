package algo

import (
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

	/*
		func() {
			// init
			expected := "1;2;fizz;4;buzz;fizz;7;8;fizz;buzz;11;fizz;13;14;fizzbuzz;16;17;fizz;19;buzz;fizz;22;23;fizz;buzz;26;fizz;28;29;fizzbuzz;31;32;fizz;34;buzz;fizz;37;38;fizz;buzz;41;fizz;43;44;fizzbuzz;46;47;fizz;49;buzz;fizz;52;53;fizz;buzz;56;fizz;58;59;fizzbuzz;61;62;fizz;64;buzz;fizz;67;68;fizz;buzz;71;fizz;73;74;fizzbuzz;76;77;fizz;79;buzz;fizz;82;83;fizz;buzz;86;fizz;88;89;fizzbuzz;91;92;fizz;94;buzz;fizz;97;98;fizz;buzz"
			ret := Run(3, 5, "fizz", "buzz")

			// assert
			assert.Equal(t, expected, ret)
		}()

		func() {
			// init
			expected := "1;2;fizz;4;buzz;fizz;7;8;fizz;buzz;11;fizz;13;14;fizzbuzz;16;17;fizz;19;buzz;fizz;22;23;fizz;buzz;26;fizz;28;29;fizzbuzz;31;32;fizz;34;buzz;fizz;37;38;fizz;buzz;41;fizz;43;44;fizzbuzz;46;47;fizz;49;buzz;fizz;52;53;fizz;buzz;56;fizz;58;59;fizzbuzz;61;62;fizz;64;buzz;fizz;67;68;fizz;buzz;71;fizz;73;74;fizzbuzz;76;77;fizz;79;buzz;fizz;82;83;fizz;buzz;86;fizz;88;89;fizzbuzz;91;92;fizz;94;buzz;fizz;97;98;fizz;buzz"
			ret := Run(-3, -5, "fizz", "buzz")

			// assert
			assert.Equal(t, expected, ret)
		}()

		// should return an error because v1 == 0 and can't be use to modulo operator
		func() {
			// init
			expected := "error"
			ret := Run(0, 5, "fizz", "buzz")

			// assert
			assert.Equal(t, expected, ret)
		}()

	*/
}

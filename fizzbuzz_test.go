package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	// default: should be ok with positives values
	func() {
		// init
		expected := "1;2;fizz;4;buzz;fizz;7;8;fizz;buzz;11;fizz;13;14;fizzbuzz;16;17;fizz;19;buzz;fizz;22;23;fizz;buzz;26;fizz;28;29;fizzbuzz;31;32;fizz;34;buzz;fizz;37;38;fizz;buzz;41;fizz;43;44;fizzbuzz;46;47;fizz;49;buzz;fizz;52;53;fizz;buzz;56;fizz;58;59;fizzbuzz;61;62;fizz;64;buzz;fizz;67;68;fizz;buzz;71;fizz;73;74;fizzbuzz;76;77;fizz;79;buzz;fizz;82;83;fizz;buzz;86;fizz;88;89;fizzbuzz;91;92;fizz;94;buzz;fizz;97;98;fizz;buzz"
		ret := Run(3, 5, "fizz", "buzz")

		// assert
		assert.Equal(t, expected, ret)
	}()

	// default: should be ok with negatives values
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

	// should return an error because v2 == 0 and can't be use to modulo operator
	func() {
		// init
		expected := "error"
		ret := Run(3, 0, "fizz", "buzz")

		// assert
		assert.Equal(t, expected, ret)
	}()
}

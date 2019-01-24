package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	// Default: should be OK
	func() {
		// init
		expectedHash := "10035fizzbuzz"
		i := InputFB{
			Int1:  3,
			Int2:  5,
			Limit: 100,
			Str1:  "fizz",
			Str2:  "buzz",
		}
		hash := i.Hash()

		// assert
		assert.Equal(t, expectedHash, hash)
	}()
}

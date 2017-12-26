package mathematics_test

import (
	"testing"

	"github.com/goncharovnikita/going-go/algoritms/mathematics"
	"github.com/stretchr/testify/assert"
)

func TestFactorial(t *testing.T) {
	test := 12
	expected := 479001600

	assert.Equal(t, expected, mathematics.Factorial(test))
}

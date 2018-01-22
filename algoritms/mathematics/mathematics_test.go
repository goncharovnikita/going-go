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

func TestSquareMatrixSum(t *testing.T) {
	m1 := mathematics.SquareMatrix{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}
	expected := mathematics.SquareMatrix{{2, 4, 6}, {4, 6, 8}, {6, 8, 10}}

	assert.Equal(t, expected, m1.Sum(m1))
}

func TestSquareMatrixMultiplication(t *testing.T) {
	m1 := mathematics.SquareMatrix{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}
	expected := mathematics.SquareMatrix{{14, 20, 26}, {20, 29, 38}, {26, 38, 50}}

	assert.Equal(t, expected, m1.Mult(m1))
}

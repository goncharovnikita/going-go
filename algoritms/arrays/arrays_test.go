package arrays_test

import (
	"testing"

	"github.com/goncharovnikita/going-go/algoritms/arrays"
	"github.com/stretchr/testify/assert"
)

func TestFindMaxCrossingSlice(t *testing.T) {
	testCase := []int{1, -1, -45, 123, 21, -56, 23, 42}
	expectedLeft, expectedRight, expectedSum := 3, 7, 153
	actualLeft, actualRight, actualSum := arrays.FindMaxCrossingSlice(testCase)

	assert.Equal(t, expectedLeft, actualLeft)
	assert.Equal(t, expectedRight, actualRight)
	assert.Equal(t, expectedSum, actualSum)
}

func TestFindMaxSlice(t *testing.T) {
	testCase := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	eLeft, eRight, eSum := 7, 10, 43
	aLeft, aRight, aSum := arrays.FindMaxSlice(testCase)

	assert.Equal(t, eLeft, aLeft)
	assert.Equal(t, eRight, aRight)
	assert.Equal(t, eSum, aSum)
}

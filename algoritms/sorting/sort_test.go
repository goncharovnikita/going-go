package sorting_test

import (
	"testing"

	"github.com/goncharovnikita/going-go/algoritms/sorting"
	"github.com/stretchr/testify/assert"
)

func TestInsertion(t *testing.T) {
	arr := []int{2, 4, 1, 3, 6, 5}
	expected := []int{1, 2, 3, 4, 5, 6}

	assert.Equal(t, expected, sorting.Insertion(arr))
}

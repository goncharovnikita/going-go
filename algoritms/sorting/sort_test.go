package sorting_test

import (
	"testing"

	"github.com/goncharovnikita/going-go/algoritms/sorting"
	"github.com/stretchr/testify/assert"
)

var testCases = []map[string][]int{
	{
		"test":     {2, 4, 1, 3, 6, 5},
		"expected": {1, 2, 3, 4, 5, 6},
	},
}

func TestInsertion(t *testing.T) {
	for _, v := range testCases {
		assert.Equal(t, v["expected"], sorting.Insertion(v["test"]))
	}
}

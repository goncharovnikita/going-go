package sorting_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/goncharovnikita/going-go/algoritms/sorting"
	"github.com/stretchr/testify/assert"
)

var testCases = []map[string][]int{
	{
		"test":     {2, 4, 1, 3, 6, 5},
		"expected": {1, 2, 3, 4, 5, 6},
	},
	{
		"test":     {2, 4, 1, 3, 6, 5, 12, 124, 21251, 1231254, 3453, 123, 123, 321, 321, 1432, 156, 1, 2, 45, 5, 56},
		"expected": {1, 1, 2, 2, 3, 4, 5, 5, 6, 12, 45, 56, 123, 123, 124, 156, 321, 321, 1432, 3453, 21251, 1231254},
	},
}

func TestInsertion(t *testing.T) {
	start := time.Now()
	for _, v := range testCases {
		assert.Equal(t, v["expected"], sorting.Insertion(v["test"]))
	}
	fmt.Printf("Insertion sort takes %s\n", time.Since(start))
}

func TestSelection(t *testing.T) {
	start := time.Now()
	for _, v := range testCases {
		assert.Equal(t, v["expected"], sorting.Selection(v["test"]))
	}
	fmt.Printf("Selection sort takes %s\n", time.Since(start))
}

func TestMerge(t *testing.T) {
	start := time.Now()
	for _, v := range testCases {
		assert.Equal(t, v["expected"], sorting.Merge(v["test"]))
	}
	fmt.Printf("Merge sort takes %s\n", time.Since(start))
}

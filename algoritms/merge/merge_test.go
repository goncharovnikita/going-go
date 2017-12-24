package merge_test

import (
	"testing"

	"github.com/goncharovnikita/going-go/algoritms/merge"
	"github.com/stretchr/testify/assert"
)

var cases = []map[string][]int{
	{
		"a":        {1, 2, 3, 4, 5},
		"b":        {2, 3, 5, 6},
		"expected": {1, 2, 2, 3, 3, 4, 5, 5, 6},
	},
	{
		"a":        {1},
		"b":        {2, 3, 5, 6},
		"expected": {1, 2, 3, 5, 6},
	},
}

func TestMerge(t *testing.T) {
	for _, v := range cases {
		assert.Equal(t, v["expected"], merge.Merge(v["a"], v["b"]))
	}
}

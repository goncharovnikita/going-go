package sorting

import (
	"github.com/goncharovnikita/going-go/algoritms/merge"
)

// Merge sort
func Merge(arr []int) (result []int) {
	if len(arr) > 1 {
		q := len(arr) / 2
		return merge.Merge(Merge(arr[0:q]), Merge(arr[q:]))
	}
	result = arr
	return
}

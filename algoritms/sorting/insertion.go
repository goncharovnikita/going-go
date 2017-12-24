package sorting

// Insertion sort
func Insertion(arr []int) (result []int) {
	result = make([]int, len(arr))
	copy(result, arr)
	if len(result) < 2 {
		return result
	}
	for j := 1; j < len(result); j++ {
		key := result[j]
		i := j - 1
		for i >= 0 && result[i] > key {
			result[i+1] = result[i]
			i--
		}
		result[i+1] = key
	}
	return
}

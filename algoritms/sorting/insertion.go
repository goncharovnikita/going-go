package sorting

// Insertion sort
func Insertion(arr []int) (result []int) {
	if len(arr) < 2 {
		return arr
	}
	result = arr
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

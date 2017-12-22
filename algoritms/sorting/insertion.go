package sorting

// Insertion sort
func Insertion(arr []int) (result []int) {
	if len(arr) < 2 {
		return arr
	}
	result = arr
	for i := 1; i < len(result); i++ {
		key := result[i]
		j := i - 1
		for j > 0 && result[j] > key {
			result[j+1] = result[j]
			j--
		}
		result[j+1] = key
	}
	return
}

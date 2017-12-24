package sorting

// Bubble sort
func Bubble(arr []int) (result []int) {
	result = make([]int, len(arr))
	copy(result, arr)
	for i := 0; i < len(result)-1; i++ {
		for j := len(result) - 1; j > i; j-- {
			if result[j] < result[j-1] {
				result[j-1], result[j] = result[j], result[j-1]
			}
		}
	}
	return result
}

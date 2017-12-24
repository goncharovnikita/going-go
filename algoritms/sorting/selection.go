package sorting

// Selection sort
func Selection(arr []int) (result []int) {
	result = make([]int, len(arr))
	copy(result, arr)
	for i := 0; i < len(result)-1; i++ {
		minIndex := i
		for k := i + 1; k < len(result); k++ {
			if result[k] < result[minIndex] {
				minIndex = k
			}
		}
		result[i], result[minIndex] = result[minIndex], result[i]
	}
	return result
}

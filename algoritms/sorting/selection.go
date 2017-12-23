package sorting

// Selection sort
func Selection(arr []int) (result []int) {
	result = make([]int, len(arr))
	for i := 0; i < len(arr)-1; i++ {
		j := arr[i]
		for k := i; k < len(arr); k++ {
			key := arr[k]
			if key < j {
				j = key
			}
		}
		result[i] = j
	}
	result[len(arr)-1] = arr[len(arr)-1]
	return result
}

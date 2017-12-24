package merge

// Merge func
func Merge(a []int, b []int) (result []int) {
	result = make([]int, len(a)+len(b))
	j := 0
	k := 0
	for i := 0; i < len(a)+len(b); i++ {
		if j == len(a) {
			for index, v := range b[k:] {
				result[i+index] = v
			}
			break
		}
		if k == len(b) {
			for index, v := range a[j:] {
				result[i+index] = v
			}
			break
		}
		if a[j] <= b[k] {
			result[i] = a[j]
			j++
		} else {
			result[i] = b[k]
			k++
		}
	}
	return result
}

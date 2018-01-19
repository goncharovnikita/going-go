package arrays

import "math"

// FindMaxCrossingSlice func
func FindMaxCrossingSlice(arr []int) (leftMax int, rightMax int, maxSum int) {
	if len(arr) < 2 {
		return 0, 0, arr[0]
	}

	if len(arr) == 2 {
		return 0, 1, arr[0] + arr[1]
	}
	divider := len(arr) / 2
	sumLeft := math.MinInt64
	sum := 0
	for i := divider; i >= 0; i-- {
		sum += arr[i]
		if sum > sumLeft {
			leftMax = i
			sumLeft = sum
		}
	}

	sumRight := math.MinInt64
	sum = 0
	for i := divider + 1; i < len(arr); i++ {
		sum += arr[i]
		if sum > sumRight {
			rightMax = i
			sumRight = sum
		}
	}
	maxSum = sumLeft + sumRight
	return
}

// FindMaxSlice func
func FindMaxSlice(arr []int) (leftMax int, rightMax int, sum int) {
	if len(arr) < 2 {
		return 0, 0, arr[0]
	}
	divider := len(arr) / 2
	lMin, lMax, lSum := FindMaxSlice(arr[:divider])
	rMin, rMax, rSum := FindMaxSlice(arr[divider:])
	crossMin, crossMax, crossSum := FindMaxCrossingSlice(arr)
	if lSum >= rSum && lSum >= crossSum {
		leftMax, rightMax, sum = lMin, lMax, lSum
	} else if rSum >= lSum && rSum >= crossSum {
		leftMax, rightMax, sum = rMin, rMax, rSum
	} else {
		leftMax, rightMax, sum = crossMin, crossMax, crossSum
	}

	return
}

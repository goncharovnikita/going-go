package mathematics

// Factorial func
func Factorial(n int) (result int) {
	if n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}

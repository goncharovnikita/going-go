package mathematics

// SquareMatrix is the rectangular array of numbers
// with columns count equal to rows count
type SquareMatrix [][]int

// NewSquareMatrix creates new square matrix with specified side length
func NewSquareMatrix(side int) SquareMatrix {
	a := make(SquareMatrix, side)
	for i := 0; i < side; i++ {
		b := make([]int, side)
		a[i] = b
	}
	return a
}

// Sum is the addition operation of square matrices
func (s SquareMatrix) Sum(m SquareMatrix) SquareMatrix {
	result := NewSquareMatrix(len(s))
	s.toSquare()
	m.toSquare()
	for j := range s {
		for k := range s[j] {
			result[j][k] = s[j][k] + m[j][k]
		}
	}

	return result
}

// Mult is the multiplication operation of square matrices
func (s SquareMatrix) Mult(m SquareMatrix) SquareMatrix {
	result := NewSquareMatrix(len(s))
	s.toSquare()
	m.toSquare()
	for j := range s {
		for k := range s[j] {
			a := 0
			for z := 0; z < len(s); z++ {
				a += s[j][z] * m[z][k]
			}
			result[j][k] = a
		}
	}

	return result
}

// toSquare converts matrix to be sure, that matrix is square
func (s SquareMatrix) toSquare() {
	result := NewSquareMatrix(len(s))
	for j := range s {
		for k := range s[j] {
			result[j][k] = s[j][k]
		}
	}
	s = result
}

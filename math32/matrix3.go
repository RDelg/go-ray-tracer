package math32

// Matrix3 represents a 3x3 float32 matrix
type Matrix3 struct {
	Values [3][3]float32
}

// Matrix3Ones returns a pointer to a Matrix3 full of 1s
func Matrix3Ones() *Matrix3 {
	oneVector := [3]float32{1.0, 1.0, 1.0}
	return &Matrix3{[3][3]float32{oneVector, oneVector, oneVector}}
}

// Matrix3Zeros returns a pointer to a Matrix3 full of 0s
func Matrix3Zeros() *Matrix3 {
	return &Matrix3{}
}

// Matrix3Identity returns a pointer to a Matrix3 with the identity values
func Matrix3Identity() *Matrix3 {
	return &Matrix3{}
}

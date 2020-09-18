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

// Add adds the other Matrix3 to this one
// Returns a pointer to the updated Matrix3
func (m *Matrix3) Add(other *Matrix3) *Matrix3 {
	m.Values[0][0] += other.Values[0][0]
	m.Values[0][1] += other.Values[0][1]
	m.Values[0][2] += other.Values[0][2]

	m.Values[1][0] += other.Values[1][0]
	m.Values[1][1] += other.Values[1][1]
	m.Values[1][2] += other.Values[1][2]

	m.Values[2][0] += other.Values[0][0]
	m.Values[2][1] += other.Values[0][1]
	m.Values[2][2] += other.Values[0][2]
	return m
}

// Mult multiplies the other Matrix3 to this one
// Returns a pointer to the updated Matrix3
func (m *Matrix3) Mult(other *Matrix3) *Matrix3 {

	t := Matrix3{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			t.Values[i][j] += m.Values[i][j] * other.Values[j][i]
		}
	}
	m = &t
	return m
}

// Transpose transpose the matrix
// Returns a pointer to the updated Matrix3
func (m *Matrix3) Transpose() *Matrix3 {
	m.Values[1][0], m.Values[0][1] = m.Values[0][1], m.Values[1][0]
	m.Values[2][0], m.Values[0][2] = m.Values[0][2], m.Values[2][0]
	m.Values[2][1], m.Values[1][2] = m.Values[1][2], m.Values[2][1]
	return m
}

// TODO
// func (m *Matrix3) SubMatrix2(n, m, int) math32.Matrix2 {
// }
// func (m *Matrix3) Determinant() float32 {
// }
// func (m *Matrix3) Equal(other Matrix3) bool {
// }
// func (m *Matrix3) Determinant() float32 {
// }
//  func (m *Matrix3) AddVector3(vector math.Vector3) *Matrix3 {
// }
//  func (m *Matrix3) MultVector3(vector math.Vector3) *Matrix3 {
// }
//  func (m *Matrix3) AddScalar3(scalar float32) *Matrix3 {
// }
//  func (m *Matrix3) MultScalar3(scalar float32) *Matrix3 {
// }

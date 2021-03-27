package math32

import (
	"errors"
	"fmt"
)

// Matrix4 represents a 3x3 float32 matrix
type Matrix4 struct {
	Values [4][4]float32
}

// NewMatrix4Ones returns a pointer to a new Matrix4 full of 1s
func NewMatrix4Ones() *Matrix4 {
	oneVector := [4]float32{1., 1., 1., 1.}
	return &Matrix4{[4][4]float32{oneVector, oneVector, oneVector, oneVector}}
}

// NewMatrix4Zeros returns a pointer to a new Matrix4 full of zeros
func NewMatrix4Zeros() *Matrix4 {
	return &Matrix4{}
}

// NewMatrix4Identity returns a pointer to a new Matrix4 containingn the Identity_3 matrix
func NewMatrix4Identity() *Matrix4 {
	return &Matrix4{
		[4][4]float32{
			{1.},
			{0., 1.},
			{0., 0., 1.},
			{0., 0., 0., 1.},
		}}
}

// Add adds the other Matrix4 to this one
// Returns a pointer to the updated Matrix4
func (m *Matrix4) Add(other *Matrix4) *Matrix4 {
	m.Values[0][0] += other.Values[0][0]
	m.Values[0][1] += other.Values[0][1]
	m.Values[0][2] += other.Values[0][2]
	m.Values[0][3] += other.Values[0][3]

	m.Values[1][0] += other.Values[1][0]
	m.Values[1][1] += other.Values[1][1]
	m.Values[1][2] += other.Values[1][2]
	m.Values[1][3] += other.Values[1][3]

	m.Values[2][0] += other.Values[2][0]
	m.Values[2][1] += other.Values[2][1]
	m.Values[2][2] += other.Values[2][2]
	m.Values[2][3] += other.Values[2][3]

	m.Values[3][0] += other.Values[3][0]
	m.Values[3][1] += other.Values[3][1]
	m.Values[3][2] += other.Values[3][2]
	m.Values[3][3] += other.Values[3][3]
	return m
}

// Mult multiplies the other Matrix4 to this one
// Returns a pointer to the updated Matrix4
func (m *Matrix4) Mult(other *Matrix4) *Matrix4 {
	a11 := m.Values[0][0]
	a12 := m.Values[0][1]
	a13 := m.Values[0][2]
	a14 := m.Values[0][3]

	a21 := m.Values[1][0]
	a22 := m.Values[1][1]
	a23 := m.Values[1][2]
	a24 := m.Values[1][3]

	a31 := m.Values[2][0]
	a32 := m.Values[2][1]
	a33 := m.Values[2][2]
	a34 := m.Values[2][3]

	a41 := m.Values[3][0]
	a42 := m.Values[3][1]
	a43 := m.Values[3][2]
	a44 := m.Values[3][3]

	b11 := other.Values[0][0]
	b12 := other.Values[0][1]
	b13 := other.Values[0][2]
	b14 := other.Values[0][3]

	b21 := other.Values[1][0]
	b22 := other.Values[1][1]
	b23 := other.Values[1][2]
	b24 := other.Values[1][3]

	b31 := other.Values[2][0]
	b32 := other.Values[2][1]
	b33 := other.Values[2][2]
	b34 := other.Values[2][3]

	b41 := other.Values[3][0]
	b42 := other.Values[3][1]
	b43 := other.Values[3][2]
	b44 := other.Values[3][3]

	m.Values[0][0] = a11*b11 + a12*b21 + a13*b31 + a14*b41
	m.Values[0][1] = a11*b12 + a12*b22 + a13*b32 + a14*b42
	m.Values[0][2] = a11*b13 + a12*b23 + a13*b33 + a14*b43
	m.Values[0][3] = a11*b14 + a12*b24 + a13*b34 + a14*b44

	m.Values[1][0] = a21*b11 + a22*b21 + a23*b31 + a24*b41
	m.Values[1][1] = a21*b12 + a22*b22 + a23*b32 + a24*b42
	m.Values[1][2] = a21*b13 + a22*b23 + a23*b33 + a24*b43
	m.Values[1][3] = a21*b14 + a22*b24 + a23*b34 + a24*b44

	m.Values[2][0] = a31*b11 + a32*b21 + a33*b31 + a34*b41
	m.Values[2][1] = a31*b12 + a32*b22 + a33*b32 + a34*b42
	m.Values[2][2] = a31*b13 + a32*b23 + a33*b33 + a34*b43
	m.Values[2][3] = a31*b14 + a32*b24 + a33*b34 + a34*b44

	m.Values[3][0] = a41*b11 + a42*b21 + a43*b31 + a44*b41
	m.Values[3][1] = a41*b12 + a42*b22 + a43*b32 + a44*b42
	m.Values[3][2] = a41*b13 + a42*b23 + a43*b33 + a44*b43
	m.Values[3][3] = a41*b14 + a42*b24 + a43*b34 + a44*b44

	return m
}

// Transpose transpose the matrix
// Returns a pointer to the updated Matrix4
func (m *Matrix4) Transpose() *Matrix4 {
	m.Values[1][0], m.Values[0][1] = m.Values[0][1], m.Values[1][0]
	m.Values[2][0], m.Values[0][2] = m.Values[0][2], m.Values[2][0]
	m.Values[3][0], m.Values[0][3] = m.Values[0][3], m.Values[3][0]
	m.Values[2][1], m.Values[1][2] = m.Values[1][2], m.Values[2][1]
	m.Values[3][1], m.Values[1][3] = m.Values[1][3], m.Values[3][1]
	m.Values[3][2], m.Values[2][3] = m.Values[2][3], m.Values[3][2]
	return m
}

// Determinant return the determinant of the matrix
func (m *Matrix4) Determinant() float32 {
	a := m.Values[2][2]*m.Values[3][3] - m.Values[3][2]*m.Values[2][3]
	b := m.Values[2][1]*m.Values[3][3] - m.Values[3][1]*m.Values[2][3]
	c := m.Values[2][1]*m.Values[3][2] - m.Values[3][1]*m.Values[2][2]
	d := m.Values[2][0]*m.Values[3][3] - m.Values[3][0]*m.Values[2][3]
	e := m.Values[2][0]*m.Values[3][2] - m.Values[3][0]*m.Values[2][2]
	f := m.Values[2][0]*m.Values[3][1] - m.Values[3][0]*m.Values[2][1]

	return m.Values[0][0]*(m.Values[1][1]*a-m.Values[1][2]*b+m.Values[1][3]*c) -
		m.Values[0][1]*(m.Values[1][0]*a-m.Values[1][2]*d+m.Values[1][3]*e) +
		m.Values[0][2]*(m.Values[1][0]*b-m.Values[1][1]*d+m.Values[1][3]*f) -
		m.Values[0][3]*(m.Values[1][0]*c-m.Values[1][1]*e+m.Values[1][2]*f)
}

// Inverse returns the inverse of the matrix
func (m *Matrix4) Inverse() (*Matrix4, error) {
	det := m.Determinant()
	if det == 0. {
		return nil, errors.New("the matrix has 0 determinant. inverse cannot be calculated")
	}
	r := m.Cofactor().Transpose().DivScalar(det)
	return r, nil
}

// Cofactor returns the cofactor matrix
func (m *Matrix4) Cofactor() *Matrix4 {
	r := Matrix4{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			x, _ := m.SubMatrix3(i, j)
			r.Values[i][j] = x.Determinant() * (1. + float32((-i-j)%2)*2.)
		}
	}
	return &r
}

// SubMatrix3 returns a Matrix3 ignoring the row and column indexes (starting from 0) of the original Matrix4
// if the indexes are outside the range of the matrix returns an error
func (m *Matrix4) SubMatrix3(row, column int) (*Matrix3, error) {
	if row > 3 || column > 3 || row < 0 || column < 0 {
		return nil, fmt.Errorf("row and columns must be lower than 3 and greater than 0. Got row %s and column %s", fmt.Sprint(row), fmt.Sprint(column))
	}
	r := Matrix3{}
	for i, ii := 0, 0; i < 3; i, ii = i+1, ii+1 {
		if ii == row {
			ii++
		}
		for j, jj := 0, 0; j < 3; j, jj = j+1, jj+1 {
			if jj == column {
				jj++
			}
			r.Values[i][j] = m.Values[ii][jj]
		}
	}
	return &r, nil
}

// Equal returns true if the matrices are equal, else return false
func (m *Matrix4) Equal(other *Matrix4) bool {
	return m.Values[0][0] == other.Values[0][0] &&
		m.Values[0][1] == other.Values[0][1] &&
		m.Values[0][2] == other.Values[0][2] &&
		m.Values[0][3] == other.Values[0][3] &&
		m.Values[1][0] == other.Values[1][0] &&
		m.Values[1][1] == other.Values[1][1] &&
		m.Values[1][2] == other.Values[1][2] &&
		m.Values[1][3] == other.Values[1][3] &&
		m.Values[2][0] == other.Values[2][0] &&
		m.Values[2][1] == other.Values[2][1] &&
		m.Values[2][2] == other.Values[2][2] &&
		m.Values[2][3] == other.Values[2][3] &&
		m.Values[3][0] == other.Values[3][0] &&
		m.Values[3][1] == other.Values[3][1] &&
		m.Values[3][2] == other.Values[3][2] &&
		m.Values[3][3] == other.Values[3][3]
}

// AddVector4 adds a Vector4 to each column if columnWise is true
// else adds the values to each row
func (m *Matrix4) AddVector4(vector *Vector4, columnWise bool) *Matrix4 {
	if columnWise {
		m.Values[0][0] = vector.X
		m.Values[0][1] = vector.X
		m.Values[0][2] = vector.X
		m.Values[0][3] = vector.X
		m.Values[1][0] = vector.Y
		m.Values[1][1] = vector.Y
		m.Values[1][2] = vector.Y
		m.Values[1][3] = vector.Y
		m.Values[2][0] = vector.Z
		m.Values[2][1] = vector.Z
		m.Values[2][2] = vector.Z
		m.Values[2][3] = vector.Z
		m.Values[3][0] = vector.W
		m.Values[3][1] = vector.W
		m.Values[3][2] = vector.W
		m.Values[3][3] = vector.W
	} else {
		m.Values[0][0] = vector.X
		m.Values[1][0] = vector.X
		m.Values[2][0] = vector.X
		m.Values[3][0] = vector.X
		m.Values[0][1] = vector.Y
		m.Values[1][1] = vector.Y
		m.Values[2][1] = vector.Y
		m.Values[3][1] = vector.Y
		m.Values[0][2] = vector.Z
		m.Values[1][2] = vector.Z
		m.Values[2][2] = vector.Z
		m.Values[3][2] = vector.Z
		m.Values[0][3] = vector.W
		m.Values[1][3] = vector.W
		m.Values[2][3] = vector.W
		m.Values[3][3] = vector.W
	}
	return m
}

// MultVector4 multiplies each row with the vector
func (m *Matrix4) MultVector4(vector *Vector4, columnWise bool) *Vector4 {
	var x, y, z, w float32
	if columnWise {
		x = m.Values[0][0]*vector.X + m.Values[1][0]*vector.Y + m.Values[2][0]*vector.Z + m.Values[3][0]*vector.W
		y = m.Values[0][1]*vector.X + m.Values[1][1]*vector.Y + m.Values[2][1]*vector.Z + m.Values[3][1]*vector.W
		z = m.Values[0][2]*vector.X + m.Values[1][2]*vector.Y + m.Values[2][2]*vector.Z + m.Values[3][2]*vector.W
		w = m.Values[0][3]*vector.X + m.Values[1][3]*vector.Y + m.Values[2][3]*vector.Z + m.Values[3][3]*vector.W
	} else {
		x = m.Values[0][0]*vector.X + m.Values[0][1]*vector.Y + m.Values[0][2]*vector.Z + m.Values[0][3]*vector.W
		y = m.Values[1][0]*vector.X + m.Values[1][1]*vector.Y + m.Values[1][2]*vector.Z + m.Values[1][3]*vector.W
		z = m.Values[2][0]*vector.X + m.Values[2][1]*vector.Y + m.Values[2][2]*vector.Z + m.Values[2][3]*vector.W
		w = m.Values[3][0]*vector.X + m.Values[3][1]*vector.Y + m.Values[3][2]*vector.Z + m.Values[3][3]*vector.W
	}
	return &Vector4{x, y, z, w}
}

// AddScalar adds a scalar to each value of the matrix
func (m *Matrix4) AddScalar(scalar float32) *Matrix4 {
	m.Values[0][0] += scalar
	m.Values[0][1] += scalar
	m.Values[0][2] += scalar
	m.Values[0][3] += scalar
	m.Values[1][0] += scalar
	m.Values[1][1] += scalar
	m.Values[1][2] += scalar
	m.Values[1][3] += scalar
	m.Values[2][0] += scalar
	m.Values[2][1] += scalar
	m.Values[2][2] += scalar
	m.Values[2][3] += scalar
	m.Values[3][0] += scalar
	m.Values[3][1] += scalar
	m.Values[3][2] += scalar
	m.Values[3][3] += scalar
	return m
}

// MultScalar multiplies each value of the matrix by the scalar
func (m *Matrix4) MultScalar(scalar float32) *Matrix4 {
	m.Values[0][0] *= scalar
	m.Values[0][1] *= scalar
	m.Values[0][2] *= scalar
	m.Values[0][3] *= scalar
	m.Values[1][0] *= scalar
	m.Values[1][1] *= scalar
	m.Values[1][2] *= scalar
	m.Values[1][3] *= scalar
	m.Values[2][0] *= scalar
	m.Values[2][1] *= scalar
	m.Values[2][2] *= scalar
	m.Values[2][3] *= scalar
	m.Values[3][0] *= scalar
	m.Values[3][1] *= scalar
	m.Values[3][2] *= scalar
	m.Values[3][3] *= scalar
	return m
}

// DivScalar divides each value of the matrix by the scalar
func (m *Matrix4) DivScalar(scalar float32) *Matrix4 {
	m.Values[0][0] /= scalar
	m.Values[0][1] /= scalar
	m.Values[0][2] /= scalar
	m.Values[0][3] /= scalar
	m.Values[1][0] /= scalar
	m.Values[1][1] /= scalar
	m.Values[1][2] /= scalar
	m.Values[1][3] /= scalar
	m.Values[2][0] /= scalar
	m.Values[2][1] /= scalar
	m.Values[2][2] /= scalar
	m.Values[2][3] /= scalar
	m.Values[3][0] /= scalar
	m.Values[3][1] /= scalar
	m.Values[3][2] /= scalar
	m.Values[3][3] /= scalar
	return m
}

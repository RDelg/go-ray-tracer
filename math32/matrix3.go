package math32

// Matrix3 represents a 3x3 float32 matrix
type Matrix3 struct {
	Values [3][3]float32
}

// NewMatrix3Ones returns a pointer to a new Matrix3 full of 1s
func NewMatrix3Ones() *Matrix3 {
	oneVector := [3]float32{1.0, 1.0, 1.0}
	return &Matrix3{[3][3]float32{oneVector, oneVector, oneVector}}
}

// NewMatrix3Zeros returns a pointer to a new Matrix3 full of zeros
func NewMatrix3Zeros() *Matrix3 {
	return &Matrix3{}
}

// NewMatrix3Identity returns a pointer to a new Matrix3 containingn the Identity_3 matrix
func NewMatrix3Identity() *Matrix3 {
	return &Matrix3{
		[3][3]float32{
			[3]float32{1.0, 0.0, 0.0},
			[3]float32{0.0, 1.0, 0.0},
			[3]float32{0.0, 0.0, 1.0},
		}}
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

	m.Values[2][0] += other.Values[2][0]
	m.Values[2][1] += other.Values[2][1]
	m.Values[2][2] += other.Values[2][2]
	return m
}

// Mult multiplies the other Matrix3 to this one
// Returns a pointer to the updated Matrix3
func (m *Matrix3) Mult(other *Matrix3) *Matrix3 {
	a11 := m.Values[0][0]
	a12 := m.Values[0][1]
	a13 := m.Values[0][2]
	a21 := m.Values[1][0]
	a22 := m.Values[1][1]
	a23 := m.Values[1][2]
	a31 := m.Values[2][0]
	a32 := m.Values[2][1]
	a33 := m.Values[2][2]

	b11 := other.Values[0][0]
	b12 := other.Values[0][1]
	b13 := other.Values[0][2]
	b21 := other.Values[1][0]
	b22 := other.Values[1][1]
	b23 := other.Values[1][2]
	b31 := other.Values[2][0]
	b32 := other.Values[2][1]
	b33 := other.Values[2][2]

	m.Values[0][0] = a11*b11 + a12*b21 + a13*b31
	m.Values[0][1] = a11*b12 + a12*b22 + a13*b32
	m.Values[0][2] = a11*b13 + a12*b23 + a13*b33

	m.Values[1][0] = a21*b11 + a22*b21 + a23*b31
	m.Values[1][1] = a21*b12 + a22*b22 + a23*b32
	m.Values[1][2] = a21*b13 + a22*b23 + a23*b33

	m.Values[2][0] = a31*b11 + a32*b21 + a33*b31
	m.Values[2][1] = a31*b12 + a32*b22 + a33*b32
	m.Values[2][2] = a31*b13 + a32*b23 + a33*b33

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

// Determinant return the determinant of the matrix
func (m *Matrix3) Determinant() float32 {
	return m.Values[0][0]*m.Values[1][1]*m.Values[2][2] -
		m.Values[0][0]*m.Values[2][1]*m.Values[1][2] -
		m.Values[0][1]*m.Values[1][0]*m.Values[2][2] +
		m.Values[0][1]*m.Values[2][0]*m.Values[1][2] +
		m.Values[0][2]*m.Values[1][0]*m.Values[2][1] -
		m.Values[0][2]*m.Values[2][0]*m.Values[1][1]
}

// Equal returns true if the matrices are equal, else return false
func (m *Matrix3) Equal(other *Matrix3) bool {
	return m.Values[0][0] == other.Values[0][0] &&
		m.Values[0][1] == other.Values[0][1] &&
		m.Values[0][2] == other.Values[0][2] &&
		m.Values[1][0] == other.Values[1][0] &&
		m.Values[1][1] == other.Values[1][1] &&
		m.Values[1][2] == other.Values[1][2] &&
		m.Values[2][0] == other.Values[2][0] &&
		m.Values[2][1] == other.Values[2][1] &&
		m.Values[2][2] == other.Values[2][2]
}

// AddVector3 adds a Vector3 to each column if columnWise is true
// else adds the values to each row
func (m *Matrix3) AddVector3(vector *Vector3, columnWise bool) *Matrix3 {
	if columnWise {
		m.Values[0][0] = vector.X
		m.Values[1][0] = vector.X
		m.Values[2][0] = vector.X
		m.Values[0][1] = vector.Y
		m.Values[1][1] = vector.Y
		m.Values[2][1] = vector.Y
		m.Values[0][2] = vector.Z
		m.Values[1][2] = vector.Z
		m.Values[2][2] = vector.Z
	} else {
		m.Values[0][0] = vector.X
		m.Values[0][1] = vector.X
		m.Values[0][2] = vector.X
		m.Values[1][0] = vector.Y
		m.Values[1][1] = vector.Y
		m.Values[1][2] = vector.Y
		m.Values[2][0] = vector.Z
		m.Values[2][1] = vector.Z
		m.Values[2][2] = vector.Z
	}
	return m
}

// MultVector3 multiplies each row with the vector
func (m *Matrix3) MultVector3(vector *Vector3, columnWise bool) *Vector3 {
	var x, y, z float32
	if columnWise {
		x = m.Values[0][0]*vector.X + m.Values[0][1]*vector.Y + m.Values[0][2]*vector.Z
		y = m.Values[1][0]*vector.X + m.Values[1][1]*vector.Y + m.Values[1][2]*vector.Z
		z = m.Values[2][0]*vector.X + m.Values[2][1]*vector.Y + m.Values[2][2]*vector.Z
	} else {
		x = m.Values[0][0]*vector.X + m.Values[1][0]*vector.Y + m.Values[2][0]*vector.Z
		y = m.Values[0][1]*vector.X + m.Values[1][1]*vector.Y + m.Values[2][1]*vector.Z
		z = m.Values[0][2]*vector.X + m.Values[1][2]*vector.Y + m.Values[2][2]*vector.Z
	}
	return &Vector3{x, y, z}
}

// MultPoint3 multiplies each row with the vector
func (m *Matrix3) MultPoint3(vector *Point3, columnWise bool) *Point3 {
	var x, y, z float32
	if columnWise {
		x = m.Values[0][0]*vector.X + m.Values[0][1]*vector.Y + m.Values[0][2]*vector.Z
		y = m.Values[1][0]*vector.X + m.Values[1][1]*vector.Y + m.Values[1][2]*vector.Z
		z = m.Values[2][0]*vector.X + m.Values[2][1]*vector.Y + m.Values[2][2]*vector.Z
	} else {
		x = m.Values[0][0]*vector.X + m.Values[1][0]*vector.Y + m.Values[2][0]*vector.Z
		y = m.Values[0][1]*vector.X + m.Values[1][1]*vector.Y + m.Values[2][1]*vector.Z
		z = m.Values[0][2]*vector.X + m.Values[1][2]*vector.Y + m.Values[2][2]*vector.Z
	}
	return &Point3{x, y, z}
}

// AddScalar adds a scalar to each value of the matrix
func (m *Matrix3) AddScalar(scalar float32) *Matrix3 {
	m.Values[0][0] += scalar
	m.Values[0][1] += scalar
	m.Values[0][2] += scalar
	m.Values[1][0] += scalar
	m.Values[1][1] += scalar
	m.Values[1][2] += scalar
	m.Values[2][0] += scalar
	m.Values[2][1] += scalar
	m.Values[2][2] += scalar
	return m
}

// MultScalar multiplies each value of the matrix by the scalar
func (m *Matrix3) MultScalar(scalar float32) *Matrix3 {
	m.Values[0][0] *= scalar
	m.Values[0][1] *= scalar
	m.Values[0][2] *= scalar
	m.Values[1][0] *= scalar
	m.Values[1][1] *= scalar
	m.Values[1][2] *= scalar
	m.Values[2][0] *= scalar
	m.Values[2][1] *= scalar
	m.Values[2][2] *= scalar
	return m
}

// DivScalar divides each value of the matrix by the scalar
func (m *Matrix3) DivScalar(scalar float32) *Matrix3 {
	m.Values[0][0] /= scalar
	m.Values[0][1] /= scalar
	m.Values[0][2] /= scalar
	m.Values[1][0] /= scalar
	m.Values[1][1] /= scalar
	m.Values[1][2] /= scalar
	m.Values[2][0] /= scalar
	m.Values[2][1] /= scalar
	m.Values[2][2] /= scalar
	return m
}

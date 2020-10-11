package math32

import "math"

// Vector3 represents a 3-dimensional vector with values {X, Y, Z}
type Vector3 struct {
	X, Y, Z float32
}

// Vector3FromArray creates a Vector3 from a float array of lenth 3
func Vector3FromArray(array [3]float32) *Vector3 {
	return &Vector3{array[0], array[1], array[2]}
}

// AsArray returs the Vector3 values as a float array of lenth 3
func (vector *Vector3) AsArray() [3]float32 {
	return [3]float32{vector.X, vector.Y, vector.Z}
}

// SetFromArray returs the Vector3 values as a float array of lenth 3
func (vector *Vector3) SetFromArray(array [3]float32) *Vector3 {
	vector.X = array[0]
	vector.Y = array[1]
	vector.Z = array[2]
	return vector
}

// Neg multiplies each vector value by -1
// Returns a pointer to the updated vector
func (vector *Vector3) Neg() *Vector3 {
	vector.X *= -1.0
	vector.Y *= -1.0
	vector.Z *= -1.0
	return vector
}

// Add adds the values of other to vector
// Returns a pointer to the updated vector
func (vector *Vector3) Add(other *Vector3) *Vector3 {
	vector.X += other.X
	vector.Y += other.Y
	vector.Z += other.Z
	return vector
}

// Substract substracts the values of other to vector
// Returns a pointer to the updated vector
func (vector *Vector3) Substract(other *Vector3) *Vector3 {
	return vector.Add(other.Neg())
}

// MultScalar multiplies the values of vector by scalar
// Returns a pointer to the updated vector
func (vector *Vector3) MultScalar(scalar float32) *Vector3 {
	vector.X *= scalar
	vector.Y *= scalar
	vector.Z *= scalar
	return vector
}

// DivideScalar divides the values of vector by scalar
// Returns a pointer to the updated vector
func (vector *Vector3) DivideScalar(scalar float32) *Vector3 {
	vector.X /= scalar
	vector.Y /= scalar
	vector.Z /= scalar
	return vector
}

// Magnitude returns the euclidean norm of the vector
func (vector *Vector3) Magnitude() float32 {
	return float32(math.Sqrt(float64(vector.X*vector.X + vector.Y*vector.Y + vector.Z*vector.Z)))
}

// Normalize applies a normalization to the vector using the euclidean norm
// Returns a pointer to the updated vector
func (vector *Vector3) Normalize() *Vector3 {
	return vector.DivideScalar(vector.Magnitude())
}

// Dot applies the dot product between vector and other
// Returns a float32 with the result value
func (vector *Vector3) Dot(other *Vector3) float32 {
	return vector.X*other.X + vector.Y*other.Y + vector.Z*other.Z
}

// Cross returns a pointer to a new Vector3 containing the cross product between the two Vector3
func (vector *Vector3) Cross(other *Vector3) *Vector3 {
	return &Vector3{vector.Y*other.Z - vector.Z*other.Y, vector.Z*other.X - vector.X*other.Z, vector.X*other.Y - vector.Y*other.X}
}

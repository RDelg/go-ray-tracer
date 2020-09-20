package math32

import "math"

// Vector4 represents a 3-dimensional vector with values {X, Y, Z}
type Vector4 struct {
	X, Y, Z, W float32
}

// Vector4FromArray creates a Vector4 from a float array of lenth 3
func Vector4FromArray(array [4]float32) *Vector4 {
	return &Vector4{array[0], array[1], array[2], array[3]}
}

// AsArray returs the Vector4 values as a float array of lenth 3
func (vector *Vector4) AsArray() [4]float32 {
	return [4]float32{vector.X, vector.Y, vector.Z, vector.W}
}

// SetFromArray returs the Vector4 values as a float array of lenth 3
func (vector *Vector4) SetFromArray(array [4]float32) *Vector4 {
	vector.X = array[0]
	vector.Y = array[1]
	vector.Z = array[2]
	vector.W = array[3]
	return vector
}

// Neg multiplies each vector value by -1
// Returns a pointer to the updated vector
func (vector *Vector4) Neg() *Vector4 {
	vector.X *= -1.0
	vector.Y *= -1.0
	vector.Z *= -1.0
	vector.W *= -1.0
	return vector
}

// Add adds the values of other to vector
// Returns a pointer to the updated vector
func (vector *Vector4) Add(other *Vector4) *Vector4 {
	return &Vector4{vector.X + other.X, vector.Y + other.Y, vector.Z + other.Z, vector.W + other.W}
}

// Substract substracts the values of other to vector
// Returns a pointer to the updated vector
func (vector *Vector4) Substract(other *Vector4) *Vector4 {
	return vector.Add(other.Neg())
}

// MultScalar multiplies the values of vector by scalar
// Returns a pointer to the updated vector
func (vector *Vector4) MultScalar(scalar float32) *Vector4 {
	vector.X *= scalar
	vector.Y *= scalar
	vector.Z *= scalar
	vector.W *= scalar
	return vector
}

// DivideScalar divides the values of vector by scalar
// Returns a pointer to the updated vector
func (vector *Vector4) DivideScalar(scalar float32) *Vector4 {
	vector.X /= scalar
	vector.Y /= scalar
	vector.Z /= scalar
	vector.W /= scalar
	return vector
}

// Magnitude returns the euclidean norm of the vector
func (vector *Vector4) Magnitude() float32 {
	return float32(math.Sqrt(float64(vector.X*vector.X + vector.Y*vector.Y + vector.Z*vector.Z + vector.W*vector.W)))
}

// Normalize applies a normalization to the vector using the euclidean norm
// Returns a pointer to the updated vector
func (vector *Vector4) Normalize() *Vector4 {
	return vector.DivideScalar(vector.Magnitude())
}

// Dot applies the dot product between vector and other
// Returns a float32 with the result value
func (vector *Vector4) Dot(other *Vector4) float32 {
	return vector.X*other.X + vector.Y*other.Y + vector.Z*other.Z + vector.W*other.W
}

// Cross product of vector only exists for 3D and 7D
// https://www.jstor.org/stable/2323537
// func (vector *Vector4) Cross(other *Vector4) *Vector4 {
// 	return &Vector4{}
// }

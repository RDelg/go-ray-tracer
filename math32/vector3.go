package vector3

import "math"

// Vector3 represents a 3-dimensional vector with values {X, Y, Z}
type Vector3 struct {
	X, Y, Z float32
}

// FromArray creates a Vector3 from an array of 3 values
func FromArray(array [3]float32) *Vector3 {
	return &Vector3{array[0], array[1], array[2]}
}

// AsArray returs the Vector3 values as an array
func (vA *Vector3) AsArray() [3]float32 {
	return [3]float32{vA.X, vA.Y, vA.Z}
}

// Neg returns a reference to a new Vector3 with values (-1)* {X, Y, Z}
func (vA *Vector3) Neg() *Vector3 {
	return &Vector3{-1.0 * vA.X, -1.0 * vA.Y, -1.0 * vA.Z}
}

// Add returns a reference to a new Vector3 with the sum of the two Vector3
func (vA *Vector3) Add(vb *Vector3) *Vector3 {
	return &Vector3{vA.X + vb.X, vA.Y + vb.Y, vA.Z + vb.Z}
}

// Substract returns a reference to a new Vector3 with the substraction of the two Vector3
func (vA *Vector3) Substract(vb *Vector3) *Vector3 {
	return vA.Add(vb.Neg())
}

// MultScalar returns a reference to a new Vector3 with the multiplication between the Vector3 and the scalar
func (vA *Vector3) MultScalar(scalar float32) *Vector3 {
	return &Vector3{scalar * vA.X, scalar * vA.Y, scalar * vA.Z}
}

// DivideScalar returns a reference to a new Vector3 with the division between the Vector3 and the scalar
func (vA *Vector3) DivideScalar(scalar float32) *Vector3 {
	return &Vector3{vA.X / scalar, vA.Y / scalar, vA.Z / scalar}
}

// Magnitude returns the euclidean norm of the vector
func (vA *Vector3) Magnitude() float32 {
	return float32(math.Sqrt(float64(vA.X*vA.X + vA.Y*vA.Y + vA.Z*vA.Z)))
}

// Normalize returns a reference to a new Vector3 with its values normalized using the euclidean norm
func (vA *Vector3) Normalize() *Vector3 {
	return vA.DivideScalar(vA.Magnitude())
}

// Dot returns the dot product between the two Vector3
func (vA *Vector3) Dot(vB *Vector3) float32 {
	return vA.X*vB.X + vA.Y*vB.Y + vA.Z*vB.Z
}

// Cross returns a reference to a new Vector3 containing the cross product between the two Vector3
func (vA *Vector3) Cross(vB *Vector3) *Vector3 {
	return &Vector3{vA.Y*vB.Z - vA.Z*vB.Y, vA.Z*vB.X - vA.X*vB.Z, vA.X*vB.Y - vA.Y*vB.X}
}

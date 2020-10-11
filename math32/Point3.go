package math32

// Point3 represents a 3-dimensional point with values {X, Y, Z}
type Point3 struct {
	X, Y, Z float32
}

// Point3FromArray creates a Point3 from a float array of lenth 3
func Point3FromArray(array [3]float32) *Point3 {
	return &Point3{array[0], array[1], array[2]}
}

// AsArray returs the Point3 values as a float array of lenth 3
func (point *Point3) AsArray() [3]float32 {
	return [3]float32{point.X, point.Y, point.Z}
}

// SetFromArray returs the Point3 values as a float array of lenth 3
func (point *Point3) SetFromArray(array [3]float32) *Point3 {
	point.X = array[0]
	point.Y = array[1]
	point.Z = array[2]
	return point
}

// Neg multiplies each point value by -1
// Returns a pointer to the updated point
func (point *Point3) Neg() *Point3 {
	point.X *= -1.0
	point.Y *= -1.0
	point.Z *= -1.0
	return point
}

// Add adds the values of other to point
// Returns a pointer to the updated point
func (point *Point3) Add(other *Point3) *Point3 {
	point.X += other.X
	point.Y += other.Y
	point.Z += other.Z
	return point
}

// Substract substracts the values of other to point
// Returns a vector to the updated point
func (point *Point3) Substract(other *Point3) *Vector3 {
	return Vector3FromArray(point.Add(other.Neg()).AsArray())
}

// MultScalar multiplies the values of point by scalar
// Returns a pointer to the updated point
func (point *Point3) MultScalar(scalar float32) *Point3 {
	point.X *= scalar
	point.Y *= scalar
	point.Z *= scalar
	return point
}

// DivideScalar divides the values of point by scalar
// Returns a pointer to the updated point
func (point *Point3) DivideScalar(scalar float32) *Point3 {
	point.X /= scalar
	point.Y /= scalar
	point.Z /= scalar
	return point
}

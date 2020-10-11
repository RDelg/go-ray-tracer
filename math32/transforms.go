package math32

import "math"

func cos32(r float32) float32 {
	return float32(math.Cos(float64(r)))
}

func sin32(r float32) float32 {
	return float32(math.Sin(float64(r)))
}

// newScale3 return a new Matrix3 representing a scale transform
// the first 3 dimensions are scaled by x, y, z correspontly
func newScale3(x, y, z float32) *Matrix3 {
	return &Matrix3{
		[3][3]float32{
			[3]float32{x, 0., 0.},
			[3]float32{0., y, 0.},
			[3]float32{0., 0., z},
		}}
}

// newRotateX3 return a new Matrix3 representing a rotation over X axis transform
// r is the radian value of the rotation
func newRotateX3(r float32) *Matrix3 {
	return &Matrix3{
		[3][3]float32{
			[3]float32{1., 0., 0.},
			[3]float32{0., cos32(r), -sin32(r)},
			[3]float32{0., sin32(r), cos32(r)},
		}}
}

// newRotateY3 return a new Matrix3 representing a rotation over Y axis transform
// r is the radian value of the rotation
func newRotateY3(r float32) *Matrix3 {
	return &Matrix3{
		[3][3]float32{
			[3]float32{cos32(r), 0., sin32(r)},
			[3]float32{0., 1., 0.},
			[3]float32{-sin32(r), 0., cos32(r)},
		}}
}

// newRotateZ3 return a new Matrix3 representing a rotation over Z axis transform
// r is the radian value of the rotation
func newRotateZ3(r float32) *Matrix3 {
	return &Matrix3{
		[3][3]float32{
			[3]float32{cos32(r), -sin32(r), 0.},
			[3]float32{-sin32(r), cos32(r), 0.},
			[3]float32{0., 0., 1.},
		}}
}

// newShear3 return a new Matrix3 representing a rotation over Z axis transform
// xy, xz, yx, yz, zx, zy are the shaering coefficients of each dimension respect to the other
func newShear3(xy, xz, yx, yz, zx, zy float32) *Matrix3 {
	return &Matrix3{
		[3][3]float32{
			[3]float32{1., xy, xz},
			[3]float32{yz, 1., yz},
			[3]float32{zx, zy, 1.},
		}}
}

// Scale scales the Vector3 by x, y, z factors
// returns a pointer to the updated Vector3
func (vector *Vector3) Scale(x, y, z float32) *Vector3 {
	return newScale3(x, y, z).MultVector3(vector, true /*columnWise*/)
}

// RotateX rotates the Vector3 around the X axis in r radians
// returns a pointer to the updated Vector3
func (vector *Vector3) RotateX(r float32) *Vector3 {
	return newRotateX3(r).MultVector3(vector, true /*columnWise*/)
}

// RotateY rotates the Vector3 around the Y axis in r radians
// returns a pointer to the updated Vector3
func (vector *Vector3) RotateY(r float32) *Vector3 {
	return newRotateY3(r).MultVector3(vector, true /*columnWise*/)
}

// RotateZ rotates the Vector3 around the Z axis in r radians
// returns a pointer to the updated Vector3
func (vector *Vector3) RotateZ(r float32) *Vector3 {
	return newRotateZ3(r).MultVector3(vector, true /*columnWise*/)
}

// Shear shears the Vector3 by xy, xz, yx, yz, zx, zy factors
// returns a pointer to the updated Vector3
func (vector *Vector3) Shear(xy, xz, yx, yz, zx, zy float32) *Vector3 {
	return newShear3(xy, xz, yx, yz, zx, zy).MultVector3(vector, true /*columnWise*/)
}

// Translate translates the Point3 by x, y, z factors
// returns a pointer to the updated Point3
func (point *Point3) Translate(x, y, z float32) *Point3 {
	point.X += x
	point.Y += y
	point.Z += z
	return point
}

// Scale scales the Point3 by x, y, z factors
// returns a pointer to the updated Point3
func (point *Point3) Scale(x, y, z float32) *Point3 {
	return newScale3(x, y, z).MultPoint3(point, true /*columnWise*/)
}

// RotateX rotates the Point3 around the X axis in r radians
// returns a pointer to the updated Point3
func (point *Point3) RotateX(r float32) *Point3 {
	return newRotateX3(r).MultPoint3(point, true /*columnWise*/)
}

// RotateY rotates the Point3 around the Y axis in r radians
// returns a pointer to the updated Point3
func (point *Point3) RotateY(r float32) *Point3 {
	return newRotateY3(r).MultPoint3(point, true /*columnWise*/)
}

// RotateZ rotates the Point3 around the Z axis in r radians
// returns a pointer to the updated Point3
func (point *Point3) RotateZ(r float32) *Point3 {
	return newRotateZ3(r).MultPoint3(point, true /*columnWise*/)
}

// Shear shears the Point3 by xy, xz, yx, yz, zx, zy factors
// returns a pointer to the updated Point3
func (point *Point3) Shear(xy, xz, yx, yz, zx, zy float32) *Point3 {
	return newShear3(xy, xz, yx, yz, zx, zy).MultPoint3(point, true /*columnWise*/)
}

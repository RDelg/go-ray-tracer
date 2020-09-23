package math32

import "math"

func cos32(r float32) float32 {
	return float32(math.Cos(float64(r)))
}

func sin32(r float32) float32 {
	return float32(math.Sin(float64(r)))
}

// NewTranslation4 return a new Matrix4 representing a translation transform
// the first 3 dimensions are translated by x, y, z correspontly
func NewTranslation4(x, y, z float32) *Matrix4 {
	return &Matrix4{
		[4][4]float32{
			[4]float32{1., 0., 0., x},
			[4]float32{0., 1., 0., y},
			[4]float32{0., 0., 1., z},
			[4]float32{0., 0., 0., 1.},
		}}
}

// NewScaling4 return a new Matrix4 representing a scaling transform
// the first 3 dimensions are scaled by x, y, z correspontly
func NewScaling4(x, y, z float32) *Matrix4 {
	return &Matrix4{
		[4][4]float32{
			[4]float32{x, 0., 0., 1.},
			[4]float32{0., y, 0., 1.},
			[4]float32{0., 0., z, 1.},
			[4]float32{0., 0., 0., 1.},
		}}
}

// NewRotateX4 return a new Matrix4 representing a rotation over X axis transform
// r is the radian value of the rotation
func NewRotateX4(r float32) *Matrix4 {

	return &Matrix4{
		[4][4]float32{
			[4]float32{1., 0., 0., 0.},
			[4]float32{0., cos32(r), -sin32(r), 0.},
			[4]float32{0., sin32(r), cos32(r), 0.},
			[4]float32{0., 0., 0., 1.},
		}}
}

// NewRotateY4 return a new Matrix4 representing a rotation over Y axis transform
// r is the radian value of the rotation
func NewRotateY4(r float32) *Matrix4 {
	return &Matrix4{
		[4][4]float32{
			[4]float32{cos32(r), 0., sin32(r), 0.},
			[4]float32{0., 1., 0., 0.},
			[4]float32{-sin32(r), 0., cos32(r), 0.},
			[4]float32{0., 0., 0., 1.},
		}}
}

// NewRotateZ4 return a new Matrix4 representing a rotation over Z axis transform
// r is the radian value of the rotation
func NewRotateZ4(r float32) *Matrix4 {
	return &Matrix4{
		[4][4]float32{
			[4]float32{cos32(r), -sin32(r), 0., 0.},
			[4]float32{-sin32(r), cos32(r), 0., 0.},
			[4]float32{0., 0., 1., 0.},
			[4]float32{0., 0., 0., 1.},
		}}
}

// NewShearing4 return a new Matrix4 representing a rotation over Z axis transform
// xy, xz, yx, yz, zx, zy are the shaering coefficients of each dimension respect to the other
func NewShearing4(xy, xz, yx, yz, zx, zy float32) *Matrix4 {
	return &Matrix4{
		[4][4]float32{
			[4]float32{1., xy, xz, 0.},
			[4]float32{yz, 1., yz, 0.},
			[4]float32{zx, zy, 1., 0.},
			[4]float32{0., 0., 0., 1.},
		}}
}

package space

// Tuple4DFloat is a 3D arr
type Tuple4DFloat struct {
	Arr [4]float32
}

// Point3D is a Tuple4DFloat with a 1.0 in the last position
type Point3D struct {
	Arr Tuple4DFloat
}

// Vector3D is a Tuple4DFloat  with a 0.0 in the last position
type Vector3D struct {
	Arr Tuple4DFloat
}

// NewPoint3D creates a new Point3D with values x,y,z
func NewPoint3D(x, y, z float32) Point3D {
	return Point3D{Tuple4DFloat{[4]float32{x, y, z, 1.0}}}
}

// NewVector3D creates a new Vector3D with values x,y,z
func NewVector3D(x, y, z float32) Vector3D {
	return Vector3D{Tuple4DFloat{[4]float32{x, y, z, 0.0}}}
}

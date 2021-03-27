package math32

type Object3D interface {
	Intersect(*Ray) (int, []float32)
}

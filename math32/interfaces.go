package math32

type Object interface {
	Intersect(*Ray) (int, []float32)
}

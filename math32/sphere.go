package math32

type Sphere struct {
	Center *Point3
	Radius float32
}

func NewSphere(center *Point3, radius float32) *Sphere {
	return &Sphere{center, radius}
}

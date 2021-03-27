package math32

import "math"

type Sphere struct {
	Center *Point3
	Radius float32
}

func NewSphere(center *Point3, radius float32) *Sphere {
	return &Sphere{center, radius}
}

func (s *Sphere) Intersect(r *Ray) (int32, []float32) {
	sphereToRay := r.Origin.Substract(s.Center)

	a := float64(r.Direction.Dot(r.Direction))
	b := float64(2 * r.Direction.Dot(sphereToRay))
	c := float64(sphereToRay.Dot(sphereToRay) - 1)
	discriminant := b*b - 4.*a*c
	if discriminant < 0 {
		return 0, nil
	}
	t1 := float32((-b - math.Sqrt(discriminant)) / 2 * a)
	t2 := float32((-b + math.Sqrt(discriminant)) / 2 * a)
	return 2, []float32{t1, t2}
}

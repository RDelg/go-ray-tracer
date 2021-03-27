package math32

import "math"

type Ray struct {
	Origin    *Point3
	Direction *Vector3
}

func NewRay(origin *Point3, direction *Vector3) *Ray {
	return &Ray{origin, direction}
}

func (r *Ray) Position(time float32) *Vector3 {
	return &Vector3{
		r.Origin.X + r.Direction.X*time,
		r.Origin.Y + r.Direction.Y*time,
		r.Origin.Z + r.Direction.Z*time,
	}
}

func (r *Ray) Intersect(sphere *Sphere) (int32, []float32) {
	sphereToRay := r.Origin.Substract(sphere.Center)

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

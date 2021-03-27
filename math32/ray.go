package math32

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
	return 2, make([]float32, 2)
}

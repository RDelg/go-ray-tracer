package math32

type Ray struct {
	Origin    *Point3
	Direction *Vector3
}

func NewRay(origin *Point3, direction *Vector3) *Ray {
	return &Ray{origin, direction}
}

func (r *Ray) Position(time float32) *Point3 {
	return &Point3{
		r.Origin.X + r.Direction.X*time,
		r.Origin.Y + r.Direction.Y*time,
		r.Origin.Z + r.Direction.Z*time,
	}
}

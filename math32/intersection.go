package math32

type Intersection struct {
	Time float32
	Obj  Object
}

type Intersections struct {
	Count  uint32
	Memory []Intersection
}

func NewIntersections(size int) *Intersections {
	return &Intersections{0, make([]Intersection, size)}
}

func (i *Intersections) Add(new *Intersection) *Intersections {
	i.Count += 1
	i.Memory = append(i.Memory, *new)
	return i
}

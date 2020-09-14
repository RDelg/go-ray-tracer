package main

import (
	"fmt"

	"github.com/RDelg/go-ray-tracer/space/space"
)

func main() {

	s := space.NewPoint3D(0.1, 0.2, 0.3)

	fmt.Println(s)
}

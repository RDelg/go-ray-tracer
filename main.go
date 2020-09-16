package main

import (
	"fmt"

	"github.com/RDelg/go-ray-tracer/math32"
)

func main() {

	s := math32.Vector3FromArray([3]float32{0.1, 0.2, 0.3})

	fmt.Println(s)
}

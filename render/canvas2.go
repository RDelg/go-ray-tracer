package render

import "github.com/RDelg/go-ray-tracer/math32"

// Canvas2 is a 2D canvas with width and height WxH
type Canvas2 struct {
	W      int
	H      int
	screen [][]math32.Color
}

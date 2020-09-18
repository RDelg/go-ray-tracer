package render

import "github.com/RDelg/go-ray-tracer/math32"

// Canvas2 is a 2D canvas with width and height whH
type Canvas2 struct {
	width  int
	height int
	screen [][]math32.Color
}

// CreateCanvas2 creates a Canvas2 with the given width and height
func CreateCanvas2(width, height int) *Canvas2 {
	matrix := make([][]math32.Color, width)
	for i := 0; i < width; i++ {
		matrix[i] = make([]math32.Color, height)
	}
	return &Canvas2{width: width, height: height, screen: matrix}
}

// WritePixel writes the color in the position x,y in the canvas
// Returns a pointer to the updated canvas
func (c *Canvas2) WritePixel(x, y int, color *math32.Color) *Canvas2 {
	c.screen[x][y] = *color
	return c
}

// ToPPM returns the canvas values as PPM format
// func (c *Canvas2) ToPPM() string {
// 	return
// }

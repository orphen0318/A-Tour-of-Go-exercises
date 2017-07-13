package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	height int
	width  int
}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

// func Rect(x0, y0, x1, y1 int) Rectangle, in image package
func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.width, m.height)
}

/*
type Color interface {
        RGBA() (r, g, b, a uint32)
}
*/
func (m Image) At(x, y int) color.Color {
	img_func := func(x, y int) uint8 {
		return uint8(x ^ y)
	}
	v := img_func(x, y)
	return color.RGBA{v, v, 207, 177}
}

func main() {
	m := Image{256, 64}
	pic.ShowImage(m)
}

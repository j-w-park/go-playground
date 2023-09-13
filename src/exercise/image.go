package exercise

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w int
	h int
}

func (i *Image) At(x, y int) color.Color {
	r := uint8(255.0 * float64(x) / float64(i.w))
	g := uint8(255.0 * float64(y) / float64(i.h))
	b := uint8(255.0 * float64(x+y) / float64(i.w+i.h))
	return color.RGBA{r, g, b, 255}
}

func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func ExerciseImage() {
	m := Image{300, 300}
	pic.ShowImage(&m)
}

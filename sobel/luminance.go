package sobel

import (
	"image"
	"image/color"
)

func luminanceOp(img image.Image, x, y int) color.Gray16 {
	return luminanceGrayScale(img.At(x, y))
}

// luminanceGrayScale applies luminance gray scale conversion
func luminanceGrayScale(c color.Color) color.Gray16 {
	r, g, b, _ := c.RGBA()

	const (
		redFactor   = 0.299
		greenFactor = 0.587
		blueFactor  = 0.114
	)
	intensity := redFactor*float64(r) + greenFactor*float64(g) + blueFactor*float64(b)

	return color.Gray16{Y: uint16(intensity)}
}

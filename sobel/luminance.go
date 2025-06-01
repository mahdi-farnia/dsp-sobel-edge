package sobel

import (
	"image"
	"image/color"
)

func luminanceOp(img image.Image, x, y int) color.Color {
	return luminanceGrayScale(getPixelColor(img, x, y))
}

// luminanceGrayScale applies luminance gray scale conversion
func luminanceGrayScale(c color.Color) color.Color {
	r, g, b, _ := c.RGBA()

	const (
		redFactor   = 0.299
		greenFactor = 0.587
		blueFactor  = 0.114
	)
	intensity := uint32(redFactor*float64(r)+greenFactor*float64(g)+blueFactor*float64(b)) >> 8

	return color.Gray{Y: uint8(intensity)}
}

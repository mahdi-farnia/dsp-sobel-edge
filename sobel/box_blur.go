package sobel

import (
	"image"
	"image/color"
)

// boxBlurOp applies box blur kernel on image
//
// box blur kernel: [[1, 1, 1], [1, 1, 1], [1, 1, 1]] * 1/9
func boxBlurOp(img image.Image, x, y int) color.Color {
	currentWindow := getCurrentWindow(img, x, y)

	// r & g & b are equal because we applied luminance but i keep this because this
	// algorithm works for all images
	acc := struct{ r, g, b uint32 }{}

	for _, v := range currentWindow {
		r, g, b, _ := v.RGBA()
		acc.r += r
		acc.g += g
		acc.b += b
	}

	return color.RGBA{R: uint8(acc.r / 9 >> 8), G: uint8(acc.g / 9 >> 8), B: uint8(acc.b / 9 >> 8), A: 255}
}

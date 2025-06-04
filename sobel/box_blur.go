package sobel

import (
	"image"
	"image/color"
)

// boxBlurOp applies box blur kernel on image
//
// box blur kernel: [[1, 1, 1], [1, 1, 1], [1, 1, 1]] * 1/9
func boxBlurOp(img image.Image, x, y int) color.Gray16 {
	gray := img.(*image.Gray16)
	currentWindow := getCurrentWindow(gray, x, y)

	var acc int
	for _, v := range currentWindow {
		acc += v
	}

	return color.Gray16{Y: uint16(acc / 9)}
}

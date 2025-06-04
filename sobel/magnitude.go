package sobel

import (
	"image"
	"image/color"
)

// calcMagnitudeX calculates the magnitude of X-axis using the following kernel
//
//	kernelX := [9]int32{
//		-1, 0, 2,
//		-2, 0, 2,
//		-1, 0, 1,
//	}
func calcMagnitudeX(img image.Image, x, y int) color.Gray16 {
	gray := img.(*image.Gray16)

	intensities := getCurrentWindow(gray, x, y)

	magnitudeX := -1*int(intensities[0]) + 2*int(intensities[2]) - 2*int(intensities[3]) + 2*int(intensities[5]) -
		1*int(intensities[6]) + int(intensities[8])

	return color.Gray16{Y: uint16(magnitudeX)}
}

// calcMagnitudeY calculates the magnitude of Y-axis using the following kernel
//
//	kernelY := [9]int32{
//		-1, -2, -1,
//		0, 0, 0,
//		1, 2, 1,
//	}
func calcMagnitudeY(img image.Image, x, y int) color.Gray16 {
	gray := img.(*image.Gray16)

	intensities := getCurrentWindow(gray, x, y)

	magnitudeY := -1*intensities[0] - 2*intensities[1] - 1*intensities[2] + intensities[6] + 2*intensities[7] + intensities[8]

	return color.Gray16{Y: uint16(magnitudeY)}
}

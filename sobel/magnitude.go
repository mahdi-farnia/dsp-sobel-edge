package sobel

import (
	"image"
	"image/color"
	"math"
)

func mapMagnitudeOp(img image.Image, x, y int) color.Color {
	// kernelX := [9]int32{
	// 	-1, 0, 2,
	// 	-2, 0, 2,
	// 	-1, 0, 1,
	// }
	// kernelY := [9]int32{
	// 	-1, -2, -1,
	// 	0, 0, 0,
	// 	1, 2, 1,
	// }

	currentWindow := getCurrentWindow(img, x, y)
	r0, _, _, _ := currentWindow[0].RGBA()
	r0 >>= 8
	r1, _, _, _ := currentWindow[1].RGBA()
	r1 >>= 8
	r2, _, _, _ := currentWindow[2].RGBA()
	r2 >>= 8

	r3, _, _, _ := currentWindow[3].RGBA()
	r3 >>= 8
	// r4, _, _, _ := currentWindow[4].RGBA() <-- pixel itself doesn't count!
	r5, _, _, _ := currentWindow[5].RGBA()
	r5 >>= 8

	r6, _, _, _ := currentWindow[6].RGBA()
	r6 >>= 8
	r7, _, _, _ := currentWindow[7].RGBA()
	r7 >>= 8
	r8, _, _, _ := currentWindow[8].RGBA()
	r8 >>= 8

	magnitudeX := int64(-1*int32(r0) + 2*int32(r2) - 2*int32(r3) + 2*int32(r5) - 1*int32(r6) + int32(r8))
	magnitudeY := int64(-1*int32(r0) - 2*int32(r1) - 1*int32(r2) + int32(r6) + 2*int32(r7) + int32(r8))

	magnitude := math.Hypot(float64(magnitudeX*magnitudeX), float64(magnitudeY*magnitudeY))

	return color.Gray{Y: uint8(magnitude)}
}

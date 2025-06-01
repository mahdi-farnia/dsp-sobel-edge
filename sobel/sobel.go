package sobel

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

// SobelOnJpeg applies sobel operator on jpeg images
func SobelOnJpeg(imageFile *os.File) (writer func(outfile *os.File) error, _ error) {
	img, err := jpeg.Decode(imageFile)
	if err != nil {
		return nil, err
	}

	// new image with same size
	newImage := image.NewRGBA(img.Bounds())

	readImageFn(img, newImage, luminanceOp)
	readImageFn(newImage, newImage, boxBlurOp)
	readImageFn(newImage, newImage, mapMagnitudeOp)
	readImageFn(newImage, newImage, sobelOp)

	return func(outfile *os.File) error {
		return jpeg.Encode(outfile, newImage, &jpeg.Options{Quality: jpeg.DefaultQuality})
	}, nil
}

type OperatorFn = func(img image.Image, x, y int) color.Color

// readImageFn maps img pixel color to the corresponding newImg pixel color via op function
func readImageFn(img image.Image, newImg *image.RGBA, op OperatorFn) {
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			newColor := op(img, x, y)
			newImg.Set(x, y, newColor)
		}
	}
}

var maxMag uint8

func sobelOp(img image.Image, x, y int) color.Color {
	if maxMag == 0 {
		for _, m := range img.(*image.RGBA).Pix {
			if m > maxMag {
				maxMag = m
			}
		}
	}
	currentWindow := getCurrentWindow(img, x, y)

	// var acc int
	// for _, c := range currentWindow {
	// 	r, _, _, _ := c.RGBA()
	// 	acc += int(r >> 8) // sum intensities
	// }

	// threshold := uint32(float64(acc) / 9.0) // get mean
	const threshold uint8 = 120

	// current pixel intensity
	magnitude, _, _, _ := currentWindow[4].RGBA()
	magnitude >>= 8

	tmpMag := uint8(float64(magnitude) / float64(maxMag) * 255)
	if threshold > tmpMag {
		return color.Gray{Y: 0}
	}

	return color.Gray{Y: 255}
}

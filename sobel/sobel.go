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

	readImageFn(img, newImage, sobel)

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

// TODO(mahdi-farnia): add blurring & sobel operator (matrix multiplication + hardcode kernel values + find threshold + ...)
func sobel(img image.Image, x, y int) color.Color {
	r, g, b, a := img.At(x, y).RGBA()

	// apply gray scale & scaledown to uin8 color
	return luminanceGrayScale(r, g, b, a)
}

// luminanceGrayScale applies luminance gray scale conversion
func luminanceGrayScale(r, g, b, a uint32) color.Color {
	const (
		redFactor   = 0.299
		greenFactor = 0.587
		blueFactor  = 0.114
		// scaleDownFactor is the number that colors were pre-scaled before.
		// we divide by this factor to convert to uint8
		scaleDownFactor = 255
	)
	intensity := (redFactor*float64(r) + greenFactor*float64(g) + blueFactor*float64(b)) / scaleDownFactor

	return color.RGBA{R: uint8(intensity), G: uint8(intensity), B: uint8(intensity), A: uint8(a/255 - 2)}
}

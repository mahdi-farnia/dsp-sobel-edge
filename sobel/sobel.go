package sobel

import (
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"os"
)

// SobelOnJpeg applies sobel operator on jpeg images
func SobelOnJpeg(imageFile *os.File, threshold uint8) (writer func(outfile *os.File) error, _ error) {
	img, err := jpeg.Decode(imageFile)
	if err != nil {
		return nil, err
	}

	// used for filtering & magnitude matrix calculation
	filterImage := image.NewGray16(img.Bounds())
	magnitudeMatrixX := image.NewGray16(img.Bounds())
	magnitudeMatrixY := image.NewGray16(img.Bounds())
	// used as final image output
	finalImage := image.NewGray16(img.Bounds())

	mapImageFn(img, filterImage, luminanceOp)
	mapImageFn(filterImage, filterImage, boxBlurOp)
	mapImageFn(filterImage, magnitudeMatrixX, calcMagnitudeX)
	mapImageFn(filterImage, magnitudeMatrixY, calcMagnitudeY)

	sobel(magnitudeMatrixX, magnitudeMatrixY, finalImage, threshold)

	return func(outfile *os.File) error {
		return jpeg.Encode(outfile, finalImage, &jpeg.Options{Quality: jpeg.DefaultQuality})
	}, nil
}

type OperatorFn = func(img image.Image, x, y int) color.Gray16

// mapImageFn maps img pixel color to the corresponding newImg pixel color via op function
func mapImageFn(img image.Image, newImg *image.Gray16, op OperatorFn) {
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			newColor := op(img, x, y)
			newImg.Set(x, y, newColor)
		}
	}
}

func sobel(magMatrixX, magMatrixY *image.Gray16, outImage *image.Gray16, threshold uint8) {
	bounds := outImage.Bounds()
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			magX := magMatrixX.At(x, y).(color.Gray16).Y
			magY := magMatrixY.At(x, y).(color.Gray16).Y

			magnitude := uint8(uint16(math.Hypot(float64(magX), float64(magY))) >> 8)
			if magnitude > threshold {
				outImage.Set(x, y, color.Black)
			} else {
				outImage.Set(x, y, color.White)
			}
		}
	}
}

package sobel

import (
	"image"
	"image/color"
)

// getPixelColor returns the pixel if is in range or zero (black)
func getPixelColor(img image.Image, x, y int) color.Color {
	bounds := img.Bounds()
	// the requested color is not in range
	if bounds.Min.X > x || bounds.Min.Y > y || bounds.Max.X < x || bounds.Max.Y < y {
		return color.Black
	}

	return img.At(x, y)
}

// getCurrentWindow gets the current window with radius = 1
func getCurrentWindow(img image.Image, x, y int) [9]color.Color {
	return [9]color.Color{
		getPixelColor(img, x-1, y-1), getPixelColor(img, x, y-1), getPixelColor(img, x+1, y-1),
		getPixelColor(img, x-1, y), getPixelColor(img, x, y), getPixelColor(img, x+1, y),
		getPixelColor(img, x-1, y+1), getPixelColor(img, x, y+1), getPixelColor(img, x+1, y+1),
	}
}

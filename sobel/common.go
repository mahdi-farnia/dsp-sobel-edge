package sobel

import (
	"image"
	"image/color"
)

// getPixelIntensity16 returns the pixel if is in range or zero (black)
func getPixelIntensity16(img image.Image, x, y int) int {
	gray := img.(*image.Gray16)
	bounds := gray.Bounds()
	// the requested color is not in range
	if bounds.Min.X > x || bounds.Min.Y > y || bounds.Max.X < x || bounds.Max.Y < y {
		return 0
	}

	return int(gray.At(x, y).(color.Gray16).Y)
}

// getCurrentWindow gets the current window with radius = 1
func getCurrentWindow(img image.Image, x, y int) [9]int {
	return [9]int{
		getPixelIntensity16(img, x-1, y-1), getPixelIntensity16(img, x, y-1), getPixelIntensity16(img, x+1, y-1),
		getPixelIntensity16(img, x-1, y), getPixelIntensity16(img, x, y), getPixelIntensity16(img, x+1, y),
		getPixelIntensity16(img, x-1, y+1), getPixelIntensity16(img, x, y+1), getPixelIntensity16(img, x+1, y+1),
	}
}

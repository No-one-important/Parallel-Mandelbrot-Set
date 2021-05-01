package main

//https://github.com/esimov/gobrot/blob/b383d69bb3e19484e38cfb8785684dade6862c7f/mandelbrot.go#L149
import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"os"
	"time"
)

const (
	imageWidth    = 4096
	imageHeight   = 4096
	maxIterations = 256
)

var (
	startX     float64
	startY     float64
	endX       float64
	endY       float64
)

func main() {
	startX = -1.4
	startY = -0.8
	endX = 0.4
	endY = 0.8
	startTime := time.Now()
	render()
	endTime := time.Now()
	fmt.Printf("Total time: %d(millis)\n", endTime.Sub(startTime).Milliseconds())
}

func calculateColour(cx, cy float64) uint8 {
	var x, y, xx, yy = 0.0, 0.0, 0.0, 0.0

	for i := 0; i < maxIterations; i++ {
		xy := x * y
		xx = x * x
		yy = y * y
		if xx+yy > 4 {
			return uint8(math.MaxUint8 - i)
		}
		x = xx - yy + cx
		y = 2*xy + cy
	}
	return 0 // black
}

func render() {
	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	for x := 0; x < imageWidth; x++ {
		for y := 0; y < imageHeight; y++ {
			cx := startX + (endX-startX)*float64(x)/float64(imageWidth-1)
			cy := startY + (endY-startY)*float64(y)/float64(imageHeight-1)
			col := calculateColour(cx, cy)
			img.Set(x, y, color.RGBA{
				R: 0,
				G: col,
				B: 0,
				A: math.MaxUint8,
			})
			//log.Printf("x: %d, y: %d, cx: %f, cy: %f, colour: %d", x, y, cx, cy, col)
		}
	}

	output, _ := os.Create("result.jpg")
	_ = jpeg.Encode(output, img, &jpeg.Options{Quality: 100})
}

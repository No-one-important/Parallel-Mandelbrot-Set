package main

//https://github.com/esimov/gobrot/blob/b383d69bb3e19484e38cfb8785684dade6862c7f/mandelbrot.go#L149
import (
	"flag"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math"
	"os"
	"time"
)

const (
	imageWidth    = 4096
	imageHeight   = 4096
	maxIterations = 256

	startX = -1.4
	startY = -0.8
	endX   = 0.4
	endY   = 0.8
)

var (
	pixels      [imageWidth][imageHeight]uint8
	parallelism uint
)

func main() {
	flag.UintVar(&parallelism, "parallelism", 1, "")
	flag.Parse()
	log.Printf("Starting with parallelism %d\n", parallelism)

	startTime := time.Now()
	calculateMandelbrotSet()
	log.Printf("Done calculating Mandelbrot set. Time elapsed: %d (millis)\n", time.Now().Sub(startTime).Milliseconds())
	render()
	log.Printf("Done rendering the picture. Time elapsed (total): %d (millis)\n", time.Now().Sub(startTime).Milliseconds())
}

func calculateMandelbrotSet() {
	for x := 0; x < imageWidth; x++ {
		for y := 0; y < imageHeight; y++ {
			cx := startX + (endX-startX)*float64(x)/float64(imageWidth-1)
			cy := startY + (endY-startY)*float64(y)/float64(imageHeight-1)
			pixels[x][y] = calculateColour(cx, cy)
		}
	}
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
	return 0
}

func render() {
	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	for x := 0; x < imageWidth; x++ {
		for y := 0; y < imageHeight; y++ {
			img.Set(x, y, color.RGBA{
				R: 0,
				G: pixels[x][y],
				B: 0,
				A: math.MaxUint8,
			})
		}
	}

	output, _ := os.Create("result.jpg")
	_ = jpeg.Encode(output, img, &jpeg.Options{Quality: 100})
}

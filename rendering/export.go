package rendering

import (
	"image"
	"image/png"
	"log"
	"os"
)

// Export exports an image as [filename].png.
func Export(img *image.RGBA, filename string) {
	output, err := os.Create(filename + ".png")
	if err != nil {
		log.Fatalf("Could not create result.png. Reason: %s\n", err.Error())
	}
	if err = png.Encode(output, img); err != nil {
		log.Fatalf("Could not construct the result fractal. Reason: %s\n", err.Error())
	}
}

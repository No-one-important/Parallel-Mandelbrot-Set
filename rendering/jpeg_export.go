package rendering

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

func ExportAsJPG(img *image.RGBA, filename string) {
	output, err := os.Create(fmt.Sprintf("%s.png", filename))
	if err != nil {
		log.Fatalf("Could not create result.png. Reason: %s\n", err.Error())
	}
	if err = png.Encode(output, img); err != nil {
		log.Fatalf("Could not construct the result fractal. Reason: %s\n", err.Error())
	}
}

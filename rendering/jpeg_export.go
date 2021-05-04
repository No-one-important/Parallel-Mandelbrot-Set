package rendering

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
)

func ExportAsJPG(img *image.RGBA, filename string) {
	output, err := os.Create(fmt.Sprintf("%s.jpg", filename))
	if err != nil {
		log.Fatalf("Could not create result.jpg. Reason: %s\n", err.Error())
	}
	if err = jpeg.Encode(output, img, &jpeg.Options{Quality: 100}); err != nil {
		log.Fatalf("Could not construct the result fractal. Reason: %s\n", err.Error())
	}
}

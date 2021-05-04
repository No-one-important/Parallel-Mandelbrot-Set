package rendering

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"sync"
)

func GraphThreadSegments(parallelism, granularity, imageWidth, imageHeight int) *image.RGBA {
	threadsColours := make([]color.RGBA, parallelism)
	for tN := 0; tN < parallelism; tN++ {
		threadsColours[tN] = color.RGBA{
			R: uint8((tN + 1) * (math.MaxUint8 / parallelism)),
			G: uint8(math.MaxUint8/2 + powerOfMinus1(tN)*math.MaxUint8*tN/6),
			B: uint8(math.MaxUint8/2 + powerOfMinus1(tN+1)*tN*3),
			A: math.MaxUint8,
		}

		fmt.Printf("Colour %d is %+v\n", tN, threadsColours[tN])
	}

	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	var wg sync.WaitGroup
	wg.Add(parallelism)

	for i := 0; i < parallelism; i++ {
		go func(i int) {
			taskSize := imageHeight / granularity / parallelism

			for taskN := 1; taskN <= granularity; taskN++ {
				startingRow := i*taskSize + (taskN-1)*imageHeight/granularity // small offset + big offset
				endingRow := startingRow + taskSize
				if i == (parallelism - 1) { // If it's the last thread
					// Finish the remainder of this section's rows (needed due to integer arithmetic)
					endingRow = taskN * imageHeight / granularity
				}
				for x := startingRow; x < endingRow; x++ {
					for y := 0; y < imageWidth; y++ {
						img.Set(y, x, threadsColours[i])
					}
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return img
}

func powerOfMinus1(exp int) int {
	if exp%2 == 0 {
		return 1
	}
	return -1
}

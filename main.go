package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/kbinani/screenshot"
)

func save(img *image.RGBA, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
}

func main() {
	n := screenshot.NumActiveDisplays()
	if n <= 0 {
		panic("Active display not found")
	}

	var all image.Rectangle = image.Rect(0, 0, 0, 0)

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		all = bounds.Union(all)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		fileName := time.Now().Format("2006-01-02_15:04:05") + ".png"
		save(img, fileName)
		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
	}

	fmt.Printf("%v\n", all)
	img, err := screenshot.Capture(all.Min.X, all.Min.Y, all.Dx(), all.Dy())
	if err != nil {
		panic(err)
	}
	fileName := time.Now().Format("2006-01-02_15:04:05") + ".png"
	save(img, fileName)
}

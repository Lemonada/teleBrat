package main

import (
	"image/png"
	"os"

	"github.com/kbinani/screenshot"
)

func getscreenshot() {
	n := screenshot.NumActiveDisplays()

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		fileName := "tmp.png"
		file, _ := os.Create(fileName)
		defer file.Close()
		png.Encode(file, img)
	}
}

package app

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func toRGBA(b []byte, w, h int, name string) {
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			px := w*i*3 + j*3
			img.Set(j, h-i, color.RGBA{R: b[px], G: b[px+1], B: b[px+2], A: 255})
		}
	}

	out, _ := os.Create(name)
	defer func() { _ = out.Close() }()

	err := png.Encode(out, img)
	if err != nil {
		log.Println(err)
	}

}

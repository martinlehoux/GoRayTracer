package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

// Pixel contains color between 0 and 255
type Pixel struct {
	r, g, b int
}

// Frame contains image data
type Frame [Height][Width]Pixel

// Save saves frame to PPM file
func (frame Frame) Save(filename string) {
	upLeft, bottomRight := image.Point{0, 0}, image.Point{Width, Height}
	image := image.NewRGBA(image.Rectangle{upLeft, bottomRight})
	for x, row := range frame {
		for y, pixel := range row {
			color := color.RGBA{uint8(pixel.r), uint8(pixel.g), uint8(pixel.b), 0xff}
			image.Set(y, x, color)
		}
	}
	f, _ := os.Create(filename)
	defer f.Close()
	png.Encode(f, image)
}

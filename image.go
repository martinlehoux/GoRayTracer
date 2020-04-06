package main

import (
	"fmt"
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
	f, err := os.Create(filename)
	if err != nil {
		panic(nil)
	}
	defer f.Close()
	f.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", Width, Height))
	for _, row := range frame {
		for _, pixel := range row {
			f.WriteString(fmt.Sprintf("%d %d %d ", pixel.r, pixel.g, pixel.b))
		}
	}
}

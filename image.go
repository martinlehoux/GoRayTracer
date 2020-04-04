package main

import (
	"fmt"
	"os"
)

type Pixel struct {
	r, g, b int
}

type Frame [HEIGHT][WIDTH]Pixel

func (frame Frame) Save(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(nil)
	}
	defer f.Close()
	f.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", WIDTH, HEIGHT))
	for _, row := range frame {
		for _, pixel := range row {
			f.WriteString(fmt.Sprintf("%d %d %d ", pixel.r, pixel.g, pixel.b))
		}
	}
}

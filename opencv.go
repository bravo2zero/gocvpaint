package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

// Get Mat from image file path string
func GetImg(path string) *gocv.Mat {
	img := gocv.IMRead(path, gocv.IMReadColor)
	if img.Empty() {
		panic(fmt.Sprintf("Error reading image from: %s", path))
	}
	return &img
}

type Color []uint8

// Get color of Mat pixel
func GetColor(m gocv.Mat, x, y int) Color {
	ch := m.Channels()
	v := make(Color, ch)
	for c := 0; c < ch; c++ {
		v[c] = m.GetUCharAt(y, x*ch+c)
	}
	return v
}

// Set Mat pixel to color
func SetColor(m gocv.Mat, x, y int, color Color) {
	ch := m.Channels()
	for c := 0; c < ch; c++ {
		m.SetUCharAt(y, x*ch+c, color[c])
	}
}

// Return inverted color
func Inverse(c Color) Color {
	inv := make(Color, len(c))
	for idx, value := range c {
		inv[idx] = 255 - value
	}
	return inv
}

// Show Mats list in separate windows (useful to check image processing steps visually)
func Display(imgs ...*gocv.Mat) {
	var windows = make([]*gocv.Window, len(imgs))

	for idx, img := range imgs {
		windows[idx] = gocv.NewWindow(fmt.Sprintf("window-%v", idx))
		windows[idx].IMShow(*img)
		defer windows[idx].Close()
	}
	for {
		for _, window := range windows {
			if window.WaitKey(1) >= 0 {
				return
			}
		}
	}
}

package main

import (
	"fmt"
	"image"

	"gocv.io/x/gocv"
)

func main() {

	orig := GetImg("img/battleblocktheater.jpeg")
	defer orig.Close()

	work := orig.Clone()
	gocv.CvtColor(*orig, &work, gocv.ColorBGRAToGray)
	gocv.GaussianBlur(work, &work, image.Point{5, 5}, 0, 0, gocv.BorderDefault)
	//gocv.Threshold(work, &work, 60, 255, gocv.ThresholdToZero)
	// kernel := gocv.GetStructuringElement(gocv.MorphRect, image.Pt(3, 3))
	// gocv.MorphologyEx(work, &work, gocv.MorphGradient, kernel)
	gocv.Canny(work, &work, 100, 200)

	Display(orig, &work)
}

func GetImg(path string) *gocv.Mat {
	img := gocv.IMRead(path, gocv.IMReadColor)
	if img.Empty() {
		panic(fmt.Sprintf("Error reading image from: %s", path))
	}
	return &img
}

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

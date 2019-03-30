package main

import (
	"flag"
	"image"

	"gocv.io/x/gocv"
)

const (
	KMeansClusters = 12
)

func main() {

	inputArg := flag.String("i", "", "input image path")
	outputArg := flag.String("o", "", "output image path")
	displayArg := flag.Bool("d", false, "display input and output image windows (bool)")

	flag.Parse()

	orig := GetImg(*inputArg)
	defer orig.Close()

	work := orig.Clone()
	gocv.CvtColor(*orig, &work, gocv.ColorBGRAToGray)
	gocv.GaussianBlur(work, &work, image.Point{5, 5}, 0, 0, gocv.BorderDefault)
	gocv.Canny(work, &work, 100, 200)

	size := work.Size()

	for y := 0; y < size[0]; y++ {
		for x := 0; x < size[1]; x++ {
			c := GetColor(work, x, y)
			SetColor(work, x, y, Inverse(c))
		}
	}

	if len(*outputArg) > 0 {
		gocv.IMWrite(*outputArg, work)
	}

	if *displayArg {
		Display(orig, &work)
	}
}

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			if math.IsInf(ay, 0) {
				continue
			}
			bx, by := corner(i, j)
			if math.IsInf(by, 0) {
				continue
			}
			cx, cy := corner(i, j+1)
			if math.IsInf(cy, 0) {
				continue
			}
			dx, dy := corner(i+1, j+1)
			if math.IsInf(dy, 0) {
				continue
			}
			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g,'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	funcs := [3]func(float64, float64) float64{chickenEggBox, mogle, horseSaddle}

	var f func(float64, float64) float64
	if len(os.Args) < 2 {
		f = chickenEggBox
	} else {
		i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		if i >= len(funcs) {
			fmt.Fprint(os.Stderr, "You can specify a number 0 or 1, 2.")
			os.Exit(1)
		}
		f = funcs[i]
	}

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func mogle(x, y float64) float64 {
	return math.Sin(x) * math.Sin(y) / 10.
}

func chickenEggBox(x, y float64) float64 {
	return (math.Sin(x) + math.Sin(y)) / 20.
}

func horseSaddle(x, y float64) float64 {
	return (0.1*x*x - 0.2*y*y) / 25.
}

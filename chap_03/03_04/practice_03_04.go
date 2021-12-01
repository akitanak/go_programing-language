package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	var (
		width  = 600
		height = 320
	)
	queries := r.URL.Query()
	if v, err := strconv.Atoi(queries.Get("width")); err == nil {
		width = v
	}
	if v, err := strconv.Atoi(queries.Get("height")); err == nil {
		height = v
	}
	writeChickenEggBox(w, width, height)
}

const (
	cells = 100
	angle = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func writeChickenEggBox(w io.Writer, width, height int) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j, width, height)
			if math.IsInf(ay, 0) {
				continue
			}
			bx, by, bz := corner(i, j, width, height)
			if math.IsInf(by, 0) {
				continue
			}
			cx, cy, cz := corner(i, j+1, width, height)
			if math.IsInf(cy, 0) {
				continue
			}
			dx, dy, dz := corner(i+1, j+1, width, height)
			if math.IsInf(dy, 0) {
				continue
			}
			color := color((az + bz + cz + dz) / 4.)
			fmt.Fprintf(w, "<polygon style='fill: #%s' points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				color, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j, width, height int) (float64, float64, float64) {
	xyrange := 30.0
	xyscale := float64(width) / 2. / xyrange
	zscale := float64(height) * 0.4
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	return (math.Sin(x) + math.Sin(y)) / 20.
}

const (
	colorMax   uint32  = 0xff0000
	colorMin   uint32  = 0x0000ff
	colorWidth uint32  = colorMax - colorMin
	zRange     float64 = 0.1
)

func color(z float64) string {
	ratio := (z + zRange) / (zRange * 2)
	r := uint8(255. * ratio)
	b := 255 - uint8(255.*ratio)
	return fmt.Sprintf("%02x00%02x", r, b)
}

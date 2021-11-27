package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xFF},
	color.RGBA{0x00, 0x22, 0x00, 0xFF},
	color.RGBA{0x00, 0x44, 0x00, 0xFF},
	color.RGBA{0x00, 0x66, 0x00, 0xFF},
	color.RGBA{0x00, 0x88, 0x00, 0xFF},
	color.RGBA{0x00, 0xAA, 0x00, 0xFF},
	color.RGBA{0x00, 0xCC, 0x00, 0xFF},
	color.RGBA{0x00, 0xFF, 0x00, 0xFF},
}

const (
	backgroundIndex = 0
	lineIndex       = 1
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	var (
		cycle = 5
		size  = 100
	)
	queries := r.URL.Query()
	if v, err := strconv.Atoi(queries.Get("cycle")); err == nil {
		cycle = v
	}
	if v, err := strconv.Atoi(queries.Get("size")); err == nil {
		size = v
	}
	lissajous(w, cycle, size)
}

func lissajous(out io.Writer, cycles int, size int) {
	const (
		res     = 0.001
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			paletteSize := len(palette)
			lineColorIndex := int((y+1.0)*0.5*float64(paletteSize)) % paletteSize
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), uint8(lineColorIndex))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

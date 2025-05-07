package main

// go build lissajous.go
// ./lissajous > out.gif

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.Black, color.RGBA{R: 0, G: 255, B: 0, A: 255}, color.RGBA{R: 0, G: 0, B: 255, A: 255}, color.RGBA{R: 100, G: 100, B: 100, A: 255}, color.RGBA{R: 0, G: 255, B: 255, A: 255}, color.RGBA{R: 255, G: 255, B: 0, A: 255}}

const (
	whiteIndex = 0 // Первый цвет палитры
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		color_index := uint8(rand.Uint32() * 5)
		color_index = color_index % 5
		if color_index == 0 {
			color_index += 1
		}

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// println(color_index)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), color_index)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

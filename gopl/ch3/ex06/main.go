package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"

	"github.com/Phrog195/Golang/gopl/ch3/ex06/supersampling"
)

var palette = []color.Color{
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0xff, 0xa5, 0x00, 0xff},
	color.RGBA{0xff, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0x80, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
	color.RGBA{0x4b, 0x00, 0x82, 0xff},
	color.RGBA{0xee, 0x82, 0xee, 0xff},
	// color.RGBA{0xf4, 0x43, 0x36, 0xff}, // Material Red
	// color.RGBA{0xff, 0xff, 0xff, 0xff}, // White
	// color.RGBA{0x4c, 0xaf, 0x50, 0xff}, // Material Green
	// color.RGBA{0xff, 0xff, 0xff, 0xff}, // White
	// color.RGBA{0x21, 0x96, 0xf3, 0xff}, // Material Blue
	// color.RGBA{0xff, 0xff, 0xff, 0xff}, // White
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y0 := (float64(py)+0.5)/height*(ymax-ymin) + ymin
		y1 := (float64(py)-0.5)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x0 := (float64(px)+0.5)/width*(xmax-xmin) + xmin
			x1 := (float64(px)-0.5)/width*(xmax-xmin) + xmin
			z1 := complex(x0, y0)
			z2 := complex(x1, y0)
			z3 := complex(x0, y1)
			z4 := complex(x1, y1)
			color := supersampling.Supersampling([]color.Color{
				mandelbrot(z1),
				mandelbrot(z2),
				mandelbrot(z3),
				mandelbrot(z4),
			})
			// Image point (px, py) represents complex value z.
			img.Set(px, py, color)
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200

	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return palette[n%len(palette)]
		}
	}
	return color.Black

	// const iterations = 200
	// const contrast = 15

	// var v complex128
	// for n := uint8(0); n < iterations; n++ {
	// 	v = v*v + z
	// 	if cmplx.Abs(v) > 2 {
	// 		return color.Gray{255 - contrast*n}
	// 	}
	// }
	// return color.Black
}

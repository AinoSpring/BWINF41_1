package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

type Display struct {
	image *image.NRGBA
	size  Vector
}

func NewDisplay(size Vector) (display Display) {
	display.size = size
	display.image = image.NewNRGBA(image.Rect(0, 0, display.size.x, display.size.y))
	for y := 0; y < display.size.y; y++ {
		for x := 0; x < display.size.x; x++ {
			display.image.Set(x, y, color.NRGBA{
				R: 0,
				G: 0,
				B: 255,
				A: 255,
			})
		}
	}
	return
}

func (display Display) Save(path string) {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, display.image); err != nil {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

func (display *Display) SetPixel(position Vector, color color.Color) {
	display.image.Set(position.x, position.y, color)
}

func DisplayField(field Field) (display Display) {
	display = NewDisplay(field.size)
	for y := 0; y < field.size.y; y++ {
		for x := 0; x < field.size.x; x++ {
			if positionPointer := field.Position(Vector{x: x, y: y}); positionPointer.notNull {
				display.SetPixel(Vector{x: x, y: y}, color.Gray{Y: positionPointer.seed.orientation})
			}
		}
	}
	return
}

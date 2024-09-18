package main

import (
	// "fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"golang.org/x/tour/pic"
)

type Image struct {
	x, y, width, height int
	color               uint8
}

func (p Image) Bounds() image.Rectangle {
	return image.Rect(p.x, p.y, p.width, p.height)
}

func (p Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (p Image) At(x, y int) color.Color {
	return color.RGBA{p.color + uint8(x), p.color + uint8(y), 255, 255}
}

func TryImage() {
	// m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	// fmt.Println(m.Bounds())
	// fmt.Println(m.At(10, 0).RGBA())
	m := Image{0, 0, 100, 100, 100}
	pic.ShowImage(m)

	file, err := os.Create("image.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Encode writes the Image m to w in PNG format.
	err = png.Encode(file, m)
	if err != nil {
		panic(err)
	}
}

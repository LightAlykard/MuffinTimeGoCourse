package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	green := color.RGBA{0, 255, 0, 255}
	red := color.RGBA{200, 30, 30, 255}
	rectImg := image.NewRGBA(image.Rect(0, 0, 500, 500))

	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)
	for x := 20; x < 480; x++ {
		y := 20
		rectImg.Set(x, y, red)
		y = 480
		rectImg.Set(x, y, red)
	}
	for y := 20; y < 480; y++ {
		x := 20
		rectImg.Set(x, y, red)
		x = 480
		rectImg.Set(x, y, red)
	}

	file, err := os.Create("rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}

	defer file.Close()
	png.Encode(file, rectImg)
}

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

const (
	rectX    = 1000
	rectY    = 1000
	fileName = "chessboard.png"
)

func main() {
	border := color.RGBA{R: 110, G: 50, B: 50, A: 255}
	blackk := color.RGBA{R: 0, G: 0, B: 0, A: 255}
	whitee := color.RGBA{R: 255, G: 255, B: 255, A: 255}

	dX, dY := rectX/10, rectY/10

	rectImg := image.NewRGBA(image.Rect(0, 0, rectX, rectY))
	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{C: border}, image.ZP, draw.Src)
	draw.Draw(rectImg, image.Rect(dX, dY, rectX-dX, rectY-dY), &image.Uniform{C: blackk}, image.ZP, draw.Src)

	var k int
	for j := 1; j < 9; j++ {
		for i := 1; i < 9; i += 2 {
			if j%2 == 0 {
				k = i + 1
			} else {
				k = i
			}
			draw.Draw(rectImg, image.Rect(k*dX, j*dY, k*dX+dX, j*dY+dY),
				&image.Uniform{C: whitee}, image.ZP, draw.Src)
		}
	}

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()

	err = png.Encode(file, rectImg)
	if err != nil {
		log.Fatalf("Failed encode file: %s", err)
	}

	fmt.Println("Image ", fileName, "done!")
}

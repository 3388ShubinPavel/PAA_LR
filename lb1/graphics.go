package main

import (
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"
)

func showGraphic(N int, squares []Square) {
	cellSize := 50
	imgWidth, imgHeight := N*cellSize, N*cellSize
	img := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))

	for i := 0; i <= N; i++ {
		for x := 0; x < imgWidth; x++ {
			img.Set(x, i*cellSize, color.Black)
		}
		for y := 0; y < imgHeight; y++ {
			img.Set(i*cellSize, y, color.Black)
		}
	}

	rand.Seed(time.Now().UnixNano())
	for _, square := range squares {
		x, y, size := square.x*cellSize, square.y*cellSize, square.size*cellSize
		r := uint8(rand.Intn(256))
		g := uint8(rand.Intn(256))
		b := uint8(rand.Intn(256))
		col := color.RGBA{R: r, G: g, B: b, A: 255}

		for dx := 0; dx < size; dx++ {
			for dy := 0; dy < size; dy++ {
				img.Set(x+dx, y+dy, col)
			}
		}

		borderColor := color.RGBA{R: 0, G: 0, B: 0, A: 255}
		for dx := 0; dx < size; dx++ {
			img.Set(x+dx, y, borderColor)
			img.Set(x+dx, y+size-1, borderColor)
		}
		for dy := 0; dy < size; dy++ {
			img.Set(x, y+dy, borderColor)
			img.Set(x+size-1, y+dy, borderColor)
		}
	}

	outFile, err := os.Create("./lb1/images/squares.png")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	png.Encode(outFile, img)
}

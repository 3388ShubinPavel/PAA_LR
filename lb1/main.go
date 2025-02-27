package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"
)

type Square struct {
	x, y, size int
}

func (s Square) String() string {
	return fmt.Sprintf("%d %d %d", s.x+1, s.y+1, s.size)
}

var minSquares = 999999
var bestResult []Square
var iterationsCnt int

func main() {
	benchmark := flag.Bool("benchmark", false, "Run mode")
	flag.Parse()
	if *benchmark {
		Benchmark()
		return
	}
	var N int
	fmt.Print("Enter N: ")
	fmt.Scan(&N)

	start := time.Now()
	occupied := make([][]bool, N)
	for i := range occupied {
		occupied[i] = make([]bool, N)
	}

	newGridSize, squareSize := ScaleSize(N)

	if newGridSize != N {
		fmt.Printf("Scaled grid size: %d, Square size: %d\n", newGridSize, squareSize)
		occupied := make([][]bool, newGridSize)
		for i := range occupied {
			occupied[i] = make([]bool, newGridSize)
		}
		Solve(occupied, []Square{}, newGridSize, squareSize)

		finalResult := []Square{}
		for _, square := range bestResult {
			finalResult = append(finalResult, Square{
				x:    square.x * squareSize,
				y:    square.y * squareSize,
				size: square.size * squareSize,
			})
		}
		bestResult = finalResult
	} else {
		occupied := make([][]bool, N)
		for i := range occupied {
			occupied[i] = make([]bool, N)
		}
		initialSquare := placeInitialSquares(N, occupied)
		Solve(occupied, initialSquare, N, 1)
	}
	duration := time.Since(start)
	fmt.Println("Time to solve:", duration)
	fmt.Println("Iterations:", iterationsCnt)
	fmt.Println(minSquares)
	for _, square := range bestResult {
		fmt.Println(square.String())
	}

	showGraphic(N, bestResult)
}

func ScaleSize(gridSize int) (int, int) {
	maxDivisor := 1
	for i := gridSize / 2; i >= 1; i-- {
		if gridSize%i == 0 {
			maxDivisor = i
			break
		}
	}
	squareSize := maxDivisor
	return gridSize / maxDivisor, squareSize
}

func placeInitialSquares(N int, occupied [][]bool) []Square {
	squares := []Square{}
	size1 := (N + 1) / 2
	squares = append(squares, placeSquare(0, 0, size1, occupied))
	size2 := N / 2
	squares = append(squares, placeSquare(0, size1, size2, occupied))
	squares = append(squares, placeSquare(size1, 0, size2, occupied))
	return squares
}

func Solve(occupied [][]bool, current []Square, gridSize, scale int) {
	iterationsCnt++
	pos := findFirstFreePosition(occupied, gridSize)
	if pos == -1 {
		if len(current) < minSquares {
			minSquares = len(current)
			bestResult = append([]Square{}, current...)

			fmt.Println("--- New Best Result ---")
			for _, square := range bestResult {
				fmt.Println(square.String())
			}
			fmt.Println("-----------------------")
		}

		return
	}

	x, y := pos/gridSize, pos%gridSize
	maxSz := min(min(gridSize-x, gridSize-y), gridSize-1)

	for size := maxSz; size >= 1; size-- {
		if canPlace(x, y, size, occupied) {
			square := placeSquare(x, y, size, occupied)
			current = append(current, square)
			if len(current) < minSquares {
				Solve(occupied, current, gridSize, scale)
			}
			current = current[:len(current)-1]
			removeSquare(square, occupied)
		}
		if len(current) >= minSquares {
			break
		}
	}
}

func findFirstFreePosition(occupied [][]bool, N int) int {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if !occupied[i][j] {
				return i*N + j
			}
		}
	}
	return -1
}

func canPlace(x, y, size int, occupied [][]bool) bool {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if occupied[x+i][y+j] {
				return false
			}
		}
	}
	return true
}

func placeSquare(x, y, size int, occupied [][]bool) Square {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			occupied[x+i][y+j] = true
		}
	}
	return Square{x, y, size}
}

func removeSquare(square Square, occupied [][]bool) {
	for i := 0; i < square.size; i++ {
		for j := 0; j < square.size; j++ {
			occupied[square.x+i][square.y+j] = false
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

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

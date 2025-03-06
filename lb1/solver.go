package main

import (
	"fmt"
	"strings"
)

func solveAndDisplay(N int) {
	newGridSize, squareSize := ScaleSize(N)

	if newGridSize != N {
		fmt.Printf("Scaled grid size: %d, Square size: %d\n", newGridSize, squareSize)
		solveScaled(newGridSize, squareSize)
	} else {
		solveOriginal(N)
	}

	showGraphic(N, bestResult)
}

func solveScaled(gridSize, scale int) {
	occupied := initializeGrid(gridSize)
	initialSquares := placeInitialSquares(gridSize, occupied)
	Solve(occupied, initialSquares, gridSize, scale, 0)

	finalResult := upscaleSquares(bestResult, scale)
	bestResult = finalResult
}

func solveOriginal(gridSize int) {
	occupied := initializeGrid(gridSize)
	initialSquares := placeInitialSquares(gridSize, occupied)
	Solve(occupied, initialSquares, gridSize, 1, 0)
}

func upscaleSquares(squares []Square, scale int) []Square {
	result := []Square{}
	for _, square := range squares {
		result = append(result, Square{
			x:    square.x * scale,
			y:    square.y * scale,
			size: square.size * scale,
		})
	}
	return result
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

func Solve(occupied [][]bool, current []Square, gridSize, scale, depth int) {
	iterationsCnt++
	pos := findFirstFreePosition(occupied, gridSize)

	if pos == -1 {
		indent := strings.Repeat("  ", depth)
		fmt.Printf("%sCompleted configuration with %d squares\n", indent, len(current))

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
	indent := strings.Repeat("  ", depth)
	fmt.Printf("%sFound free position at (%d, %d)\n", indent, x, y)
	maxSz := Min(Min(gridSize-x, gridSize-y), gridSize-1)

	for size := maxSz; size >= 1; size-- {
		indent := strings.Repeat("  ", depth)
		fmt.Printf("%sAttempting square at (%d, %d) size %d\n", indent, x, y, size)

		if canPlace(x, y, size, occupied) {
			square := placeSquare(x, y, size, occupied)
			current = append(current, square)

			fmt.Printf("%sPlaced square at (%d, %d) size %d\n", indent, x, y, size)

			if len(current) < minSquares {
				Solve(occupied, current, gridSize, scale, depth+1)
			}

			fmt.Printf("%sRemoving square at (%d, %d) size %d\n", indent, x, y, size)
			current = current[:len(current)-1]
			removeSquare(square, occupied)
		}

		if len(current) >= minSquares {
			break
		}
	}
}

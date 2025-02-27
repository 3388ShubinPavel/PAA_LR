package main

import "fmt"

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
	Solve(occupied, initialSquares, gridSize, scale)

	finalResult := upscaleSquares(bestResult, scale)
	bestResult = finalResult
}

func solveOriginal(gridSize int) {
	occupied := initializeGrid(gridSize)
	initialSquares := placeInitialSquares(gridSize, occupied)
	Solve(occupied, initialSquares, gridSize, 1)
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

func Solve(occupied [][]bool, current []Square, gridSize, scale int) {
	iterationsCnt++
	pos := findFirstFreePosition(occupied, gridSize)
	if pos == -1 {
		if len(current) < minSquares {
			minSquares = len(current)
			bestResult = append([]Square{}, current...)
		}
		return
	}

	x, y := pos/gridSize, pos%gridSize
	maxSz := Min(Min(gridSize-x, gridSize-y), gridSize-1)
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

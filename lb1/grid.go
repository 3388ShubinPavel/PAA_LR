package main

func initializeGrid(size int) [][]bool {
	grid := make([][]bool, size)
	for i := range grid {
		grid[i] = make([]bool, size)
	}
	return grid
}

func ScaleSize(gridSize int) (int, int) {
	maxDivisor := 1
	for i := gridSize / 2; i >= 1; i-- {
		if gridSize%i == 0 {
			maxDivisor = i
			break
		}
	}
	return gridSize / maxDivisor, maxDivisor
}

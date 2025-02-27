package main

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
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

package main

import (
	"flag"
	"fmt"
	"time"
)

var minSquares = 999999
var bestResult []Square
var iterationsCnt int

type Square struct {
	x, y, size int
}

func (s Square) String() string {
	return fmt.Sprintf("%d %d %d", s.x+1, s.y+1, s.size)
}

func main() {
	benchmark := flag.Bool("benchmark", false, "Run mode")
	flag.Parse()

	if *benchmark {
		Benchmark()
		return
	}

	N := getGridSizeFromUser()
	start := time.Now()

	solveAndDisplay(N)

	duration := time.Since(start)
	fmt.Println("Time to solve:", duration)
	fmt.Println("Iterations:", iterationsCnt)
	fmt.Println(minSquares)
	for _, square := range bestResult {
		fmt.Println(square.String())
	}
}

func getGridSizeFromUser() int {
	var N int
	fmt.Print("Enter N: ")
	fmt.Scan(&N)
	return N
}

package main

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"log"
)

func Benchmark() {
	var data []struct {
		N          int
		Iterations int
	}

	for N := 2; N <= 40; N++ {
		if !IsPrime(N) {
			continue
		}
		/**
		Update variables for each N
		*/
		iterationsCnt = 0
		minSquares = 999999
		bestResult = []Square{}

		occupied := make([][]bool, N)
		for i := range occupied {
			occupied[i] = make([]bool, N)
		}
		newGridSize, squareSize := ScaleSize(N)
		if newGridSize != N {
			occupied := make([][]bool, newGridSize)
			for i := range occupied {
				occupied[i] = make([]bool, newGridSize)
			}
			Solve(occupied, []Square{}, newGridSize, squareSize)
		} else {
			initialSquare := placeInitialSquares(N, occupied)
			Solve(occupied, initialSquare, N, 1)
		}

		data = append(data, struct {
			N          int
			Iterations int
		}{N: N, Iterations: iterationsCnt})

		fmt.Printf("Processed N=%d, Iterations=%d\n", N, iterationsCnt)
	}

	p := plot.New()
	points := make(plotter.XYs, len(data))
	for i, d := range data {
		points[i].X = float64(d.N)
		points[i].Y = float64(d.Iterations)
	}

	scatter, err := plotter.NewScatter(points)
	if err != nil {
		log.Fatal(err)
	}
	scatter.GlyphStyle.Color = plotutil.Color(0)
	scatter.GlyphStyle.Radius = vg.Length(1)

	line, err := plotter.NewLine(points)
	if err != nil {
		log.Fatal(err)
	}
	line.LineStyle.Color = plotutil.Color(1)
	line.LineStyle.Width = vg.Points(1)

	p.Add(scatter, line)
	p.Legend.Add("Iterations", scatter)

	p.Title.Text = "Growth of Iterations vs N (Prime Numbers Only)"
	p.X.Label.Text = "N (Prime Numbers)"
	p.Y.Label.Text = "Number of Iterations"

	if err := p.Save(8*vg.Inch, 8*vg.Inch, "./lb1/images/iterations.png"); err != nil {
		log.Fatal(err)
	}
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

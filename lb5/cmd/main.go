package main

import (
	"fmt"
)

func main() {
	var task int
	fmt.Println("Choose task:")
	fmt.Println("1. Multiple pattern search")
	fmt.Println("2. Wildcard pattern search")
	fmt.Println("3. Non-overlapping pattern search")
	fmt.Scan(&task)

	verbose := true

	switch task {
	case 1:
		handleMultiplePatternSearch(verbose)
	case 2:
		handleWildcardSearch(verbose)
	case 3:
		handleNonOverlappingSearch(verbose)
	default:
		fmt.Println("Invalid task choice")
	}
}

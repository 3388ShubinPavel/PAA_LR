package main

import (
	"fmt"
	"sort"
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
		var text string
		var n int
		fmt.Print("Enter text: ")
		fmt.Scan(&text)
		fmt.Print("Number of patterns: ")
		fmt.Scan(&n)

		patterns := make([]string, n)
		for i := 0; i < n; i++ {
			fmt.Printf("Pattern %d: ", i+1)
			fmt.Scan(&patterns[i])
		}

		ac := NewAhoCorasick(patterns, verbose)
		if verbose {
			ac.PrintAutomaton()
		}

		results := ac.Search(text)

		fmt.Println("\nResults:")
		for _, res := range results {
			fmt.Printf("%d %d\n", res.Pos, res.PatternIndex)
		}

	case 2:
		var text, pattern, wildcardStr string
		fmt.Print("Enter text: ")
		fmt.Scan(&text)
		fmt.Print("Enter pattern: ")
		fmt.Scan(&pattern)
		fmt.Print("Enter wildcard character: ")
		fmt.Scan(&wildcardStr)
		wildcard := []rune(wildcardStr)[0]

		substrings, positions := splitPattern(pattern, wildcard)
		if verbose {
			fmt.Println("\nSplit pattern:")
			for i, s := range substrings {
				fmt.Printf("  Part %d: '%s' at position %d\n", i+1, s, positions[i])
			}
		}

		if len(substrings) == 0 {
			fmt.Println("No valid subpatterns found")
			return
		}

		m := len(text)
		n := len(pattern)
		if n > m {
			fmt.Println("Pattern longer than text")
			return
		}

		ac := NewAhoCorasick(substrings, verbose)
		if verbose {
			ac.PrintAutomaton()
		}

		occurrences := ac.Search(text)
		if verbose {
			fmt.Println("\nFound subpatterns:")
			for _, occ := range occurrences {
				fmt.Printf("  Pattern %d at position %d\n", occ.PatternIndex, occ.Pos)
			}
		}

		C := make([]int, m-n+1)
		for _, occ := range occurrences {
			idx := occ.PatternIndex - 1
			li := positions[idx]
			j := occ.Pos - 1
			iStart := j - li
			if iStart >= 0 && iStart <= m-n {
				C[iStart]++
				if verbose {
					fmt.Printf("Increment C[%d] (now %d)\n", iStart+1, C[iStart])
				}
			}
		}

		var results []int
		for i := 0; i < len(C); i++ {
			if C[i] == len(substrings) {
				results = append(results, i+1)
			}
		}

		sort.Ints(results)
		fmt.Println("\nResults:")
		for _, pos := range results {
			fmt.Println(pos)
		}
	case 3:
		var text string
		var n int
		fmt.Print("Enter text: ")
		fmt.Scan(&text)
		fmt.Print("Number of patterns: ")
		fmt.Scan(&n)

		patterns := make([]string, n)
		for i := 0; i < n; i++ {
			fmt.Printf("Pattern %d: ", i+1)
			fmt.Scan(&patterns[i])
		}

		ac := NewAhoCorasick(patterns, verbose)
		if verbose {
			ac.PrintAutomaton()
		}

		results := ac.Search(text)
		nonOverlapping := filterNonOverlapping(results, ac.patternLengths)

		fmt.Println("\nNon-overlapping results:")
		var positions []int
		for _, res := range nonOverlapping {
			positions = append(positions, res.Pos)
		}
		sort.Ints(positions)
		for _, pos := range positions {
			fmt.Println(pos)
		}
	default:
		fmt.Println("Invalid task choice")
	}
}

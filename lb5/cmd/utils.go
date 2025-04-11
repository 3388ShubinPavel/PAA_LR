package main

import (
	"fmt"
	"sort"
)

func handleMultiplePatternSearch(verbose bool) {
	var text string
	var n int
	fmt.Print("Enter text: ")
	fmt.Scan(&text)
	fmt.Print("Number of patterns: ")
	fmt.Scan(&n)

	patterns := inputPatterns(n)
	ac := NewAhoCorasick(patterns, verbose)
	printAutomatonIfVerbose(ac, verbose)
	results := ac.Search(text)
	printResults(results)
}

func handleWildcardSearch(verbose bool) {
	var text, pattern, wildcardStr string
	fmt.Print("Enter text: ")
	fmt.Scan(&text)
	fmt.Print("Enter pattern: ")
	fmt.Scan(&pattern)
	fmt.Print("Enter wildcard character: ")
	fmt.Scan(&wildcardStr)
	wildcard := []rune(wildcardStr)[0]

	substrings, positions := splitPattern(pattern, wildcard)
	printSplitPattern(substrings, positions, verbose)

	if len(substrings) == 0 {
		fmt.Println("No valid subpatterns found")
		return
	}

	if len(pattern) > len(text) {
		fmt.Println("Pattern longer than text")
		return
	}

	ac := NewAhoCorasick(substrings, verbose)
	printAutomatonIfVerbose(ac, verbose)
	occurrences := ac.Search(text)
	printSubpatterns(occurrences, verbose)

	matches := findWildcardMatches(occurrences, positions, len(text), len(pattern), len(substrings), verbose)
	printWildcardResults(matches)
}

func handleNonOverlappingSearch(verbose bool) {
	var text string
	var n int
	fmt.Print("Enter text: ")
	fmt.Scan(&text)
	fmt.Print("Number of patterns: ")
	fmt.Scan(&n)

	patterns := inputPatterns(n)
	ac := NewAhoCorasick(patterns, verbose)
	printAutomatonIfVerbose(ac, verbose)
	results := ac.Search(text)
	nonOverlapping := filterNonOverlapping(results, ac.patternLengths)
	printNonOverlappingResults(nonOverlapping)
}

func inputPatterns(n int) []string {
	patterns := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Pattern %d: ", i+1)
		fmt.Scan(&patterns[i])
	}
	return patterns
}

func printAutomatonIfVerbose(ac *AhoCorasick, verbose bool) {
	if verbose {
		ac.PrintAutomaton()
	}
}

func printResults(results []Pair) {
	fmt.Println("\nResults:")
	for _, res := range results {
		fmt.Printf("%d %d\n", res.Pos, res.PatternIndex)
	}
}

func printSplitPattern(substrings []string, positions []int, verbose bool) {
	if verbose {
		fmt.Println("\nSplit pattern:")
		for i, s := range substrings {
			fmt.Printf("  Part %d: '%s' at position %d\n", i+1, s, positions[i])
		}
	}
}

func printSubpatterns(occurrences []Pair, verbose bool) {
	if verbose {
		fmt.Println("\nFound subpatterns:")
		for _, occ := range occurrences {
			fmt.Printf("  Pattern %d at position %d\n", occ.PatternIndex, occ.Pos)
		}
	}
}

func findWildcardMatches(occurrences []Pair, positions []int, textLen, patternLen, subsCount int, verbose bool) []int {
	C := make([]int, textLen-patternLen+1)
	for _, occ := range occurrences {
		idx := occ.PatternIndex - 1
		li := positions[idx]
		j := occ.Pos - 1
		iStart := j - li
		if iStart >= 0 && iStart <= textLen-patternLen {
			C[iStart]++
			if verbose {
				fmt.Printf("Increment C[%d] (now %d)\n", iStart+1, C[iStart])
			}
		}
	}

	var results []int
	for i := 0; i < len(C); i++ {
		if C[i] == subsCount {
			results = append(results, i+1)
		}
	}
	sort.Ints(results)
	return results
}

func printWildcardResults(matches []int) {
	fmt.Println("\nResults:")
	for _, pos := range matches {
		fmt.Println(pos)
	}
}

func printNonOverlappingResults(nonOverlapping []Pair) {
	fmt.Println("\nNon-overlapping results:")
	var positions []int
	for _, res := range nonOverlapping {
		positions = append(positions, res.Pos)
	}
	sort.Ints(positions)
	for _, pos := range positions {
		fmt.Println(pos)
	}
}

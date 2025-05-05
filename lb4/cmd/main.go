package main

import (
	"bufio"
	"fmt"
	"os"
)

func computePrefixFunction(p string) []int {
	m := len(p)
	pi := make([]int, m)
	k := 0
	for q := 1; q < m; q++ {
		for k > 0 && p[k] != p[q] {
			k = pi[k-1]
		}
		if p[k] == p[q] {
			k++
		}
		pi[q] = k
	}
	return pi
}

func kmpSearch(text, pattern string) []int {
	pi := computePrefixFunction(pattern)
	q := 0
	var result []int
	for i := 0; i < len(text); i++ {
		for q > 0 && pattern[q] != text[i] {
			q = pi[q-1]
		}
		if pattern[q] == text[i] {
			q++
		}
		if q == len(pattern) {
			result = append(result, i-len(pattern)+1)
			q = pi[q-1]
		}
	}
	return result
}

func solveTask1(reader *bufio.Reader) {
	var pattern, text string
	fmt.Fscanln(reader, &pattern)
	fmt.Fscanln(reader, &text)
	matches := kmpSearch(text, pattern)
	if len(matches) == 0 {
		fmt.Println(-1)
		return
	}
	for i, pos := range matches {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print(pos)
	}
	fmt.Println()
}

func solveTask2(reader *bufio.Reader) {
	var A, B string
	fmt.Fscanln(reader, &A)
	fmt.Fscanln(reader, &B)
	if len(A) != len(B) {
		fmt.Println(-1)
		return
	}
	doubleA := A + A
	matches := kmpSearch(doubleA, B)
	for _, pos := range matches {
		if pos < len(A) {
			fmt.Println(pos)
			return
		}
	}
	fmt.Println(-1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var task int
	fmt.Fscanln(reader, &task)
	switch task {
	case 1:
		solveTask1(reader)
	case 2:
		solveTask2(reader)
	default:
		fmt.Fprintln(os.Stderr, "Unknown task. Use 1 or 2.")
	}
}

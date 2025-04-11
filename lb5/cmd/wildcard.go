package main

func splitPattern(pattern string, wildcard rune) (substrings []string, positions []int) {
	current := []rune{}
	for i, c := range pattern {
		if c != wildcard {
			current = append(current, c)
		} else {
			if len(current) > 0 {
				substrings = append(substrings, string(current))
				positions = append(positions, i-len(current))
				current = []rune{}
			}
		}
	}
	if len(current) > 0 {
		substrings = append(substrings, string(current))
		positions = append(positions, len(pattern)-len(current))
	}
	return
}

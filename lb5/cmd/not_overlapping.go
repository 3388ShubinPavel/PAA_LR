package main

import "sort"

func filterNonOverlapping(pairs []Pair, patternLengths []int) []Pair {
	if len(pairs) == 0 {
		return nil
	}

	type Interval struct {
		start, end int
		pair       Pair
	}

	intervals := make([]Interval, len(pairs))
	for i, p := range pairs {
		length := patternLengths[p.PatternIndex-1]
		intervals[i] = Interval{
			start: p.Pos,
			end:   p.Pos + length - 1,
			pair:  p,
		}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].end < intervals[j].end
	})

	var result []Pair
	lastEnd := -1

	for _, iv := range intervals {
		if iv.start > lastEnd {
			result = append(result, iv.pair)
			lastEnd = iv.end
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Pos < result[j].Pos
	})

	return result
}

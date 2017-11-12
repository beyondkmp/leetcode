package main

import "sort"

type Interval struct {
	Start int
	End   int
}

type Intervals []Interval

func (in Intervals) Len() int {
	return len(in)
}
func (in Intervals) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}
func (in Intervals) Less(i, j int) bool {
	return in[i].Start < in[j].Start
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func merge(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return []Interval{}
	} else if len(intervals) == 1 {
		return intervals
	}

	sort.Sort(Intervals(intervals))
	var result []Interval
	for _, interval := range intervals {
		if len(result) >= 1 && interval.Start <= result[len(result)-1].End {
			result[len(result)-1].End = max(interval.End, result[len(result)-1].End)
		} else {
			result = append(result, interval)
		}
	}
	return result
}

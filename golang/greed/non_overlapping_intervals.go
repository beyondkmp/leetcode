package main

import (
	"fmt"
	"sort"
)

// Definition for an interval.
type Interval struct {
	Start int
	End   int
}

type ByEndIntervals []Interval

func (e ByEndIntervals) Len() int {
	return len(e)
}

func (e ByEndIntervals) Swap(x, y int) {
	e[x], e[y] = e[y], e[x]
}

func (e ByEndIntervals) Less(x, y int) bool {
	return e[x].End < e[y].End
}

func eraseOverlapIntervals(intervals []Interval) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Sort(ByEndIntervals(intervals))
	nums := 1
	compare := intervals[0]
	for _, v := range intervals {
		if v.Start > compare.End {
			nums += 1
			compare = v
		}
	}
	return len(intervals) - nums
}

func main() {
	intervals := []Interval{Interval{1, 2}, Interval{2, 3}, Interval{2, 4}}
	fmt.Println(eraseOverlapIntervals(intervals))
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	begin, end int
}

func (r Range) Length() int {
	return r.end - r.begin + 1
}

func ReadRange(s string) Range {
	split := strings.Split(s, "-")
	begin, _ := strconv.Atoi(split[0])
	end, _ := strconv.Atoi(split[1])

	return Range{
		begin: begin,
		end: end,
	}
}

func InRange(n int, r Range) bool {
	return r.begin <= n && n <= r.end
}

func SortRangesAndMerge(rs *[]Range) {
	sort.Slice(*rs, func(i, j int) bool {
	    return (*rs)[i].begin < (*rs)[j].begin
	})

	var merged []Range
	top_index := -1
	for _, r := range *rs {
		if top_index >= 0 && merged[top_index].end >= r.begin - 1 {
			merged[top_index].end = max(merged[top_index].end, r.end)
		} else {
			merged = append(merged, r)
			top_index++
		}
	}
	*rs = merged
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	var ranges []Range
	for i := 0; scanner.Scan(); i++ {
		input := scanner.Text()
		if input == "" {
			break
		}
		ranges = append(ranges, ReadRange(input))
	}
	SortRangesAndMerge(&ranges)

	var count int
	for _, r := range ranges {
		count += r.Length()
	}

	fmt.Println("Result " + strconv.Itoa(count))
}

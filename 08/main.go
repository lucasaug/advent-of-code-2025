package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"container/heap"
)

const MAX_COUNT = 1000

type Point struct {
	x int
	y int
	z int
}

func (p Point) distance(other Point) float64 {
	return math.Pow(float64(p.x - other.x), 2.) +
 	    math.Pow(float64(p.y - other.y), 2.) +
 	    math.Pow(float64(p.z - other.z), 2.)
}

type DistanceEntry struct {
	value [2]uint
	d float64
	index int
}

type PriorityQueue []*DistanceEntry

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].d < pq[j].d
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*DistanceEntry)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *DistanceEntry, value [2]uint, d float64) {
	item.value = value
	item.d = d
	heap.Fix(pq, item.index)
}

func main() {
	var points []Point

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		coords := strings.Split(input, ",")

		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])

		points = append(points, Point{x, y, z})
	}

	pq := make(PriorityQueue, (len(points)*len(points) - len(points))/2)
	i := 0
	for i_a, point_a := range points {
		for i_b, point_b := range points {
			if i_b >= i_a {
				break
			}

			dist := point_a.distance(point_b)

			pq[i] = &DistanceEntry{
				value: [2]uint{uint(i_a), uint(i_b)},
				d: dist,
				index: i,
			}
			i++
		}	
	}
	heap.Init(&pq)

	circuits := make([]int, len(points))
	for i := range circuits {
		circuits[i] = -1
	}

	var result int
	var circuit_index int
	for {
		if len(pq) == 0 {
			break
		}
		item := heap.Pop(&pq).(*DistanceEntry)

		if circuits[item.value[0]] == -1 && circuits[item.value[1]] == -1 {
			circuits[item.value[0]] = circuit_index
			circuits[item.value[1]] = circuit_index
			circuit_index++
		} else if circuits[item.value[0]] == -1 {
			circuits[item.value[0]] = circuits[item.value[1]]
		} else if circuits[item.value[1]] == -1 {
			circuits[item.value[1]] = circuits[item.value[0]]
		} else {
			value_to_replace := circuits[item.value[1]]
			for i, v := range circuits {
				if v == value_to_replace {
					circuits[i] = circuits[item.value[0]]
				}
			}
		}

		if circuits[0] != -1 {
			common_circuit := circuits[0]
			fully_connected := true

			for _, v := range circuits {
				if v != common_circuit {
					fully_connected = false
					break
				}
			}

			if fully_connected {
				result = points[item.value[0]].x * points[item.value[1]].x
				break
			}
		}
	}

	// circuit_counts := make([]int, circuit_index)
	// for _, v := range circuits {
	// 	if v != -1 {
	// 		circuit_counts[v]++
	// 	}
	// }
	//
	// slices.Sort(circuit_counts)
	// slices.Reverse(circuit_counts)
	// result := circuit_counts[0] * circuit_counts[1] * circuit_counts[2]

	fmt.Println("Result " + strconv.Itoa(result))
}

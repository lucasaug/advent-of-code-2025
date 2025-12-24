package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func main() {
	var points []Point

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		coords := strings.Split(input, ",")

		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])

		points = append(points, Point{x, y})
	}

	min_coords := points[0]
	max_coords := points[0]

	for _, p := range points {
		if p.x < min_coords.x {
			min_coords.x = p.x
		}
		if p.y < min_coords.y {
			min_coords.y = p.y
		}
		if p.x > max_coords.x {
			max_coords.x = p.x
		}
		if p.y > max_coords.y {
			max_coords.y = p.y
		}
	}

	var result int
	for _, p_1 := range points {
		for _, p_2 := range points {
			if p_1 == p_2 {
				continue
			}

			signed_area := float64((p_2.x - p_1.x + 1) * (p_2.y - p_1.y + 1))
			area := int(math.Abs(signed_area))
			if area > result {
				result = int(area)
			}
		}
	}

	fmt.Println("Result " + strconv.Itoa(result))
}

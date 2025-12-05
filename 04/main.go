package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MOVABLE_COUNT = 4

type Grid struct {
	has_roll [][]bool 
	max_x int
	max_y int
}

func (g Grid) get(x, y int) bool {
	if x < 0 || y < 0 || x >= g.max_x || y >= g.max_y {
		return false
	}

	return g.has_roll[x][y]
}

func (g *Grid) set_row(x int, s string) {
	var row []bool
	for _, c := range s {
		value := false
		if c == '@' {
			value = true
		}
		row = append(row, value)
	}

	g.has_roll[x] = row
}

func (g Grid) CountNeighborhood(x, y int) (count int){
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if g.get(x + i, y + j) {
				count += 1
			}
		}
	}
	return
}

func (g Grid) IsMovable(x, y int) bool {
	if g.get(x, y) && g.CountNeighborhood(x, y) <= MOVABLE_COUNT {
		return true
	}
	return false
}

func NewGrid(max_x, max_y int) Grid {
	roll_grid := make([][]bool, max_x)
	for range max_x {
		roll_grid = append(roll_grid, make([]bool, max_y))
	}

	return Grid{
		has_roll: roll_grid,
		max_x: max_x,
		max_y: max_y,
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	var grid Grid
	for i := 0; scanner.Scan(); i++ {
		input := scanner.Text()
		if i == 0 {
			grid = NewGrid(len(input), len(input))
		}
		grid.set_row(i, input)
	}

	should_continue := true
	var count int

	for should_continue {
		new_grid := NewGrid(grid.max_x, grid.max_y)
		copy(new_grid.has_roll, grid.has_roll)
		should_continue = false

		for i := range grid.max_x {
			for j := range grid.max_y {
				new_grid.has_roll[i][j] = grid.has_roll[i][j]

				if grid.IsMovable(i, j) {
					new_grid.has_roll[i][j] = false
					count += 1
					should_continue = true
				}
			}
		}
	}

	fmt.Println("Result " + strconv.Itoa(count))
}

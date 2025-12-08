package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	var beam_count []int

	for i := 0; scanner.Scan(); i++ {
		input := scanner.Text()

		if len(beam_count) == 0 {
			starting_index := strings.Index(input, "S")
			beam_count = make([]int, len(input))
			beam_count[starting_index] = 1
		} else {
			new_count := make([]int, len(beam_count))
			copy(new_count, beam_count)

			for beam_index, value := range beam_count {
				if value == 0 || input[beam_index] != '^' {
					continue
				}

				new_count[beam_index] -= value
				if beam_index > 0 {
					new_count[beam_index-1] += value
				}
				if beam_index < len(beam_count) - 1 {
					new_count[beam_index+1] += value
				}
			}

			beam_count = new_count
		}
	}

	var count int
	for _, value := range beam_count {
		count += value
	}

	fmt.Println("Result " + strconv.Itoa(count))
}

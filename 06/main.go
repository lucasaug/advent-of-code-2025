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
	
	lines := []string{}
	max_line_length := 0

	for i := 0; scanner.Scan(); i++ {
		input := scanner.Text()
		lines = append(lines, input)
		if len(input) > max_line_length {
			max_line_length = len(input)
		}
	}

	var sum int
	current_acc := []int{}

	last_line_index := len(lines) - 1
	for i := range max_line_length {
		index := max_line_length - 1 - i

		value_str := ""
		for col := range last_line_index {
			if index < len(lines[col]) {
				value_str += string(lines[col][index])
			}
		}

		value_str = strings.TrimSpace(value_str)
		if value_str == "" {
			continue
		} else {
			value, _ := strconv.Atoi(value_str)
			current_acc = append(current_acc, value)
		}

		char := byte(' ')
		if index < len(lines[last_line_index]) {
			char = lines[last_line_index][index]
		}

		if char != ' ' {
			value := 0
			if char == '+' {
				for _, entry := range current_acc {
					value += entry
				}
			} else {
				value = 1
				for _, entry := range current_acc {
					value *= entry
				}
			}

			sum += value
			current_acc = []int{}
		}

	}

	fmt.Println("Result " + strconv.Itoa(sum))
}

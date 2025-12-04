package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const BATTERY_COUNT = 12

func MaxJoltage(entry string) (result int) {
	last_index := -1

	for current_index := range BATTERY_COUNT {
		max_digit_index := last_index + 1
		remaining_digit_count := BATTERY_COUNT - current_index - 1

		for i := max_digit_index; i < len(entry) - remaining_digit_count; i++ {
			if entry[i] > entry[max_digit_index] {
				max_digit_index = i
			}
		}

		last_index = max_digit_index
		current_digit := entry[max_digit_index] - '0'
		result += int(math.Pow(10, float64(remaining_digit_count))) * int(current_digit)
	}

	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	sum := 0
	for scanner.Scan() {
		input := scanner.Text()
		sum += MaxJoltage(input)
	}

	fmt.Println("Result: " + strconv.Itoa(sum))
}

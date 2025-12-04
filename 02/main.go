package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ExtractRange(entry string) (begin, end int) {
	interval := strings.Split(entry, "-")
	begin, _ = strconv.Atoi(interval[0])
	end, _ = strconv.Atoi(interval[1])

	return
}

func IsSillyPattern(value string) bool {
	for stride := 1; stride <= len(value)/2; stride++ {
		if len(value) % stride != 0 {
			continue
		}

		match := true
		for offset := 0; offset + stride <= len(value); offset += stride {
			if value[offset:offset+stride] != value[:stride] {
				match = false
				break
			}
		}

		if !match {
			continue
		}

		return true
	}

	return false
}

func SumRange(entry string) (sum int) {
	begin, end := ExtractRange(entry)

	current := begin
	for current <= end {
		if IsSillyPattern(strconv.Itoa(current)) {
			sum += current
		}
		current += 1
	}

	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	entries := strings.Split(input, ",")
	sum := 0
	for _, entry := range entries {
		sum += SumRange(entry)
	}

	fmt.Println("Result: " + strconv.Itoa(sum))
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func ReadRotation(s string) int {
	value, _ := strconv.Atoi(s[1:])

	if s[0] == byte('L') {
		return -value
	}

	return value
}

func Sign(v int) int {
	if v > 0 {
		return 1
	} else if v < 0 {
		return -1
	}

	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	rotation := 50
	password := 0

	for scanner.Scan() {
		text := scanner.Text()
		delta := ReadRotation(text)
		fmt.Println("----")
		fmt.Println("Rotating by " + strconv.Itoa(delta))
		intermediate_rotation := rotation + delta

		previous_sign := Sign(rotation)
		current_sign := Sign(intermediate_rotation)

		password_delta := int(math.Abs(float64(intermediate_rotation))) / 100

		if previous_sign != 0 && previous_sign != current_sign {
			password_delta += 1
		}

		password += int(password_delta)
		fmt.Println("Pass increased by " + strconv.Itoa(int(password_delta)))

		// Update dial value
		rotation = (intermediate_rotation % 100)
		fmt.Println("Dial currently at " + strconv.Itoa(rotation))
	}

	fmt.Println(password)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

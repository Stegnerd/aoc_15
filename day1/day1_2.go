package day1

import (
	"fmt"
	"os"
)

func Part2Run() *int {
	data, err := os.ReadFile("./day1/day1_input.txt")
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		return nil
	}

	count := 0

	for pos, char := range data {
		if char == '(' {
			count++
		} else if char == ')' {
			count--
		}

		if count < 0 {
			fmt.Printf("\n hits the negative for the first time at pos: %d \n", pos)
			break
		}
	}

	fmt.Printf("\n output: %d \n", count)

	return &count
}

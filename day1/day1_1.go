package day1

import (
	"fmt"
	"os"
)

func Part1Run() *int {
	data, err := os.ReadFile("./day1/day1_input.txt")
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		return nil
	}

	count := 0
	for _, char := range data {
		if char == '(' {
			count++
		} else if char == ')' {
			count--
		}
	}

	fmt.Printf("\n output: %d \n", count)

	return &count
}

package day3

import (
	"fmt"
	"os"
)

type point struct {
	x int
	y int
}

func Part1Run() *int {
	data, err := os.ReadFile("./day3/day3_input.txt")
	if err != nil {
		fmt.Printf("Failed to read file: %+v", err)
	}

	startingPoint := point{0, 0}
	houses := map[point]int{}
	for _, char := range data {
		if char == 'v' {
			startingPoint.y = startingPoint.y - 1
		} else if char == '<' {
			startingPoint.x = startingPoint.x - 1
		} else if char == '>' {
			startingPoint.x = startingPoint.x + 1
		} else if char == '^' {
			startingPoint.y = startingPoint.y + 1
		}

		val, found := houses[startingPoint]
		if !found {
			houses[startingPoint] = val + 1
		} else {
			houses[startingPoint] = 1
		}
	}

	houseCount := len(houses)
	fmt.Printf("\n house count: %+v \n", houseCount)
	return &houseCount
}

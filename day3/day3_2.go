package day3

import (
	"fmt"
	"os"
)

func Part2Run() *int {
	data, err := os.ReadFile("./day3/day3_input.txt")
	if err != nil {
		fmt.Printf("Failed to read file: %+v", err)
	}

	normalSanta := point{0, 0}
	roboSanta := point{0, 0}
	houses := map[point]int{}
	houses[point{0, 0}] = 2
	isRoboSanta := false
	for _, char := range data {
		if char == 'v' {
			if isRoboSanta {
				roboSanta.y = roboSanta.y - 1
			} else {
				normalSanta.y = normalSanta.y - 1
			}
		} else if char == '<' {
			if isRoboSanta {
				roboSanta.x = roboSanta.x - 1
			} else {
				normalSanta.x = normalSanta.x - 1
			}
		} else if char == '>' {
			if isRoboSanta {
				roboSanta.x = roboSanta.x + 1
			} else {
				normalSanta.x = normalSanta.x + 1
			}
		} else if char == '^' {
			if isRoboSanta {
				roboSanta.y = roboSanta.y + 1
			} else {
				normalSanta.y = normalSanta.y + 1
			}
		}

		var pointToUse point
		if isRoboSanta {
			pointToUse = roboSanta
		} else {
			pointToUse = normalSanta
		}

		val, found := houses[pointToUse]
		if !found {
			houses[pointToUse] = val + 1
		} else {
			houses[pointToUse] = 1
		}

		isRoboSanta = !isRoboSanta
	}

	houseCount := len(houses)
	fmt.Printf("\n house count: %+v \n", houseCount)
	return &houseCount
}

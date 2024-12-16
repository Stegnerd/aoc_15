package day2

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Part1Run() *int {
	data, err := os.Open("./day2/day2_input.txt")
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		return nil
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		splits := strings.Split(line, "x")

		length, _ := strconv.Atoi(splits[0])
		width, _ := strconv.Atoi(splits[1])
		height, _ := strconv.Atoi(splits[2])

		dimensions := []int{length, width, height}
		sort.Ints(dimensions)

		area := dimensions[0] * dimensions[1]
		sa := 2*length*width + 2*width*height + 2*height*length

		result := sa + area
		total += result

	}

	fmt.Printf("\n day2 part 1 total: %d \n", total)

	return &total
}

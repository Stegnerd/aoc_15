/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"aoc15/day1"
	"aoc15/day2"
	"aoc15/day3"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Which day to test")

	day, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Failed to read input: %+v", err)
	}
	day = strings.TrimSpace(day)

	fmt.Println("which part to test")
	part, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Failed to read input: %+v", err)
	}
	part = strings.TrimSpace(part)

	fmt.Printf("running day: %s and part: %s\n", day, part)

	match(day, part)
}

func match(day string, part string) {
	switch day {
	case "1":
		switch part {
		case "1":
			day1.Part1Run()
		case "2":
			day1.Part2Run()
		}
	case "2":
		switch part {
		case "1":
			day2.Part1Run()
		case "2":
			day2.Part2Run()
		}
	case "3":
		switch part {
		case "1":
			day3.Part1Run()
		case "2":
			day3.Part2Run()
		}
	}
}

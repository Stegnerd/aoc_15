/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"aoc15/day1"
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

	fmt.Printf("running day: %s and part: %s", day, part)

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
	}
}

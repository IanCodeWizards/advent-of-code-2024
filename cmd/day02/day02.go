package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"aoc/internal/utils"
)

func main() {
	// read form file
	input, err := utils.ReadFile("resources/day02.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", Part1(input))
	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", Part2(input))

	os.Exit(0)
}

func splitStrings(input string) []int {
	result := strings.Split(input, "\r\n")

	var list1 []int

	for _, y := range result {
		res := strings.Split(y, "   ")

		num1, _ := strconv.Atoi(res[0])

		list1 = append(list1, num1)
	}

	return list1
}

// part one
func Part1(input string) int {
	nums := splitStrings(input)

	prev := 0
	for _, y := range nums {

	}

	return 0
}

// part two
func Part2(input string) int {
	return 0
}

package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"aoc/internal/utils"
)

func main() {
	input, err := utils.ReadFile("resources/day01.txt")
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

// part one
func Part1(input string) int {
	list1, list2 := splitStrings(input)

	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})
	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	var runningTotal float64 = 0
	for index, _ := range list1 {
		runningTotal = runningTotal + math.Abs(float64(list1[index])-float64(list2[index]))
	}

	return int(runningTotal)
}

func splitStrings(input string) ([]int, []int) {
	result := strings.Split(input, "\r\n")

	var list1, list2 []int

	for _, y := range result {
		res := strings.Split(y, "   ")

		num1, _ := strconv.Atoi(res[0])
		num2, _ := strconv.Atoi(res[1])

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	return list1, list2
}

// part two
func Part2(input string) int {
	list1, list2 := splitStrings(input)

	total := 0
	for _, y := range list1 {
		total = total + calcScore(y, list2)
	}

	return total
}

func calcScore(num int, slice []int) int {
	count := 0
	for _, s := range slice {
		if s == num {
			count++
		}

	}
	return num * count
}

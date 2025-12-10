package main

import (
	"aoc/helpers"
	"fmt"
	"strings"
	"time"
)

func main() {
	data := helpers.ReadFile("input/2025/day02.txt")
	var idRanges []string = strings.Split(data, ",")

	var invalidSum1 int
	var invalidSum2 int

	start := time.Now()
	for _, idRange := range idRanges {
		firstId := helpers.StringToInt(strings.Split(idRange, "-")[0])
		lastId := helpers.StringToInt(strings.Split(idRange, "-")[1])

		for id := firstId; id <= lastId; id++ {
			if isRepeatedId(id) {
				invalidSum1 += id
			}

			if isRepeatedSequenceId(id) {
				invalidSum2 += id
			}
		}
	}

	fmt.Printf("time: %v\n", time.Since(start))
	fmt.Printf("invalid id sum 1: %s\n", fmt.Sprint(invalidSum1))
	fmt.Printf("invalid id sum 2: %s\n", fmt.Sprint(invalidSum2))
}

func isRepeatedId(id int) bool {
	idString := fmt.Sprint(id)
	length := len(idString)

	if length%2 != 0 {
		return false
	}

	midpoint := length / 2
	firstHalf := idString[:midpoint]
	secondHalf := idString[midpoint:]

	return firstHalf == secondHalf
}

func isRepeatedSequenceId(id int) bool {
	idString := fmt.Sprint(id)
	length := len(idString)

	if length < 2 {
		return false
	}

	doubled := idString + idString
	return strings.Contains(doubled[1:length*2-1], idString)
}

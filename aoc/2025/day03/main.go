package main

import (
	"aoc/helpers"
	"fmt"
	"strings"
)

func main() {
	data := helpers.ReadFile("input/2025/day03.txt")
	var banks []string = strings.Split(data, "\n")

	totalJoltage1 := 0
	totalJoltage2 := 0

	for _, bank := range banks {
		if len(bank) == 0 {
			continue
		}

		totalJoltage1 += getHighestBankVoltage(bank, 2)
		totalJoltage2 += getHighestBankVoltage(bank, 12)
	}

	fmt.Printf("total joltage 1: %s\n", fmt.Sprint(totalJoltage1))
	fmt.Printf("total joltage 2: %s\n", fmt.Sprint(totalJoltage2))
}

func getHighestBankVoltage(bank string, digits int) int {
	if len(bank) < digits {
		panic("wtf")
	}

	result := 0
	startPos := 0

	for digitsLeft := digits; digitsLeft > 0; digitsLeft-- {
		maxAllowedPos := len(bank) - digitsLeft
		maxDigit := -1
		maxDigitPos := -1

		for i := startPos; i <= maxAllowedPos; i++ {
			digit := int(bank[i] - '0')
			if digit > maxDigit {
				maxDigit = digit
				maxDigitPos = i
			}
		}

		result = joinInts(result, maxDigit)
		startPos = maxDigitPos + 1
	}

	return result
}

func joinInts(i int, j int) int {
	return i*10 + j
}

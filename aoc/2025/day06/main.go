package main

import (
	"aoc/helpers"
	"fmt"
	"strings"
)

type Problem struct {
	numbers  []string
	operator string
}

func main() {
	data := helpers.ReadFile("input/2025/day06.txt")

	problemData := strings.Split(data, "\n")
	problems := parseProblems(problemData)

	grandTotal := 0
	grandTotalCephalopod := 0

	for _, problem := range problems {
		grandTotal += calculateProblem(problem)
		grandTotalCephalopod += calculateCephalopodProblem(problem)
	}

	fmt.Printf("grandTotal: %v\n", grandTotal)
	fmt.Printf("grandTotalCephalopod: %v\n", grandTotalCephalopod)
}

func calculateProblem(problem Problem) int {
	toInt := func(s string) int {
		return helpers.StringToInt(strings.TrimSpace(s))
	}

	result := toInt(problem.numbers[0])

	switch problem.operator {
	case "+":
		for i := 1; i < len(problem.numbers); i++ {
			result += toInt(problem.numbers[i])
		}
	case "*":
		for i := 1; i < len(problem.numbers); i++ {
			result *= toInt(problem.numbers[i])
		}
	default:
		panic("wtf")
	}

	return result
}

func calculateCephalopodProblem(problem Problem) int {
	if len(problem.numbers) == 0 {
		return 0
	}

	width := len(problem.numbers[0])
	var newNumbers []string

	for col := width - 1; col >= 0; col-- {
		var str string

		for _, row := range problem.numbers {
			if col < len(row) && row[col] != ' ' {
				str += string(row[col])
			}
		}

		if str != "" {
			newNumbers = append(newNumbers, str)
		}
	}

	return calculateProblem(Problem{
		numbers:  newNumbers,
		operator: problem.operator,
	})
}

func parseProblems(lines []string) []Problem {
	var problems []Problem

	maxWidth := 0
	for _, l := range lines {
		if len(l) > maxWidth {
			maxWidth = len(l)
		}
	}

	startCol := 0

	for col := 0; col <= maxWidth; col++ {
		isSeparator := true

		if col < maxWidth {
			for _, line := range lines {
				if col < len(line) && line[col] != ' ' {
					isSeparator = false
					break
				}
			}
		}

		if isSeparator {
			if col > startCol {
				problems = append(problems, extractBlock(lines, startCol, col))
			}
			startCol = col + 1
		}
	}

	return problems
}

func extractBlock(lines []string, start, end int) Problem {
	var problem Problem
	lastLineIdx := len(lines) - 1

	if lastLineIdx >= 0 {
		opLine := lines[lastLineIdx]
		s, e := start, end

		if s < len(opLine) {
			if e > len(opLine) {
				e = len(opLine)
			}

			problem.operator = strings.TrimSpace(opLine[s:e])
		}
	}

	for i := range lines[:lastLineIdx] {
		line := lines[i]

		val := ""
		if start < len(line) {
			safeEnd := end
			if safeEnd > len(line) {
				safeEnd = len(line)
			}
			val = line[start:safeEnd]
		}

		requiredLen := end - start
		if len(val) < requiredLen {
			val = fmt.Sprintf("%-*s", requiredLen, val)
		}

		problem.numbers = append(problem.numbers, val)
	}

	return problem
}

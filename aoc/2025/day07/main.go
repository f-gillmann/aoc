package main

import (
	"aoc/helpers"
	"fmt"
	"slices"
	"strings"
)

type Direction int

const (
	None Direction = iota
	Right
	Left
)

func main() {
	data := helpers.ReadFile("input/2025/day07.txt")
	lines := strings.Split(data, "\n")

	totalTimeLines := 0

	totalSplits := 0
	for i := range lines {
		if len(lines) > i+1 {
			nextLine, splits := getNextLineWithBeams(lines[i], lines[i+1], None)
			lines[i+1] = nextLine
			totalSplits += splits
		}
	}

	fmt.Printf("totalTimeLines: %v\n", totalTimeLines)
	fmt.Printf("totalSplits: %v\n", totalSplits)
}

func getNextLineWithBeams(line string, nextLine string, direction Direction) (string, int) {
	lineArr := strings.Split(line, "")
	nextLineArr := strings.Split(nextLine, "")
	splits := 0

	startPosition := slices.Index(lineArr, "S")
	if startPosition != -1 {
		nextLineArr[startPosition] = "|"
		return strings.Join(nextLineArr, ""), splits
	}

	beamsInLine := strings.Count(line, "|")
	for range beamsInLine {
		currentBeamPosition := slices.Index(lineArr, "|")

		switch nextLineArr[currentBeamPosition] {
		case ".":
			nextLineArr[currentBeamPosition] = "|"
		case "^":
			if direction != Left {
				nextLineArr[currentBeamPosition+1] = "|"
			}

			if direction != Right {
				nextLineArr[currentBeamPosition-1] = "|"
			}

			splits++
		}

		lineArr[currentBeamPosition] = "x"
	}

	return strings.Join(nextLineArr, ""), splits
}

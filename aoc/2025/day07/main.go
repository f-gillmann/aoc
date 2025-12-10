package main

import (
	"aoc/helpers"
	"fmt"
	"slices"
	"strings"
)

type Beam struct {
	value int
}

type Position struct {
	row int
	col int
}

func main() {
	data := helpers.ReadFile("input/2025/day07.txt")
	lines := strings.Split(data, "\n")

	totalSplits, totalTimelines := getSplitsAndTimelines(lines)

	fmt.Printf("totalSplits: %v\n", totalSplits)
	fmt.Printf("totalTimelines: %v\n", totalTimelines)
}

func getSplitsAndTimelines(lines []string) (int, int) {
	var beams = make(map[Position]Beam)
	var timelines int
	splits := 0

	for i, line := range lines {
		if i == 0 {
			startPosition := slices.Index(strings.Split(line, ""), "S")

			if startPosition != -1 {
				beams[Position{row: 1, col: startPosition}] = Beam{value: 1}
			}

			continue
		}

		if i > len(lines) {
			break
		}

		var nextBeams = make(map[Position]Beam)

		for position, beam := range beams {
			if line[position.col] == '^' { // split
				for _, newCol := range []int{position.col + 1, position.col - 1} {
					newPos := Position{row: position.row + 1, col: newCol}

					if existingBeam, ok := nextBeams[newPos]; ok {
						existingBeam.value += beam.value
						nextBeams[newPos] = existingBeam
					} else {
						nextBeams[newPos] = Beam{value: beam.value}
					}
				}

				splits++
			} else { // move down
				newPos := Position{row: position.row + 1, col: position.col}

				if existingBeam, ok := nextBeams[newPos]; ok {
					existingBeam.value += beam.value
					nextBeams[newPos] = existingBeam
				} else {
					nextBeams[newPos] = beam
				}
			}
		}

		beams = nextBeams
	}

	for _, beam := range beams {
		timelines += beam.value
	}

	return splits, timelines
}

package main

import (
	"aoc/helpers"
	"fmt"
	"strings"
	"time"
)

func main() {
	data := helpers.ReadFile("input/2025/day04.txt")
	var rows []string = strings.Split(data, "\n")

	matrix := make([][]string, len(rows))

	for i, row := range rows {
		matrix[i] = make([]string, len(row))
		for j, col := range row {
			matrix[i][j] = string(col)
		}
	}

	totalRemoved := 0

	start := time.Now()
	for {
		_matrix, removedPositions := removeValidPositions(matrix)
		matrix = _matrix
		totalRemoved += removedPositions

		fmt.Printf("valid positions removed: %s\n", fmt.Sprint(removedPositions))

		if removedPositions == 0 {
			break
		}
	}

	fmt.Printf("time: %v\n", time.Since(start))
	fmt.Printf("total removed: %s\n", fmt.Sprint(totalRemoved))
}

func removeValidPositions(matrix [][]string) ([][]string, int) {
	potionsRemoved := 0

	removedMatrix := make([][]string, len(matrix))
	for i := range matrix {
		removedMatrix[i] = make([]string, len(matrix[i]))
		copy(removedMatrix[i], matrix[i])
	}

	for i, row := range matrix {
		for j, char := range row {
			if char == "@" {
				nearbyPaperRolls := getNearbyPaperCount(matrix, i, j)

				if nearbyPaperRolls <= 4 {
					potionsRemoved++
					removedMatrix[i][j] = "x"
				}
			}
		}
	}

	return removedMatrix, potionsRemoved
}

func getNearbyPaperCount(matrix [][]string, row int, col int) int {
	paperRolls := 0

	maxRows := len(matrix)
	maxCols := len(matrix[0])

	fromRow := max(0, row-1)
	fromCol := max(0, col-1)
	toRow := min(maxRows-1, row+1)
	toCol := min(maxCols-1, col+1)

	for i := fromRow; i <= toRow; i++ {
		for j := fromCol; j <= toCol; j++ {
			if matrix[i][j] == "@" {
				paperRolls++
			}
		}
	}

	return paperRolls
}

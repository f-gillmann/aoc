package main

import (
	"aoc/helpers"
	"fmt"
	"strings"
	"time"
)

func main() {
	data := helpers.ReadFile("input/2025/day01.txt")
	var rotations []string = strings.Split(data, "\n")

	dial := 50
	password1 := 0
	password2 := 0

	start := time.Now()
	for _, rotation := range rotations {
		direction := string(rotation[0])
		distance := helpers.StringToInt(string(rotation[1:]))

		dial, password1, password2 = turnDial(dial, direction, distance, password1, password2)
	}

	fmt.Printf("time: %v\n", time.Since(start))
	fmt.Printf("dial position: %s\n", fmt.Sprint(dial))
	fmt.Printf("password1: %s\n", fmt.Sprint(password1))
	fmt.Printf("password2: %s\n", fmt.Sprint(password2))
}

func turnDial(dial int, direction string, distance int, password1 int, password2 int) (int, int, int) {
	var pos int

	switch direction {
	case "L":
		pos = dial - distance

		startFloor := (dial - 100) / 100
		endFloor := (pos - 100) / 100
		password2 += startFloor - endFloor
	case "R":
		pos = dial + distance
		password2 += pos / 100
	default:
		panic("wtf")
	}

	pos = pos % 100

	if pos < 0 {
		pos += 100
	}

	if pos == 0 {
		password1++
	}

	return pos, password1, password2
}

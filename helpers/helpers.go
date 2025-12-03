package helpers

import (
	"os"
	"path/filepath"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(where string) string {
	dat, err := os.ReadFile(filepath.Join(where))
	check(err)
	return string(dat)
}

func StringToInt(input string) int {
	output, err := strconv.Atoi(input)
	check(err)
	return output
}

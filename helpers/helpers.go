package helpers

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(where string) string {
	dat, err := os.ReadFile(filepath.Join(where))
	check(err)
	return strings.TrimSpace(string(dat))
}

func StringToInt(input string) int {
	output, err := strconv.Atoi(input)
	check(err)
	return output
}

package helpers

import (
	"os"
	"path/filepath"
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

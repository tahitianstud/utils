package io
import (
	"os"
)

func DirectoryExists(directory string) bool {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		return false
	}

	return true
}

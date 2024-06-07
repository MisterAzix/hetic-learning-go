package utils

import (
	"fmt"
	"os"
)

func CreateFile(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, fmt.Errorf("Error creating file: %w", err)
	}
	return file, nil
}

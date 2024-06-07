package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput(prompt string) string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(prompt)
	if scanner.Scan() {
		return strings.TrimSpace(scanner.Text())
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading input:", scanner.Err())
	}

	return ""
}

func ReadInt(prompt string) int {
	for {
		input := ReadInput(prompt)
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a valid number.")
		}
		return number
	}
}

func ReadFloat(prompt string) float64 {
	for {
		input := ReadInput(prompt)
		number, err := strconv.ParseFloat(input, 64)
		if err == nil {
			return number
		}
		fmt.Println("Please enter a valid number.")
	}
}

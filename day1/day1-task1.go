package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run filename")
		return
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		calibrationValue := calculateCalibrationValue(line)
		total += calibrationValue
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Total Calibration Value:", total)
}

func calculateCalibrationValue(line string) int {
	first := findFirstDigit(line)
	last := findLastDigit(line)

	// Combine the first and last digits to form a two-digit number
	value := first*10 + last

	return value
}

func findFirstDigit(s string) int {
	for _, char := range s {
		if unicode.IsDigit(char) {
			return int(char - '0') // Convert the digit character to its integer value
		}
	}
	return 0
}

func findLastDigit(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		char := rune(s[i])
		if unicode.IsDigit(char) {
			return int(char - '0') // Convert the digit character to its integer value
		}
	}
	return 0
}


package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

var words = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

var words3 = map[string]string{
	"one": "one",
	"two": "two",
	"thr": "three",
	"fou": "four",
	"fiv": "five",
	"six": "six",
	"sev": "seven",
	"eig": "eight",
	"nin": "nine",
}

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
	num := 0
	var tmplLine string

out:
	for pos, char := range line {
		switch {
		case unicode.IsDigit(char):
			iChar, _ := strconv.Atoi(string(char))
			num = iChar * 10
			break out
		default:
			tmplLine += string(char)
			if fword, ok := words3[tmplLine]; ok {
				if (pos-len(tmplLine))+len(fword) <= len(line) {
					cword := line[pos-(len(tmplLine)-1) : pos-(len(tmplLine)-1)+(len(fword))]
					if v, ok := words[cword]; ok {
						num = v * 10
						break out
					}
				}
			}
			if len(tmplLine) >= 3 {
				tmplLine = tmplLine[1:]
			}
		}
	}
	tmplLine = ""
rout:
	for rpos := range line {
		pos := len(line) - 1 - rpos
		char := line[pos]
		switch {
		case unicode.IsDigit(rune(char)):
			iChar, _ := strconv.Atoi(string(char))
			num += iChar
			break rout
		default:
			tmplLine += string(char)
			if len(tmplLine) == 3 {
				rline := string(tmplLine[2]) + string(tmplLine[1]) + string(tmplLine[0])
				if fword, ok := words3[rline]; ok {
					if (pos-len(tmplLine))+len(fword) <= len(line) {
						cword := line[pos : pos+(len(fword))]
						if v, ok := words[cword]; ok {
							num += v
							break rout
						}
					}
				}
			}
			if len(tmplLine) >= 3 {
				tmplLine = tmplLine[1:]
			}
		}

	}
	return num
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

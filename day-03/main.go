package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	banks, err := getInput()
	if err != nil {
		panic(err)
	}

	var joltageSum int
	for _, bank := range banks {
		largestJoltageStr := getLargestJoltage2(bank, 12)
		largestJoltageInt, err := strconv.Atoi(largestJoltageStr)
		if err != nil {
			panic(err)
		}
		joltageSum += largestJoltageInt
	}

	fmt.Println(joltageSum)
}

func getInput() ([]string, error) {
	bytes, err := os.ReadFile("day-03/input.txt")
	if err != nil {
		return nil, err
	}

	ranges := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	return ranges, nil
}

func getLargestJoltage(s string) string {
	largest := s[:2]
	for first := 0; first < len(s)-1; first++ {
		for second := first + 1; second < len(s); second++ {
			firstDigit := s[first]
			secondDigit := s[second]
			joltage := string([]byte{firstDigit, secondDigit})

			if joltage > largest {
				largest = joltage
			}
		}
	}
	return largest
}

func getLargestJoltage2(s string, subsequenceLength int) string {
	// create a 2D map that stores all the digits and its corresponding index
	digits := []rune{'9', '8', '7', '6', '5', '4', '3', '2', '1', '0'}
	m := make(map[rune]map[int]bool)
	for i, digit := range s {
		if m[digit] == nil {
			m[digit] = make(map[int]bool)
		}
		m[digit][i] = true
	}

	indices := make([]int, 0, subsequenceLength)
	currentLatestIndex := 0

outer:
	for range subsequenceLength {
		for _, digit := range digits {
			for i := currentLatestIndex; i < len(s); i++ {
				if len(indices) == subsequenceLength {
					break outer
				}

				seen := m[digit][i]
				if !seen {
					continue
				}

				remainingChars := len(s) - i
				neededChars := subsequenceLength - len(indices)
				if neededChars > remainingChars {
					continue
				}

				indices = append(indices, i)
				currentLatestIndex = i + 1
				continue outer
			}
		}
	}

	builder := make([]byte, len(indices))
	for i, e := range indices {
		builder[i] = s[e]
	}

	return string(builder)
}

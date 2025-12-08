package main

import (
	"fmt"
	"os"
	"sort"
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

func getLargestJoltage2(inputString string, subsequenceLength int) string {
	digits := []rune{'9', '8', '7', '6', '5', '4', '3', '2', '1', '0'}
	posMap := make(map[rune][]int)

	for i, digit := range inputString {
		posMap[digit] = append(posMap[digit], i)
	}

	resultIndices := make([]int, 0, subsequenceLength)
	nextStart := 0

	for len(resultIndices) < subsequenceLength {
		for _, digit := range digits {
			positions := posMap[digit]
			if len(positions) == 0 {
				continue
			}

			i := sort.SearchInts(positions, nextStart)
			if i == len(positions) {
				continue
			}

			position := positions[i]

			remainingChars := len(inputString) - position
			neededChars := subsequenceLength - len(resultIndices)
			if neededChars > remainingChars {
				continue
			}

			resultIndices = append(resultIndices, position)
			nextStart = position + 1
			break
		}
	}

	builder := make([]byte, 0, len(resultIndices))
	for _, index := range resultIndices {
		builder = append(builder, inputString[index])
	}

	return string(builder)
}

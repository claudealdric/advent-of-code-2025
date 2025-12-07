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
		largestJoltageStr := getLargestJoltage(bank)
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

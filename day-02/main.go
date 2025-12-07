package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ranges, err := getInput()
	if err != nil {
		panic(err)
	}

	var invalidIdsSum int
	for _, r := range ranges {
		split := strings.Split(r, "-")
		lowerBound, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		upperBound, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		for id := lowerBound; id <= upperBound; id++ {
			if !isInvalidId(id) {
				continue
			}
			invalidIdsSum += id
		}
	}

	fmt.Println(invalidIdsSum)
}

func getInput() ([]string, error) {
	bytes, err := os.ReadFile("day-02/input.txt")
	if err != nil {
		return nil, err
	}

	ranges := strings.Split(strings.TrimSpace(string(bytes)), ",")

	return ranges, nil
}

func isInvalidId(id int) bool {
	idString := strconv.Itoa(id)
	if len(idString)%2 != 0 {
		return false
	}

	midpoint := len(idString) / 2

	return idString[:midpoint] == idString[midpoint:]
}

package main

import (
	"errors"
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
			if !isInvalidId2(id) {
				continue
			}
			invalidIdsSum += id
		}
	}

	fmt.Println(invalidIdsSum)
}

func allElementsAreEqual(s []string) bool {
	if len(s) == 0 {
		return true
	}

	firstElement := s[0]

	for i := 1; i < len(s); i++ {
		if s[i] != firstElement {
			return false
		}
	}

	return true
}

func getInput() ([]string, error) {
	bytes, err := os.ReadFile("day-02/input.txt")
	if err != nil {
		return nil, err
	}

	ranges := strings.Split(strings.TrimSpace(string(bytes)), ",")

	return ranges, nil
}

func getChunks(s string, totalChunks int) ([]string, error) {
	if len(s)%totalChunks != 0 {
		return nil, errors.New("cannot form equal-sized chunks")
	}

	chunks := make([]string, 0, totalChunks)
	chunkSize := len(s) / totalChunks
	for n := range totalChunks {
		chunks = append(chunks, s[n*chunkSize:(n+1)*chunkSize])
	}

	return chunks, nil
}

func isInvalidId(id int) bool {
	idString := strconv.Itoa(id)
	if len(idString)%2 != 0 {
		return false
	}

	midpoint := len(idString) / 2

	return idString[:midpoint] == idString[midpoint:]
}

func isInvalidId2(id int) bool {
	idString := strconv.Itoa(id)

	for totalChunks := 2; totalChunks <= len(idString); totalChunks++ {
		chunks, err := getChunks(idString, totalChunks)
		if err != nil {
			continue
		}

		if allElementsAreEqual(chunks) {
			return true
		}
	}

	return false
}

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const modulus = 100

func main() {
	var err error

	input, err := getInput()
	if err != nil {
		panic(err)
	}

	currentValue := 50
	counter := 0
	for _, rotation := range input {
		currentValue, err = applyRotation(currentValue, rotation)
		if err != nil {
			panic(err)
		}

		if currentValue == 0 {
			counter++
		}
	}

	fmt.Printf("times the value has reached 0: %d\n", counter)
}

func applyRotation(currentValue int, rotation string) (newValue int, err error) {
	rotationValue, err := parseRotationValue(rotation)
	if err != nil {
		return 0, err
	}

	newValue = currentValue + rotationValue
	newValue = (newValue + modulus) % modulus
	return newValue, nil
}

func getInput() ([]string, error) {
	bytes, err := os.ReadFile("day-01/input.txt")
	if err != nil {
		return nil, err
	}

	return strings.Split(strings.TrimSpace(string(bytes)), "\n"), nil
}

func parseRotationValue(rotation string) (int, error) {
	splits := strings.SplitN(rotation, "", 2)
	direction := splits[0]
	magnitude, err := strconv.Atoi(splits[1])
	if err != nil {
		return 0, err
	}

	switch direction {
	case "R":
		return magnitude, nil
	case "L":
		return -magnitude, nil
	default:
		return 0, errors.New("invalid direction")
	}
}

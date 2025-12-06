package main

import (
	"errors"
	"fmt"
	"math"
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
		results, err := applyRotation(currentValue, rotation)
		if err != nil {
			panic(err)
		}

		currentValue = results.newValue
		counter += results.zeroCount
	}

	fmt.Printf("times the value has reached 0: %d\n", counter)
}

type applyRotationResults struct {
	newValue  int
	zeroCount int
}

func applyRotation(currentValue int, rotation string) (results applyRotationResults, err error) {
	rotationValue, err := parseRotationValue(rotation)
	if err != nil {
		return results, err
	}

	var newValue, zeroCount int
	newValue = currentValue
	zeroCount += getAbsInt(rotationValue) / modulus
	rotationValue = rotationValue % modulus
	newValue += rotationValue
	if currentValue != 0 && (newValue >= 100 || newValue <= 0) {
		zeroCount++
	}
	newValue = (newValue + modulus) % modulus

	return applyRotationResults{
		newValue:  newValue,
		zeroCount: zeroCount,
	}, nil
}

func getAbsInt(n int) int {
	return int(math.Abs(float64(n)))
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

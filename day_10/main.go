package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func BuildInteger(number string) int {

	returnInteger := 0

	for position, digit := range number {

		integer := (int(digit) - 48) * int(math.Pow(10, float64(len(number)-1-position)))

		returnInteger += integer

	}

	return returnInteger
}

func main() {

	fileReader, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileReader.Close()

	fileScanner := bufio.NewScanner(fileReader)

	fileScanner.Split(bufio.ScanLines)

	registerX := 1
	cycleTracker := 0

	specifiedCycles := []int{20, 60, 100, 140, 180, 220}

	signalStrengthSum := 0

	for fileScanner.Scan() {

		line := strings.Split(fileScanner.Text(), " ")

		for i := 0; i < 2; i++ {
			cycleTracker++

			if Contains(specifiedCycles, cycleTracker) {
				signalStrengthSum += registerX * cycleTracker
			}

			if line[0] == "noop" {
				break
			}
		}

		if line[0] == "addx" {
			if line[1][0] == uint8('-') {
				registerX -= BuildInteger(line[1][1:])
			} else {
				registerX += BuildInteger(line[1])
			}
		}

	}

	fmt.Println("Answer:", signalStrengthSum)

}

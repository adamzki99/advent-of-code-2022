package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func BuildInteger(number string) int {

	returnInteger := 0

	for position, digit := range number {

		integer := (int(digit) - 48) * int(math.Pow(10, float64(len(number)-1-position)))

		returnInteger += integer

	}

	return returnInteger
}

func DrawScreen(screenCRT [6][40]int) {

	for _, row := range screenCRT {

		line := ""
		for _, pixel := range row {

			if pixel == 1 {
				line += "#"
			} else {
				line += "."
			}
		}
		fmt.Println(line)
	}

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

	var screenCRT [6][40]int

	drawingRow := 0

	for fileScanner.Scan() {

		line := strings.Split(fileScanner.Text(), " ")

		for i := 0; i < 2; i++ {

			if registerX-1 <= cycleTracker && cycleTracker <= registerX+1 {
				screenCRT[drawingRow][cycleTracker] = 1
			} else {
				screenCRT[drawingRow][cycleTracker] = 0
			}

			if cycleTracker == 39 && cycleTracker > 0 {
				drawingRow++
				cycleTracker = -1

				if drawingRow == 6 {
					DrawScreen(screenCRT)
					return
				}
			}

			cycleTracker++

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

}

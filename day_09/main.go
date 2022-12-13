package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func PerformHeadMove(ropeTailPositionsVisited *map[string]bool, headCoordinate *[]int, tailCoordinate *[]int, move []int) {

	headCoordinateShadow := []int{0, 0}

	if move[0] == 76 { // move left
		for i := 0; i < move[1]; i++ {

			copy(headCoordinateShadow, *headCoordinate)
			(*headCoordinate)[0]--

			if CalculateDistanceBetweenCoordinates(*headCoordinate, *tailCoordinate) > 1 {

				copy(*tailCoordinate, headCoordinateShadow)
				coordinate := fmt.Sprintln((*tailCoordinate)[0], (*tailCoordinate)[1])
				(*ropeTailPositionsVisited)[coordinate] = true
			}
		}
	} else if move[0] == 82 { // move right
		for i := 0; i < move[1]; i++ {

			copy(headCoordinateShadow, *headCoordinate)
			(*headCoordinate)[0]++

			if CalculateDistanceBetweenCoordinates(*headCoordinate, *tailCoordinate) > 1 {

				copy(*tailCoordinate, headCoordinateShadow)
				coordinate := fmt.Sprintln((*tailCoordinate)[0], (*tailCoordinate)[1])
				(*ropeTailPositionsVisited)[coordinate] = true
			}
		}
	} else if move[0] == 85 { // move up
		for i := 0; i < move[1]; i++ {

			copy(headCoordinateShadow, *headCoordinate)
			(*headCoordinate)[1]++

			if CalculateDistanceBetweenCoordinates(*headCoordinate, *tailCoordinate) > 1 {

				copy(*tailCoordinate, headCoordinateShadow)
				coordinate := fmt.Sprintln((*tailCoordinate)[0], (*tailCoordinate)[1])
				(*ropeTailPositionsVisited)[coordinate] = true
			}
		}
	} else if move[0] == 68 { // move down
		for i := 0; i < move[1]; i++ {

			copy(headCoordinateShadow, *headCoordinate)
			(*headCoordinate)[1]--

			if CalculateDistanceBetweenCoordinates(*headCoordinate, *tailCoordinate) > 1 {

				copy(*tailCoordinate, headCoordinateShadow)
				coordinate := fmt.Sprintln((*tailCoordinate)[0], (*tailCoordinate)[1])
				(*ropeTailPositionsVisited)[coordinate] = true
			}
		}
	}

}

func CalculateDistanceBetweenCoordinates(headCoordinate []int, tailCoordinate []int) int {

	diffX := headCoordinate[0] - tailCoordinate[0]
	diffY := headCoordinate[1] - tailCoordinate[1]

	diffXSquared := math.Pow(float64(diffX), 2)
	diffYSquared := math.Pow(float64(diffY), 2)

	return int(math.Sqrt(diffXSquared + diffYSquared))

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

	var seriesOfMoves [][]int

	for fileScanner.Scan() {

		s := strings.Split(fileScanner.Text(), " ")

		moveDirection := int(rune(s[0][0]))
		moveSteps, err := strconv.Atoi(s[1])

		if err != nil {
			fmt.Println(err)
			return
		}

		seriesOfMoves = append(seriesOfMoves, []int{moveDirection, moveSteps})
	}

	headCoordinate := []int{0, 0}
	tailCoordinate := []int{0, 0}
	ropeTailPositionsVisited := make(map[string]bool)

	for _, move := range seriesOfMoves {
		PerformHeadMove(&ropeTailPositionsVisited, &headCoordinate, &tailCoordinate, move)
	}

	fmt.Println("Answer:", len(ropeTailPositionsVisited))

}

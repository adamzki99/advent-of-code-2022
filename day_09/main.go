package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func moveKnot(leadingKnot []int, knot *[]int) {

	if CalculateDistanceBetweenCoordinates(leadingKnot, *knot) < 2 {
		return
	}

	if leadingKnot[0] == (*knot)[0] {
		if leadingKnot[1] > (*knot)[1] {
			(*knot)[1]++
		} else {
			(*knot)[1]--
		}
	} else if leadingKnot[1] == (*knot)[1] {
		if leadingKnot[0] > (*knot)[0] {
			(*knot)[0]++
		} else {
			(*knot)[0]--
		}
	} else {

		if leadingKnot[0] > (*knot)[0] {
			(*knot)[0]++
		} else {
			(*knot)[0]--
		}

		if leadingKnot[1] > (*knot)[1] {
			(*knot)[1]++
		} else {
			(*knot)[1]--
		}
	}

}

func PerformHeadMove(headCoordinate *[]int, direction int) []int {

	headCoordinateShadow := []int{0, 0}

	if direction == 76 { // move left

		copy(headCoordinateShadow, *headCoordinate)
		(*headCoordinate)[0]--

	} else if direction == 82 { // move right

		copy(headCoordinateShadow, *headCoordinate)
		(*headCoordinate)[0]++

	} else if direction == 85 { // move up

		copy(headCoordinateShadow, *headCoordinate)
		(*headCoordinate)[1]++

	} else if direction == 68 { // move down

		copy(headCoordinateShadow, *headCoordinate)
		(*headCoordinate)[1]--
	}

	return headCoordinateShadow
}

func CalculateDistanceBetweenCoordinates(headCoordinate []int, tailCoordinate []int) float64 {

	diffX := headCoordinate[0] - tailCoordinate[0]
	diffY := headCoordinate[1] - tailCoordinate[1]

	diffXSquared := math.Pow(float64(diffX), 2)
	diffYSquared := math.Pow(float64(diffY), 2)

	return math.Sqrt(diffXSquared + diffYSquared)
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

	ropeTailPositionsVisited := make(map[string]bool)

	ropeKnots := [][]int{
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
	}

	for _, move := range seriesOfMoves {
		fmt.Println("==", string(move[0]), move[1], "==")
		for i := 0; i < move[1]; i++ {

			PerformHeadMove(&ropeKnots[0], move[0])

			for k := 1; k < 10; k++ {

				moveKnot(ropeKnots[k-1], &ropeKnots[k])
			}

			coordinate := fmt.Sprint(ropeKnots[9][0], ropeKnots[9][1])
			(ropeTailPositionsVisited)[coordinate] = true

		}
	}

	fmt.Println("Answer:", len(ropeTailPositionsVisited))
}

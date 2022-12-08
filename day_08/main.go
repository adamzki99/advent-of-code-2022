package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func CheckVisibility(mapOfTrees [][]int, row int, column int, treeHeight int) bool {

	leftSide := 0
	rightSide := 0

	treeTracker := &leftSide

	// check left and right
	for i, tree := range mapOfTrees[row] {

		if i == column {
			treeTracker = &rightSide
			continue
		}

		if tree >= treeHeight {
			*treeTracker++
		}

	}

	above := 0
	below := 0

	treeTracker = &above
	// check top to bottom
	for i, treeLine := range mapOfTrees {

		if i == row {
			treeTracker = &below
			continue
		}

		if treeLine[column] >= treeHeight {
			*treeTracker++
		}

	}

	if leftSide > 0 && rightSide > 0 && above > 0 && below > 0 {
		return false
	}

	return true

}

func main() {

	fileReader, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(fileReader)

	fileScanner.Split(bufio.ScanLines)

	var mapOfTrees [][]int
	var lineOfTrees []int

	for fileScanner.Scan() {

		for _, number := range fileScanner.Text() {

			treeHeight, err := strconv.Atoi(string(number))

			if err != nil {
				fmt.Println(err)
				return
			}

			lineOfTrees = append(lineOfTrees, treeHeight)
		}

		mapOfTrees = append(mapOfTrees, lineOfTrees)
		lineOfTrees = []int{}

	}

	var visibleTrees []int

	sumOfVisibleTrees := 0

	for i := 1; i < len(mapOfTrees)-1; i++ {

		for i2, treeHeight := range mapOfTrees[i] {

			if i2 == 0 || i2 == len(mapOfTrees[i])-1 {
				continue
			}

			if CheckVisibility(mapOfTrees, i, i2, treeHeight) == true {
				sumOfVisibleTrees++

				visibleTrees = append(visibleTrees, i, i2)
			}

		}
	}

	// add top and bottom rows
	sumOfVisibleTrees += 2 * len(mapOfTrees[0])

	// add left and right side
	sumOfVisibleTrees += 2 * (len(mapOfTrees) - 2)

	fmt.Println("Answer:", sumOfVisibleTrees)

}

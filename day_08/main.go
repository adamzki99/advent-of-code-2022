package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func GetScenicScore(mapOfTrees [][]int, row int, column int, treeHeight int) int {

	// check left side
	leftScenicScore := 0
	for i := column - 1; i > -1; i-- {

		leftScenicScore++

		if mapOfTrees[row][i] >= treeHeight {
			break
		}
	}

	// check right side
	rightScenicScore := 0
	for i := column + 1; i < len(mapOfTrees[row]); i++ {

		rightScenicScore++

		if mapOfTrees[row][i] >= treeHeight {
			break
		}
	}

	// check above
	aboveScenicScore := 0
	for i := row - 1; i > -1; i-- {
		aboveScenicScore++

		if mapOfTrees[i][column] >= treeHeight {
			break
		}
	}

	// check below
	belowScenicScore := 0
	for i := row + 1; i < len(mapOfTrees); i++ {

		belowScenicScore++

		if mapOfTrees[i][column] >= treeHeight {
			break
		}
	}

	return leftScenicScore * rightScenicScore * aboveScenicScore * belowScenicScore

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

	var scenicScores []int

	for i := 1; i < len(mapOfTrees)-1; i++ {

		for i2, treeHeight := range mapOfTrees[i] {

			if i2 == 0 || i2 == len(mapOfTrees[i])-1 {
				continue
			}

			scenicScores = append(scenicScores, GetScenicScore(mapOfTrees, i, i2, treeHeight))

		}
	}

	sort.Ints(scenicScores)

	fmt.Println("Answer: ", scenicScores[len(scenicScores)-1])

}

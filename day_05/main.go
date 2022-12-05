package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetCratesFromLine(line string) []string {

	var crateLayer []string

	for i := 1; i < len(line); i = i + 4 {
		crateLayer = append(crateLayer, string(line[i]))
	}

	return crateLayer
}

func GetMovesFromLine(line string) [3]int {

	var moves [3]int

	s := strings.Split(line, " ")

	moves[0], _ = strconv.Atoi(s[1])
	moves[1], _ = strconv.Atoi(s[3])
	moves[2], _ = strconv.Atoi(s[5])

	return moves
}

func ReOrderCrateStack(crateStack [][]string) [][]string {

	var reOrderedCrateStacks [][]string
	var a []string

	longestRowsLength := 0
	for _, row := range crateStack {
		if len(row) > longestRowsLength {
			longestRowsLength = len(row)
		}
	}

	// prepare new crate stacks
	for i := 0; i < longestRowsLength; i++ {
		reOrderedCrateStacks = append(reOrderedCrateStacks, a)
	}

	for _, stack := range crateStack {

		for index, crate := range stack {
			if crate != " " {
				reOrderedCrateStacks[index] = append(reOrderedCrateStacks[index], crate)
			}
		}

	}

	// reverse the stacks so that the crate on top is the last element
	for s := 0; s < len(reOrderedCrateStacks); s++ {
		for i, j := 0, len(reOrderedCrateStacks[s])-1; i < j; i, j = i+1, j-1 {
			reOrderedCrateStacks[s][i], reOrderedCrateStacks[s][j] = reOrderedCrateStacks[s][j], reOrderedCrateStacks[s][i]
		}
	}

	return reOrderedCrateStacks
}

func CrateMover9001(move [3]int, crateStack *[][]string) {

	homeStack := move[1] - 1
	destinationStack := move[2] - 1

	var tempStack []string

	for i := 0; i < move[0]; i++ {
		tempStack = append(tempStack, (*crateStack)[homeStack][len((*crateStack)[homeStack])-1])

		(*crateStack)[homeStack] = (*crateStack)[homeStack][:len((*crateStack)[homeStack])-1]
	}

	for i := 0; i < move[0]; i++ {
		(*crateStack)[destinationStack] = append((*crateStack)[destinationStack], tempStack[len(tempStack)-1-i])
	}
}

func main() {

	fileReader, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(fileReader)

	fileScanner.Split(bufio.ScanLines)

	var crateStacks [][]string
	var arrayOfMoves [][3]int

	buildingCrateStack := true

	for fileScanner.Scan() {

		if len(fileScanner.Text()) == 0 {
			buildingCrateStack = false
			crateStacks = crateStacks[:len(crateStacks)-1]
			continue
		}

		if buildingCrateStack {
			// Get current stacks
			crateStacks = append(crateStacks, GetCratesFromLine(fileScanner.Text()))
		}

		if !buildingCrateStack {
			arrayOfMoves = append(arrayOfMoves, GetMovesFromLine(fileScanner.Text()))
		}
	}
	crateStacks = ReOrderCrateStack(crateStacks)

	for _, move := range arrayOfMoves {
		CrateMover9001(move, &crateStacks)
	}

	answer := ""

	for _, stack := range crateStacks {

		answer = answer + stack[len(stack)-1]
	}

	fmt.Printf("Awnser: %s", answer)

}

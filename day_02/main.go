package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func translateInput(input string) int {

	if input == "X" || input == "A" {
		return 1
	}
	if input == "Y" || input == "B" {
		return 2
	}
	if input == "Z" || input == "C" {
		return 3
	}

	return -1

}

func calculateRound(elfSelection string, mySelection int) int {

	elfsTranslatedInput := translateInput(elfSelection)

	differance := mySelection - elfsTranslatedInput

	if differance == 1 || differance == -2 {
		return mySelection + 6
	}

	if differance == 0 {
		return mySelection + 3
	}

	return mySelection
}

func makeSelection(elfSelection string, roundResult string) int {

	winMatrix := [3][2]int{
		{1, 3},
		{2, 1},
		{3, 2},
	}

	elfSelectionTranslatedToInt := translateInput(elfSelection)

	if roundResult == "X" {
		for _, element := range winMatrix {

			if element[0] == elfSelectionTranslatedToInt {
				return element[1]
			}
		}
	}

	if roundResult == "Y" {
		return elfSelectionTranslatedToInt
	}

	if roundResult == "Z" {
		for _, element := range winMatrix {

			if element[1] == elfSelectionTranslatedToInt {
				return element[0]
			}
		}
	}

	return 0
}

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	gameResult := 0

	for fileScanner.Scan() {

		round := strings.Split(fileScanner.Text(), " ")

		mySelection := makeSelection(round[0], round[1])

		roundResult := calculateRound(round[0], mySelection)

		gameResult = gameResult + roundResult
	}

	fmt.Println(gameResult)

}

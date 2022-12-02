package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
X = 1 = Rock
Y = 2 = Paper
Z = 3 = Scissor

A = Rock = 1
B = Paper = 2
C = Scissor = 3
*/

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

func calculateRound(elf string, me string) int {

	myTranslatedInput := translateInput(me)
	elfsTranslatedInput := translateInput(elf)

	differance := myTranslatedInput - elfsTranslatedInput

	if differance == 1 || differance == -2 {
		return myTranslatedInput + 6
	}

	if differance == 0 {
		return myTranslatedInput + 3
	}

	return myTranslatedInput
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

		roundResult := calculateRound(round[0], round[1])

		gameResult = gameResult + roundResult
	}

	fmt.Println(gameResult)

}

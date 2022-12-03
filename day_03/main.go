package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

/*
- Lowercase item types a through z have priorities 1 through 26.
- Uppercase item types A through Z have priorities 27 through 52.
*/

func CalulateSum(arrayOfElfItems []rune) int {

	totalSum := 0

	for _, r := range arrayOfElfItems {

		priority := 0
		if unicode.IsLower(r) {

			priority = int(r) - 96

		} else { // If not lowercase, it must be uppercase

			priority = int(r) - 38
		}

		fmt.Println("Char", string(r), "-> Priority", priority)

		totalSum = totalSum + priority

	}

	return totalSum
}

func RugsackSplit(rugsack string) [2]string {

	middleIndex := (len(rugsack)) / 2

	firstHalf := rugsack[:middleIndex]
	secondHalf := rugsack[middleIndex:]

	return [2]string{firstHalf, secondHalf}
}

func FindItemsInRugsack(splitRugsack [2]string) []rune {

	var arrayOfElfItems []rune

	var backPack []rune

	for _, r := range splitRugsack[0] {

		if strings.ContainsRune(splitRugsack[1], r) && strings.ContainsRune(string(backPack), r) == false {
			backPack = append(backPack, r)
			continue
		}
	}

	for _, r := range backPack {
		arrayOfElfItems = append(arrayOfElfItems, r)
	}

	return arrayOfElfItems
}

func main() {

	inputFileName := os.Args[1]

	fileReader, err := os.Open(inputFileName)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer fileReader.Close()

	fileScanner := bufio.NewScanner(fileReader)

	fileScanner.Split(bufio.ScanLines)

	var arrayOfElfItems []rune

	amountOfLines := 0

	for fileScanner.Scan() {

		splitRugsack := RugsackSplit(fileScanner.Text())

		itemsInRugSack := FindItemsInRugsack(splitRugsack)

		for _, r := range itemsInRugSack {

			arrayOfElfItems = append(arrayOfElfItems, r)
		}

		amountOfLines++
	}

	fmt.Println("Amount of lines:", amountOfLines)
	fmt.Println("Amount of elf items", len(arrayOfElfItems))
	fmt.Println("Awnser:", CalulateSum(arrayOfElfItems))

}

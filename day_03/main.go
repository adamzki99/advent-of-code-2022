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

		totalSum = totalSum + priority

	}

	return totalSum
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

func FindCommenItemBetweenThreeElfs(arrayOfRugsacks [3]string) rune {

	commonItems := FindItemsInRugsack([2]string{arrayOfRugsacks[0], arrayOfRugsacks[1]})

	commonItems = FindItemsInRugsack([2]string{string(commonItems), arrayOfRugsacks[2]})

	return commonItems[0]
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

	var tripletOfElfRugsacks [3]string

	amountOfLines := 0

	for fileScanner.Scan() {

		tripletOfElfRugsacks[amountOfLines%3] = fileScanner.Text()

		amountOfLines++

		if amountOfLines%3 == 0 {
			commonItem := FindCommenItemBetweenThreeElfs(tripletOfElfRugsacks)
			arrayOfElfItems = append(arrayOfElfItems, commonItem)
		}
	}

	fmt.Println("Awnser:", CalulateSum(arrayOfElfItems))

}

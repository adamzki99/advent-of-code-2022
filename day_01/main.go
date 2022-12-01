package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var allElvesTotalCalories []int
	currentElvesTotalCalories := 0

	for fileScanner.Scan() {
		if fileScanner.Text() == "" {

			allElvesTotalCalories = append(allElvesTotalCalories, currentElvesTotalCalories)
			currentElvesTotalCalories = 0

		} else {

			calories, err := strconv.Atoi(fileScanner.Text())

			if err != nil {
				fmt.Println(err)
			}

			currentElvesTotalCalories = currentElvesTotalCalories + calories
		}
	}

	sort.Ints(allElvesTotalCalories)

	sumOfTopThreeElves := allElvesTotalCalories[len(allElvesTotalCalories)-1] + allElvesTotalCalories[len(allElvesTotalCalories)-2] + allElvesTotalCalories[len(allElvesTotalCalories)-3]

	fmt.Println(sumOfTopThreeElves)

	readFile.Close()
}

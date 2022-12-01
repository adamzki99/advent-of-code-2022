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

	var allGnomes []int
	currentGnome := 0

	for fileScanner.Scan() {
		if fileScanner.Text() == "" {

			allGnomes = append(allGnomes, currentGnome)
			currentGnome = 0

		} else {

			calories, err := strconv.Atoi(fileScanner.Text())

			if err != nil {
				fmt.Println(err)
			}

			currentGnome = currentGnome + calories
		}
	}

	sort.Ints(allGnomes)

	fmt.Println(allGnomes[len(allGnomes)-1])

	readFile.Close()
}

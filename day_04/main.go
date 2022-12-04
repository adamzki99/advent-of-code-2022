package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ExpandElfPair(puzzelInputLine string) [2][]int {

	var returnArray [2][]int

	twoElfs := strings.Split(puzzelInputLine, ",")
	//Exampel: ["2-4","6-8"]

	//Extract first and last idex from both elfs
	firstElfInfo := strings.Split(twoElfs[0], "-")
	// Exampel: ["2","4"]

	secondElfInfo := strings.Split(twoElfs[1], "-")

	//Build first elf
	lowerLimit, _ := strconv.Atoi(firstElfInfo[0])
	upperLimit, _ := strconv.Atoi(firstElfInfo[1])

	for i := lowerLimit; i <= upperLimit; i++ {
		returnArray[0] = append(returnArray[0], i)
	}

	//Build second elf
	lowerLimit, _ = strconv.Atoi(secondElfInfo[0])
	upperLimit, _ = strconv.Atoi(secondElfInfo[1])

	for i := lowerLimit; i <= upperLimit; i++ {
		returnArray[1] = append(returnArray[1], i)
	}

	return returnArray
}

func Contains(value int, arrayOfInts []int) bool {

	for _, v := range arrayOfInts {

		if v == value {
			return true
		}
	}

	return false
}

func DoesOneRangePartiallyContainTheOther(twoElfsAssignments [2][]int) bool {

	var shortestAsignmentListIndex int
	var longestAsignmentListIndex int
	if len(twoElfsAssignments[0]) < len(twoElfsAssignments[1]) {
		shortestAsignmentListIndex = 0
		longestAsignmentListIndex = 1
	} else {
		shortestAsignmentListIndex = 1
		longestAsignmentListIndex = 0
	}

	for _, id := range twoElfsAssignments[shortestAsignmentListIndex] {
		if Contains(id, twoElfsAssignments[longestAsignmentListIndex]) == true {
			return true
		}
	}

	return false
}

func DoesOneRangePartiallyContainTheOther(twoElfsAssignments [2][]int) bool {

	var shortestAsignmentListIndex int
	var longestAsignmentListIndex int
	if len(twoElfsAssignments[0]) < len(twoElfsAssignments[1]) {
		shortestAsignmentListIndex = 0
		longestAsignmentListIndex = 1
	} else {
		shortestAsignmentListIndex = 1
		longestAsignmentListIndex = 0
	}

	for _, id := range twoElfsAssignments[shortestAsignmentListIndex] {
		if Contains(id, twoElfsAssignments[longestAsignmentListIndex]) == true {
			return true
		}
	}

	return false
}

func main() {

	fileName := os.Args[1]
	fmt.Println("Using file:", fileName)

	fileReader, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileReader.Close()

	fileScanner := bufio.NewScanner(fileReader)

	fileScanner.Split(bufio.ScanLines)

	score := 0

	for fileScanner.Scan() {

		twoElfsAssignments := ExpandElfPair(fileScanner.Text())

		if DoesOneRangePartiallyContainTheOther(twoElfsAssignments) {
			score++
		}
	}

	fmt.Println("Answer:", score)
}

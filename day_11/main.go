package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items     []int
	operation string
	test      int
	ifTrue    int
	ifFalse   int
}

func (m Monkey) PerformOperation(itemIndex int) {

	operationParts := strings.Split(m.operation, " ")
	var integers [2]int

	if operationParts[3] == "old" {
		integers[0] = m.items[itemIndex]
	} else {
		var err error
		integers[0], err = strconv.Atoi(operationParts[3])

		if err != nil {
			fmt.Println(err)
		}
	}

	if operationParts[5] == "old" {
		integers[1] = m.items[itemIndex]
	} else {
		var err error
		integers[1], err = strconv.Atoi(operationParts[5])

		if err != nil {
			fmt.Println(err)
		}
	}

	if operationParts[4] == "*" {

		m.items[itemIndex] = integers[0] * integers[1]
		//fmt.Printf("\t\tWorry level is multiplied by itself to %d.\n", m.items[itemIndex])

	} else if operationParts[4] == "+" {

		m.items[itemIndex] = integers[0] + integers[1]
		//fmt.Printf("\t\tWorry level increases by %d to %d.\n", integers[1], m.items[itemIndex])

	} else {

		m.items[itemIndex] = integers[0] - integers[1]
		//fmt.Printf("\t\tWorry level decreases by %d to %d.\n", integers[1], m.items[itemIndex])

	}

}

func (m Monkey) RemoveSpecificItemFromMonkey(item int) {

	itemIndex := -1

	for i, monkeyItem := range m.items {
		if monkeyItem == item {
			itemIndex = i
		}
	}

	m.items = append(m.items[:itemIndex], m.items[itemIndex+1:]...)

}

func (m Monkey) PerformTest(itemIndex int) bool {

	return m.items[itemIndex]%m.test == 0
}

func PerformThrow(dst *Monkey, item int) {

	(*dst).items = append((*dst).items, item)

}

func main() {

	//read puzzle input
	fileReader, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileReader.Close()

	fileScanner := bufio.NewScanner(fileReader)

	fileScanner.Split(bufio.ScanLines)

	var monkeys []Monkey

	for fileScanner.Scan() {
		fileScanner.Text()

		if strings.Contains(fileScanner.Text(), "Monkey") {
			monkeys = append(monkeys, Monkey{})
			continue
		}

		if strings.Contains(fileScanner.Text(), "Starting items:") {

			items := strings.Split(fileScanner.Text(), ":")

			items = strings.Split(items[1], " ")

			for _, item := range items {

				item = strings.ReplaceAll(item, ",", "")

				integer, err := strconv.Atoi(item)

				if err != nil {
					fmt.Println(err)
					continue
				}

				monkeys[len(monkeys)-1].items = append(monkeys[len(monkeys)-1].items, integer)

			}

		}

		if strings.Contains(fileScanner.Text(), "Operation:") {

			splitLine := strings.Split(fileScanner.Text(), ":")

			monkeys[len(monkeys)-1].operation = splitLine[1]

		}

		if strings.Contains(fileScanner.Text(), "Test:") {

			splitLine := strings.Split(fileScanner.Text(), " ")

			integer, err := strconv.Atoi(splitLine[len(splitLine)-1])

			if err != nil {
				fmt.Println(err)
				return
			}

			monkeys[len(monkeys)-1].test = integer

		}

		if strings.Contains(fileScanner.Text(), "If") {

			splitLine := strings.Split(fileScanner.Text(), " ")

			integer, err := strconv.Atoi(splitLine[len(splitLine)-1])

			if err != nil {
				fmt.Println(err)
				return
			}

			if strings.Contains(fileScanner.Text(), "true:") {

				monkeys[len(monkeys)-1].ifTrue = integer

			}

			if strings.Contains(fileScanner.Text(), "false:") {

				monkeys[len(monkeys)-1].ifFalse = integer

			}

		}

	}

	monkeyInspections := make([]int, len(monkeys))

	var commonDenominator int
	commonDenominator = 1

	for _, monkey := range monkeys {
		commonDenominator *= monkey.test
	}

	//perform 10000 rounds

	for r := 0; r < 10000; r++ {

		for i, monkey := range monkeys {

			//fmt.Printf("Monkey %d:\n", i)

			for itemIndex, _ := range monkey.items {
				monkeyInspections[i]++

				//fmt.Printf("\tMonkey inspects an item with a worry level of %d.\n", monkey.items[itemIndex])

				monkey.PerformOperation(itemIndex)

				monkey.items[itemIndex] = monkey.items[itemIndex] % commonDenominator
				//fmt.Printf("\t\tMonkey gets bored with item. Worry level is divided by 3 to %d.\n", monkey.items[itemIndex])

				if monkey.PerformTest(itemIndex) == true {
					//fmt.Printf("\t\tCurrent worry level is divisible by %d.\n", monkey.test)
					PerformThrow(&monkeys[monkey.ifTrue], monkey.items[itemIndex])
					//fmt.Printf("\t\tItem with worry level %d is thrown to monkey %d.\n", monkey.items[itemIndex], monkey.ifTrue)
				} else {
					//fmt.Printf("\t\tCurrent worry level is not divisible by %d.\n", monkey.test)
					PerformThrow(&monkeys[monkey.ifFalse], monkey.items[itemIndex])
					//fmt.Printf("\t\tItem with worry level %d is thrown to monkey %d.\n", monkey.items[itemIndex], monkey.ifFalse)

				}

			}
			monkeys[i].items = nil
		}

	}

	sort.Ints(monkeyInspections)
	fmt.Println("Answer:", monkeyInspections[len(monkeyInspections)-1]*monkeyInspections[len(monkeyInspections)-2])

}

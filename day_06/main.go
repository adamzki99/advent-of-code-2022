package main

import (
	"bufio"
	"fmt"
	"os"
)

func PacketExaminations(subString string) bool {

	for i := 0; i < 3; i++ {
		for j := i + 1; j < 4; j++ {
			if subString[i] == subString[j] {
				return false
			}
		}
	}

	return true
}

func FindStartOfPacketMarker(line string) int {

	for i := 0; i < len(line)-4; i++ {

		if PacketExaminations(line[i:i+4]) == true {
			return i
		}
	}

	return -1
}

func main() {

	fileReader, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(fileReader)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fmt.Println("Answer:", FindStartOfPacketMarker(fileScanner.Text())+4)
	}

}

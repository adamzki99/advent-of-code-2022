package main

import (
	"bufio"
	"fmt"
	"os"
)

func PacketExaminations(subString string) bool {

	for i := 0; i < len(subString)-1; i++ {
		for j := i + 1; j < len(subString); j++ {
			if subString[i] == subString[j] {
				return false
			}
		}
	}

	return true
}

func FindStartOfPacketMarker(line string, stride int) int {

	for i := 0; i < len(line)-stride; i++ {

		if PacketExaminations(line[i:i+stride]) == true {
			return i + stride
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
		fmt.Println("Answer:", FindStartOfPacketMarker(fileScanner.Text(), 14))
	}

}

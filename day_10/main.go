package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fileReader, err := os.Open("example-input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileReader.Close()

	fileScanner := bufio.NewScanner(fileReader)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		fileScanner.Text()
	}

}

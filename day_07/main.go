package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type file struct {
	name           string
	size           int
	subDirectories []*file
	parentFile     *file
}

func FindSumOfDirectoriesBelow100000(dir *file, result *int) {

	if dir.size < 100000 {
		*result = *result + dir.size
	}

	if dir.subDirectories != nil {
		for _, directory := range dir.subDirectories {
			FindSumOfDirectoriesBelow100000(directory, result)
		}
	}
}

// UpdateDirectorySizes
// returns the updated size of the given directory
// /*
func UpdateDirectorySizes(dir *file) int {

	subDirectorySizeSum := 0

	if dir.subDirectories == nil {
		return dir.size
	}

	for _, directory := range dir.subDirectories {
		subDirectorySizeSum = subDirectorySizeSum + UpdateDirectorySizes(directory)
	}

	dir.size = dir.size + subDirectorySizeSum

	return dir.size
}

func main() {

	fileReader, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(fileReader)

	fileScanner.Split(bufio.ScanLines)

	fileSystem := file{name: "/"}

	currentFile := &fileSystem

	for fileScanner.Scan() {

		s := strings.Split(fileScanner.Text(), " ")

		if s[0] == "$" {

			if s[1] == "cd" {

				if s[2] == "/" {
					continue
				}

				if s[2] == ".." {

					currentFile = currentFile.parentFile

				} else {
					newFile := &file{name: s[2], size: 0, parentFile: currentFile}

					currentFile.subDirectories = append(currentFile.subDirectories, newFile)

					currentFile = newFile

				}

			}

			if s[1] == "ls" {

			}

		} else if s[0] != "dir" {
			fileSize, err := strconv.Atoi(s[0])

			if err != nil {
				fmt.Println(err)
				return
			}

			currentFile.size = currentFile.size + fileSize
		}

	}

	UpdateDirectorySizes(&fileSystem)

	result := 0

	FindSumOfDirectoriesBelow100000(&fileSystem, &result)

	fmt.Println("Answer:", result)
}

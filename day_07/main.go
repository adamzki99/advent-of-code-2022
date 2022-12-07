package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type file struct {
	name           string
	size           int
	subDirectories []*file
	parentFile     *file
}

func FindSizesThatCanBeDeleted(dir *file, optionsForDeletion *[]int, sizeRequired int) {

	if dir.size > sizeRequired {
		*optionsForDeletion = append(*optionsForDeletion, dir.size)
	}

	if dir.subDirectories != nil {
		for _, directory := range dir.subDirectories {
			FindSizesThatCanBeDeleted(directory, optionsForDeletion, sizeRequired)
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

	diskCapacity := 70000000

	spaceNeeded := 30000000

	sizeToBeRemoved := spaceNeeded - (diskCapacity - fileSystem.size)

	fmt.Println("Space needed:", sizeToBeRemoved)

	var optionsForDeletion []int

	FindSizesThatCanBeDeleted(&fileSystem, &optionsForDeletion, sizeToBeRemoved)

	sort.Ints(optionsForDeletion)

	fmt.Println("Answer:", optionsForDeletion[0])
}

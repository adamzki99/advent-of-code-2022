package main

import "testing"

func TestCalulateSum(t *testing.T) {

	testData := []rune{rune('a'), rune('z')}

	sum := CalulateSum(testData)
	if sum != 27 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", sum, 27)
	}

	testData = []rune{rune('A'), rune('Z')}

	sum = CalulateSum(testData)
	if sum != 79 {
		t.Errorf("calulateSum was incorrect, got: %d, want: %d.", sum, 79)
	}
}

func TestRugsackSplit(t *testing.T) {

	testData := "vJrwpWtwJgWrhcsFMMfFFhFp"

	splitRugsack := RugsackSplit(testData)
	if splitRugsack[0] != "vJrwpWtwJgWr" || splitRugsack[1] != "hcsFMMfFFhFp" {
		t.Errorf("RugsackSplit was incorrect, got: [\"%s\",\"%s\"], want %s", splitRugsack[0], splitRugsack[1], testData)
	}

	testData = "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL"

	splitRugsack = RugsackSplit(testData)
	if splitRugsack[0] != "jqHRNqRjqzjGDLGL" || splitRugsack[1] != "rsFMfFZSrLrFZsSL" {
		t.Errorf("RugsackSplit was incorrect, got: [\"%s\",\"%s\"], want %s", splitRugsack[0], splitRugsack[1], testData)
	}

	testData = "PmmdzqPrVvPwwTWBwg"

	splitRugsack = RugsackSplit(testData)
	if splitRugsack[0] != "PmmdzqPrV" || splitRugsack[1] != "vPwwTWBwg" {
		t.Errorf("RugsackSplit was incorrect, got: [\"%s\",\"%s\"], want %s", splitRugsack[0], splitRugsack[1], testData)
	}
}

func TestFindItemsInRugsack(t *testing.T) {
	testData := [2]string{"vJrwpWtwJgWr", "hcsFMMfFFhFp"}

	itemsInRugsack := FindItemsInRugsack(testData)
	if itemsInRugsack[0] != rune('p') {
		t.Errorf("RugsackSplit was incorrect, got: %d, want %d", itemsInRugsack[0], rune('p'))
	}

	testData = [2]string{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"}

	itemsInRugsack = FindItemsInRugsack(testData)
	if itemsInRugsack[0] != rune('L') {
		t.Errorf("RugsackSplit was incorrect, got: %d, want %d", itemsInRugsack[0], rune('L'))
	}

	testData = [2]string{"PmmdzqPrV", "vPwwTWBwg"}

	itemsInRugsack = FindItemsInRugsack(testData)
	if itemsInRugsack[0] != rune('P') {
		t.Errorf("RugsackSplit was incorrect, got: %d, want %d", itemsInRugsack[0], rune('P'))
	}

	testData = [2]string{"wMqvLMZHhHMvwLH", "jbvcjnnSBnvTQFn"}

	itemsInRugsack = FindItemsInRugsack(testData)
	if itemsInRugsack[0] != rune('v') {
		t.Errorf("RugsackSplit was incorrect, got: %d, want %d", itemsInRugsack[0], rune('v'))
	}

	testData = [2]string{"ttgJtRGJ", "QctTZtZT"}

	itemsInRugsack = FindItemsInRugsack(testData)
	if itemsInRugsack[0] != rune('t') {
		t.Errorf("RugsackSplit was incorrect, got: %d, want %d", itemsInRugsack[0], rune('t'))
	}

	testData = [2]string{"CrZsJsPPZsGz", "wwsLwLmpwMDw"}

	itemsInRugsack = FindItemsInRugsack(testData)
	if itemsInRugsack[0] != rune('s') {
		t.Errorf("RugsackSplit was incorrect, got: %d, want %d", itemsInRugsack[0], rune('s'))
	}

	testData = [2]string{"PmmdzqPrVAA", "vPwwATWABwg"}

	itemsInRugsack = FindItemsInRugsack(testData)
	if itemsInRugsack[0] != rune('P') && itemsInRugsack[1] != rune('A') {
		t.Errorf("RugsackSplit was incorrect, got: %d and %d, want %d and %d", itemsInRugsack[0], itemsInRugsack[1], rune('P'), rune('A'))
	}
}

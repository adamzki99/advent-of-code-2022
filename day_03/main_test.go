package main

import "testing"

func TestCalulateSum(t *testing.T) {

	testData := []rune{rune('a'), rune('z')}

	sum := CalulateSum(testData)
	if sum != 27 {
		t.Errorf("calulateSum was incorrect, got: %d, want: %d.", sum, 27)
	}

	testData = []rune{rune('A'), rune('Z')}

	sum = CalulateSum(testData)
	if sum != 79 {
		t.Errorf("calulateSum was incorrect, got: %d, want: %d.", sum, 79)
	}
}

func TestFindItemsInRugsack(t *testing.T) {
	testData := [2]string{"vJrwpWtwJgWr", "hcsFMMfFFhFp"}

	itemsInRugsack := FindItemsInRugsack(testData)
	if itemsInRugsack[0] != rune('p') {
		t.Errorf("FindItemsInRugsack was incorrect, got: %d, want %d", itemsInRugsack[0], rune('p'))
	}

	testData = [2]string{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"}

	itemsInRugsack = FindItemsInRugsack(testData)
	if itemsInRugsack[0] != rune('L') {
		t.Errorf("FindItemsInRugsack was incorrect, got: %d, want %d", itemsInRugsack[0], rune('L'))
	}

	testData = [2]string{"PmmdzqPrV", "vPwwTWBwg"}

	itemsInRugsack = FindItemsInRugsack(testData)
	if itemsInRugsack[0] != rune('P') {
		t.Errorf("FindItemsInRugsack was incorrect, got: %d, want %d", itemsInRugsack[0], rune('P'))
	}

	testData = [2]string{"wMqvLMZHhHMvwLH", "jbvcjnnSBnvTQFn"}

	itemsInRugsack = FindItemsInRugsack(testData)
	if itemsInRugsack[0] != rune('v') {
		t.Errorf("FindItemsInRugsack was incorrect, got: %d, want %d", itemsInRugsack[0], rune('v'))
	}

	testData = [2]string{"ttgJtRGJ", "QctTZtZT"}

	itemsInRugsack = FindItemsInRugsack(testData)
	if itemsInRugsack[0] != rune('t') {
		t.Errorf("FindItemsInRugsack was incorrect, got: %d, want %d", itemsInRugsack[0], rune('t'))
	}

	testData = [2]string{"CrZsJsPPZsGz", "wwsLwLmpwMDw"}

	itemsInRugsack = FindItemsInRugsack(testData)
	if itemsInRugsack[0] != rune('s') {
		t.Errorf("FindItemsInRugsack was incorrect, got: %d, want %d", itemsInRugsack[0], rune('s'))
	}

	testData = [2]string{"PmmdzqPrVAA", "vPwwATWABwg"}

	itemsInRugsack = FindItemsInRugsack(testData)
	if itemsInRugsack[0] != rune('P') && itemsInRugsack[1] != rune('A') {
		t.Errorf("FindItemsInRugsack was incorrect, got: %d and %d, want %d and %d", itemsInRugsack[0], itemsInRugsack[1], rune('P'), rune('A'))
	}
}

func TestFindCommenItemBetweenThreeElfs(t *testing.T) {

	testData := [3]string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
	}

	commonItem := FindCommenItemBetweenThreeElfs(testData)
	if commonItem != rune('r') {
		t.Errorf("FindCommenItemBetweenThreeElfs was incorrect, got: %d, want %d", commonItem, rune('r'))
	}

	testData = [3]string{
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}

	commonItem = FindCommenItemBetweenThreeElfs(testData)
	if commonItem != rune('Z') {
		t.Errorf("FindCommenItemBetweenThreeElfs was incorrect, got: %d, want %d", commonItem, rune('Z'))
	}

}

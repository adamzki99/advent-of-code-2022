package main

import "testing"

func TestGetCratesFromLine(t *testing.T) {
	result := GetCratesFromLine("    [D]    ")
	if result[0] != "" || result[1] != "D" || result[2] != "" {
		t.Errorf("TestGetCratesFromLine FAILED! Expected P,M,Z but got %s,%s,%s", result[0], result[1], result[2])
	}

	result = GetCratesFromLine("[N] [C]    ")
	if result[0] != "N" || result[1] != "C" || result[2] != "" {
		t.Errorf("TestGetCratesFromLine FAILED! Expected N,C but got %s,%s", result[0], result[1])
	}

	result = GetCratesFromLine("[Z] [M] [P]")
	if result[0] != "Z" || result[1] != "M" || result[2] != "P" {
		t.Errorf("TestGetCratesFromLine FAILED! Expected P,M,Z but got %s,%s,%s", result[0], result[1], result[2])
	}
}

func TestReOrderCrateStack(t *testing.T) {

	result := ReOrderCrateStack([][]string{
		{"", "D", ""},
		{"N", "C", ""},
		{"Z", "M", "P"},
	})
	// Check first stack
	if len(result[0]) != 2 {
		t.Errorf("ReOrderCrateStack FAILED! Expected first stack lengt to be 2 but got %d", len(result[0]))
	}
	if result[0][0] != "Z" || result[0][1] != "N" {
		t.Errorf("ReOrderCrateStack FAILED! Expected first stack to be Z,N but got %s,%s", result[0][0], result[0][1])
	}
	// Check second stack
	if len(result[1]) != 3 {
		t.Errorf("ReOrderCrateStack FAILED! Expected second stack lengt to be 3 but got %d", len(result[0]))
	}
	if result[1][0] != "M" || result[1][1] != "C" || result[1][2] != "D" {
		t.Errorf("ReOrderCrateStack FAILED! Expected second stack to be M,C,D but got %s,%s,%s", result[1][0], result[1][1], result[1][2])
	}
	// Check third stack
	if len(result[2]) != 1 {
		t.Errorf("ReOrderCrateStack FAILED! Expected third stack lengt to be 1 but got %d", len(result[0]))
	}
	if result[2][0] != "P" {
		t.Errorf("ReOrderCrateStack FAILED! Expected third stack to be P but got %s", result[2][0])
	}

}

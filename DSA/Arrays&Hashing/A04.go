package main

import (
	"fmt"
	"sort"
)

// sortString sorts the characters in a string and returns the sorted string
func sortString(s string) string {
	// Convert string to rune slice for sorting
	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}

// groupAnagrams groups strings that are anagrams of each other
func groupAnagrams(strs []string) [][]string {
	// Map to store groups of anagrams
	anagramMap := make(map[string][]string)

	// Group strings by their sorted character form
	for _, str := range strs {
		// Sort the current string
		sortedStr := sortString(str)

		// Add the original string to its group
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}

	// Convert map to result slice
	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}

func A04() {
	// Example usage
	strs := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	result := groupAnagrams(strs)
	fmt.Println(result)

	// The result will be something like:
	// [["hat"], ["act", "cat"], ["pots", "stop", "tops"]]
}

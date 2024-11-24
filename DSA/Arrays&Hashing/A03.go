// anagram
package main

import "fmt"

// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}
// 	var scount = make(map[rune]int)
// 	var tcount = make(map[rune]int)

// 	for _, char := range s {
// 		scount[char]++
// 	}
// 	for _, char := range t {
// 		tcount[char]++
// 	}

// 	if len(scount) != len(tcount) {
// 		return false
// 	}

// 	for key, value := range scount {
// 		if tcount[key] != value {
// 			return false
// 		}
// 	}

// 	return true
// }

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var count = make(map[rune]int)

	for i, value := range s {
		count[value]++
		count[rune(t[i])]--
	}

	for _, value := range count {
		if value != 0 {
			return false
		}
	}

	return true
}

func A03() {
	var s string = "racecar"
	var t string = "carrace"

	fmt.Println(isAnagram(s, t))

}

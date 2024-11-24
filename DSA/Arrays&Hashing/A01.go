// Duplicate int
package main

import "fmt"

func ContinsDuplicate(nums []int) bool {

	var seen = make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return true
		} else {
			seen[num] = true
		}
	}
	return false
}

func A01() {
	var num1 = []int{1, 2, 3, 3}
	var num2 = []int{1, 2, 3, 4}

	fmt.Println(ContinsDuplicate(num1))
	fmt.Println(ContinsDuplicate(num2))
}

// two sum
package main

import "fmt"

func twoSum(nums []int, target int) []int {

	var seen = make(map[int]int)

	for i, num := range nums {
		var complement int = target - num
		if j, exists := seen[complement]; exists {
			return []int{j, i}
		}
		seen[num] = i

	}

	return []int{}
}

func A02() {

	nums1 := []int{2, 7, 11, 15}
	target1 := 9

	nums2 := []int{3, 2, 4}
	target2 := 6

	fmt.Println(twoSum(nums1, target1))
	fmt.Println(twoSum(nums2, target2))
}

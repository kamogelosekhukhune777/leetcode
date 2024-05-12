package main

import "fmt"

/*
Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color
are adjacent, with the colors in the order red, white, and blue.

We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function.

Example 1:
	Input: nums = [2,0,2,1,1,0]
	Output: [0,0,1,1,2,2]

Example 2:
	Input: nums = [2,0,1]
	Output: [0,1,2]

Constraints:
	n == nums.length
	1 <= n <= 300
	nums[i] is either 0, 1, or 2.

*/

func sortColors(nums []int) {
	// Initialize pointers for red (0), white (1), and blue (2)
	red, white, blue := 0, 0, len(nums)-1

	// Perform cyclic sort to partition the array into red, white, and blue sections
	for white <= blue {
		switch nums[white] {
		case 0: // If element is red (0), swap with red pointer and move both pointers
			nums[red], nums[white] = nums[white], nums[red]
			red++
			white++
		case 1: // If element is white (1), move white pointer
			white++
		case 2: // If element is blue (2), swap with blue pointer and move blue pointer
			nums[white], nums[blue] = nums[blue], nums[white]
			blue--
		}
	}
}

func main() {
	//Example 2:
	nums1 := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums1)
	for i := range nums1 {
		fmt.Print(nums1[i], ",")
	}
	fmt.Println()

	//Example 2:
	nums2 := []int{2, 0, 1}
	sortColors(nums2)
	for i := range nums2 {
		fmt.Print(nums2[i], ",")
	}
}

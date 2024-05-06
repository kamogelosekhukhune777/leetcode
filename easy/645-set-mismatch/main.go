package main

import "fmt"

/*
You have a set of integers s, which originally contains all the numbers from 1 to n.
Unfortunately, due to some error, one of the numbers in s got duplicated to another number
in the set, which results in repetition of one number and loss of another number.

You are given an integer array nums representing the data status of this set after the error.

Find the number that occurs twice and the number that is missing and return them in the form of an array.

Example 1:
	Input: nums = [1,2,2,4]
	Output: [2,3]

Example 2:
	Input: nums = [1,1]
	Output: [1,2]

Constraints:
	2 <= nums.length <= 10^4
	1 <= nums[i] <= 10^4
*/

func findErrorNums(nums []int) []int {
	result := make([]int, 2)

	// Step 1: Perform cyclic sort to find the number that occurs twice
	for i := 0; i < len(nums); {
		if nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}

	// Step 2: Find the missing number
	for i, num := range nums {
		if num != i+1 {
			result[0] = num
			result[1] = i + 1
			break
		}
	}

	return result
}

func main() {
	//Example 1:
	nums1 := []int{1, 2, 2, 4}
	fmt.Println(nums1)

	//Example 2:
	nums2 := []int{1, 1}
	fmt.Println(nums2)
}

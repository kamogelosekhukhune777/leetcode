package main

import "fmt"

/*
Given an array nums containing n distinct numbers in the range [0, n], return the only number in
the range that is missing from the array.

Example 1:
	Input: nums = [3,0,1]
	Output: 2
	Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3]. 2 is the
	missing number in the range since it does not appear in nums.

Example 2:
	Input: nums = [0,1]
	Output: 2
	Explanation: n = 2 since there are 2 numbers, so all numbers are in the range [0,2]. 2 is the missing
	number in the range since it does not appear in nums.

Example 3:
	Input: nums = [9,6,4,2,3,5,7,0,1]
	Output: 8
	Explanation: n = 9 since there are 9 numbers, so all numbers are in the range [0,9]. 8 is the missing number in
	the range since it does not appear in nums.

Constraints:
	n == nums.length
	1 <= n <= 104
	0 <= nums[i] <= n
	All the numbers of nums are unique.

*/

/*
To find the only missing number in the given array nums, which contains n distinct numbers
in the range [0, n], we can use cyclic sort to rearrange the elements and then iterate
through the sorted array to identify the missing number. Here's how you can implement this approach in Go:

*/
func missingNumber(nums []int) int {
	// Perform cyclic sort to rearrange the elements
	i := 0
	for i < len(nums) {
		if nums[i] < len(nums) && nums[i] != nums[nums[i]] {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		} else {
			i++
		}
	}
	//finding the missing number
	for i, num := range nums {
		if num != i {
			return i
		}
	}

	return len(nums)
}

func main() {
	//Example 1:
	nums1 := []int{3, 0, 1}
	fmt.Println(missingNumber(nums1))

	//Example 2:
	nums2 := []int{0, 1}
	fmt.Println(missingNumber(nums2))

	//Example 3:
	nums3 := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber(nums3))

}

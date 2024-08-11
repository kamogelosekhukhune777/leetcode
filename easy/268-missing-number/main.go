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
// Function to find the missing number in the range [0, n]
func missingNumber(nums []int) int {
	n := len(nums)

	i := 0
	for i < n {
		correctIdx := nums[i]
		// Only swap if the current number is within the range [0, n-1]
		if nums[i] < n && nums[i] != nums[correctIdx] {
			nums[i], nums[correctIdx] = nums[correctIdx], nums[i]
		} else {
			i++
		}
	}

	// After the cyclic sort, find the missing number
	for i = 0; i < n; i++ {
		if nums[i] != i {
			return i
		}
	}

	// If all numbers are at their correct positions, the missing number is n
	return n
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

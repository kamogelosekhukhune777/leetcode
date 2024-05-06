package main

import "fmt"

/*
Given an integer array nums of length n where all the integers of nums are in the range [1, n] and
each integer appears once or twice, return an array of all the integers that appears twice.

You must write an algorithm that runs in O(n) time and uses only constant extra space.

Example 1:
	Input: nums = [4,3,2,7,8,2,3,1]
	Output: [2,3]

Example 2:
	Input: nums = [1,1,2]
	Output: [1]

Example 3:
	Input: nums = [1]
	Output: []

Constraints:
	n == nums.length
	1 <= n <= 105
	1 <= nums[i] <= n
	Each element in nums appears once or twice.
*/

/*
To find the integers that appear twice in the given array nums using cyclic sort and meet the time
complexity constraint of O(n) with constant extra space, we can modify the cyclic sort technique
slightly. Instead of simply swapping elements to their correct positions, we can identify the duplicate
numbers during the cyclic sort process. Here's how you can implement this in Go:
*/

func findDuplicates(nums []int) []int {
	duplicates := []int{}

	// Perform cyclic sort to identify duplicate numbers
	i := 0
	for i < len(nums) {
		if nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else if nums[i] != i+1 {
			duplicates = append(duplicates, nums[i])
			i++
		} else {
			i++
		}
	}

	return duplicates
}

func main() {
	//Example 1:
	nums1 := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(nums1))

	//Example 2:
	nums2 := []int{1, 1, 2}
	fmt.Println(findDuplicates(nums2))

	//Example 3:
	nums3 := []int{1}
	fmt.Println(findDuplicates(nums3))

}

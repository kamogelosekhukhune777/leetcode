package main

import "fmt"

/*
Given an array arr of positive integers sorted in a strictly increasing
order, and an integer k.

Return the kth positive integer that is missing from this array.



Example 1:

	Input: arr = [2,3,4,7,11], k = 5
	Output: 9
	Explanation: The missing positive integers are [1,5,6,8,9,10,12,13,...]. The 5th missing positive integer is 9.

Example 2:

	Input: arr = [1,2,3,4], k = 2
	Output: 6
	Explanation: The missing positive integers are [5,6,7,...]. The 2nd missing positive integer is 6.


Constraints:

	1 <= arr.length <= 1000
	1 <= arr[i] <= 1000
	1 <= k <= 1000
	arr[i] < arr[j] for 1 <= i < j <= arr.length


Follow up:
	Could you solve this problem in less than O(n) complexity?

*/

// Function to find the kth missing positive integer
func findKthPositive(arr []int, k int) int {
	// Define the missingCount function
	missingCount := func(index int) int {
		return arr[index] - (index + 1)
	}

	// Binary search to find the right index
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)/2
		if missingCount(mid) < k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	// Calculate the kth missing positive integer
	return left + k
}

func main() {
	// Test cases
	fmt.Println(findKthPositive([]int{2, 3, 4, 7, 11}, 5)) // Output: 9
	fmt.Println(findKthPositive([]int{1, 2, 3, 4}, 2))     // Output: 6
}

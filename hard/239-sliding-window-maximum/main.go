package main

import (
	"container/list"
	"fmt"
)

/*
You are given an array of integers nums, there is a sliding window of size k which is moving from
the very left of the array to the very right. You can only see the k numbers in the window.
Each time the sliding window moves right by one position.

Return the max sliding window.


Example 1:
	Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
	Output: [3,3,5,5,6,7]
	Explanation:
	Window position                Max
	---------------               -----
	[1  3  -1] -3  5  3  6  7       3
	1 [3  -1  -3] 5  3  6  7       3
	1  3 [-1  -3  5] 3  6  7       5
	1  3  -1 [-3  5  3] 6  7       5
	1  3  -1  -3 [5  3  6] 7       6
	1  3  -1  -3  5 [3  6  7]      7

Example 2:
	Input: nums = [1], k = 1
	Output: [1]


Constraints:
	1 <= nums.length <= 105
	-104 <= nums[i] <= 104
	1 <= k <= nums.length
*/

func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, 0, len(nums)-k+1)

	deque := list.New()

	for i := 0; i < len(nums); i++ {
		//remove elements not within the sliding window
		if deque.Len() > 0 && deque.Front().Value.(int) < i-k {
			deque.Remove(deque.Front())
		}

		//remove elements from the back of the deque which are less than the current element
		//since they are not useful anymore
		for deque.Len() > 0 && nums[deque.Back().Value.(int)] < nums[i] {
			deque.Remove(deque.Back())
		}

		//Add the current element's index to the deque
		deque.PushBack(i)

		//The maximum of the current window is at the front of the deque
		if i >= k-1 {
			result = append(result, nums[deque.Front().Value.(int)])
		}
	}

	return result
}

func main() {

	//example 1
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k)) //Output: [3,3,5,5,6,7]

	//example 2
	nums = []int{1}
	k = 1
	fmt.Println(maxSlidingWindow(nums, k)) //Output: [1]

}

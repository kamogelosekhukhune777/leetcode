package main

import "fmt"

/*
Given an array of integers arr of even length n and an integer k.

We want to divide the array into exactly n / 2 pairs such that the sum of each pair is divisible by k.

Return true If you can find a way to do that or false otherwise.

Example 1:
	Input: arr = [1,2,3,4,5,10,6,7,8,9], k = 5
	Output: true
	Explanation: Pairs are (1,9),(2,8),(3,7),(4,6) and (5,10).

Example 2:
	Input: arr = [1,2,3,4,5,6], k = 7
	Output: true
	Explanation: Pairs are (1,6),(2,5) and(3,4).

Example 3:
	Input: arr = [1,2,3,4,5,6], k = 10
	Output: false
	Explanation: You can try all possible pairs to see that there is no way to divide arr
	into 3 pairs each with sum divisible by 10.

Constraints:
	arr.length == n
	1 <= n <= 105
	n is even.
	-109 <= arr[i] <= 109
	1 <= k <= 105
*/

func canArrange(arr []int, k int) bool {
	// Create a frequency map to store the count of remainders
	remainderFreq := make(map[int]int)

	// Step 1: Count the occurrences of remainders when dividing each element by k
	for _, num := range arr {
		remainder := num % k
		if remainder < 0 {
			remainder += k // Ensure positive remainders
		}
		remainderFreq[remainder]++
	}

	// Step 2: Pair up elements with complementary remainders
	for remainder, count := range remainderFreq {
		// If remainder is 0, ensure it's paired with another 0 remainder to form a valid pair
		if remainder == 0 {
			if count%2 != 0 {
				return false
			}
		} else if remainder*2 == k { // Check if the remainder is half of k
			if count%2 != 0 {
				return false
			}
		} else { // Check if the complementary remainder exists
			compRemainder := k - remainder
			if remainderFreq[compRemainder] != count {
				return false
			}
		}
	}

	// If all remainders are paired up, return true
	return true
}

func main() {
	//Example 1:
	arr := []int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}
	k := 5
	output := canArrange(arr, k)
	if output {
		fmt.Println("can rearrange example:1")
	} else {
		fmt.Println("can't rearrange example:1")
	}

	//Example 2:
	arr = []int{1, 2, 3, 4, 5, 6}
	k = 7
	output = canArrange(arr, k)
	if output {
		fmt.Println("can rearrange example:2")
	} else {
		fmt.Println("can't rearrange example:2")
	}

	//Example 3:
	arr = []int{1, 2, 3, 4, 5, 6}
	k = 10
	output = canArrange(arr, k)
	if output {
		fmt.Println("can rearrange example:3")
	} else {
		fmt.Println("can't rearrange example:3")
	}
}

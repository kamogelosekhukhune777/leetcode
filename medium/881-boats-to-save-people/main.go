package main

import (
	"fmt"
	"sort"
)

//https://leetcode.com/problems/boats-to-save-people/description/
/*

You are given an array people where people[i] is the weight of the ith person, and an infinite
number of boats where each boat can carry a maximum weight of limit. Each boat carries at most
two people at the same time, provided the sum of the weight of those people is at most limit.

Return the minimum number of boats to carry every given person.

Example 1:

	Input: people = [1,2], limit = 3
	Output: 1
	Explanation: 1 boat (1, 2)

Example 2:

	Input: people = [3,2,2,1], limit = 3
	Output: 3
	Explanation: 3 boats (1, 2), (2) and (3)

Example 3:

	Input: people = [3,5,3,4], limit = 5
	Output: 4
	Explanation: 4 boats (3), (3), (4), (5)


Constraints:

	1 <= people.length <= 5 * 10^4
	1 <= people[i] <= limit <= 3 * 10^4
*/

// numRescueBoats returns the minimum number of boats required to carry all people
func numRescueBoats(people []int, limit int) int {
	// Sort the people by their weights
	sort.Ints(people)

	// Initialize pointers and boat counter
	i, j := 0, len(people)-1
	boats := 0

	// Iterate while there are people left to place on boats
	for i <= j {
		boats++ // Every iteration means one boat is used
		if people[i]+people[j] <= limit {
			i++ // If the lightest and heaviest can share a boat, move the pointer for the lightest
		}
		j-- // Always move the pointer for the heaviest person (they either share a boat or go alone)
	}

	return boats
}

func main() {
	people1 := []int{1, 2}
	limit1 := 3
	fmt.Println(numRescueBoats(people1, limit1)) // Output: 1

	people2 := []int{3, 2, 2, 1}
	limit2 := 3
	fmt.Println(numRescueBoats(people2, limit2)) // Output: 3

	people3 := []int{3, 5, 3, 4}
	limit3 := 5
	fmt.Println(numRescueBoats(people3, limit3)) // Output: 4
}

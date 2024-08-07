package main

import (
	"fmt"
	"sort"
)

//https://leetcode.com/problems/two-city-scheduling/description/
/*
A company is planning to interview 2n people. Given the array costs where costs[i] = [aCosti, bCosti], the cost of
flying the ith person to city a is aCosti, and the cost of flying the ith person to city b is bCosti.

Return the minimum cost to fly every person to a city such that exactly n people arrive in each city.


Example 1:

	Input: costs = [[10,20],[30,200],[400,50],[30,20]]
	Output: 110
	Explanation:
		The first person goes to city A for a cost of 10.
		The second person goes to city A for a cost of 30.
		The third person goes to city B for a cost of 50.
		The fourth person goes to city B for a cost of 20.

		The total minimum cost is 10 + 30 + 50 + 20 = 110 to have half the people interviewing in each city.
Example 2:

	Input: costs = [[259,770],[448,54],[926,667],[184,139],[840,118],[577,469]]
	Output: 1859

Example 3:

	Input: costs = [[515,563],[451,713],[537,709],[343,819],[855,779],[457,60],[650,359],[631,42]]
	Output: 3086


Constraints:

	2 * n == costs.length
	2 <= costs.length <= 100
	costs.length is even.
	1 <= aCosti, bCosti <= 1000
*/

// twoCitySchedCost returns the minimum cost to fly exactly n people to each city
func twoCitySchedCost(costs [][]int) int {
	// Sort costs based on the difference of cost[i][0] - cost[i][1]
	sort.Slice(costs, func(i, j int) bool {
		return (costs[i][0] - costs[i][1]) < (costs[j][0] - costs[j][1])
	})

	totalCost := 0
	n := len(costs) / 2

	// Sum up the cost for the first n people to city A and the rest to city B
	for i := 0; i < n; i++ {
		totalCost += costs[i][0] // First n people to city A
	}
	for i := n; i < 2*n; i++ {
		totalCost += costs[i][1] // Remaining n people to city B
	}

	return totalCost
}

func main() {
	costs1 := [][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}
	fmt.Println(twoCitySchedCost(costs1)) // Output: 110

	costs2 := [][]int{{259, 770}, {448, 54}, {926, 667}, {184, 139}, {840, 118}, {577, 469}}
	fmt.Println(twoCitySchedCost(costs2)) // Output: 1859

	costs3 := [][]int{{515, 563}, {451, 713}, {537, 709}, {343, 819}, {855, 779}, {457, 60}, {650, 359}, {631, 42}}
	fmt.Println(twoCitySchedCost(costs3)) // Output: 3086
}

package main

import "fmt"

/*
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are
given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must
take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return true if you can finish all courses. Otherwise, return false.


Example 1:

	Input: numCourses = 2, prerequisites = [[1,0]]
	Output: true
	Explanation: There are a total of 2 courses to take.
		To take course 1 you should have finished course 0. So it is possible.

Example 2:

	Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
	Output: false
	Explanation: There are a total of 2 courses to take.
		To take course 1 you should have finished course 0, and to take course 0 you should also
		have finished course 1. So it is impossible.


Constraints:

	1 <= numCourses <= 2000
	0 <= prerequisites.length <= 5000
	prerequisites[i].length == 2
	0 <= ai, bi < numCourses
	All the pairs prerequisites[i] are unique.
*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	// Step 1: Initialize the graph and in-degree array
	graph := make(map[int][]int)
	inDegree := make([]int, numCourses)

	// Step 2: Build the graph and in-degree array from prerequisites
	for _, prereq := range prerequisites {
		course, pre := prereq[0], prereq[1]
		graph[pre] = append(graph[pre], course)
		inDegree[course]++
	}

	// Step 3: Initialize the queue with courses having zero in-degrees
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// Step 4: Process the queue for topological sorting
	processedCourses := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		processedCourses++

		// Decrease the in-degree of neighbors
		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// Step 5: Check if all courses were processed (no cycles)
	return processedCourses == numCourses
}

func main() {
	numCourses1 := 2
	prerequisites1 := [][]int{{1, 0}}
	fmt.Println(canFinish(numCourses1, prerequisites1)) // Output: true

	numCourses2 := 2
	prerequisites2 := [][]int{{1, 0}, {0, 1}}
	fmt.Println(canFinish(numCourses2, prerequisites2)) // Output: false
}

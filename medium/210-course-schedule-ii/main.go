package main

import "fmt"

/*
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You
are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take
course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return the ordering of courses you should take to finish all courses. If there are many valid
answers, return any of them. If it is impossible to finish all courses, return an empty array.


Example 1:

	Input: numCourses = 2, prerequisites = [[1,0]]
	Output: [0,1]
	Explanation: There are a total of 2 courses to take. To take course 1 you should have
		finished course 0. So the correct course order is [0,1].

Example 2:

	Input: numCourses = 4, prerequisites = [[1,0],[2,0],[3,1],[3,2]]
	Output: [0,2,1,3]
	Explanation: There are a total of 4 courses to take. To take course 3 you should have finished
		both courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
		So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3].

Example 3:

	Input: numCourses = 1, prerequisites = []
	Output: [0]


Constraints:

	1 <= numCourses <= 2000
	0 <= prerequisites.length <= numCourses * (numCourses - 1)
	prerequisites[i].length == 2
	0 <= ai, bi < numCourses
	ai != bi
	All the pairs [ai, bi] are distinct

*/

func findOrder(numCourses int, prerequisites [][]int) []int {
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
	var topoOrder []int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		topoOrder = append(topoOrder, node)

		// Decrease the in-degree of neighbors
		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// Step 5: Check if topological sort is possible (i.e., no cycles)
	if len(topoOrder) == numCourses {
		return topoOrder
	}

	return []int{} // Cycle detected, topological sort not possible
}

func main() {
	numCourses1 := 2
	prerequisites1 := [][]int{{1, 0}}
	fmt.Println(findOrder(numCourses1, prerequisites1)) // Output: [0, 1]

	numCourses2 := 4
	prerequisites2 := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	fmt.Println(findOrder(numCourses2, prerequisites2)) // Output: [0, 2, 1, 3] or [0, 1, 2, 3]

	numCourses3 := 1
	prerequisites3 := [][]int{}
	fmt.Println(findOrder(numCourses3, prerequisites3)) // Output: [0]
}

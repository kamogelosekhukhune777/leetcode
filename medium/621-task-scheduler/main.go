package main

import (
	"fmt"
)

/*
You are given an array of CPU tasks, each represented by letters A to Z, and a cooling time, n. Each cycle
or interval allows the completion of one task. Tasks can be completed in any order, but there's a constraint:
identical tasks must be separated by at least n intervals due to cooling time.

​Return the minimum number of intervals required to complete all tasks.


Example 1:

	Input: tasks = ["A","A","A","B","B","B"], n = 2
	Output: 8
	Explanation: A possible sequence is: A -> B -> idle -> A -> B -> idle -> A -> B.

	After completing task A, you must wait two cycles before doing A again. The same applies to task B. In the
	3rd interval, neither A nor B can be done, so you idle. By the 4th cycle, you can do A again as 2 intervals have passed.

Example 2:

	Input: tasks = ["A","C","A","B","D","B"], n = 1
	Output: 6
	Explanation: A possible sequence is: A -> B -> C -> D -> A -> B.

	With a cooling interval of 1, you can repeat a task after just one other task.

Example 3:

	Input: tasks = ["A","A","A", "B","B","B"], n = 3
	Output: 10
	Explanation: A possible sequence is: A -> B -> idle -> idle -> A -> B -> idle -> idle -> A -> B.

	There are only two types of tasks, A and B, which need to be separated by 3 intervals. This leads to idling twice
	between repetitions of these tasks.


Constraints:

	1 <= tasks.length <= 104
	tasks[i] is an uppercase English letter.
	0 <= n <= 100
*/

func leastInterval(tasks []byte, n int) int {
	// Step 1: Count the frequency of each task
	frequency := make(map[byte]int)
	for _, task := range tasks {
		frequency[task]++
	}

	// Step 2: Find the task with the maximum frequency
	maxFreq := 0
	for _, freq := range frequency {
		if freq > maxFreq {
			maxFreq = freq
		}
	}

	// Step 3: Count how many tasks have the maximum frequency
	maxFreqCount := 0
	for _, freq := range frequency {
		if freq == maxFreq {
			maxFreqCount++
		}
	}

	// Step 4: Calculate the minimum number of intervals
	partCount := maxFreq - 1
	partLength := n - (maxFreqCount - 1)
	emptySlots := partCount * partLength
	availableTasks := len(tasks) - (maxFreq * maxFreqCount)
	idleSlots := max(0, emptySlots-availableTasks)

	// The result is the total length including the idle slots
	return len(tasks) + idleSlots
}

// Helper function to find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example 1
	tasks1 := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n1 := 2
	result1 := leastInterval(tasks1, n1)
	fmt.Println("Output:", result1) // Output: 8

	// Example 2
	tasks2 := []byte{'A', 'C', 'A', 'B', 'D', 'B'}
	n2 := 1
	result2 := leastInterval(tasks2, n2)
	fmt.Println("Output:", result2) // Output: 6

	// Example 3
	tasks3 := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n3 := 3
	result3 := leastInterval(tasks3, n3)
	fmt.Println("Output:", result3) // Output: 10
}

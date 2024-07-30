package main

import (
	"container/heap"
	"fmt"
)

//https://leetcode.com/problems/minimum-number-of-refueling-stops/description/
/*
A car travels from a starting position to a destination which is target miles east of the
starting position.

There are gas stations along the way. The gas stations are represented as an array stations
where stations[i] = [positioni, fueli] indicates that the ith gas station is positioni miles
east of the starting position and has fueli liters of gas.

The car starts with an infinite tank of gas, which initially has startFuel liters of fuel
in it. It uses one liter of gas per one mile that it drives. When the car reaches a gas
station, it may stop and refuel, transferring all the gas from the station into the car.

Return the minimum number of refueling stops the car must make in order to reach its
destination. If it cannot reach the destination, return -1.

Note that if the car reaches a gas station with 0 fuel left, the car can still refuel
there. If the car reaches the destination with 0 fuel left, it is still
considered to have arrived.


Example 1:

	Input: target = 1, startFuel = 1, stations = []
	Output: 0
	Explanation: We can reach the target without refueling.

Example 2:

	Input: target = 100, startFuel = 1, stations = [[10,100]]
	Output: -1
	Explanation: We can not reach the target (or even the first gas station).

Example 3:

	Input: target = 100, startFuel = 10, stations = [[10,60],[20,30],[30,30],[60,40]]
	Output: 2
	Explanation: We start with 10 liters of fuel.
		We drive to position 10, expending 10 liters of fuel.  We refuel from 0 liters to 60 liters of gas.
		Then, we drive from position 10 to position 60 (expending 50 liters of fuel),
		and refuel from 10 liters to 50 liters of gas.  We then drive to and reach the target.
		We made 2 refueling stops along the way, so we return 2.


Constraints:

	1 <= target, startFuel <= 10^9
	0 <= stations.length <= 500
	1 <= positioni < positioni+1 < target
	1 <= fueli < 10^9
*/

// A max-heap of integers
type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]

}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	// Max-heap to store fuel amounts from stations we passed
	fuelHeap := &MaxHeap{}
	heap.Init(fuelHeap)

	curFuel := startFuel
	stops := 0
	prevPosition := 0

	for _, station := range stations {
		position, fuel := station[0], station[1]

		// Travel to the current station
		curFuel -= position - prevPosition

		// If we don't have enough fuel to reach the station, refuel from the heap
		for curFuel < 0 && fuelHeap.Len() > 0 {
			curFuel += heap.Pop(fuelHeap).(int)
			stops++
		}

		// If we still can't reach the station, return -1
		if curFuel < 0 {
			return -1
		}

		// Add the current station's fuel to the heap
		heap.Push(fuelHeap, fuel)
		prevPosition = position
	}

	// Check if we can reach the target from the last station
	curFuel -= target - prevPosition
	for curFuel < 0 && fuelHeap.Len() > 0 {
		curFuel += heap.Pop(fuelHeap).(int)
		stops++
	}

	if curFuel < 0 {
		return -1
	}

	return stops
}

func main() {
	fmt.Println(minRefuelStops(1, 1, [][]int{}))                                          // Output: 0
	fmt.Println(minRefuelStops(100, 1, [][]int{{10, 100}}))                               // Output: -1
	fmt.Println(minRefuelStops(100, 10, [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}})) // Output: 2
}

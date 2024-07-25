package main

import (
    "fmt"
)

// Function to search for a target element in a sorted array using binary search
func search(nums []int, target int) int {
    low, high := 0, len(nums)-1
    
    for low <= high {
        mid := low + (high - low) / 2
        
        if nums[mid] == target {
            return mid // Target found, return its index
        } else if nums[mid] < target {
            low = mid + 1 // Search in the right half
        } else {
            high = mid - 1 // Search in the left half
        }
    }
    
    return -1 // Target not found
}

func main() {
    nums1 := []int{-1, 0, 3, 5, 9, 12}
    target1 := 9
    fmt.Printf("Index of %d in array: %d\n", target1, search(nums1, target1)) // Output: 4

    nums2 := []int{-1, 0, 3, 5, 9, 12}
    target2 := 2
    fmt.Printf("Index of %d in array: %d\n", target2, search(nums2, target2)) // Output: -1
}

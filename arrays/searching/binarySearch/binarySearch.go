/* requires already ascending order sorted array */
package main

import "fmt"

func main() {
	searchArray := []int{1, 2, 5, 7}
	searchItem := 1
	result := binarySearch(searchArray, searchItem)
	if result != -1 {
		fmt.Printf("found at index %d \n", result)
	} else {
		fmt.Println("not found")
	}
}
func binarySearch(arr []int, item int) int {

	leftPointer := 0
	rightPointer := len(arr) - 1
	for leftPointer <= rightPointer {
		midPointer := (leftPointer + rightPointer) / 2
		midValue := arr[midPointer]
		if midValue == item {
			return midPointer
		} else if midValue < item {
			leftPointer = midPointer + 1
		} else {
			rightPointer = midPointer - 1
		}

	}
	return -1
}

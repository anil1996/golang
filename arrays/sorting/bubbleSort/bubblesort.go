package main

import "fmt"

func main() {
	mySlice := []int{2, 5, 4, 3, 1}
	sortedSlice := bubbleSortNormal(mySlice)
	for _, v := range sortedSlice {
		fmt.Println(v)
	}
	sortedSlice = bubbleSortEfficient(mySlice)
	for _, v := range sortedSlice {
		fmt.Println(v)
	}
}

func bubbleSortNormal(mySlice []int) []int {
	for i := 0; i < len(mySlice); i++ {
		for j := 0; j < len(mySlice)-1; j++ {
			if mySlice[j] > mySlice[j+1] {
				mySlice[j], mySlice[j+1] = mySlice[j+1], mySlice[j]
			}
		}
	}
	return mySlice
}
func bubbleSortEfficient(mySlice []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(mySlice)-1; i++ {
			if mySlice[i] > mySlice[i+1] {
				mySlice[i], mySlice[i+1] = mySlice[i+1], mySlice[i]
				swapped = true
			}
		}
	}
	return mySlice
}

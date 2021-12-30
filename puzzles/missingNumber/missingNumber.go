package main

import "fmt"

func main() {
	arr := []int{1, 5, 3, 4, 6, 8, 7, 9, 10}
	n := 10
	fmt.Println("missing number is ", missingNumber(arr, n))
	fmt.Println("missing number using map way", missingNumberHardway(arr, n))
}
func missingNumber(arr []int, n int) int {
	sum := n * (n + 1) / 2
	arrSum := 0
	for _, v := range arr {
		arrSum += v
	}
	return sum - arrSum
}
func missingNumberHardway(arr []int, n int) int {
	m := make(map[int]bool)
	for _, v := range arr {
		m[v] = true
	}
	for i := 1; i <= len(arr); i++ {
		if _, ok := m[i]; !ok {
			return i

		}
	}
	return 0
}

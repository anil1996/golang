package main

import "fmt"

func main() {
	space := 5

	fmt.Println("Count of numbers of perfect square less than this space is", countPerfectSquare(space))
}
func countPerfectSquare(n int) int {
	count := 0
	for i := 1; i < n; i++ {
		if i*i < n {
			count++
		} else {
			continue
		}
	}
	return count
}

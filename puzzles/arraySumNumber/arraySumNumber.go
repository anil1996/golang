/* How to find all pairs of elements in an array whose sum is equal to given number? */
package main

import "fmt"

type sub struct {
	a int
	b int
}

func main() {
	a := []int{4, 5, 7, 11, 9, 13, 8, 12}
	n := 20
	subStruct := findSub(a, n)
	fmt.Println(subStruct)

}
func findSub(arr []int, n int) []sub {
	temp := make([]sub, 0)
	for i := 0; i < len(arr); i++ {
		for j := 1 + 1; j < len(arr); j++ {
			if sum := arr[i] + arr[j]; sum == n {
				temp = append(temp, sub{arr[i], arr[j]})
			}
		}
	}
	return temp

}

/* improve program to consider {11,9} and {9,11} as same- */

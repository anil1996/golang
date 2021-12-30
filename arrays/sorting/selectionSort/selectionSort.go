package main

import "fmt"

func main() {
	fmt.Println("Selection Sort")
	slice := []int{1, 5, 4, 2, 3}
	for i := 0; i < len(slice); i++ {
		min := i
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[min] {
				min = j
			}

		}
		slice[i], slice[min] = slice[min], slice[i]
	}
	fmt.Println(slice)
}

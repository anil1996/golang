package main

import (
	"fmt"
)

func main() {
	fmt.Println("hoo")
	str := "Hello This is My Interview"
	reversed := reverse(str)
	fmt.Println("reversed string", reversed)
}

func reverse(str string) string {
	splitted := []rune(str)

	for i, j := 0, len(splitted)-1; i < j; i, j = i+1, j-1 {
		splitted[i], splitted[j] = splitted[j], splitted[i]

	}
	return string(splitted)
}

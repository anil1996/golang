package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello To The Interviewer"
	fmt.Println("reversed words in string", reverseWords(str))
}
func reverseWords(str string) string {
	words := strings.Fields(str)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

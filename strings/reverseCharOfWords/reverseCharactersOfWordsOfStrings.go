package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello Good Evening"
	fmt.Println("reversed characters of words in string is ", reverseCharactersOfWords(str))
}
func reverseCharactersOfWords(str string) string {
	var word1 []string
	words := strings.Fields(str)
	for _, words := range words {
		chars := []rune(words)
		for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
			chars[i], chars[j] = chars[j], chars[i]
		}
		word1 = append(word1, string(chars))
	}
	return strings.Join(word1, " ")
}

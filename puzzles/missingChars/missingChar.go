package main

import (
	"fmt"
)

func main() {
	str1 := "ac"
	str2 := "ab"
	str3 := "abcdeg"
	str4 := "afc"
	fmt.Println("whats missing in str1", missingChar(str1))
	fmt.Println("whats missing in str1", missingChar(str2))
	fmt.Println("whats missing in str1", missingChar(str3))
	fmt.Println("different characters in two strings are", missingCharInString(str1, str4))
}
func missingChar(str string) string {
	chars := []rune(str)
	j := 0
	count := 0
	for i := 0; i < len(chars)-1; i++ {
		if chars[i] == chars[i+1]-1 {
			continue
		} else {
			count = count + 1
			j = i
			break
		}
	}
	if count != 0 {
		return string(chars[j] + 1)
	} else {
		return ""
	}

}
func missingCharInString(str1, str2 string) string {
	length := 0
	if len(str1) > len(str2) {
		length = len(str2)
	} else {
		length = len(str1)
	}
	chars1 := []rune(str1)
	chars2 := []rune(str2)
	i := 0
	var smallStrSum rune
	var largeStrSum rune
	for ; i < length; i++ {
		smallStrSum += chars1[i]
		largeStrSum += chars2[i]
	}
	largeStrSum += chars2[i]
	extra := largeStrSum - smallStrSum
	return string(extra)
}

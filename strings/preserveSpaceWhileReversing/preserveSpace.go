package main

import (
	"fmt"
)

func main() {
	str := "Reverse this string with space preserved"
	reverseStr := preserveSpace(str)
	fmt.Println(str)
	fmt.Println(reverseStr)
}

func preserveSpace(str string) string {
	chars := []rune(str)
	tempChars := make([]rune, len(chars))
	for i := 0; i < len(chars); i++ {
		if chars[i] == ' ' {
			tempChars[i] = ' '
		}
	}
	for i, j := 0, len(tempChars)-1; i < len(chars); i++ {
		if chars[i] != ' ' {
			if tempChars[j] == ' ' {
				j--
			}
			tempChars[j] = chars[i]
			j--
		}
	}
	return string(tempChars)
}

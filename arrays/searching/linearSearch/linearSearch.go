package main

import "fmt"

func main() {
	searchArray := []string{"abc", "def", "ghi", "jkl"}
	searchString := "gi"
	result := linearSearch(searchArray, searchString)
	fmt.Println(result)

}
func linearSearch(str []string, search string) bool {
	var result bool
	for _, val := range str {
		if val == search {
			result = true
		}
	}
	return result
}

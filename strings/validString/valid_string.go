/*
Question: Given a string, Sherlock considers it valid if all the characters in the string occur the same number of time. However, a string is also valid if the frequencies are same after removing any one character.

Example 1:
Input: str = “aabbcd”
Output: NO

Example 2:
Input: str = “abcc”
Output: YES
*/
package main

import (
	"fmt"
	"sort"
)

/*
 * Complete the 'isValid' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func isValid(s string) string {
	str := []rune(s)
	m := make(map[rune]int)

	for _, v := range str {
		m[v]++
	}
	fmt.Println("map is", m)
	v := make([]int, 0, len(m))

	for _, value := range m {
		v = append(v, value)
	}
	fmt.Println("array", v)
	sort.Ints(v)
	fmt.Println("sorted occurance", v)
	first := v[0]
	secondlast := v[len(v)-2]
	last := v[len(v)-1]
	second := v[1]
	if first == last {
		return "YES"
	}
	if first == 1 && second == last {
		return "YES"
	}
	if first == second && second == secondlast && secondlast == (last)-1 {
		return "YES"
	}

	return "NO"
}

func main() {
	// reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	// s := readLine(reader)

	result := isValid("aabbccddeefghi")
	fmt.Println("result is", result)

	// fmt.Fprintf(writer, "%s\n", result)

	// writer.Flush()
}

// func readLine(reader *bufio.Reader) string {
//     str, _, err := reader.ReadLine()
//     if err == io.EOF {
//         return ""
//     }

//     return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
//     if err != nil {
//         panic(err)
//     }
// }

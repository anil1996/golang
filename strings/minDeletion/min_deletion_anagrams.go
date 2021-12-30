package main

import (
	"fmt"
	"math"
)

/*
 * Complete the 'makeAnagram' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING a
 *  2. STRING b
 */

func makeAnagram(a string, b string) int32 {
	// Write your code here
	var c1 [26]int
	var c2 [26]int
	var deletion int32
	s1 := []rune(a)
	s2 := []rune(b)
	for i := 0; i < len(s1); i++ {
		position := s1[i] - 'a'
		c1[position]++
	}
	for i := 0; i < len(s2); i++ {
		//fmt.Println("char", s1[i])
		position := s2[i] - 'a'
		//fmt.Println("position", position)
		c2[position]++
	}
	for i := 0; i < 26; i++ {
		diff1 := c1[i] - c2[i]
		diff := math.Abs(float64(diff1))
		deletion += int32(diff)
	}
	return deletion

}

func main() {
	// reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	// a := readLine(reader)

	// b := readLine(reader)

	res := makeAnagram("fcrxzwscanmligyxyvym", "jxwtrhvujlmrpdoqbisbwhmgpmeoke")

	fmt.Println("required deletion", res)
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

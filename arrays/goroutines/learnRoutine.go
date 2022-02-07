package main

import (
	"fmt"
)

func main() {
	oddChanel := make(chan int)
	evenChanel := make(chan int)
	go func() {
		for i := 0; true; i += 2 {
			evenChanel <- i
		}

	}()
	go func() {
		for i := 1; true; i += 2 {
			oddChanel <- i
		}
	}()
	for {
		select {
		case oddmessage := <-oddChanel:
			fmt.Println(oddmessage)
		case evenMessage := <-evenChanel:
			fmt.Println(evenMessage)
		}
	}

}

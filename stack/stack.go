package main

import "fmt"

type Stack []string

func (s *Stack) push(item string) {
	*s = append(*s, item)
}
func (s *Stack) pop() {
	if len(*s) == 0 {
		fmt.Printf("Stack empty")
		return
	}
	*s = (*s)[:len(*s)-1]
}
func (s *Stack) display() {
	fmt.Println("stack elements are")
	for i := 0; i < len(*s); i++ {
		fmt.Println((*s)[i])
	}
}

func main() {
	var stack Stack
	stack.push("ap")
	stack.push("ps")
	stack.push("pa")
	stack.display()
	stack.pop()
	stack.display()
	stack.pop()
	stack.display()
}

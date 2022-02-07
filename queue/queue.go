package main

import "fmt"

type queue []int

func (q *queue) enqueue(number int) {
	*q = append(*q, number)
}
func (q *queue) dequeue() {
	if len(*q) == 0 {
		fmt.Println("queue empty")
		return
	}
	*q = (*q)[1:]
}
func (q *queue) display() {
	fmt.Println("queue is", (*q))

}
func main() {
	var q queue
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.display()
	q.dequeue()
	q.display()
	q.dequeue()
	q.display()

}

/*
  Queues
  Queues are a data structure that follow the "first in, first out"
  pattern. It is essentially a line. This queue implementation is a
  linked list, with each node containiner a pointer to the next node
  in the queue. Go conveniently initialize to nil, so we don't have to
  set a null terminator at the end of the list.
*/
package main

import (
	"fmt"
)

// A node has data and a pointer to the next node or null if it is the last node
type Node struct {
	data string
	next *Node
}

// A queue is a series of nodes, linked together in order of arrival
type Queue struct {
	head *Node
	tail *Node
	size int
}

// Add node to queue
func (q *Queue) enqueue(n *Node) {
	if q.head == nil {
		// Empty queue
		q.head = n
		q.tail = n
		q.size++
	} else {
		q.tail.next = n
		q.tail = n
		q.size++
	}

	fmt.Println("enqueued: ", n.data)
}

// Remove first node from queue
func (q *Queue) dequeue() {
	fmt.Println("dequeue: ", q.head.data)
	q.head = q.head.next
	q.size--
}

// Peek returns the current head of the queue
func (q *Queue) peek() *Node {
	return q.head
}

// Print the queue
func (q *Queue) print() {
	current := q.head
	result := "[ "
	for current.next != nil {
		result += current.data + ", "
		current = current.next
	}
	result += current.data + " ]"

	fmt.Println(result)
}

func main() {
	// Create some nodes
	a := &Node{data: "A"}
	b := &Node{data: "B"}
	c := &Node{data: "C"}
	d := &Node{data: "D"}

	// Create a queue and start adding nodes to it
	q := new(Queue)
	q.enqueue(a)
	q.enqueue(b)
	q.enqueue(c)

	fmt.Println("peek: ", q.peek().data)
	fmt.Printf("Queue size: %d\n", q.size)

	// Remove the first node from the queue (FIFO)
	q.dequeue()
	fmt.Println("peek: ", q.peek().data)

	fmt.Printf("Queue size: %d\n", q.size)

	q.enqueue(d)

	fmt.Printf("Queue size: %d\n", q.size)

	// Print the result
	q.print()
	fmt.Printf("Queue size: %d\n", q.size)

}

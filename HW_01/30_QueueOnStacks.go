package main

import "fmt"

type Queue struct {
	tail []int
	head []int
}

func Rotate(queue *Queue) {
	if len(queue.head) != 0 {
		return
	}

	for i := len(queue.tail) - 1; i >= 0; i-- {
		queue.head = append(queue.head, queue.tail[i])
		queue.tail = queue.tail[:i]
	}
}

func Push(queue *Queue, value int) {
	queue.tail = append(queue.tail, value)
}

func Front(queue *Queue) int {
	Rotate(queue)
	return queue.head[len(queue.head)-1]
}

func Pop(queue *Queue) {
	Rotate(queue)
	queue.head = queue.head[:len(queue.head)-1]
}

func Size(queue *Queue) int {
	return len(queue.head) + len(queue.tail)
}

func TestQueue() {
	var queue Queue
	Push(&queue, 4)
	Push(&queue, 3)
	fmt.Println(Front(&queue))
	Pop(&queue)
	Push(&queue, 7)
	Push(&queue, 8)
	fmt.Println(Front(&queue))
	Pop(&queue)
	fmt.Println(Front(&queue))
	Pop(&queue)
	fmt.Println(Front(&queue), Size(&queue))

}

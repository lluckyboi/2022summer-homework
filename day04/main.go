package main

import (
	"2022summer-homework/day04/queue"
	"log"
)

func main() {
	t := queue.TSlice{
		I:    make([]interface{}, 1),
		Head: 0,
		Tail: -1,
	}
	t.Push(2)
	t.Push("da")
	t.Push(3.21)
	t.Pop()
	log.Println(t.Top())
	log.Println(t.Length())
}

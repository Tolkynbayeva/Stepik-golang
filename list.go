package main

import (
	"container/list"
	"fmt"
)

func Push(elem interface{}, queue *list.List) {
	queue.PushBack(elem)
}

func Pop(queue *list.List) interface{} {
	return queue.Remove(queue.Front())

}

func printQueue(queue *list.List) {
	for temp := queue.Front(); temp != nil; temp = temp.Next() {
		fmt.Print(temp.Value)
	}
}

func ReverseList(l *list.List) *list.List {
	reversedList := list.New()
	for temp := l.Back(); temp != nil; temp = temp.Prev() {
		reversedList.PushBack(temp.Value)
	}
	return reversedList
}

package main

import (
	"container/list"
	"fmt"
)

func main() {
	myList := list.New()
	//myList.PushBack(2)
	//myList.PushBack(5)
	//myList.PushBack(3)
	//myList.PushBack(11)
	//myList.PushBack(12)

	Push(1, myList)
	Push(2, myList)
	Push(3, myList)
	Push(4, myList)
	Push(5, myList)

	myList2 := ReverseList(myList)
	printQueue(myList)
	fmt.Println()
	printQueue(myList2)
}

// ReverseList - функция для реверса списка
func ReverseList(l *list.List) *list.List {
	// Здесь ваш код
	ls := list.New()
	for val := l.Front(); val != nil; val = val.Next() {
		ls.PushFront(val.Value)
	}
	return ls
}
func Push(elem interface{}, queue *list.List) {
	queue.PushBack(elem)
}

func printQueue(queue *list.List) {
	for e := queue.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
}

package main

import (
	"container/list"
	"fmt"
)

//Задача FIFO
//Реализуйте функции для очереди FIFO (First In, First Out (ПЕРВЫЙ пришел, ПЕРВЫЙ вышел)) с помощью списков.
//Должны быть данные функции:
//
//Push (добавление элемента)
//Pop (удаление элемента и его возврат)
//printQueue (печать очереди в одну строку без пробелов)
//Функцию main писать не нужно! Писать код вне функций не нужно.

func main() {
	// Создаем контейнер list и добавляем в него элементы
	myList := list.New()
	//myList.PushBack(2)
	//myList.PushBack(5)
	//myList.PushBack(3)
	//myList.PushBack(11)
	//myList.PushBack(12)

	Push(1, myList)
	Push(1, myList)
	Push(1, myList)
	Push(1, myList)
	Push(1, myList)
	Push(1, myList)
	printQueue(myList)
}

func Push(elem interface{}, queue *list.List) {
	queue.PushBack(elem)
}

func Pop(queue *list.List) interface{} {

	if queue != nil {
		return nil
	}
	elem := queue.Front()

	queue.Remove(elem)
	return elem.Value
}

func printQueue(queue *list.List) {
	for e := queue.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
}

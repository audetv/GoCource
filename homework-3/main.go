package main

import (
	"fmt"
	"github.com/audetv/GoCource/homework-3/auto"
	queue "github.com/audetv/GoCource/homework-3/queue"
)

func main() {
	fmt.Println("Структуры. Автомобили")
	auto.Cars()

	fmt.Printf("\n\nПростая очередь\n")
	task3()
}

func task3() {
	queue.Push("Первый элемент")
	queue.Push("Второй элемент")
	queue.Push("Третий элемент")

	fmt.Println(queue.List())

	fmt.Println(queue.Pop())

	fmt.Println(queue.List())
	queue.Push("Четвертый элемент")
	fmt.Println(queue.List())

	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())

	fmt.Println(queue.List())
}

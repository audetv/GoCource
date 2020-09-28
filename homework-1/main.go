package main

import (
	"fmt"
	"github.com/audetv/GoCource/homework-1/bankdeposit"
	"github.com/audetv/GoCource/homework-1/converter"
	"github.com/audetv/GoCource/homework-1/triangle"
)

func main() {
	fmt.Println("Задача №1")
	converter.Calculate()
	fmt.Println("Задача №2")
	triangle.Calculate()
	fmt.Println("Задача №3")
	bankdeposit.Calculate()
}

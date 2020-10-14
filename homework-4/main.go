package main

import (
	"fmt"
	"github.com/audetv/GoCource/homework-4/calculator"
	"github.com/audetv/GoCource/homework-4/chess"
	"github.com/audetv/GoCource/homework-4/phones"
	"github.com/audetv/GoCource/homework-4/shape"
)

func main() {
	shape.Process()

	phones.Process()

	calc()

	chess.Process()
}

func calc() {
	input := ""
	for {
		fmt.Print("> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}

		if input == "exit" {
			break
		}

		if input == "help" {
			calculator.ShowHelp()
			continue
		}

		if res, err := calculator.Calculate(input); err == nil {
			fmt.Printf("Результат: %v\n", res)
		} else {
			fmt.Println("Не удалось произвести вычисление")
		}
	}

}

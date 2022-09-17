package main

import (
	"fmt"
	"oop/books"
)

type game struct {
	title string
	price float64
}

func (g game) print() {
	fmt.Printf("%-15s: $%.2f\n", g.title, g.price)
}

func main() {
	mobydick := books.Book{
		Title: "moby dick",
		Price: 10,
	}

	minecraft := game{
		title: "minecraft",
		price: 20,
	}

	tetris := game{
		title: "tetris",
		price: 5,
	}

	mobydick.Print()
	minecraft.print()
	tetris.print()

}

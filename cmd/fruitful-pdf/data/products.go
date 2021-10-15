package data

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v6"
)

type Fruit struct {
	Name        string  `fake:"{fruit}"`
	Description string  `fake:"{loremipsumsentence:10}"`
	Price       float64 `fake:"{price:1,10}"`
}

func generateFruit() []string {
	var f Fruit
	gofakeit.Struct(&f)

	froot := []string{}
	froot = append(froot, f.Name, f.Description, fmt.Sprintf("%.2f", f.Price))

	return froot
}

func FruitList(length int) [][]string {
	var fruits [][]string

	for i := 0; i < length; i++ {
		fruit := generateFruit()
		fruits = append(fruits, fruit)
	}

	return fruits
}

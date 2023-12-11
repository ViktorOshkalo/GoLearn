package main

import (
	"fmt"
)

var products = map[int]Product{
	1: {Id: 1, Name: "Book1", Color: "Blue", Weight: 10},
	2: {Id: 2, Name: "Book2", Color: "Red", Weight: 20},
	3: {Id: 3, Name: "Book3", Color: "Orange", Weight: 30},
}

// model
type Product struct {
	Id     int
	Name   string
	Color  string
	Weight float32
}

func GetProductById(id int) Product {
	return products[id]
}

// controller
func DisplayProduct(id int) {
	product := GetProductById(id)
	var productView = GetView(product)
	fmt.Println(productView)
}

// view
func GetView(product Product) string {
	return fmt.Sprintf("Product: %s, color: %s, weight: %f", product.Name, product.Color, product.Weight)
}

func PrintProducts(products map[int]Product) {
	fmt.Println("Your products:")
	for i := range products {
		DisplayProduct(i)
	}
}

func main() {
	fmt.Println("Hello!")
	PrintProducts(products)
}

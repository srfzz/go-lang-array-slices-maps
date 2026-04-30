package main

import "fmt"

type Product struct {
	title string
	id    int64
	price float64
}

func main() {
	productNames := [4]string{"apple", "mango", "guava", "litchi"}
	prices := [4]float64{1.0, 15.45, 16.78, 89.20}
	fmt.Println(prices[2], productNames[2])
}

package main

import "fmt"

type Product struct {
	title string
	id    int64
	price float64
}

func main() {
	prices := [4]float64{1.0, 15.45, 16.78, 89.20}
	fmt.Println(prices)
}

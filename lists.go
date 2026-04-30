package main

import (
	"fmt"
)

func main() {
	prices := []float64{12.22, 80.99}
	updatedPrices := append(prices, 200.56)
	fmt.Printf("length of the pricce is %d and capacity is %d and lemenets are %v addreess is %p \n", len(prices), cap(prices), prices, &prices)
	fmt.Printf("length of the updatedPrice is %d and capacity is %d and lemenets are %v address is %p \n", len(updatedPrices), cap(updatedPrices), updatedPrices, &updatedPrices)
}

// func main() {
// 	productNames := [4]string{"apple", "mango", "guava", "litchi"}
// 	prices := [4]float64{1.0, 15.45, 16.78, 89.20}
// 	fmt.Println(len(prices))
// 	featuredprice := prices[1:3]
// 	fmt.Println(featuredprice)
// 	fmt.Println(len(featuredprice), cap(featuredprice))
// 	slices.Sort(prices[:])
// 	fmt.Println(prices, productNames[2])
// }

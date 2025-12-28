package main

import "fmt"

func main() {
	fmt.Println("Paari has made changes")
}

func CalculateTotal(item_count int, item_price float64) float64 {
	return float64(item_count) * item_price
}

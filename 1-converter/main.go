package main

import "fmt"

func main() {
	const USDtoRUB = 88  // 1 USD = 88 RUB
	const EURtoUSD = 1.3 // 1 EUR = 1.3 USD

	EURtoRUB := USDtoRUB * EURtoUSD

	fmt.Println(EURtoRUB)
}

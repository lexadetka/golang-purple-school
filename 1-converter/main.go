package main

import "fmt"

func main() {
	const RUBtoUSD = 88  // 1 USD = 88 RUB
	const EURtoUSD = 1.3 // 1 EUR = 1.3 USD

	EURtoRUB := RUBtoUSD * EURtoUSD

	fmt.Println(EURtoRUB)
}

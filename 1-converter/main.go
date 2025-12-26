package main

import (
	"fmt"
)

func main() {
	startCurrency, amount, targetCurrency := getUserInput()

	res := calculate(amount, startCurrency, targetCurrency)

	fmt.Println(res)
}

func getUserInput() (string, float64, string) {
	var startCurrency string
	var targetCurrency string
	var amount float64
	fmt.Println("Введите исходную валюту [usd, eur, rub]")

	for {
		fmt.Scan(&startCurrency)

		if startCurrency != "usd" && startCurrency != "eur" && startCurrency != "rub" {
			fmt.Println("Исходная валюта введена неверно")
			continue
		}

		break
	}

	fmt.Println("Введите целевую валюту [usd, eur, rub]")

	for {
		_, err := fmt.Scan(&targetCurrency)

		if err != nil || (targetCurrency != "usd" && targetCurrency != "eur" && targetCurrency != "rub") || targetCurrency == startCurrency {
			fmt.Println("Целевая валюта введена неверно")
			continue
		}

		break
	}

	fmt.Println("Введите объем валюты")

	for {
		_, err := fmt.Scan(&amount)

		if err != nil {
			fmt.Println("Объем валюты должен быть числом")
			continue
		}

		break
	}

	return startCurrency, amount, targetCurrency
}

func calculate(amount float64, startCurrency string, targetCurrency string) float64 {

	rate := getExchangeRate(startCurrency, targetCurrency)

	if rate == -1 {
		fmt.Println("Нет подходящей валютной пары")
	}

	return amount * rate
}

func getExchangeRate(startCurrency string, targetCurrency string) float64 {
	const USDtoRUB float64 = 88  // 1 USD = 88 RUB
	const EURtoUSD float64 = 1.3 // 1 EUR = 1.3 USD
	const EURtoRUB float64 = USDtoRUB * EURtoUSD

	m := map[string]float64{
		"usdrub": USDtoRUB,
		"eurrub": EURtoRUB,
		"eurusd": EURtoUSD,
		"rubusd": 1 / USDtoRUB,
		"rubeur": 1 / EURtoRUB,
		"usdeur": 1 / EURtoUSD,
	}

	value, exists := m[startCurrency+targetCurrency]

	if !exists {
		return -1
	}

	return value
}

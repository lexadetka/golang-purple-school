package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	operation, numbers := getInput()

	res := calculate(operation, numbers)

	fmt.Println(res)
}

func getInput() (string, []float64) {

	var operation string
	var numbers string

	fmt.Println("Operation")

	_, err := fmt.Scan(&operation)

	if err != nil {
		fmt.Println("Error reading input")
		panic(err)
	}

	fmt.Println("Numbers")

	_, err = fmt.Scan(&numbers)

	if err != nil {
		fmt.Println("Error reading input")
		panic(err)
	}

	numbersArr := strings.Split(numbers, ",")
	floatSlice := make([]float64, 0, len(numbersArr))

	for _, s := range numbersArr {
		num, err := strconv.ParseFloat(s, 64)
		if err != nil {
			fmt.Println("Ошибка преобразования:", err)
			panic(err)
		}
		floatSlice = append(floatSlice, num)
	}
	return operation, floatSlice
}

func calculate(operation string, numbers []float64) float64 {

	switch strings.ToLower(operation) {
	case "avg":
		return getAverage(numbers)
	case "sum":
		return getSum(numbers)
	case "med":
		return getMedian(numbers)
	default:
		fmt.Println("Invalid operation")
	}

	return 0.0
}

func getAverage(numbers []float64) float64 {
	sum := 0.0

	for _, number := range numbers {
		sum += number
	}

	return sum / float64(len(numbers))
}

func getSum(numbers []float64) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func getMedian(nums []float64) float64 {
	sort.Float64s(nums) // Сортируем срез

	n := len(nums)
	if n%2 == 0 {
		mid1 := nums[n/2-1]
		mid2 := nums[n/2]
		return (mid1 + mid2) / 2
	}

	return nums[n/2]
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	price := 1397848.75775525555555

	price = 134494.75775525555555

	fmt.Println(price)

	priceStr := Float64ToString(price)

	fmt.Println(priceStr)

	price = StringToFloat64(priceStr)

	fmt.Println(price)
}


func Float64ToString(value float64) string {
	return fmt.Sprintf("%v", value)
}

func StringToFloat64(str string) float64 {
	res, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return res
}
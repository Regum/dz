package main

import "fmt"

func main() {
	const (
		RUB_PER_USD = 90.59
		RUB_PER_EUR = 77.96
	)

	var (
		USD_TO_EUR = RUB_PER_USD / RUB_PER_EUR
		USD_TO_RUB = RUB_PER_USD
	)
	fmt.Printf("1 USD = %.2f EUR", USD_TO_EUR)
	fmt.Println()
	fmt.Printf("1 USD = %.2f RUB", USD_TO_RUB)

}

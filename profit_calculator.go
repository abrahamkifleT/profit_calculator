package main

import (
	"fmt"
)

func main() {

	revenue, err := getUserInput("Revenue: ")
	if err != nil || revenue <= 0 {
		fmt.Println("Invalid input provided")
		return
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil || expenses <= 0 {
		fmt.Println("Invalid input provided")
		return
	}

	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil || taxRate <= 0 {
		fmt.Println("Invalid input provided")
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)

}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	_, err := fmt.Scan(&userInput)
	if err != nil {
		return 0, err
	}

	return userInput, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

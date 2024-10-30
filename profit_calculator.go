package main

import (
	"errors"
	"fmt"
)

func main() {

	revenue, err1 := getUserInput("Revenue: ")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	expenses, err2 := getUserInput("Expenses: ")
	// if err != nil{
	// 	fmt.Println(err)
	// 	return
	// }

	taxRate, err3 := getUserInput("Tax Rate: ")
	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
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
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("value must be positive number")
	}
	return userInput, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

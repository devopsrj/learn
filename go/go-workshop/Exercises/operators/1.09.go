package main

import (
	"fmt"
)

func main() {
	//Main Course
	var total float64 = 2 * 13
	fmt.Println("Sub :", total)

	// Drinks
	total = total + (4 * 2.25)
	fmt.Println("Sub :", total)

	//Discount
	total = total - 5
	fmt.Println("Sub :", total)

	// 10% tip
	tip := total * 0.1
	fmt.Println("Tip :", tip)

	//Split bill
	split := total / 2
	fmt.Println("Split : ", split)

	//Reward every 5th visit
	visitcount := 24
	visitcount = visitcount + 1

	remainder := visitcount % 5
	if remainder == 0 {
		fmt.Println("With this visit, you've earned a reward")
	}
}

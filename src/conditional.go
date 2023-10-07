package main

import (
	"fmt"
	"time"
)

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	// switch case
	today := time.Now()

	switch today.Day(){
	case 5:
		fmt.Println("Today is 5th. Clean your house")
	case 10:
		fmt.Println("Today is 10th. Buy som wine.")
	case 15:
		fmt.Println("Today is 15th. Visit the doctor.")
	case 25:
		fmt.Println("Today is 25th. Buy some food")
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day")
	}

	// multiple switch case
	today1 := time.Now()
	var t int = today1.Day()

	switch t {
	case 5, 10, 15:
		fmt.Println("Clean your house")
	case 25, 26, 27:
		fmt.Println("Buy some food")
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No informatio available for that day.")
	}

	// switch case fallthrough
	today2 := time.Now()
	switch today2.Day() {
	case 5:
		fmt.Println("Clean your house.")
		fallthrough
	case 10:
		fmt.Println("Buy some wine.")
		fallthrough
	case 15:
		fmt.Println("Visit a doctor.")
		fallthrough
	case 25:
		fmt.Println("Buy some food.")
		fallthrough
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
}
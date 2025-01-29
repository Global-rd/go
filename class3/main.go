package main

import (
	"fmt"
	"time"
)

// Iterator function that generates numbers from 0 to n-1
func generate(n int) func(func(int) bool) {
	return func(yield func(int) bool) {
		for i := 0; i < n; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func switchExamples() {
	switch today := time.Now().Weekday(); today {
	case time.Monday:
		fmt.Println("Monday")
	case time.Tuesday:
		fmt.Println("Tuesday")
	case time.Wednesday:
		fmt.Println("Wednesday")
	case time.Thursday:
		fmt.Println("Thursday")
	case time.Friday:
		fmt.Println("Friday")
	default:
		fmt.Println("Weekend!")
	}
}

func branchExamples() {
	balance := 0

	if balance < 0 {
		fmt.Println("Balance is below 0, add funds now or you will be charged a penalty.")
	} else if balance == 0 {
		fmt.Println("Balance is equal to 0, add funds soon.")
	} else {
		fmt.Println("Your balance is greater than 0.")
	}
}

func loopExamples() {
	list := []string{"Follow", "The", "Pattern"}

	for index, value := range list {
		fmt.Println(index, value)
	}
}

func main() {

}

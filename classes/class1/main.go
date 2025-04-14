package main

import (
	"fmt"
)

func main() {
	var target string
	value := "follow the pattern"

	Scan(&target, value)

	fmt.Println("hello", target)
}

func ScanValue(target string, value string) {
	target = value
}

func Scan(target *string, value string) {
	*target = value
}

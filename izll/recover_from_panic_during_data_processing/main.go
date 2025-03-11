package main

import (
	"fmt"
	tm "github.com/buger/goterm"
	"os"
	"recover_from_panic_during_data_processing/level_1"
	"recover_from_panic_during_data_processing/level_2"
	"recover_from_panic_during_data_processing/level_3"
)

func main() {
	handleMenu()
}

func handleMenu() {
	for {
		showMenu()
		result := scanOption()
		runOption(result)
	}
}

func showMenu() {
	tm.Clear()
	tm.MoveCursor(1, 1)
	tm.Flush()
	fmt.Println("Choose a variant of program:")
	fmt.Println("---------------------------")
	fmt.Println("1. Source: random string")
	fmt.Println("2. Source: file")
	fmt.Println("3. Source: third-party server")
	fmt.Println("0. Exit")
	fmt.Println("---------------------------")
}

func scanOption() int {
	var option int
	fmt.Scanln(&option)
	return option
}

func runOption(option int) {
	switch option {
	case 1:
		runLevel1()
	case 2:
		runLevel2()
	case 3:
		runLevel3()
	case 0:
		exit()
	default:
		fmt.Println("Invalid option")
	}
}

func printTitle(title string) {
	fmt.Println()
	fmt.Println(title)
	fmt.Println("----------------")
	fmt.Println()
}

func runLevel1() {
	printTitle("Run random string")
	level_1.Run()

	pause()
}

func runLevel2() {
	printTitle("Run file")
	level_2.Run()

	pause()
}

func runLevel3() {
	printTitle("Run third-party server")
	level_3.Run()

	pause()
}

func pause() {
	fmt.Println()
	fmt.Println("Press enter to continue")
	fmt.Scanln()
}

func exit() {
	fmt.Println("Exit")
	tm.Flush()
	os.Exit(0)
}

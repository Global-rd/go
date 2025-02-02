package main

import (
	"fmt"
	tm "github.com/buger/goterm"
	"os"
	"os/exec"
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
	fmt.Println("1. Hardcoded input in")
	fmt.Println("2. Scanning from input")
	fmt.Println("3. Get value from env vars and convert it")
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
		runHardcoded()
	case 2:
		runScan()
	case 3:
		runEnv()
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

func runHardcoded() {
	printTitle("Run hardcoded")
	runProgram("main_hardcoded.go")
}

func runEnv() {
	printTitle("Run env")
	runProgram("main_env.go")
}

func runScan() {
	printTitle("Run scan")
	runProgram("main_scan.go")
}

func exit() {
	fmt.Println("Exit")
	tm.Flush()
	os.Exit(0)
}

func runProgram(path string) {
	cmd := exec.Command("go", "run", path)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nPress enter to continue")
	fmt.Scanln()
}

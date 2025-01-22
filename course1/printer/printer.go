package printer

import (
	"fmt"
)

type Test struct {
	alma string
}

func PrintInfo(msg string) {
	fmt.Println("[info]", msg)
}

func PrintError(msg string) {
	fmt.Println("[error]", msg)
}

package main

import (
	"fmt"
)

type User struct {
	Name string
}

func main() {
	GenInterface()
}

func GenInterface() {
	var db Cache[string] = &BTree[string]{
		value: "hello",
	}

	var dbInt Cache[int] = &BTree[int]{
		value: 1,
	}

	fmt.Println("db:", db.GetData())
	fmt.Println("db:", dbInt.GetData())

}

func CreateGenericStruct() {
	apple := BTree[string]{
		value: "root",
		Left: &BTree[string]{
			value: "hello",
		},
	}

	fmt.Println("apple", apple)
}

func RunGenericFilter() {
	genResult, ok := First([]User{
		{
			Name: "hello1",
		},
		{
			Name: "hello2",
		},
		{
			Name: "hello3",
		},
	},
		func(s User) bool { return s.Name == "hello2" },
	)

	if ok {
		fmt.Println("gen found", genResult)
	}

	result, ok := FirstString([]string{
		"hello",
		"follow",
		"the",
		"pattern",
	},
		func(s string) bool { return s == "follow" },
	)

	if ok {
		fmt.Println("found", result)
	}

	resultInt, ok := FirstInt([]int{
		0,
		1,
		2,
		3,
	},
		func(i int) bool { return i == 1 },
	)

	if ok {
		fmt.Println("found int", resultInt)
	}

	fmt.Println("finished")
}

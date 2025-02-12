package main

type Int interface {
	GetInt() int
}

type GraphQLInt int

func (g GraphQLInt) GetInt() int {
	return int(g)
}

func example() {
	var customInt GraphQLInt = 1

	var Iint Int = customInt
	// var Iint2 Int = GraphQLInt(1)

	Iint.GetInt() // 1
}

package main

import (
	"fmt"
	"unsafe"
)

type UnoptimizedBase struct {
	A int8
	O []int8
	B int64
	C int8
}

type Unoptimized struct {
	UnoptimizedBase
	D int64
	E int8
	F int8
}

type OptimizedBase struct {
	B int64
	A int8
	C int8
}

type Optimized struct {
	OptimizedBase
	O []int8
	D int64
	E int8
	F int8
}

func layoutExample() {
	fmt.Println("Size of Unoptimized struct:", unsafe.Sizeof(Unoptimized{}))
	fmt.Println("Size of Optimized struct:", unsafe.Sizeof(Optimized{}))
}

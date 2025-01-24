package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	valueTrue := true // bool

	// returns with the reflection type
	tvalue1 := reflect.TypeOf(valueTrue)
	fmt.Println(tvalue1)

	var valueFalse bool = 1 >= 0 // bool

	// returns with the reflection type
	tvalue2 := reflect.TypeOf(valueFalse)
	fmt.Println(tvalue2)

	// whole number type will be int by default
	value1 := 1
	fmt.Println(reflect.TypeOf(value1))

	// different int types must be explicitly declared
	var value2 int8 = 1
	fmt.Println(reflect.TypeOf(value2))

	// or must be casted to the type
	value3 := int16(1)
	tvalue3 := reflect.TypeOf(value3)
	fmt.Println(tvalue3)

	value4 := int32(1)
	tvalue4 := reflect.TypeOf(value4)
	fmt.Println(tvalue4)

	value5 := int64(1)
	tvalue5 := reflect.TypeOf(value5)
	fmt.Println(tvalue5)

	var value6 uint = 1
	tvalue6 := reflect.TypeOf(value6)
	fmt.Println(tvalue6)

	var value7 uint8 = 1
	tvalue7 := reflect.TypeOf(value7)
	fmt.Println(tvalue7)

	var value8 uint16 = 23
	tvalue8 := reflect.TypeOf(value8)
	fmt.Println(tvalue8)

	var value9 uint32 = 23
	tvalue9 := reflect.TypeOf(value9)
	fmt.Println(tvalue9)

	var value10 uint64 = 23
	tvalue10 := reflect.TypeOf(value10)
	fmt.Println(tvalue10)

	// byte is an alias for uint8
	var value11 byte = 8
	tvalue11 := reflect.TypeOf(value11)
	fmt.Println(tvalue11)

	// float value default type is float64
	value12 := 8.7 // float64
	tvalue12 := reflect.TypeOf(value12)
	fmt.Println(tvalue12)

	var value13 float32 = 382.394
	tvalue13 := reflect.TypeOf(value13)
	fmt.Println(tvalue13)

	// returns with complex64 if the inputs are float32 types
	value14 := complex(float32(32.3), float32(18.9))
	tvalue14 := reflect.TypeOf(value14)
	fmt.Println(tvalue14)

	// returns with complex128
	value15 := complex(32.3, 18.9)
	tvalue15 := reflect.TypeOf(value15)
	fmt.Println(tvalue15)

	str := "A😊" // UTF-8 encoding: "A" is 1 byte, "😊" is 4 bytes.

	// Length of the string in bytes
	fmt.Println("String length (in bytes):", len(str)) // Output: 5 (1 byte for "A" + 4 bytes for "😊")

	// Size of a byte
	var b byte = 'A'
	fmt.Println("Size of byte:", unsafe.Sizeof(b)) // Output: 1

	// Size of a rune
	var r rune = '😊'
	fmt.Println("Size of rune:", unsafe.Sizeof(r)) // Output: 4

	// Convert string to runes
	runes := []rune(str)
	fmt.Println("Rune array length:", len(runes)) // Output: 2 (1 rune for "A", 1 rune for "😊")

	// Total memory for rune array
	fmt.Println("Memory for rune array:", len(runes)*int(unsafe.Sizeof(runes[0]))) // Output: 8

	var dec = 42     // Decimal
	var oct = 075    // Octal
	var hex = 0x1A3F // Hexadecimal
	var bin = 0b1010 // Binary

	fmt.Println(dec, oct, hex, bin)

	// Example array
	array := [3]int{10, 20, 30}

	// Get a pointer to the first element
	ptr := unsafe.Pointer(&array[0])

	// Convert to uintptr and manipulate the memory address
	for i := 0; i < len(array); i++ {
		// Use uintptr to calculate the address of the next element
		elementPtr := (*int)(unsafe.Pointer(uintptr(ptr) + uintptr(i)*unsafe.Sizeof(array[0])))

		// Print the value at the calculated address
		fmt.Println(*elementPtr)
	}
}

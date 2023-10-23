package main

import "fmt"

// https://go.dev/ref/spec#Numeric_types

func main() {
	var uint8_var uint8 = 255
	var uint16_var uint16 = 65535
	var uint32_var uint32 = 4294967295
	var uint64_var uint64 = 18446744073709551615
	var uint_var uint = 18446744073709551615

	var int8_var int8 = -127
	var int16_var int16 = -32768
	var int32_var int32 = -2147483648
	var int64_var int64 = -9223372036854775808
	var int_var int = -9223372036854775808

	var byte_var byte = 255

	fmt.Printf("uint8: %d\n", uint8_var)
	fmt.Printf("uint16: %d\n", uint16_var)
	fmt.Printf("uint32: %d\n", uint32_var)
	fmt.Printf("uint64: %d\n", uint64_var)
	fmt.Printf("uint: %d\n", uint_var)

	fmt.Printf("int8: %d\n", int8_var)
	fmt.Printf("int16: %d\n", int16_var)
	fmt.Printf("int32: %d\n", int32_var)
	fmt.Printf("int64: %d\n", int64_var)
	fmt.Printf("int: %d\n", int_var)

	fmt.Printf("byte: %d\n", byte_var)
}

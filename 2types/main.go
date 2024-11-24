package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Boolean type
	var bool1 bool = true
	fmt.Printf("\n--- Boolean ---")
	fmt.Printf("\ntype: %T | value: %v | size: %d bytes", bool1, bool1, unsafe.Sizeof(bool1))

	// Integer types
	fmt.Printf("\n\n--- Integers ---")
	var int1 int = 42
	var int81 int8 = 127
	var int161 int16 = 32767
	var int321 int32 = 2147483647
	var int641 int64 = 9223372036854775807

	fmt.Printf("\ntype: %T | value: %v | size: %d bytes", int1, int1, unsafe.Sizeof(int1))
	fmt.Printf("\ntype: %T | value: %v | size: %d bytes", int81, int81, unsafe.Sizeof(int81))
	fmt.Printf("\ntype: %T | value: %v | size: %d bytes", int161, int161, unsafe.Sizeof(int161))
	fmt.Printf("\ntype: %T | value: %v | size: %d bytes", int321, int321, unsafe.Sizeof(int321))
	fmt.Printf("\ntype: %T | value: %v | size: %d bytes", int641, int641, unsafe.Sizeof(int641))

	// Unsigned integer types
	fmt.Printf("\n\n--- Unsigned Integers ---")
	var uint1 uint = 42
	var uint81 uint8 = 255
	var uint161 uint16 = 65535
	var uint321 uint32 = 4294967295
	var uint641 uint64 = 18446744073709551615
	var uintptr1 uintptr = 18446744073709551615

	fmt.Printf("\ntype: %T | value: %v | size: %d bytes", uint1, uint1, unsafe.Sizeof(uint1))
	fmt.Printf("\ntype: %T | value: %v | size: %d bytes", uint81, uint81, unsafe.Sizeof(uint81))
	fmt.Printf("\ntype: %T | value: %v | size: %d bytes", uint161, uint161, unsafe.Sizeof(uint161))
	fmt.Printf("\ntype: %T | value: %v | size: %d bytes", uint321, uint321, unsafe.Sizeof(uint321))
	fmt.Printf("\ntype: %T | value: %v | size: %d bytes", uint641, uint641, unsafe.Sizeof(uint641))
	fmt.Printf("\ntype: %T | value: %v | size: %d bytes", uintptr1, uintptr1, unsafe.Sizeof(uintptr1))

	// Float types
	fmt.Printf("\n\n--- Floating Point Numbers ---")
	var float321 float32 = 3.14159
	var float641 float64 = 3.14159265359

	fmt.Printf("\ntype: %T | value: %v | size: %d bytes", float321, float321, unsafe.Sizeof(float321))
	fmt.Printf("\ntype: %T | value: %v | size: %d bytes", float641, float641, unsafe.Sizeof(float641))

	// Complex types
	fmt.Printf("\n\n--- Complex Numbers ---")
	var complex641 complex64 = 3.14 + 2.7i
	var complex1281 complex128 = 3.14159 + 2.71828i

	fmt.Printf("\ntype: %T | value: %v | size: %d bytes", complex641, complex641, unsafe.Sizeof(complex641))
	fmt.Printf("\ntype: %T | value: %v | size: %d bytes", complex1281, complex1281, unsafe.Sizeof(complex1281))

	// String and byte
	fmt.Printf("\n\n--- String and Byte ---")
	var string1 string = "Hello, 世界"
	var byte1 byte = 'A' // byte is an alias for uint8
	var rune1 rune = '世' // rune is an alias for int32

	fmt.Printf("\ntype: %T | value: %v | size: %d bytes", string1, string1, unsafe.Sizeof(string1))
	fmt.Printf("\ntype: %T | value: %v | size: %d bytes", byte1, byte1, unsafe.Sizeof(byte1))
	fmt.Printf("\ntype: %T | value: %v | size: %d bytes", rune1, rune1, unsafe.Sizeof(rune1))

	// Array and Slice (examples)
	fmt.Printf("\n\n--- Array and Slice ---")
	var array1 [5]int = [5]int{1, 2, 3, 4, 5}
	var slice1 []int = []int{1, 2, 3, 4, 5}

	fmt.Printf("\ntype: %T | value: %v | size: %d bytes", array1, array1, unsafe.Sizeof(array1))
	fmt.Printf("\ntype: %T | value: %v | size: %d bytes", slice1, slice1, unsafe.Sizeof(slice1))

	fmt.Printf("\n")
}

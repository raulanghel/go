package main

import (
	"fmt"
)

func main() {
	// zero value for bool is false
	var n bool
	fmt.Printf("%v %T\n", n, n)

	// numeric types
	var x uint16 = 22
	fmt.Printf("%v %T\n", x, x)

	a := 10
	b := 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	var c uint8 = 3
	fmt.Println(a + int(c))

	fmt.Printf("%.4b\n", a&b)
	fmt.Printf("%.4b\n", a|b)
	fmt.Printf("%.4b\n", a^b)
	fmt.Printf("%.4b %v\n", a&^b, a&^b)

	y := 8
	fmt.Printf("%v %.8b\n", y<<3, y<<3)
	fmt.Printf("%v %.8b\n", y>>3, y>>3)

	// floating point numbers
	pi := 3.14
	fmt.Printf("%v %T\n", pi, pi) // floats initialize to float64

	// complex types
	var com1 complex64 = 1 + 2i
	var com2 complex64 = 2 + 5.2i
	fmt.Println(com1 + com2)
	fmt.Println(com1 - com2)
	fmt.Println(com1 * com2)
	fmt.Println(com1 / com2)

	fmt.Printf("%v %T\n", real(com2), real(com2))
	fmt.Printf("%v %T\n", imag(com2), imag(com2))

	comp3 := complex(float32(a), float32(b))
	fmt.Printf("%v %T\n", comp3, comp3)

	// strings; any utf-8 char
	s := "this is a string"
	s2 := "this is also a string"
	fmt.Printf("%v %T\n", s, s)

	// it can be treated like an array of chars
	// strings in Go are aliases for bytes
	fmt.Printf("%v, %T", string(s[2]), s[2])

	// strings are generally immutable
	// s[2] = "u" is not permitted

	fmt.Println(s + s2)

	// convert the string to a byte slice
	bt := []byte(s)
	fmt.Printf("%v, %T\n", bt, bt)

	newStr := string(bt)
	fmt.Println(newStr)

	// runes; a rune represents an utf-32 char
	r := 'a' // using single quotes -> it's a rune; runes are type aliases for int32
	fmt.Printf("%v, %T\n", r, r)

}

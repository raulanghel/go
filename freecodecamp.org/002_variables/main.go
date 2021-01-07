package main

import (
	"fmt"
	"strconv"
)

// variables declared at package level; can't use the := declaration
var a int = 999

// var blocks
var (
	firstName string
	lastName  string
	age       int
)

func main() {
	var i int // variable 'i' of type integer; used when we need to declare a variable but we're not ready to initialize it yets
	i = 42
	var j float32 = 27 // initial value set; used when Go doesn't have enough information about the type we want to assign
	k := 15            //
	fmt.Println(i, j, k)
	fmt.Printf("%v %T\n", j, j)
	fmt.Printf("%v, type=%T\n", a, a)
	firstName, lastName, age := "Raul", "Anghel", 25
	fmt.Println(firstName, lastName, age)

	var a int = 1
	fmt.Println(a) // 	SHADOWING: variable with ineer most scope takes precedence

	// Conversions
	var ii int = 44
	fmt.Printf("%v, %T\n", ii, ii)

	var jj float32
	jj = float32(ii)
	fmt.Printf("%v, %T\n", jj, jj)

	// when converting to/from strings, use strconv
	qwe := strconv.Itoa(ii)
	fmt.Printf("%v %T\n", qwe, qwe)

	s := fmt.Sprintf("%f", jj)
	fmt.Printf("%v %T\n", s, s)
}

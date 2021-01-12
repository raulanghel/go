package main

import "fmt"

func main() {
	var a int = 42
	var b *int = &a
	fmt.Println(a, *b)
	a = 27
	fmt.Println(a, *b)
	*b = 14
	fmt.Println(a, *b)

	// pointer arithmetic; not allowed in Go, deemed to be unsafe
	// a := [3]int{1, 2, 3}
	// b := &a[0]
	// c := &a[1]
	// fmt.Printf("%v %p %p", a, b, c)

	//

}

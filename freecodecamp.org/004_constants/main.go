package main

import "fmt"

// enumerated constants
const (
	x = iota
	y = iota
	z = iota
)

const (
	f = iota
	g
	h
)

const (
	_ = iota
	catSpecialist
	dogSpecialist
	snakeSpecoalist
)

func main() {
	// typed const
	const a int = 42
	fmt.Println(a)

	b := 27
	fmt.Printf("%v, %T\n", a-b, a-b)

	fmt.Println(x, y, z)
	fmt.Println(f, g, h)

	var specialistType int = dogSpecialist
	fmt.Println(specialistType == catSpecialist)
}

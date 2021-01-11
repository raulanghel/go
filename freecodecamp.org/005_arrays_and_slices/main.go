package main

import "fmt"

func main() {
	// grades := [3]int{97, 85, 93}
	// fmt.Println(grades)

	// otherGrades := [...]int{77, 88, 65}
	// fmt.Printf("other grades: %v\n", otherGrades)

	// var students [3]string
	// fmt.Println(students)

	// students[0] = "First"
	// students[1] = "Second"
	// students[2] = "Third"

	// fmt.Println(students)

	// fmt.Println("Student #1:", students[0])

	// fmt.Println("Students in the array:", len(students))

	// identityMatrix := [3][3]int{{0, 0, 1}, {0, 1, 0}, {1, 0, 0}}
	// fmt.Println(identityMatrix)

	// a := [...]int{0, 2, 4}
	// b := a
	// b[1] = 9
	// println("a[1]=", a[1])
	// println("b[1]=", b[1])

	// slices
	// aa := []int{1, 2, 3}
	// fmt.Println(aa)
	// fmt.Printf("%v\n", len(aa))
	// fmt.Println("Capacity:", cap(aa))

	// slices are reference types, arrays are value types
	// bb := aa
	// bb[0] = 9
	// fmt.Printf("aa[0]=%v bb[0]=%v\n", aa[0], bb[0])

	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// b := a[:]
	// c := a[3:]
	// d := a[:6]
	// e := a[3:6]
	// a[5] = 42
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)

	// the make function
	a := make([]int, 3)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	b := make([]int, 3, 100)
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))

	c := []int{}
	c = append(c, 1)
	fmt.Println(c)
	c = append(c, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(c)

	d := []int{11, 12, 13, 14}
	e := append(c, d...)
	fmt.Println(e)
}

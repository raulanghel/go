package main

import "fmt"

func main() {
	// sayMessage("Hello there!")
	// sayGreeting("How are you", "Raul")
	// sum(1, 2, 3, 4, 5)
	// sum(3, 3, 3)

	// s := sum1(5, 6, 7)
	// fmt.Println(s)

	// s1 := sum2(9, 9, 9)
	// fmt.Println(*s1)

	// s = sum3(1, 1, 1)
	// fmt.Println(s)

	// d, err := divide(5.0, 1.0)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(d)

	// func() {
	// 	fmt.Println("hello there")
	// }()

	// f := func() {
	// 	fmt.Println("hello again")
	// }
	// f()

	// methods
	g := greeter{
		greeting: "Hello",
		name:     "Raul",
	}
	g.greet()

}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

func sayMessage(msg string) {
	fmt.Println(msg)
}

func sayGreeting(greeting, name string) {
	fmt.Println(greeting, name)
}

func sum(values ...int) {
	fmt.Println(values)
	sum := 0
	for _, v := range values {
		sum += v
	}
	fmt.Println("The sum is", sum)
}

func sum1(values ...int) int {
	fmt.Println(values)
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum

}

func sum2(values ...int) *int {
	fmt.Println(values)
	sum := 0
	for _, v := range values {
		sum += v
	}
	return &sum
}

func sum3(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

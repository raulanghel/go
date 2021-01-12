package main

import "fmt"

func main() {
	statePopulations := map[string]int{
		"California": 39,
		"Texas":      27,
	}

	if pop, ok := statePopulations["California"]; ok {
		fmt.Println(pop)
	}

	number := 50
	guess := 50

	if guess < 0 {
		fmt.Println("Number must be greater than 0")
	} else if guess > 100 {
		fmt.Println("Number must be less than 100")
	} else {
		if guess < number {
			fmt.Println("Too low")
		} else if guess > number {
			fmt.Println("Too high")
		} else {
			fmt.Println("You got it!")
		}
	}

	// switch
	// a := 8

	switch a := 2 + 3; a {
	case 1, 3, 5:
		fmt.Println("one, three or five")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("something else")
	}

	i := 7
	switch {
	case i < 10:
		fmt.Println("i is less than 10")
		fallthrough
	case i > 20:
		fmt.Println("i os less than 20")
	default:
		fmt.Println("i is more 20")
	}

	var inf interface{} = 2
	switch inf.(type) {
	case int:
		fmt.Println("int")
		// break
		fmt.Println("this will print too")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	}
}

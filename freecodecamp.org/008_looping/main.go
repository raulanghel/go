package main

import "fmt"

func main() {
	// for i, j := 0, 1; i <= 5; i, j = i+1, j+1 {
	// 	fmt.Println(i, j)
	// }

	// i := 0
	// for ; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// for i < 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for {
	// 	fmt.Println(i)
	// 	i++
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// for i := 0; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
Loop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j >= 3 {
				break Loop
			}
		}
	}

	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}
}

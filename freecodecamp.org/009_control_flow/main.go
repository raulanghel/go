package main

import "fmt"

func main() {
	// panic
	// a, b := 1, 0
	// ans := a / b
	// fmt.Println(ans)
	fmt.Println("start")
	defer fmt.Println("this was deffered")
	panic("something bad happened")
	// fmt.Println("end")
	// defered functions are executed before panic
}

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	res, err := http.Get("http://google.com/robots.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer res.Body.Close()
// 	robots, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("%s", robots)
// }

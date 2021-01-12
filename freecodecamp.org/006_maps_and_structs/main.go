package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal
	CanFly bool
	Speed  int
}

type Player struct {
	Name string `required max:"100"`
	Team string
}

func main() {
	// anonymous structs
	// someDoctor := struct {
	// 	name  string
	// 	actor string
	// }{name: "Doctor Who", actor: "Raul Anghel"}
	// fmt.Println(someDoctor)

	// type Doctor struct {
	// 	number     int
	// 	actor      string
	// 	companions []string
	// }

	// aDoctor := Doctor{
	// 	number:     1,
	// 	actor:      "Raul Anghel",
	// 	companions: []string{"AAA aaa", "BBB bbb"},
	// }

	// fmt.Printf("%v %v %v\n", aDoctor.number, aDoctor.actor, aDoctor.companions)

	// statePopulations := map[string]int{
	// 	"California": 39007654,
	// 	"Texas":      27880654,
	// }

	// englishNumbers := make(map[int]string)
	// englishNumbers[1] = "one"
	// englishNumbers[2] = "two"
	// englishNumbers[3] = "three"

	// fmt.Println(statePopulations)
	// fmt.Println(englishNumbers)

	// fmt.Printf("Population of California: %v\n", statePopulations["California"])
	// delete(englishNumbers, 2)
	// fmt.Println(englishNumbers)

	// pop, okay := statePopulations["Texas"]
	// fmt.Println(pop, okay)

	// fmt.Println(len(statePopulations))

	// structs

	// embedding
	bird := Bird{}
	bird.Name = "Pinguin"
	bird.Origin = "Antarctica"
	bird.CanFly = false
	bird.Speed = 10
	fmt.Println(bird)

	t := reflect.TypeOf(Player{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}

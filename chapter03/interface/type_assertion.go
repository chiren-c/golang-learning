package main

import (
	"fmt"
	"reflect"

	a "github.com/chiren-c/golang-learning/chapter03/animal"
)

type IAnimal interface {
	Call() string
	FavorFood() string
}

func main() {
	var animal = a.NewAnimal("中华田园犬")
	var pet = a.NewPet("泰迪")
	var ianimal IAnimal = a.NewDog(&animal, &pet)
	if dog, ok := ianimal.(a.Dog); ok {
		fmt.Println(dog.GetPetName())
		fmt.Println(dog.Call())
		fmt.Println(dog.FavorFood())
		fmt.Println(reflect.TypeOf(ianimal))
	}

}

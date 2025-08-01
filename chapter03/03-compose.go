package main

import (
	"fmt"

	"github.com/chiren-c/golang-learning/chapter03/animal"
)

func main() {
	// a := animal.Animal{Name: "中华田园犬"}
	// pet := animal.Pet{Name: "宠物狗"}
	a := animal.NewAnimal("中华田园犬")
	pet := animal.NewPet("宠物狗")
	dog := animal.NewDog(&a, &pet)
	fmt.Println(dog.GetPetName())
	fmt.Println(dog.Call())
	fmt.Println(dog.FavorFood())
}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	// myFunc(1, 2, 3, 4)
	// myFunc(s1...)
	// myFunc1(s1[1:3]...)
	myFunc2(s1)
	myPrintf(1, "1", true)
}

func myFunc1(number ...int) {
	for _, v := range number {
		fmt.Println(v)
	}
}

func myFunc2(number []int) {
	for _, v := range number {
		fmt.Println(v)
	}
}

func myPrintf(args ...interface{}) {
	for _, v := range args {
		switch reflect.TypeOf(v).Kind() {
		case reflect.Int:
			fmt.Println(v, "int")
		case reflect.String:
			fmt.Println(v, "string")
		default:
			fmt.Println(v, "default")
		}
	}
}

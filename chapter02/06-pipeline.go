package main

import (
	"fmt"
	"log"
)

type user struct {
	name string
	age  int
}

func main() {
	var users = []user{
		{
			name: "张三",
			age:  18,
		},
		{
			name: "李四",
			age:  22,
		},
		{
			name: "王五",
			age:  20,
		},
		{
			name: "赵六",
			age:  -10,
		},
		{
			name: "孙七",
			age:  60,
		},
		{
			name: "周八",
			age:  10,
		},
	}
	sum := sumAge(users, filterAge, mapAgeToSlice)
	log.Printf("用户年龄累加结果: %d\n", sum)
}

func filterAge(u []user) interface{} {
	var slice []user
	for _, v := range u {
		if v.age >= 18 {
			slice = append(slice, v)
		}
	}
	return slice
}

func mapAgeToSlice(u []user) interface{} {
	var slice []int
	for _, v := range u {
		slice = append(slice, v.age)
	}
	return slice

}

func sumAge(u []user, p ...func([]user) interface{}) int {
	var ages []int
	var sum int
	for _, f := range p {
		result := f(u)
		switch result.(type) {
		case []user:
			u = result.([]user)
			fmt.Println("--u---", u)
		case []int:
			ages = result.([]int)
			fmt.Println("--ages---", ages)
		}
	}
	if len(ages) == 0 {
		log.Fatalln("没有在管道中加入 mapAgeToSlice 方法")
	}
	for _, age := range ages {
		sum += age
	}
	return sum
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("--------面向过程思维---------")
	var users = []map[string]string{
		{
			"name":  "111",
			"age":   "16",
			"score": "70",
		},
		{
			"name":  "222",
			"age":   "18",
			"score": "25",
		},
		{
			"name":  "333",
			"age":   "26",
			"score": "95",
		},
	}
	fmt.Println("ageSum", ageSum(users), len(users))
	fmt.Println("scoreSum", scoreSum(users), len(users))
	fmt.Println("--------面向---------")
	validUsers := itemsFilter(users, func(m map[string]string) bool {
		age, ok := m["age"]
		if !ok {
			return false
		}
		intAge, _ := strconv.Atoi(age)
		if intAge <= 17 {
			return false
		}

		return true
	})
	ageSlice := mapToString(validUsers, func(m map[string]string) string {
		return m["age"]
	})
	sum := fieldSum(ageSlice, func(s string) int {
		intAge, _ := strconv.Atoi(s)
		return intAge
	})
	fmt.Printf("用户年龄累加结果: %d\n", sum)
	scoreSlice := mapToString(users, func(m map[string]string) string {
		return m["score"]
	})
	scoreSum := fieldSum(scoreSlice, func(s string) int {
		intAge, _ := strconv.Atoi(s)
		return intAge
	})
	fmt.Printf("用户分数累加结果: %d\n", scoreSum)
}

func ageSum(users []map[string]string) int {
	var sum int
	for _, v := range users {
		num, _ := strconv.Atoi(v["age"])
		sum += num
	}
	return sum
}

func scoreSum(users []map[string]string) int {
	var sum int
	for _, v := range users {
		num, _ := strconv.Atoi(v["score"])
		sum += num
	}
	return sum
}

func mapToString(items []map[string]string, f func(map[string]string) string) []string {
	newSlice := make([]string, len(items))
	for _, v := range items {
		newSlice = append(newSlice, f(v))
	}
	return newSlice
}

func fieldSum(items []string, f func(string) int) int {
	var sum int
	for _, v := range items {
		sum += f(v)
	}
	return sum
}

func itemsFilter(items []map[string]string, f func(map[string]string) bool) []map[string]string {
	newSlice := make([]map[string]string, len(items))
	for _, v := range items {
		if f(v) {
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

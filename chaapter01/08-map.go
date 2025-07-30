package main

import (
	"fmt"
	"sort"
)

func main() {
	testMap := map[string]int{"one": 1, "two": 2, "three": 3}
	testMap["four"] = 4 // 添加新元素
	fmt.Println(testMap)
	k := "two"
	v, ok := testMap[k]
	if ok {
		fmt.Printf("The element of key %q: %d\n", k, v)
	} else {
		fmt.Println("Not found!")
	}

	// var v1 map[string]int
	// v1["one"] = 1 // 运行时错误: assignment to entry in nil map
	// fmt.Println(v1)
	// 使用 make 初始化 map,第二个参数初始存储容量（超出会自动扩容）
	var v2 = make(map[string]int, 100)
	v2["one"] = 1
	fmt.Println("v2", v2)
	fmt.Println("-------------------")
	// 删除元素
	v3 := map[string]int{"one": 1, "two": 2, "three": 3}
	delete(v3, "two")
	fmt.Println("v3", v3)
	fmt.Println("-------------------")
	// 遍历 map
	for k, v := range v3 {
		fmt.Printf("k: %s, v: %d\n", k, v)
	}
	// 键值对调
	invMap := make(map[int]string, 3)
	for k, v := range v3 {
		invMap[v] = k // 键值对调
	}
	fmt.Println("invMap", invMap)
	// 排序
	v4 := map[string]int{"one": 1, "three": 3, "two": 2}
	keys := make([]string, 0, len(v4))
	values := make([]int, 0, len(v4))
	invMap = make(map[int]string, 3)
	for k, v := range v4 {
		keys = append(keys, k)
		values = append(values, v)
		invMap[v] = k
	}
	fmt.Println("keys before sort:", keys)
	// 使用 sort 包对键进行排序
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println("sort-s-k:", k, "sort-s-v:", v4[k])
	}
	sort.Ints(values)
	for _, v := range values {
		fmt.Println("sort-i-v:", v, "sort-i-k:", invMap[v])
	}
	fmt.Println("-------------------")
}

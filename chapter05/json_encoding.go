package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Id     int      `json:"id"`
	Name   string   `json:"name"`
	Age    int      `json:"age"`
	Wallet float64  `json:"wallet"`
	Skills []string `json:"skills,omitempty"` // 可选字段，omitempty表示如果为空则不输出
}

func main() {
	user := &User{
		Id:     1,
		Name:   "Alice",
		Age:    30,
		Wallet: 66.7777,
		Skills: []string{"Golang", "PHP", "C", "Java", "Python"},
	}
	fmt.Println("user", user)
	u, err := json.Marshal(user)
	if err != nil {
		fmt.Println("json.Marshal error:", err)
		return
	}
	fmt.Println("json.Marshal result:", string(u))
	var user2 User
	err = json.Unmarshal(u, &user2)
	if err != nil {
		fmt.Println("json.Unmarshal error:", err)
		return
	}
	fmt.Println("json.Unmarshal result:", user2)
	fmt.Println("----------1----------")
	u3 := []byte(`{"name": "shu-1", "website": "https://www.baidu.com", "age": 18, "male": true, "skills": ["Golang", "PHP"]}`)
	var user4 interface{}
	err = json.Unmarshal(u3, &user4)
	if err != nil {
		fmt.Printf("JSON 解码失败：%v\n", err)
		return
	}
	fmt.Println("JSON 解码结果: ", user4)
	user5, ok := user4.(map[string]interface{})
	if ok {
		for k, v := range user5 {
			switch v2 := v.(type) {
			case string:
				fmt.Println(k, "is string", v2)
			case int:
				fmt.Println(k, "is int", v2)
			case bool:
				fmt.Println(k, "is bool", v2)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, iv := range v2 {
					fmt.Println(i, iv)
				}
			default:
				fmt.Println(k, "is another type not handle yet")
			}
		}
	}
	fmt.Println("----------2流式读写----------")
	// 流式读写 JSON 数据
	dec := json.NewDecoder(os.Stdin)
	// 流式读取 JSON 数据并写入到标准输出
	enc := json.NewEncoder(os.Stdout)
	for {
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			fmt.Println(err)
			return
		}
		if err := enc.Encode(&v); err != nil {
			fmt.Println(err)
		}
	}
}

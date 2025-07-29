package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// 无符号整型，无符号数字不包含负数，所以上述转化编译时会报溢出错误
	// v1 := uint(-255)
	// fmt.Println(v1)
	v1 := uint(255)
	v2 := int8(v1)  // 转化为有符号整型
	fmt.Println(v2) // 超出 int8 表示范围且有符号的整型，所以转化结果是 -1
	v3 := 99.99
	v4 := int(v3)
	fmt.Println(v4)   // 浮点数转化为整型时会截断小数部分，所以结果是 99
	v5 := float64(v4) // 整型转化为浮点型
	fmt.Println(v5)   // 整型转化为浮点型，结果是 99
	v6 := 65
	fmt.Println("----------------------")
	fmt.Println(string(v6)) // 整型转化为字符串，结果是字符 'A'，因为 65 是 ASCII 码对应的字符
	v7 := 30028
	fmt.Println(string(v7)) // 整型转化为字符串，结果是字符 '界'
	fmt.Println("----------------------")
	v8 := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(v8)         // 字符串转化为字节切片，结果是 [104 101 108 108 111]
	fmt.Println(string(v8)) // 字节切片转化为字符串，结果是 "hello"
	fmt.Println("----------------------")
	a1 := "100"
	a2, _ := strconv.Atoi(a1)           // 将字符串转化为整型，v2 = 100
	fmt.Println(a2, reflect.TypeOf(a2)) // 输出 100
	a3 := strconv.Itoa(a2)              // 将整型转化为字符串，v3 = "100"
	fmt.Println(a3, reflect.TypeOf(a3)) // 输出 "100"
	a4 := "TRUE"
	a5, _ := strconv.ParseBool(a4)      // 将字符串转化为布尔型，v5 = true
	fmt.Println(a5, reflect.TypeOf(a5)) // 输出 true
	a6 := strconv.FormatBool(a5)        // 将布尔型转化为字符串，v6 = "true"
	fmt.Println(a6, reflect.TypeOf(a6)) // 输出 "true"
	a7 := "100"
	a8, _ := strconv.ParseInt(a7, 10, 64) // 将字符串转化为整型，第二个参数表示进制，第三个参数表示最大位数
	fmt.Println(a8, reflect.TypeOf(a8))   // 输出 100
	a7 = strconv.FormatInt(a8, 10)        // 将整型转化为字符串，第二个参数表示进制
	fmt.Println(a7, reflect.TypeOf(a7))   // 输出 "100"
	a9, _ := strconv.ParseUint(a7, 10, 64)
	fmt.Println(a9, reflect.TypeOf(a9)) // 输出 "100"，转换为无符号整型
	a7 = strconv.FormatUint(a9, 10)     // 将无符号整数型转化为字符串，参数含义同 FormatInt
	fmt.Println(a7, reflect.TypeOf(a7)) // 输出 "100"
	a10 := "3.1415926"
	a11, _ := strconv.ParseFloat(a10, 64)        // 将字符串转化为浮点型，第二个参数表示精度
	fmt.Println(a11, reflect.TypeOf(a11))        // 输出 3.1415926
	a12 := strconv.FormatFloat(a11, 'E', -1, 64) // 将浮点型转化为字符串，第二个参数表示格式，第三个参数表示小数点后位数，第四个参数表示精度
	fmt.Println(a12, reflect.TypeOf(a12))        // 输出 "3.141593"

	q := strconv.Quote("hello, 你好")         // 将字符串转化为带引号的字符串
	fmt.Println(q, reflect.TypeOf(q))       // 输出 "\"hello，你好\""
	q2 := strconv.QuoteToASCII("hello, 你好") // 将字符串转化为 ASCII 编码的带引号字符串
	fmt.Println(q2, reflect.TypeOf(q2))     // 输出 "\"hello，你好\""
}

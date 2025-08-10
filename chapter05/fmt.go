package main

import (
	"fmt"
	"os"
)

func main() {
	/**
		生产环境：使用 log 包或专门的日志库记录错误
		开发调试：使用 %+v 格式打印完整错误链
		错误包装：使用 %w 包装原始错误保留上下文
		错误检查：使用 errors.Is 和 errors.As 进行错误类型判断
		性能考虑：频繁的错误处理避免过多的字符串拼接
	**/
	fmt.Println("------------------- Fprintf 往文件中写入内容 ------------------")
	name := "枯藤"
	// 向标准输出写入内容
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("chapter05/xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)

	fmt.Println("------------------- Sprint 把传入的数据生成并返回一个字符串 ------------------")
	s1 := fmt.Sprint("枯藤1")
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("枯藤3")
	fmt.Println(s1, s2, s3)
	fmt.Println("------------------- Errorf 根据format参数生成格式化字符串并返回一个包含该字符串的错误 ------------------")
	err = fmt.Errorf("文件读取失败: %s", "权限不足")
	fmt.Println(err)
	fmt.Println(err.Error())
	err = fmt.Errorf("处理失败: %w", fmt.Errorf("内部错误: %s", "空指针引用"))
	fmt.Println(err)
	fmt.Printf("%v\n", err)
	fmt.Printf("%+v\n", err)
	fmt.Println("------------------- 从标准输入获取用户的输入 ------------------")
	var (
		bb1 string
		bb2 int
		bb3 bool
	)
	fmt.Scan(&bb1, &bb2, &bb3)
	fmt.Printf("扫描结果 bb1:%s bb2:%d bb3:%t \n", bb1, bb2, bb3)
	fmt.Println("------------------- bufio.NewReader 完整获取输入的内容,包含空格------------------")
	fmt.Println("------------------- Fscan系列 从io.Reader中读取数据------------------")
	fmt.Println("------------------- Sscan系列从指定字符串中读取数据------------------")
	fmt.Println("------------------- 格式化占位符 ------------------")
	/**
		通用占位符
			%v: 值的默认格式表示
			%+v: 类似%v，但输出结构体时会添加字段名
			%#v: Go 语法格式化
			%T: 打印值的类型
			%%: 百分号
		布尔型
			%t: 布尔值,true或false
		整型
			%b: 二进制整数
			%c: 字符，rune类型
			%d: 十进制整数
			%o: 八进制整数
			%x: 十六进制整数（小写字母）
			%X: 十六进制整数（大写字母）
			%U: Unicode格式
			%q: 字符的Go语法格式
		浮点型和复数
			%b: 十进制浮点数，指数部分以二进制表示
			%e: 科学计数法（小写e）
			%E: 科学计数法（大写E）
			%f: 十进制浮点数
			%F: 等同于%f
			%g: 根据需要使用%e或%f
			%G: 根据需要使用%E或%f
		字符串和字节切片
			%s: 直接输出字符串或者[]byte
			%q: 双引号括起来的字符串，特殊字符会转义,该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
			%x: 每个字节用两字符十六进制数表示（使用a-f
			%X: 每个字节用两字符十六进制数表示（使用A-F）
		指针
			%p: 指针地址，十六进制格式.表示为十六进制，并加上前导的0x
		时间
			%v: 默认格式
			%+v: 带有详细时间信息
			%#v: Go 语法格式化时间
		宽度标识符
			%f: 默认宽度，默认精度
			%9f: 宽度9，默认精度
			%.2f: 默认宽度，精度2
			%9.2f: 宽度9，精度2
			%-9.2f: 左对齐，宽度9，精度2
			%09.2f: 零填充，宽度9，精度2
			%9s: 宽度9，字符串右对齐
			%-9s: 宽度9，字符串左对齐
			%9d: 宽度9，整数右对齐
			%-9d: 宽度9，整数左对齐
			%09d: 宽度9，整数右对齐，前面填充0
	**/
}

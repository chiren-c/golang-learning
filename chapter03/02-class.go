package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Male  bool
	Score float64
}

func NewStudent(name string, age int, male bool, score float64) *Student {
	return &Student{
		Name:  name,
		Age:   age,
		Male:  male,
		Score: score,
	}
}

func NewStudentV2(name string, age int, score float64) Student {
	return Student{Age: age, Name: name, Score: score}
}

func main() {
	student := NewStudent("Alice", 20, true, 95.5)
	fmt.Println("student", student)
	fmt.Println("Name:", student.GetName()) // 指针调用值方法自动解引用: (*s).GetName()
	student2 := NewStudentV2("Alice", 20, 95.5)
	student2.SetName("Bob") // 是可寻址的左值，所以实际调用: (&student2).SetName("Bob")
	fmt.Println("Name-Bob:", student2.GetName())
	// NewStudentV2("Alice2", 20, 95.5).SetName("Bob") // err 值类型调用指针方法
	/**
	所谓左值就是可以出现在赋值等号左边的值，
	而右值只能出现在赋值等号右边，比如函数返回值、字面量、常量值等。左值可寻址，右值不可寻址。
	**/
}

func (s Student) GetName() string {
	return s.Name
}

func (s *Student) SetName(name string) {
	s.Name = name
}

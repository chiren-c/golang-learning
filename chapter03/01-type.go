package main

type Integer int

func (i Integer) Equal(j Integer) bool {
	return i == j
}

func (i Integer) Add(j Integer) Integer {
	return i + j
}

type Math interface {
	Equal(Integer) bool
	Add(Integer) Integer
}

func main() {
	var i Integer = 42
	println("i", i)              // Output: 42
	println("Equal", i.Equal(5)) // Output: false
	println("Add", i.Add(5))     // Output: 47
	m := &i
	println("m", m.Add(1)) // Output: 43
}

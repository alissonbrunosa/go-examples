package main

import "stack"
import "fmt"

type Person struct {
	Name string
}

func main() {
	s := stack.New()

	fmt.Println(s.IsEmpty())

	s.Push(&Person{"James"})
	s.Push(&Person{Name: "Doe"})

	s.Print()
	fmt.Printf("Stack's length is %d\n", s.Length())
}

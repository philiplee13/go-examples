package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"John Doe", 30})
	p := createperson("Bob", 25)
	fmt.Println("person is", p)
	fmt.Println("person name and age is", p.name, p.age)
	pp := &p
	fmt.Println("pointer to person is", pp)
	pp.name = "changedNamed from Bob to John"
	pp.age = 99
	fmt.Println("changed values are through the ptrs are", pp.name, pp.age)
	fmt.Println("checking the values using the original variable", p.name, p.age)

	fmt.Println("getting name of person", p.getname())
	fmt.Println("getting age of person", p.getage())
}

func createperson(name string, age int) person {
	return person{name, age}
}

func (person *person) getage() int {
	return person.age
}

func (person *person) getname() string {
	return person.name
}

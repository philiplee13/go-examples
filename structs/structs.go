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
	fmt.Println("--------------interfaces------------------")
	o := owner{"john"}
	d := dog{"dog", 2, o}
	c := cat{"cat", 3}
	makesound(d)
	fmt.Println("the dogs owner is", d.getowner())
	makesound(c)
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

// interfaces
type animal interface {
	sound() string
}

type owner struct {
	name string
}

func (o owner) getowner() string {
	return o.name
}

type dog struct {
	name string
	age  int
	owner
}

type cat struct {
	name string
	age  int
}

func (d dog) sound() string {
	return "bark"
}

func (c cat) sound() string {
	return "meow"
}

func makesound(a animal) {
	fmt.Println("animal sound for", a, "is", a.sound())
}

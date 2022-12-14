package main

import (
	"fmt"
	"strings"
)

type Person struct {
	id   string
	name string
	age  int
	sex  string
}

type Employee struct {
	title string
}

func (p Person) disp(job string) {
	fmt.Println(p.id)
	fmt.Println(p.name)
	fmt.Println(p.age)
	fmt.Println(p.sex)
	fmt.Println(p.name + " is employed as " + job + " in the " + p.sex + " gender section!")
	//return "---------"
}

func (p *Person) update(n string, a int, s string) {
	p.name = n
	p.age = a
	p.sex = s
}

func (e Employee) toLower() string {
	return strings.ToLower(e.title)
}

func main() {
	fmt.Println("Hello World")
	p := Person{
		id:   "23989",
		name: "Robert",
		age:  55,
		sex:  "male",
	}
	e := Employee{
		title: "Park land ranger",
	}

	p.disp((e.toLower()))
	p.update("Catalina", 18, "female")
	p.disp((e.toLower()))

}

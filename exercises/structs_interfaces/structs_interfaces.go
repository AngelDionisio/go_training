package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

// any function that implements the speak method,
// is implicity implementing this interface, making it also of this type
type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println(p.first, p.last, p.age)
}

func humanSpeak(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   38,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   29,
	}

	p1.speak()
	p2.speak()
	human(&p1).speak()
	humanSpeak(&p2)
}

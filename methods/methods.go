package main

import (
	"fmt"
)

// methods
// func (r receiver) identifier(parameters) (return(s)) { ... }
type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// when you have a receiver, it's going to attach this function so the type
// so any value of that type will have access to the function via the "dot" notation
func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, ", license to kill?:", s.ltk, " - the secretAgent speak")
}

// type person, now has the method "speak" attached to it. Making it available via the "dot" notation
func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

// interface is a type
// to immplement an interface in Go,
// we just need to implement all the methods in the interface
// a value can be of more than one type
// if any other type implements the speak function, one can say its is also
// of that type. So sa1 is of type secretAgent and human.
// interface can implement one or none methods. Every type is also of type interface
// this is what allows one to have varidatic params by taking ...interfacte{} as a type
type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println(
			"person type, passed into barrrrrr",
			h.(person).first,
			h.(person).last,
		)
	case secretAgent:
		fmt.Println(
			"secretAgent type passed into barrrrrr",
			h.(secretAgent).first,
			h.(secretAgent).last,
			h.(secretAgent).ltk,
		)
	}
	fmt.Println("I was passed into bar", h)
}

func methods() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Miss",
			last:  "Moneypenny",
		},
		ltk: false,
	}

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
	fmt.Printf("sa1 is of type %T\n", sa1)
	fmt.Println(p1)

	// bar is taking two types, superAgent and person
	// this is because interfaces allow values to be of more
	// than one type. Both implent a speak method
	bar(sa1)
	bar(sa2)
	bar(p1)
}

// because every type is also of type interface, this function
// takes an unlimited amount of parameters of any type
func variadic(x ...interface{}) {
	fmt.Println(x...)
}

func main() {
	methods()

	variadic("test", true, 2, 3.5, []string{"foo", "bar"})

	p := newPlayer(655454, "Angel", "NYC", 352345)

	fmt.Println(p.Greetings())
}

// example using methods implementing the interface to access the function Greetings
type User struct {
	Id             int
	Name, Location string
}

func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s", u.Name, u.Location)
}

func newPlayer(id int, name string, location string, gameid int) *Player {
	return &Player{
		User: &User{
			Id:       id,
			Name:     name,
			Location: location,
		},
		GameId: 32454,
	}
}

type Player struct {
	*User
	GameId int
}

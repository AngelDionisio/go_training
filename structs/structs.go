package main

import(
	"fmt"
)

type person struct {
	first 	string
	last	string
	age 	int
}

// embedded struct
type secretAgent struct {
	person
	license_to_kill bool
}

func embeddedStructs() {
	// when initializing an embedded type, the key is the type name
	//  initializing as a composite literal
	secretAgentOne := secretAgent{
		person: person{
			first: "james",
			last: "bond",
			age: 32,
		},
		license_to_kill: true,
	}

	fmt.Println(secretAgentOne)
	// anonymous embedded types (e.g. person is an embedeed type in secretAgent)
	// its values get promoted, so you can access the keys from the embedded value
	// directly from variables of said type, that is, se.first, se.last
	// and not having to do se.person.first, se.person.last
	fmt.Println(
		secretAgentOne.person.first, // works but no need to do this
		secretAgentOne.last,
		secretAgentOne.age,
		secretAgentOne.license_to_kill,
	)

	// anonymous struct: instead of using an identifier (person / secret agent)
	// we can replace the identifier with what it represents. e.g:
	// so instead of using type person, we replace with the struct definition
	// type person struct {
	// 	first 	string
	// 	last	string
	// 	age 	int
	// }

	person2 := struct {
		first 	string
		last	string
		age 	int
	} {
		first:	"notJames",
		last:	"notBond",
		age:	62,
	}

	fmt.Println(person2)

}

func structs() {
	p1 := person{
		first: "james",
		last: "bond",
		age: 31,
	}

	fmt.Println(p1)
	fmt.Println(p1.first, p1.last)

}

func houseKeeping() {
	// we create VALUES of a certain TYPE which are stored in VARIABLES
	// those VARIABLES have IDENTIFIERS
	// notice the similarirites

	// var x of type int
	var x int

	// var p of type struct as defined by the composite type
	type p struct {
		first	string
		last	string
	}

	// constant of a kind, when we do not let the compiler know which type
	// we declare a variable, the complier will try to figure it out itself
	// constants can only hold primitive types
	const y = 42

	fmt.Println(x)
	fmt.Printf("%v\t%T\n", y, y)

	// Named types and anonymous types
	// Anonymous types are indeterminate. They have not been declared as a type yet.
	// The compiler has flexibility with anonymous types. You can assign an anonymous type
	// to a variable declared as a certain type. If the assignment can occur, the compiler
	// will figure it out; the compiler will do an implicit conversion. You cannot assign
	// a named type to a different named type.

	// Padding an architectural alignment
	// Convention: logically organize your fields together. Readability and clarity trump performance
	// as a design concern. Go will be performant. Go for readability first. However, if you are in
	// a situation where you need to optimize for performance: lay your fields from largest to smallest,
	// e.g. int64, int32, float32, bool.
}

func main() {
	embeddedStructs()
	structs()
}
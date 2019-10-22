package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println(a)  // prints value
	fmt.Println(&a) // prints location in memory

	fmt.Printf("%T\n", a) // int
	fmt.Printf("%T\n", &a) // *int -> pointer to an int e.g. 0xc000066058

	// b is a pointer to the address of a, referencing the location of a
	b := &a
	fmt.Println(b) // address: 0xc000066058
	fmt.Println(*b) // dereferencing the pointer, * gives you the value stored at that address
	fmt.Println(*&a)

	// you can set the value at an address directly
	// because b is a pointer to the value of a
	// *b gets the value of b, so here we set the value of that address
	*b = 43
	fmt.Println(a)

	// when to use: when you do not want to pass around large amounts of data,
	// you can just pass the location around. And to change the value stored at an address

	// a pointer reciever e.g. func (c circle) area() { ... }
	// in this example (c circle).
	// a NON-POINTER receiver works with values that ar e POINTERS and NON-POINTERS
	// a POINTER RECEIVER only works with  values that are pointers

	name := "angel"
	fmt.Println("address of var name: ", &name)
	
}

func pointerReciver(y *int) {
	// changes the value of the memory location of the address y is pointing to
	*y = 43
}
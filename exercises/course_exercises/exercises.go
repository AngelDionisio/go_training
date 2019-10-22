package main

import (
	"fmt"
	"math"
)

// https://play.golang.org/p/u3ZcZuNdwe0

var x = []int{1,2,3,4,5}

func main() {
	// mapExercise()
	// exSliceOfSlice()
	// handsOnExerciseLvl5e1()
	handsOnExerciseLvl5e3()
	anonymousStruct()

	fmt.Println(
		addIntArguments(x...),
		addItemsInSliceOfInt(x),
	)

	deferEx()

	attachingMethodsToTypes()

	// hands on exercise for receivers and intefaces
	// ************************************************
	c := circle{
		radius: 12.345,
	}

	s := square{
		length: 15,
	}

	info(c)
	info(s)

	// ************************************************

	// assigning a function to a variable
	var add = func(x int, y int) int {
		return x + y
	}

	fmt.Println("calling function assigned to var:", add(2,3))

	var addTen = addCurried(10)
	fmt.Println(addTen(10))

	fmt.Println(addVariadic(1,2,3,4,5,6,7,8,9))

	nums := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(filterEven(nums))

	fmt.Println(
		"adding even nums",
		firstClassAddEvenNums(filterEven, addItemsInSliceOfInt, nums),
	)

	var p123 = person{
		first: "DEFAULT",
		last: "ORIGINAL",
		age: 45,
	}
	fmt.Println("BEFORE pointer change", p123)
	changeMe(&p123)
	fmt.Println("AFTER pointer change", p123)
	changeMe2(p123)
	fmt.Println("AFTER 2 pointer change", p123)

}

// create a person struct, create a func called changeMe, with a *person as a parameter
// change the value stored at the *person address
func changeMe(p *person) {
	(*p).first = "UPDATED"
	// p.first = "UPDATED"
}

// this will NOT change the value of the original value passed, as in GO, everything is passed by value
func changeMe2(p person) {
	p.last = "REGULAR"
	// p.first = "UPDATED"
}


func addCurried(x int) func(y int) int {
	return func(y int) int {
		return x + y
	}
}

func addVariadic(x ...int) int {
	var sum int
	for _, v := range x {
		sum += v
	}

	return sum
}

func filterEven(sliceOfInts []int) ([]int) {
	var evenNums []int
	for _, v := range sliceOfInts {
		if v % 2 == 0 {
			evenNums = append(evenNums, v)
		}
	}
	return evenNums
}

// a callback is any function that is called by another function
func firstClassAddEvenNums(
	filterEven func(nums []int) []int,
	addNums func(nums []int) int,
	input []int) int {
	even := filterEven(input)
	fmt.Println("filtered even numbers:", even)
	sum := addNums(even)

	return sum
}

// *********************************************************************************
// *********************************************************************************
type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius *  c.radius // pie r square
}

func (s square) area() float64 {
	return s.length * s.length
}

// any type that implements the area method is now also of type shape
// so both circle and square implicitly implement the area interface
// with this in mind, we can now create a generic function that can call the area
// of any type that implements the shape interface
type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

// *********************************************************************************
// *********************************************************************************

type person struct {
	first string
	last string
	age int
}

// method with a receiver of type person, making function available to any
// instantiation of type person
func (p person) speak() {
	fmt.Printf("hello I am %v, %v. And I am speaking from a method of type person\n", p.last, p.first)
}

// create a user defined struct with identifier person, has fields first, last, age
// attach a method to type person with identifier speak
// the method should print the person first and last name
// create a value of type person, call the method from value of type person
func attachingMethodsToTypes() {
	p1 := person{
		first: "james",
		last: "bond",
		age: 42,
	}

	p1.speak()
}

// *********************************************************************************
// *********************************************************************************


func deferEx() {
	fmt.Println("I am called first")
	// defer functions get executed right before the surrounding function exits
	defer fmt.Println("I am called second, but I am deferred")
	fmt.Println("I am called third")
}

func anonymousStruct() {
	s := struct {
		name string
		last string
		friends map[string]int
	} {
		name: "james",
		last: "bond",
		friends: map[string]int{
			"Moneypenny":	555,
			"Q": 			777,
			"M": 			888,
		},
	}

	fmt.Println(s)
}

func addIntArguments(nums ...int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	return sum
}

func addItemsInSliceOfInt(nums []int) int {
	var sum int
	for _, v  := range nums {
		sum += v
	}
	return sum
}

func handsOnExerciseLvl5e3() {
	type vehicle struct {
		doors	int
		color	string
	}

	type racing struct {
		isRacing bool
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		racing
		luxury bool
	}

	t := truck{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		racing: racing{
			isRacing: true,
		},
		luxury: true,
	}

	fmt.Println(t)
	fmt.Println(s)
	// showing inner type promotion
	// no need to do t.vehicle.doors
	// even though its valid
	fmt.Println(t.doors)
	fmt.Println(s.doors)
	fmt.Println("doubly nested composition isRacing:", s.isRacing)
}

// create your own type person, create two values of type person
// print out the values ranging over the elements of a slice.

func handsOnExerciseLvl5e1() {
	type person struct {
		firstName			string
		lastName			string
		favIceCreamFlavors	[]string
	}

	p1 := person{
		firstName: "James",
		lastName: "Bond",
		favIceCreamFlavors: []string{
			"vanilla",
			"chocolate",
			"rum raising",
		},
	}

	p2 := person{
		firstName: "agent",
		lastName: "47",
		favIceCreamFlavors: []string{
			"dark chocolate",
			"chocolate oreo",
		},
	}

	sliceOfPersons := []person{p1, p2}

	for _, entity := range sliceOfPersons {
		fmt.Printf("user: %v, %v:\n", entity.firstName, entity.lastName)
		for i, flavor := range entity.favIceCreamFlavors {
			fmt.Printf("\t%v\t%v\n", i, flavor)
		}
	}

	// take the code above and store the values of type person in a map
	// with the key of last name. Access each value in the map. Print out the
	// values, ranging over the slice.
	myMap := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	v, ok := myMap["Bond"]
	v2, ok2 := myMap["47"]
	fmt.Println(v, ok)
	fmt.Println(v2, ok2)

	for _, v := range myMap {
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for i, val := range v.favIceCreamFlavors {
			fmt.Println(i, val)
		}
		fmt.Println("------------")
	}
	// fmt.Println(myMap["Bond"])
	// fmt.Println(myMap["47"])
}

// create a slice of a slice of string
// store  the following data
// "Jame", "Bond", "Shaken not stireed."
// "Miss", "Moneypenny", "Hellooooo James."
// range over the records, then range over the data in each record
func exSliceOfSlice() {
	xs1 := []string{"James", "Bond", "Shaken not stireed."}
	xs2 := []string{"Miss", "Moneypenny", "Hellooooo James."}

	sliceOfSliceOfString := [][]string{xs1, xs2}

	for i, xs := range sliceOfSliceOfString {
		fmt.Println("record", i)
		for j, valueInSlice := range xs {
			fmt.Printf("\t index position: %v \t value: %v \n", j, valueInSlice)
		}
	}

}

func mapExercise() {
	users := map[string]string{
		"angel": "UFC, Soccer, Luz",
		"Luz": "graphic design, iPhones, Angel",
	}

	// add record to map
	users["zeus"] = "foo"

	// delete record from map
	delete(users, "zeus")
	
	for k, v := range users {
		fmt.Println(k, v)
	}
}

func appendingToSlices() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x, 52)
	fmt.Println(y)

	y = append(y, 53, 54, 55)
	fmt.Println(y)

	w := append(x, y...)
	fmt.Println(w)

	x = append(x[:3], x[6:]...) // [42, 43, 44, 48, 49, 50, 51]
	fmt.Println(x)
}

func ninja_level_4_ex_1() {
	/*
	using a composite literal, create an ARRAY which holds 5 values of type in,
	assing values to each index position
	*/
	a := [5]int{0, 1, 2, 3, 4} // specifiying size, creates an array
	for i := 0; i < len(a)-1; i++ {
		fmt.Println(a[i])
	}

	fmt.Printf("%T\n", a)
}

func ninja_level_4_ex_2() {
	/*
	using a composite literal, create an SLICE which holds 5 values of type in,
	assing values to each index position
	*/
	a := []int{0, 1, 2, 3, 4} // not specifying a size will create a SLICE
	for i := 0; i < len(a)-1; i++ {
		fmt.Println(a[i])
	}

	fmt.Printf("%T\n", a)
}

func usStates() []string {
	states := []string{
		"Alabama", 
		"Alaska", 
		"Arizona",
		"Arkansas",
		"California",
		"Colorado",
		"Connecticut", 
		"Delaware",
		"Florida",
		"Georgia",
		"Hawaii",
		"Idaho",
		"Illinois",
		"Indiana", 
		"Iowa",
		"Kansas",
		"Kentucky",
		"Louisiana",
		"Maine",
		"Maryland", 
		"Massachusetts", 
		"Michigan",
		"Minnesota",
		"Mississippi",
		"Missouri",
		"Montana",
		"Nebraska", 
		"Nevada",
		"New Hampshire", 
		"New Jersey",
		"New Mexico",
		"New York",
		"North Carolina", 
		"North Dakota",
		"Ohio",
		"Oklahoma", 
		"Oregon",
		"Pennsylvania",
		"Rhode Island",
		"South Carolina",
		"South Dakota",
		"Tennessee",
		"Texas",
		"Utah",
		"Vermont", 
		"Virginia",
		"Washington",
		"West Virginia", 
		"Wisconsin",
		"Wyoming",
		}
	return states
}
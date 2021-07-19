package main

import (
	"fmt"
	"time"
)

type names []string
type day struct {
	name  string
	index int
}

func (n names) printNames() {
	fmt.Println("Inside printNames method of names type")
	for _, name := range n {
		fmt.Println(name)
	}
	fmt.Println("Exiting printNames method of names type")
}

func (d *day) changeDay() {
	fmt.Println("Changing day to Sunday")
	d.name = "Sunday"
	d.index = 0
}

func main() {
	// Go doesn't have classes and objects but you can define
	// methods on predefined types

	// Accessing methods of in built type time.Duration
	const hrsInDay = 24 * time.Hour
	fmt.Printf("hrsInDay is of type %T and value %v\n", hrsInDay, hrsInDay)
	seconds := hrsInDay.Seconds()
	fmt.Printf("seconds in a day %v\n", seconds)

	// Creating method for named types
	namesType := names{"Joey", "Chandler"}
	namesType.printNames()

	// Methods for pointer types
	d := new(day) // creates a pointer to struct
	d.name = "Monday"
	d.index = 1
	fmt.Println("Before:", *d)
	d.changeDay()
	fmt.Println("After:", *d)
	dd := day{name: "Tuesday", index: 2}
	fmt.Println("Before:", dd)
	dd.changeDay() // go does the converstion of day to *day
	fmt.Println("After:", dd)

	// Method declarations are not permitted on types that are pointers
	// The following code is invalid
	/*
		type intPtrs *int

		func (i intPtrs) Print() {
			fmt.Println(*i)
		}

		i := 5
		ii := intPtrs(&i)
		ii.Print()
	*/

}

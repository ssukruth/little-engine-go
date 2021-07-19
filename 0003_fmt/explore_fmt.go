package main

import "fmt"

func main() {
	// fmt.Println can print a line
	fmt.Println("This is a line")
	// fmt.Println can append values to a line
	name := "fmt"
	fmt.Println("This is a line is printed using the package:", name)
	// fmt.Println can append results of expressions to a line
	fmt.Println("fmt.Println caclulated sum of 3 & 4 as:", 3+4)

	// fmt.Printf is similar to C's printf and is used to print
	// formatted strings. Format specifiers are as follows:-
	// %d int
	// %s string
	// %t bool
	// %f float64 or float32
	// %.3f 3 decimal points
	// %q quoted strings
	// %v print any value
	// %T type of variable
	// %b binary
	// %08b print the number in binary with 8 bits
	// %o oct
	// %O oct with 0a prefix
	// %x hex with lower case a-f
	// %X hex with upper case A-F
	// %p is for pointer
	var (
		integer   = 25
		float     = 3.14
		boolean   = true
		stringvar = "abc"
	)
	fmt.Printf("'integer' type %T value %d\n", integer, integer)
	fmt.Printf("'float' type %T value %f\n", float, float)
	fmt.Printf("'boolean' type %T value %t\n", boolean, boolean)
	fmt.Printf("'stringvar' type %T value %s\n", stringvar, stringvar)
	fmt.Printf("'stringvar' type %T value %q\n", stringvar, stringvar)
	fmt.Printf("Binary value of %d is %b\n", integer, integer)
	fmt.Printf("Binary value of %d formatted to 8 bits is %08b\n", integer, integer)
	fmt.Printf("Hex value of %d is %o. It can also be written as %O\n", integer, integer, integer)
	fmt.Printf("Hex value of %d is %x. It can also be written as %X\n", integer, integer, integer)
	fmt.Printf("'integer' %d is stored at %p\n", integer, &integer)

}

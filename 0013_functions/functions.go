package main

import (
	"fmt"
	"math"
)

func f1() {
	fmt.Println("This is f1 function")
}

func f2(m int, n int) {
	m, n = n, m
	fmt.Println("Within f2, m:", m, "and n:", n)
}

func f3(a, b, c int, d, e, f float64, g, h string) {
	fmt.Println(a, b, c, d, e, f, g, h)
}

func f4(num float64) float64 {
	return math.Pow(num, 3)
}

func f5(a, b int) (int, int) {
	return a + b, a - b
}

func f6(a, b int) (q int, r int) {
	q, r = a/b, a%b
	return
}

func f7(a ...int) {
	fmt.Printf("a is of type %T and value  %#v\n", a, a)
}

func f8(a ...int) {
	a[0] = 99
}

func f9() {
	defer f10()
	fmt.Println("Inside f9 after deferring f10")
}

func f10() {
	fmt.Println("Inside f10")
}

func counter(i int) func() int {
	fmt.Println("Initial value of counter:", i)
	f := func() int {
		i++
		fmt.Println("Counter value:", i)
		return i
	}
	return f
}

func main() {
	// It is idiomatic to use camel case for function names.
	// Within a package, function names should be unique.
	// Go functions can return multiple values.
	f1()

	// Passing arguments. Everything in functions are passed
	// by value. There's no pass by refernce in go.
	// https://dave.cheney.net/2017/04/29/there-is-no-pass-by-reference-in-go
	// More on this in pointers.go
	m, n := 5, 6
	f2(m, n)
	fmt.Println("Within main, m:", m, "and n:", n)

	// Shorthand parameter notation lets you group multiple params of the
	// same type together. Refer f3 above
	f3(1, 2, 3, 1.5, 2.5, 3.5, "aa", "bb")

	// functions can return one or more values. When a function returns one
	// or more values, you must specify the return type after the params
	// and before the opening flower bracket.
	fmt.Println("Cube of 4 is:", f4(4))

	// Receving multiple values
	sum, diff := f5(4, 2)
	fmt.Println("Sum of 4 & 2 is:", sum, "and difference of 4 & 2 is:", diff)

	// Named return values, initialize the variable and return the variable
	// upon a naked return i.e. just return. Refer f6
	// This is not recommended for larger functions as it will compromise
	// readability.
	q, r := f6(10, 3)
	fmt.Println("10/3 results in quotient:", q, "and remainder:", r)

	// Variadic functions: functions which take variable number of args
	// Only the last param of a function can have variable number of args
	f7(1, 2, 3, 4, 5)
	f7()
	// Slices can be passed to variadic functions. This avoids creating a
	// new slice inside the function and as a result changes made to slice
	// in a variadic function changes the slice
	nums := []int{1, 1, 2}
	f7(nums...)
	f8(nums...)
	fmt.Println("nums:", nums)

	// Defer statements postpones the execution of a function until the
	// surrounding function exits, either gracefully or via panic
	// In the event of multiple defers, go executes them in the reverse
	// order in which they were called i.e LIFO
	f9()

	// Go provides anonymous functions i.e. functions which do not contain
	// any name. Anonymous functions can be used to form closures.
	// A closure is a persistent scope which holds on to local variables
	// even after code execution has moved out of its block.
	func(msg string) {
		fmt.Println(msg)
	}("I'm an anonymous function")
	// Anonymous functions are typically used when a function returns another
	// function which is defined inline
	c := counter(0)
	c() // i = 1
	c() // i = 2
	c() // i = 3
	c() // i = 4
}

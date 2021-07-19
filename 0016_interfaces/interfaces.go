package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	side float64
}

type rectangle struct {
	length  float64
	breadth float64
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

func (s square) perimeter() float64 {
	return 4 * s.side
}

func (s square) area() float64 {
	return math.Pow(s.side, 2)
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

func (r rectangle) area() float64 {
	return r.length * r.breadth
}

// Interface
type shapes interface {
	area() float64
	perimeter() float64
}

func printShape(sh shapes) {
	fmt.Printf("Shape: %T\n", sh)
	fmt.Println("Area:", sh.area())
	fmt.Println("Perimeter:", sh.perimeter())
}

// Embedded interface example
type drawing interface {
	shapes // all types implementing drawing must implement all methods of shapes
	hasStraightLines() bool
}

func (c circle) hasStraightLines() bool {
	return false
}

func describe(d drawing) {
	fmt.Printf("Drawing is of type %T\n", d)
	fmt.Println("Area:", d.area())
	fmt.Println("Perimeter:", d.perimeter())
	fmt.Println("Can be drawn using a scale:", d.hasStraightLines())
}

// Empty interface
type empty interface {
}

func main() {
	// Interface is a collection of method signatures that
	// an object (usually a named type) can implement. They
	// define the behavior of an object and can implement
	// polymorphism

	// We've defined 3 structs: circle, square & rectangle
	// and implented 2 methods perimeter and area for all of
	// them. All 3 of these shapes can be represented by
	// the shape interface

	c1 := circle{radius: 5}
	s1 := square{side: 5}
	r1 := rectangle{length: 5, breadth: 2}

	printShape(c1)
	printShape(s1)
	printShape(r1)

	// Variables with type as interface have nil type
	// Intefaces implement polymorphism since the variables of interface
	// type can dynamically take many values of types that implement
	// the interface during runtime.
	var s shapes
	fmt.Printf("s is of type %T\n", s)

	s = c1 // this is valid since c1's type "circle" implements shape interface
	fmt.Printf("s is of type %T\n", s)

	s = r1
	fmt.Printf("s is of type %T\n", s)

	// Even though the dynamic type of the interface variable can be set
	// to one of the concrete types, say "circle", it cannot invoke any
	// method of circle which is not defined in the interface

	s = c1
	// The following line leads to error since shape hasn't defined diameter
	// method.
	// s.diameter()

	// To access methods of concrete type which are not defined in interface,
	// we need type assertion.
	c, ok := s.(circle)
	if !ok {
		fmt.Println("Failed to get circle from shape")
	}
	fmt.Println("Diameter is:", c.diameter())

	s = s1
	// type switch
	switch s.(type) {
	case circle:
		fmt.Println("s is a circle")
	case rectangle:
		fmt.Println("s is a rectangle")
	case square:
		fmt.Println("s is a square")
	}

	// go doesn't support inheritance by extending interfaces
	// Instead, we can embedd interfaces to achieve the same.

	var d drawing
	d = c1
	describe(d)

	// Circular interface embedding is prohibited by go
	// Ex: The following is invalid
	/*
		type interface1 interface {
			method1()
			interface2
		}
		type interface2 interface {
			method2()
			interface3
		}
		type interface3 interface {
			method3()
			interface1
		}
	*/

	// Empty interfaces do not define any methods
	// Any type implements the empty interface
	var e empty
	e = 5
	fmt.Println(e)

	e = 10.5
	fmt.Println(e)

	e = []int{1, 2, 3}
	fmt.Println(e)
	// To access the type methods, use type assertion on empty interface
	v, ok := e.([]int)
	if !ok {
		fmt.Println("failed to fetch slice from emtpy interface")
	}
	fmt.Println(len(v))
}

package main

import "fmt"

type day struct {
	name  string
	index int
}

func change(a *int) {
	*a = 10
}

func getFloat() *float64 {
	pi := 3.14
	return &pi
}

func changeVal(i int, f float64, b bool, s string) {
	i = 10
	f = 3.14
	b = false
	s = "hello"
	fmt.Println("Within function:", i, f, b, s)
}

func changeValByPtr(i *int, f *float64, b *bool, s *string) {
	*i = 10
	*f = 3.14
	*b = false
	*s = "hello"
	fmt.Println("Within function:", *i, *f, *b, *s)
}

func changeDay(d day) {
	d.name = "changedDay"
	d.index = 0
}

func changeDayPtr(d *day) {
	// pointers to structs need not be explicitly dereferenced
	d.name = "changedDay" // same as (*d).name = "changedDay"
	d.index = 0           // same as (*d).index = 0
}

func changeSlice(s []int) {
	for i := range s {
		s[i]++
	}
}

func changeMap(m map[int]int) {
	m[1] = 1
}

func main() {
	// A variable is a convenient alphanumeric nickname for
	// a memory localtion. A pointer is a variable that stores the
	// memory address of another variable.
	// Uninitialized pointers have the value as nil
	// Note: Unlike C, GO doesn't have pointer arithmetic
	foo := "bar"
	fooAddr := &foo // pointer to string
	fmt.Printf("foo has value %s and is located at %p\n", foo, fooAddr)
	fmt.Printf("fooAddr %v is of type %T\n", fooAddr, fooAddr)
	fmt.Printf("Dereferncing fooAddr provides %q\n", *fooAddr)

	// Create a pointer is to use the * operator
	var intPtr *int // initialized to nil
	i := 5
	intPtr = &i
	fmt.Println("intPtr points to:", intPtr, "address which contains value:", *intPtr)

	// Create a pointer using new
	newIntPtr := new(int)
	newIntPtr = &i
	fmt.Println("newIntPtr points to:", newIntPtr, "address which contains value:", *newIntPtr)

	// Pointers can be used to mutate the values at the address they store
	*newIntPtr = 100 // i = 100
	fmt.Println("Now, value of i is:", i)

	// &value => pointer
	// *pointer => value

	// Go allows pointers to pointers
	ptp := &newIntPtr
	fmt.Printf("ptp is of type %T and it's value contains the pointer to address %v which contains %v \n", ptp, *ptp, **ptp)
	**ptp++ // i++
	fmt.Println("Now, value of i is:", i)

	// Pointers are comparable if they're both nil or pointing to the same address
	fmt.Println("intPtr == newIntPtr:", intPtr == newIntPtr)
	var nilPtr *int
	fmt.Println("nilPtr == nil:", nilPtr == nil)

	// Passing to and receiving from functions
	// Passing pointers to a variable allows functions to change the value of the variable
	x := 5
	xPtr := &x
	fmt.Println("x before calling change:", x)
	change(xPtr)
	fmt.Println("x after calling change:", x)
	// Functions in go can return pointers to local variables
	floatPtr := getFloat()
	fmt.Println("floatPtr points to address:", floatPtr, "which has value:", *floatPtr)

	// Passing pointers to functions. Strictly speaking even when pointers are passed
	// to functions, it is not pass by reference since the pointer value is copied to
	// the arg variable within the function.

	// int, float64, bool, string are always passed by value and not pass by pointer
	// and hence these arg values when changed within the function do not reflect the
	// change on the caller.
	intVal := 1
	floatVal := 2.5
	boolVal := true
	stringVal := "hi"
	fmt.Println("Before calling the function:", intVal, floatVal, boolVal, stringVal)
	changeVal(intVal, floatVal, boolVal, stringVal)
	// same values before and after
	fmt.Println("After returning from the function:", intVal, floatVal, boolVal, stringVal)

	// To mutate the values of int, float64. bool, string we must use pointers
	// Arrays too behave similar to int, float64, bool, string etc and cannot be mutated
	// within functions. Although if you expect to change values of array within function,
	// then it's better to pass slices as opposed to pointer to an array
	changeValByPtr(&intVal, &floatVal, &boolVal, &stringVal)
	// Reflects the new values
	fmt.Println("After returning from the function:", intVal, floatVal, boolVal, stringVal)

	// Structures too are pass by value
	d1 := day{name: "Sunday", index: 1}
	fmt.Println("Before:", d1)
	changeDay(d1)
	fmt.Println("After:", d1)

	// To change structures within functions, pass the pointer
	changeDayPtr(&d1)
	fmt.Println("After:", d1)

	// Maps and Slices do not store the actual data but they store reference to
	// the address where the data is stored. They're pointers to runtime types
	slice := []int{1, 3, 5}
	fmt.Println("Before:", slice)
	changeSlice(slice)
	fmt.Println("After:", slice)

	myMap := map[int]int{2: 4, 3: 6}
	fmt.Println("Before:", myMap)
	changeMap(myMap)
	fmt.Println("After:", myMap)

}

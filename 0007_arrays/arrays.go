package main

import "fmt"

func main() {
	// Array is a composite indexable type of fixed lenght
	// containing elements of same type
	var numbers [4]int // array of 4 intergers
	fmt.Printf("numbers: %v\n", numbers)
	fmt.Printf("numbers: %#v\n", numbers)

	// Initializing using array literal
	var arrLiteralInit = [4]int{1, 2, 3, 4}
	fmt.Printf("arrLiteralInit: %#v\n", arrLiteralInit)

	// Initialize only few elements of the array
	var strArr = [4]string{"hi", "hello"}
	fmt.Printf("strArr: %#v\n", strArr)

	// Using the elipsis opertor "..." to find out length of array automatically
	var elipArr = [...]int{1, 3, 3, 4, 5, 6}
	fmt.Printf("elipArr: %#v has length %d\n", elipArr, len(elipArr))

	// Initializing array on multiple lines requires a comma for the last element too
	var mulLineArr = [...]string{
		"one",
		"two",
		"three",
		"four",
	}
	fmt.Printf("mulLineArr: %#v\n", mulLineArr)

	// Overwite element at a particular index
	mulLineArr[3] = "fourth element"
	fmt.Printf("mulLineArr: %#v\n", mulLineArr)

	// Iterating over arrays using for loop
	for index, value := range mulLineArr {
		fmt.Printf("%q is present at index %d of mulLineArr\n", value, index)
	}
	for ii := 0; ii < len(mulLineArr); ii++ {
		fmt.Println("index", ii, "value", mulLineArr[ii])
	}

	// Multidimensional array
	matrix := [3][3]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 9},
	}
	fmt.Println(matrix)

	// Array equality : Same length and values at corresponding indices
	m := [3]int{1, 2, 3}
	n := m // they do not point to same object in memory
	fmt.Println(n == m)
	n[2] = 0
	fmt.Println(n == m)

	// Keyed elements
	arrEx := [3]int{ // [1, 3, 5]
		2: 5,
		1: 3,
		0: 1,
	}
	fmt.Println(arrEx)
	// Unkeyed element gets index from last keyed element
	arrKeyEx := [...]string{ // ["", "", "hello", "world"]
		2: "hello",
		"world",
	}
	fmt.Printf("%#v\n", arrKeyEx)

}

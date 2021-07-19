package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Slices is a dynamic Array and can shrink or grow
	var defSlice []string
	fmt.Println("defSlice == nil is", defSlice == nil)
	// defSlice[0] = "abc" // run tim error. Cannot assign values to nil string

	// Use literals to initialize
	var intSlice = []int{1, 2, 3, 4}
	// intSlice[4] = 5 // run tim error. Index out of range. Cannot extend like in python
	fmt.Printf("intSlice is %#v\n", intSlice)

	// Use make to initialize
	var makeIntSlice = make([]int, 2)
	fmt.Printf("makeIntSlice is %#v\n", makeIntSlice)
	// Modifying elements
	makeIntSlice[0] = 1
	makeIntSlice[1] = 2
	fmt.Printf("makeIntSlice is %#v\n", makeIntSlice)
	// Iterating through splice
	for idx, val := range makeIntSlice {
		fmt.Println("idx:", idx, "val:", val)
	}
	// Copying slices using = operator creates a pointer to same object as the slice in RHS
	newIntSlice := makeIntSlice
	newIntSlice[0] = 0
	fmt.Printf("makeIntSlice is %#v\n", makeIntSlice) // [0, 2]
	// Comparing slices. Uninitialzed slices are comparable to nil. Initialized slices
	// cannot be compared using = operator
	var nilSpliece []int
	fmt.Println("nilSpliece is nil:", nilSpliece == nil)
	var p = []int{1, 2, 3}
	var q = []int{1, 2, 3}
	// fmt.Println(p == q) // Runtime error
	// Comparing two slices would require a for loop to compare each element
	eq := true
	if len(p) == len(q) {
		for idx_p, val_p := range p {
			if q[idx_p] != val_p {
				eq = false
				break
			}
		}
	} else {
		eq = false
	}
	fmt.Println("p:", p, "and q:", q, "are equal:", eq)
	// Extending the splice by appending elements
	p = append(p, 4)
	fmt.Printf("p is %#v\n", p)
	p = append(p, 5, 6, 7)
	fmt.Printf("p is %#v\n", p)
	p = append(p, q...)
	fmt.Printf("p is %#v\n", p)

	// Creating copy of slice
	foo := []int{1, 2, 3}
	bar := make([]int, 3)
	copy(bar, foo) // copy(dst, src) copies len(dst) number of elements from src to dst
	fmt.Println("bar:", bar, "is a copy of foo:", foo)
	bar[0] = 0
	fmt.Println("bar is now:", bar)
	bar2 := make([]int, 2)
	copy(bar2, foo) // copies only first two elements
	fmt.Println("bar2 is", bar2)
	bar3 := make([]int, 5)
	copy(bar3, foo) // copies all 3 elements of foo
	fmt.Println("bar3 is", bar3)

	// Slicing a slice
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers[1:3]) // [2, 3]
	fmt.Println(numbers[3:])  // [4, 5]
	fmt.Println(numbers[:2])  // [1, 2]
	fmt.Println(numbers[:])   // [1, 2, 3, 4, 5]

	/*
	  Go implements slice as a data structure called slice header.
	  Slice header contains 3 fields:
	  1. Address of backing array
	  2. length of slice, len() returns it // len of slice
	  3. capacity of slice, cap() returns it // len of backing array
	  New slices created from slicing an existing slice end up having same
	  backing arrays.
	*/
	sliceEx := []int{1, 2, 3, 4, 5} // backing array = [1, 2, 3, 4, 5]
	newSlice := sliceEx[2:4]        // value = [3, 4] but backing array is still [1, 2, 3, 4, 5]
	newSlice[0] = 0                 // changes value = [0, 4] and backing array to [1, 2, 0, 4, 5]
	fmt.Println("sliceEx is now", sliceEx)
	// When slices are created from array, the array becomes backing array
	arrEx := []int{1, 2, 3, 4, 5}
	s1, s2 := arrEx[0:2], arrEx[1:3] // [1, 2] , [2, 3]
	arrEx[1] = 0                     // [1, 0, 3, 4, 5]
	fmt.Println("s1 is now:", s1, "and s2 is now:", s2)
	fmt.Println("len(s1) is:", len(s1), "and cap(s1) is now:", cap(s1)) // 2, 5
	// To create a copy by slicing existing slice, we should use append
	var s3 []int
	s3 = append(s3, arrEx[0:2]...)
	arrEx[1] = 1                                           // [1, 1, 3, 4, 5]
	fmt.Println("arrEx is now:", arrEx, "but s3 is :", s3) // [1, 0]

	// Comparing size of array and slice having same number of elements
	aa := [5]int{1, 2, 3, 4, 5}
	bb := []int{1, 2, 3, 4, 5}
	fmt.Println("aa:", aa, "is of size", unsafe.Sizeof(aa), "bytes") // 5 elements and each takes 8 bytes => 40 bytes
	fmt.Println("bb:", bb, "is of size", unsafe.Sizeof(bb), "bytes") // 3 elements : backing array addr, cap, len. each takes 8 bytes => 24 bytes
}

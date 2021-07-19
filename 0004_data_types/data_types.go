package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
	  There are 4 types to represent integers:
	  int8 => 8 bit integers -128 to 127.
	  int16 => 16 bit integers.
	  int32 => 32 bit integers.
	  int64 => 62 bit integers.
	*/
	var i8 int8 = 127
	fmt.Printf("i8 is of type %T and value %d\n", i8, i8)
	/*
	  There are 4 types to represent unsigned integers:
	  uint8 => 8 bit unsigned integers.
	  uint16 => 16 bit unsigned integers.
	  uint32 => 32 bit unsigned integers.
	  uint64 => 62 bit unsignedintegers.
	*/
	var u32 uint32 = 1234567
	fmt.Printf("u32 is of type %T and value %d\n", u32, u32)
	/*
	  There are 2 types of floats:
	  float32 => 32 bit float.
	  float64 => 64 bit float. More precise. Always better to use this.
	*/
	var f32 float32 = 1.123
	fmt.Printf("f32 is of type %T and value %f\n", f32, f32)
	/*
	  byte represents uint8.
	  rune represents int32.
	  values must be enclosed by single quotes if we're using characters.
	*/
	var b byte = 'a'
	fmt.Printf("b is of type %T and value %v\n", b, b)
	var r rune = 'a'
	fmt.Printf("r is of type %T and value %v\n", r, r)
	/*
	  String is a sequence of characters enclosed by double quotes.
	*/
	var s string = "string"
	fmt.Printf("s is of type %T and value %v\n", s, s)
	/*
	  Array is a numbered sequence of elements of the same type.
	  An array has fixed length.
	*/
	var numbers = [4]int{1, 2, 3, 4}
	fmt.Printf("numbers is of type %T and value %v\n", numbers, numbers)
	/*
	  Slice is like an array but it can shrink or group dynamically.
	*/
	var nums = []int{5, 6, 7}
	fmt.Printf("nums is of type %T and value %v\n", nums, nums)
	/*
	  Map is an unordered group of elements of one type indexed by a set
	  of unique keys of another type.
	*/
	mymap := map[string]int{
		"a": 97,
		"b": 98,
	}
	fmt.Printf("mymap is of type %T and value %v\n", mymap, mymap)
	/*
	  Struct is a sequence of named elements called fields, each of which
	  has a type.
	*/
	type person struct {
		name string
		age  int
	}
	var johnDoe person
	johnDoe.name = "John Doe"
	johnDoe.age = 30
	fmt.Printf("johnDoe is of type %T and value %v\n", johnDoe, johnDoe)
	/*
	  Pointer stores address of another variable.
	  go doesn't have pointer arithmetic.
	*/
	var x int = 2
	ptr := &x
	fmt.Printf("ptr is of type %T and value %v\n", ptr, ptr)
	/*
	  functions are of the type func.
	*/
	fmt.Printf("f is of the type %T\n", f)
	/*
	  Type conversions. go is a strong typed language and operations between
	  two types are not permitted. However, we can use type conversion to
	  workaround this issue.
	*/
	var ui8 uint8 = 12
	var integer8 int8 = 2
	// fmt.Println(ui8 * integer8) // Compile time error
	fmt.Println(int8(ui8) * integer8)
	/*
	  Converting numeric values to string.
	*/
	s1 := string(99) // unicode code point for 99 i.e. c
	fmt.Println("99 is:", s1)
	// s_1 := string(3.12)  // Compile time error
	str1 := fmt.Sprintf("%f", 3.12)
	fmt.Println("str1 is:", str1)
	/*
	  Converting string to numeric values.
	*/
	s_1 := "3.12"
	var f1, _ = strconv.ParseFloat(s_1, 64) // strconv.ParseFloat(<string>, <64 or 32>)
	fmt.Println("f1 is:", f1)
	s_1 = "100"
	var f2, _ = strconv.ParseInt(s_1, 10, 64) // strconv.ParseFloat(<string>, <curr base in string: 2 or 8 or 10 or 16>, <8 or 16 or 32 or 64>)
	fmt.Println("f2 is:", f2)
	// Atoi & Itoa can be used to convert from string to int and vice versa respectively
	num_s_1, _ := strconv.Atoi(s_1)
	fmt.Printf("%q is %d in int\n", s_1, num_s_1)
	fmt.Printf("%v is %q in string\n", num_s_1, strconv.Itoa(num_s_1))
	/*
	  User defined types. Type conversion among user defined types is
	  allowed if they have same underlying type.
	*/
	type age int
	var my_age age = 28
	fmt.Printf("my_age is of type %T and value %v\n", my_age, my_age)
	type kms float64
	type miles float64
	var distKms kms = 100
	var distMiles miles = miles(distKms) * 0.62137
	fmt.Printf("%v kms is %v miles\n", distKms, distMiles)
}

func f() {

}

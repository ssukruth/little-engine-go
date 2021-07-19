package main

import "fmt"

func main() {
	// Struct is a sequence of named elements called fields.
	// Each field has a name and a type.
	// Blueprint of data the structure holds.
	type day struct {
		name  string
		index int
	}

	// Initializing via struct literal
	sunday := day{"Sunday", 1}
	monday := day{"Monday", 2}
	fmt.Printf("sunday is of type %T and value %+v\n", sunday, sunday)
	fmt.Printf("monday is of type %T and value %+v\n", monday, monday)
	fmt.Println("sunday.name", sunday.name)

	// Above method may not be useful if you're initialzing only some
	// fields or if you intend to change order of the fields in struct
	// definition. Using a field:value in struct literal is useful
	// for such cases
	tuesday := day{index: 3, name: "Tuesday"}
	fmt.Printf("tuesday is of type %T and value %+v\n", tuesday, tuesday)

	// Creating using "new" operator
	// Use the dot "." operator to access fields of the struct
	// Note that new operator returns pointer to struct
	wednesday := new(day)
	wednesday.name = "Wednesday"
	// uninitialized field values take the zero value of their type
	// wednesday.index = 0 since index is int
	fmt.Printf("wednesday is of type %T and value %+v\n", wednesday, wednesday)
	// Values of structs can be updated
	wednesday.index = 4
	fmt.Printf("wednesday is of type %T and value %+v\n", wednesday, wednesday)

	// Structs can be compared using "==" operator
	firstDec1992 := new(day)
	firstDec1992.name = "Tuesday"
	firstDec1992.index = 3
	if *firstDec1992 == tuesday {
		fmt.Println("1st dec 1992 is a Tuesday")
	} else {
		fmt.Println("1st dec 1992 is not a Tuesday")
	}

	// Copying a struct can be achieved using equals operator
	saturday := sunday
	saturday.name = "Saturday"
	saturday.index = 7
	fmt.Printf("sunday is of type %T and value %+v\n", sunday, sunday)
	fmt.Printf("saturday is of type %T and value %+v\n", saturday, saturday)

	// Anonymous struct is a structure with no explicitly defined struct type
	friday := struct {
		name  string
		index int
	}{
		name:  "Friday",
		index: 6,
	}
	fmt.Printf("friday is of type %T and value %+v\n", friday, friday)

	// Anonymous fields
	type newDay struct {
		string
		int
	}

	thursday := newDay{"Thursday", 5}
	fmt.Printf("thursday is of type %T and value %+v\n", thursday, thursday)
	fmt.Println("thursday.name is", thursday.string)

	// Embedded structs i.e. nested structs
	type date struct {
		dayInfo day
		dd      int
		mm      int
		yy      int
	}
	firstDec1992Date := date{
		dd: 1,
		mm: 12,
		yy: 92,
		dayInfo: day{
			name:  "Tuesday",
			index: 2,
		},
	}
	fmt.Printf("firstDec1992Date is of type %T and value %+v\n", firstDec1992Date, firstDec1992Date)
	fmt.Printf("firstDec1992Date falls on %s\n", firstDec1992Date.dayInfo.name)

}

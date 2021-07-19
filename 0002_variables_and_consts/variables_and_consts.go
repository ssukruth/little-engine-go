package main

import (
	"fmt"
)

func main() {
	/*
	   In traditional statically typed programming languages,
	   variables are declared as <type> <identifier> = <value>
	   Ex: int a = 10
	   In go, we use <keyword> <identifier> <type> = <value>
	   Ex: var a int = 10
	*/
	var a int = 10
	fmt.Println("Value of 'a' is:", a)
	var b bool = false
	fmt.Println("Value of 'b' is:", b)
	var c float64 = 3.1428
	fmt.Println("Value of 'c' is:", c)
	var d string = "Hello"
	fmt.Println("Value of 'd' is:", d)

	/*
	   go also provides type inference based on the value
	   provided for the variable.
	*/
	var aa = 100
	fmt.Printf("Value of variable 'aa' of type %T is: %d\n", aa, aa)
	var bb = true
	fmt.Printf("Value of variable 'bb' of type %T is: %t\n", bb, bb)
	var cc = 30.1428
	fmt.Printf("Value of variable 'cc' of type %T is: %f\n", cc, cc)
	var dd = "Hi"
	fmt.Printf("Value of variable 'dd' of type %T is: %s\n", dd, dd)

	/*
	   go also provides a shorthand variable declaration mechanism
	   using the short declaration operator ':='.
	   Note: Shorthand declaration cannot be used for package scope
	   variables.
	*/
	aaa := 1000
	fmt.Printf("Value of variable 'aaa' of type %T is: %d\n", aaa, aaa)
	bbb := true
	fmt.Printf("Value of variable 'bbb' of type %T is: %t\n", bbb, bbb)
	ccc := 300.1428
	fmt.Printf("Value of variable 'ccc' of type %T is: %f\n", ccc, ccc)
	ddd := "Greetings"
	fmt.Printf("Value of variable 'ddd' of type %T is: %s\n", ddd, ddd)

	/*
	   all variables declared except blank identifiers must be used
	   for golang to compile the program successfully.
	*/
	_ = "this is a blank identifier and it's okay to not use it"

	/*
	   Declaring multiple variables in a single statement is allowed.
	   We can use the shorthand operator to achieve this as long as
	   at least one variable is being declared newly.
	*/
	num1, num2 := 123, 456
	fmt.Println("Value of 'num1' & 'num2' are:", num1, num2)
	/*
	   The following statement leads to compilation error since
	   both variables on the left are already declared.
	*/
	// num1, num2 := 456, 123
	num2, num3 := 789, 456
	fmt.Println("Value of 'num1' 'num2' & 'num3' are:", num1, num2, num3)

	/*
	   Multiple declaration of same type can be declared using the
	   traditional declaration statement.
	*/
	var num4, num5, num6 int = 4, 5, 6
	fmt.Println("Value of 'num4' 'num5' & 'num6' are:", num4, num5, num6)

	/*
	   Multiple declaration of different types can be declared using the
	   traditional declaration statement.
	*/
	var (
		id   int    = 1234
		name string = "foobar"
	)
	fmt.Println("Value of 'id' & 'name' are:", id, name)

	/*
	  Constants are named literals. Their value isn't expected to change.
	  You may initialize a constant and not use it. However, a constant
	  must always be initialized.
	*/
	const daysInAWeek int = 7
	const m, n int = 1, 3
	const (
		sub = "-"
		add = "+"
	)
	/*
	  In a grouped constant, a constant takes value of previous const if
	  no value is provided/
	*/
	const (
		c1 = 500
		c2
		c3
	)
	fmt.Println(c1, c2, c3)

	/*
		  Some rules of constants:
		  1. You cannot change a constant
		  	const x = 5
			x = 6 ---> NOT ALLOWED
		  2. You cannot initialize constant at runtime
		  	const x = math.Pow(2, 3) ---> NOT ALLOWED
		  3. You cannot assign a variable to a constant
		  	var x = 5
			const cx = x ---> NOT ALLOWED
	*/

	/*
	  iota keyword implies successive integer constants
	*/

	const (
		c11 = iota // 0
		c22 = iota // 1
		c33 = iota // 2
	)
	fmt.Println(c11, c22, c33)

	const (
		c111 = iota // 0
		c222        // 1
		c333        // 2
	)
	fmt.Println(c111, c222, c333)

	const (
		_   = iota * 2 // 0
		c_1            // 2
		_              // 4
		c_2            // 6
		_              // 8
		c_3            // 10
	)
	fmt.Println(c_1, c_2, c_3)

	/*
	  There are 3 scopes in go:
	  	file scope
	        package scope
	        block scope
	*/

	/*
	  File Scope:
	  	1. Import statements are file scope. Reimporting same package leads to error.
	  	However, you can alias the same package while importing.
	        Wrong code:
	        	import "fmt"
	        	import "fmt" // second import leads to failure
	        Accepted (not recommended) code:
	        	import "fmt"
	        	import f "fmt" // doesn't fail because second import is aliased
	*/

	/*
	  Package Scope:
	  	Variables declared ouside functions in a file are package scope and
	        can be used by other files in the package.
	        Ex:
	        	package main
	        	import "fmt"
	        	var myVar := 5 // variable exists in package scope of "main"
	*/

	/*
	  Block Scope:
	  	Variables declared within a block
	        Ex:
	        	func f1() {
	        		var x := 5 // local or block scope variable
	        	}
	*/
}

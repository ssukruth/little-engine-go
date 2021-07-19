package main

import "fmt"

func main() {
	/*
	  Arithmetic operators apply to numeric values & are
	  used to perform common mathematical operations.
	  For ex:
	  	+ => Addition (can be used to concat strings too)
	        - => Subtraction
	        / => Division
	        * => Multiplication
	        % => modulo (only for int)
	*/
	fmt.Println("5 + 3 is ", 5+3)
	fmt.Println("5 - 3 is ", 5-3)
	fmt.Println("5 / 3 is ", float64(5)/float64(3))
	fmt.Println("5 * 3 is ", 5*3)
	fmt.Println("5 % 3 is ", 5%3)
	/*
	  Assignments operators are used to assign values.
	  For ex:
	  	= => Simple assignment
	        += => Increment assignment
	        -= => Decrement assignment
	        *= => multiplication assignment
	        /= => Division assignment
	        %= => Modulo assignment
	  	++ increments value by 1
	  	-- decrements value by 1
	  	Unlike C, it is illegal to use increment, decrements in expressions
	*/
	var x = 10
	x += 5 // 15
	x -= 1 // 14
	x *= 2 // 28
	x /= 7 // 4
	x %= 2 // 0
	x++    // 1
	x--    // 0
	fmt.Println("x is now", x)
	/*
	  Logical & comparison operators.
	  For ex:
	  	== => equals to comparison
	        >= => greater than or equal to
	        <= => less than or equal to
	        >  => greater than
	        <  => less than
	        != => not equal to comparison
	        && => conditional AND
	        || => conditional OR
	        !  => negation i.e. NOT operator
	*/
	var b1, b2 = 5, 6
	fmt.Printf("b1=%d b2=%d \n", b1, b2)
	var b3 = b1 == b2 && b1 != b2 // false AND true => false
	fmt.Println("b1 == b2 && b1 != b2 is:", b3)
	var b4 = b1 == b2 || b1 != b2 // false ||  true => true
	fmt.Println("b1 == b2 || b1 != b2 is:", b4)
	var b5 = !(b1 == b2) // NOT(false) => true
	fmt.Println("!(b1 == b2) is:", b5)
	/*
	  Bitwise operators
	  For ex:
	  	& => bitwise AND
	        | => bitwise OR
	        ^  x => bitwise XOR
	        << x => left shift, move the bits to left by x bits
	        >> x => right shift, move the bits to right by x bits
	*/
	var n1, n2 = 6, 10 // 110 , 1010
	fmt.Printf("n1=%d n2=%d\n", n1, n2)
	fmt.Println("n1 & n2 is:", n1&n2) // 110 AND 1010 = 0010 => 2
	fmt.Println("n1 | n2 is:", n1|n2) // 110 OR 1010 = 1110 => 14
	fmt.Println("n1 ^ n2 is:", n1^n2) // 110 XOR 1010 = 1100 => 12
	fmt.Println("n1 >> 1 is:", n1>>1) // 110 >> 1 => 11 => 3
	fmt.Println("n1 << 2 is:", n1<<2) // 110 << 2 => 11000 => 24
}

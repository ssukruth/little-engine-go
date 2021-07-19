package main

import (
	"fmt"
	"os"
	"strconv"
)

func ifElseIfFunc(a int, b int) {
	if a > b {
		fmt.Printf("%d is greater than %d\n", a, b)
	} else if a < b {
		fmt.Printf("%d is greater than %d\n", b, a)
	} else {
		fmt.Printf("%d is equal to %d\n", a, b)
	}
}

func ifElseShortFunc(stringvar string) {
	/*
	  Traditional if-else code is as follows:
	  	ii, err := strconv.Atoi(stringvar)
	        if err != nil {
	        	fmt.Println(err)
	        } else {
	        	fmt.Println(ii)
	        }
	  The following code is the short statement notation for the same logic.
	  Syntax:
	  	if <assignment>; <comparison1> {
	        	// comparison1 true block
	        } else if <<assignment>; <comparison2> {
	        	// comparison2 true block
	        } else {
	        	// code block when comparison1 & comparison2 both are false
	        }
	*/
	if ii, err := strconv.Atoi(stringvar); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ii)
	}
}

func forLooopFunc(mv int) {
	fmt.Println("Using for loop to print numbers from 0 to:", mv)
	for ii := 0; ii <= mv; ii++ {
		fmt.Println(ii)
	}
}

func whileLikeForLoopFunc(mv int) {
	fmt.Println("Using for loop to print numbers from", mv, "to 0 in a while loop like fashion")
	for mv >= 0 {
		fmt.Println(mv)
		mv--
	}
}

func forContinueBreakFunc(mv int) {
	fmt.Println("Using for loop to print even numbers between 0 and", mv)
	for ii := 0; ii < 100; ii++ {
		if ii%2 != 0 {
			continue
		}
		fmt.Println(ii)
		if ii >= mv {
			break
		}
	}
}

func labelsOuterLoopBreakFunc() {
	firstFivePrimes := [5]int{2, 3, 5, 7, 11}
	firstFiveEvens := [5]int{0, 2, 4, 6, 8}

outer:
	for _, prime := range firstFivePrimes {
		for _, even := range firstFiveEvens {
			if prime == even {
				fmt.Println(prime, "is both even and a prime")
				break outer
			} else {
				fmt.Println(prime, "!=", even)
			}
		}
	}
	fmt.Println("Exited outer for loop")
}

func labelsGotoFunc() {
	/*
	  goto statements cannot jump over declaration of new variables
	  Ex: Following code is illegal
	  	goto label
	        x := 5
	        label:
	        	// code
	*/
	fmt.Println("Iterating from 1 to 5 using goto & label")
	ii := 1
loop:
	if ii <= 5 {
		fmt.Println(ii)
		ii++
		goto loop
	}
}

func switchFunc() {
	day := "Wednesday"
	switch day {
	case "Sunday", "Saturday":
		fmt.Println("Weekend!!")
		// No explicit break statements required
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Alas! Weekday")
	default:
		fmt.Println("What?? is there even such a day")
	}
}

func main() {
	/*
	  Typical if-else if-else control flow.
	*/
	ifElseIfFunc(15, 10)
	ifElseIfFunc(10, 15)
	ifElseIfFunc(10, 10)
	/*
	  Command line arguments can be accessed via the slice
	  os.Args containing string elements.
	*/
	fmt.Println("os.Args:", os.Args)
	/*
	  Short statement for if-else.
	*/
	ifElseShortFunc("123")
	ifElseShortFunc("123a")
	/*
	  There's only one loop in go "for" loop
	*/
	forLooopFunc(10)
	whileLikeForLoopFunc(10)
	forContinueBreakFunc(30)
	/*
	  Labels are used with break, continue and goto statements.
	  Predominantly lables are used to break outer loops.
	*/
	labelsOuterLoopBreakFunc()
	labelsGotoFunc()
	/*
	  Switch statements can be used as an alternative to long
	  if-elseif-else statements while comparing to items of the
	  same type.
	*/
	switchFunc()
}

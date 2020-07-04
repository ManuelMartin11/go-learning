package main

// Go Tutorial

import (
	"fmt"
	"math"
	"time"
)

// Writing a Function
func add(x int, y int) int {
	return x + y*2
}

func simplerAdd(x, y int) int {
	return x + y*2
}

func returnTwoStrings(x, y string) (string, string) {
	return y, x
}

func nakedReturns(num int) (x, y int) {
	x = num / 2
	y = num/2 + 1
	return
}

// For loop
func innerForLoop(forRange int) {
	for i := 0; i < forRange; i++ {
		fmt.Printf("%v\n", i)
	}
}

// While is spelled 'for' in go
func whileWithForName(max int) {
	for max < 10 {
		max += max
	}
	fmt.Println(max)
}

func main() {
	// Basic Hello World
	fmt.Printf("hello, world!\n")

	// Format Print
	fmt.Printf("Maths Package %g\n", math.Sqrt(8))

	// Sum up 3 and 4
	var z int
	z = add(3, 4)
	fmt.Printf("%v\n", z)

	// Get 2 strings from function
	a, b := returnTwoStrings("hola", "que tal?")
	fmt.Printf("Return 2 strings from a function: %v %v\n", a, b)

	// Get 2 named variables from "naked return"
	a1, b1 := nakedReturns(23)
	fmt.Printf("%v %v\n", a1, b1)

	// Variables Declaration
	var c, python string = "C", "Python"
	fmt.Printf("%v %v\n", c, python)

	// Implicit Type Variable Declaration
	// := construct is not available outside of a function.
	// In such a case you must use 'var'
	x := 10
	fmt.Printf("%v\n", x)

	// How to make type conversions
	var i int = 23
	fmt.Printf("%v\n", float32(i)/12)

	// Run a For loop
	innerForLoop(5)

	// Conditionals
	newX := 2
	y := 0
	if newX > 3 {
		y = 23
	}
	{
		y = 12
	}
	newY := y
	fmt.Println(newY)

	// Short statement in Conditionals
	if v := math.Sqrt(float64(y)); v > float64(newX) {
		fmt.Println("Es menor!")
	}

	// If and Else
	if 2 > 3 {
		fmt.Println("2 is greater than 3")
	} else {
		fmt.Println("2 is not greather than 3!")
	}

	// Switch Case
	switch os := "Mac"; os {
	case "Windows":
		fmt.Println("you are using windows")
	case "Mac":
		fmt.Println("you are using mac")
	}

	// Switch with no specific Conditionals
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() > 12:
		fmt.Println("Good Afternoon")
	}

	// Defer: It only evaluates the function but executes
	// it only when the wrapping function ends
	func anotherFuncWithDefer(value int){
		defer fmt.Println("You'll see this when 'anotherFuncWithDefer' executes!")
		fmt.Println("After this, the `defer` statement will execute")
	}

}

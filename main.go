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

// Defer: It only evaluates the function but executes
// it only when the wrapping function ends
func anotherFuncWithDefer() {
	defer fmt.Println("You'll see this when 'anotherFuncWithDefer' executes!")

	fmt.Println("After this, the `defer` statement will execute")
	exampleVar := 2 + 3
	fmt.Println(exampleVar)
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

	/*
		POINTERS in Go. They hold the memory address of the value
	*/
	var pointer *int //the type *T is a pointer to a T value

	// The & operatior generates a pointer to its operand
	// In this case the operand is the variable `r`
	r := 42
	pointer = &r

	// Now if we print p, it will print the value of 'r'
	fmt.Println(*pointer)

	// We can change the value of the pointer at any time
	*pointer = 22

	/*
		STRUCTS: It's a collection of fields.
	*/
	type ExampleStruct struct {
		X int
		Y int
	}

	fmt.Println(ExampleStruct{1, 2})

	// How to access a struct field
	es := ExampleStruct{1, 2}
	fmt.Println(es.X)
	fmt.Println(es.Y)

	// Struct Literals
	v3 := ExampleStruct{X: 2, Y: 3}
	fmt.Println(v3)

	/*
		ARRAYS
	*/
	var exampleArray [3]string
	exampleArray[0] = "Hello"
	exampleArray[1] = ","
	exampleArray[2] = "World!"
	fmt.Println(exampleArray)

	anotherExampleArray := [4]int{1, 2, 3, 4}
	fmt.Println(anotherExampleArray)

	// Slice. They are like references to Arrays
	// Change the values of the slice, changes the underlying array
	fmt.Println(anotherExampleArray[1:3])

	// Slice Literals
	anStruct := []struct {
		i int
		b bool
	}{
		{2, true},
		{1, false},
	}

	fmt.Println(anStruct)

	// Length and Capacity
	anArray := [10]int{1, 2, 3}
	fmt.Println(len(anArray))
	fmt.Println(cap(anArray))

	anSlice := anArray[0:3]
	fmt.Println(len(anSlice))
	fmt.Println(cap(anSlice))

	// Make function and Slices
	varA := make([]int, 5)
	fmt.Println(varA)

	// Slices of Slices
	board := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}
	board[0][2] = "X"
	board[2][1] = "O"
	board[2][0] = "O"

	fmt.Println(board)

	// Using the Append Function
	justAnArray := []string{"hola"}
	fmt.Println(append(justAnArray, "adios"))

	// The Range Element
	rangeArray := [5]string{"one", "two", "three", "four", "five"}
	for i, v := range rangeArray {
		fmt.Println(i)
		fmt.Println(v)
	}

	// Maps
	m := make(map[string]string)
	m["key1"] = "value1"
	m["key2"] = "value2"
	fmt.Println(m)

	// Literal map
	type MapValue struct {
		surname string
		age     int
	}
	var newM = map[string]MapValue{
		"John": MapValue{surname: "Snow", age: 29},
	}
	fmt.Println(newM)

	// Maps Mutation
	newM["John"] = MapValue{"Snow", 30}
	elem := newM["John"]
	fmt.Println(elem)

	// Test if an element is in the map
	elem2, ok := m["John"] // if "John" is in m, ok = true
	fmt.Println(elem2, ok)

	// Test from methodName defined in newType
	newTypeInst := newType{X: 2, Y: 10}
	fmt.Println(newTypeInst.methodName(3))

	// Test with pointer struct
	fmt.Println("Testing pointers on method receivers")
	var typePointer *anotherTypeToBePointered
	typeValueA := anotherTypeToBePointered{X: 2, Y: 2}
	typeValueB := anotherTypeToBePointered{X: 3, Y: 2}

	typePointer = &typeValueA
	fmt.Println(typePointer.pointeredMethod())

	typeValueA = anotherTypeToBePointered{X: 5, Y: 2}
	fmt.Println(typePointer.pointeredMethod())

	typePointer = &typeValueB
	fmt.Println(typePointer.pointeredMethod())

	// Test on INTERFACES
	fmt.Println("Testing interface")
	var absI AbsInterface

	myF := MyOwnFloat64(-3.1416)
	absI = myF // Implementation of  AbsInterface in MyOwnFloat64
	fmt.Println(absI.Abs())
}

/*
	METHODS: Go does not have classes, but you can define methods inside types
*/
type newType struct {
	X float64
	Y float64
}

// A method has a receiver argument just before its name (methodName in this case)
func (nt newType) methodName(factor float64) float64 {
	return (nt.X * nt.Y) * factor
}

// Example of Method with a pointer receiver
type anotherTypeToBePointered struct {
	X float64
	Y float64
}

func (attbp *anotherTypeToBePointered) pointeredMethod() float64 {
	return attbp.X - attbp.Y
}

/*
	INTERFACES: Defined as a set of method signatures
*/
type AbsInterface interface {
	Abs() float64
}

type MyOwnFloat64 float64

func (f MyOwnFloat64) Abs() float64 {
	if f < 0 {
		f = f * -1
	} else {
		f = f * 1
	}
	return float64(f)
}

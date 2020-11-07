// Functions and methods example code
// Andrew Alston

package main

import "fmt"

// BasicFunction demonstrates the declaration of a basic function that takes two
// integer arguments and returns an integer.  In this case because the two
// input variables are of the same type, we can specify the type just once.
func BasicFunction(a, b int) int {
	return a * b
}

// BasicFunctionDualReturn demonstrates the declaration of a basic function that
// returns both a boolean value and an error type, as with BasicFunction, it takes two
// integer inputs and returns true with a nil error if a > b, false with a nil error if a < b,
// and false with an error if a == b
func BasicFunctionDualReturn(a, b int) (bool, error) {
	if a == b {
		return false, fmt.Errorf("error, a and b were equal")
	}
	return a > b, nil
}

// NestedAnonymous demonstrates the creation of an anonymous function inside a containing function
// The anonymous function is called by calling the variable name to which the anonymous function
// is assigned
func NestedAnonymous(a, b int) {
	Multiply := func(num1, num2 int) int {
		return num1 * num2
	}
	Divide := func(num1, num2 int) int {
		return num1 / num2
	}
	fmt.Printf("Calling Multiply on %d and %d returns %d\n", a, b, Multiply(a, b))
	fmt.Printf("Calling Divide on %d and %d returns %d\n", a, b, Divide(a, b))
}

// DemoStruct defines a structure that we will tie methods to as a demonstration of how methods
// work
type DemoStruct struct {
	a int
	b int
}

// Multiply is a method of DemoStruct and returns the two elements contained in the structure
// multiplied by each other
func (ds *DemoStruct) Multiply() int {
	return ds.a * ds.b
}

// String is a method we attach to DemoStruct to demonstrate how we can attach a method to a structure
// that performs a specific function that can be handled by things like fmt.Printf
func (ds *DemoStruct) String() string {
	return fmt.Sprintf("Struct contains A: %d and B: %d\n", ds.a, ds.b)
}

// SliceReceiver is a slice type we will use to demonstrate working with a slice that is passed as
// a function receiver
type SliceReceiver []int

// AppendSlice appends to the slice receiver.  Note - since by the golang spec you cannot assign to,
// or change the receiver using a method, we pass the slice as a pointer and append to the de-referenced
// slice.  In this case we use a varadic argument so we can append multiple integers at the same time,
// and then expand the varadic in the append function
func (sr *SliceReceiver) Append(a ...int) {
	if len(a) > 0 {
		*sr = append(*sr, a...)
	}
}

// DeleteElementNoOrder deletes an element at offset x from the receiver.  This method will result in
// the ordering of the slice changing since it does the deletion by swapping the specified element with
// the last element of the slice, and then popping the last element off the slice.  If x is larger
// than the length of the slice an error is returned.  As a note:  Since we need to refer to index entries
// within the receiver - we must de-reference the pointer and then index.  We place the receiver in
// () during this, since *sr[x] says to de-reference the element at sr[x] rather than saying return
// the element at index x of the slice
func (sr *SliceReceiver) DeleteElementNoOrder(x int) error {
	if x > len(*sr)-1 {
		return fmt.Errorf("x is out of range on the receiver")
	}
	// Swap the last element and the element we want to remove
	(*sr)[x], (*sr)[len(*sr)-1] = (*sr)[len(*sr)-1], (*sr)[x]
	// Pop the last element off the slice post swap
	*sr = (*sr)[:len(*sr)-1]
	return nil
}

func main() {
	fmt.Printf("Calling BasicFunction with arguments 2 and 3 returns %d\n", BasicFunction(2, 3))

	// Note: calling the function in this method makes both result and err scope local and as such
	// they will not be available outside of the if statement.  If we wish to use the results later
	// we can either do result, err := BasicFunctionDualReturns(2,3) and then reference the resulting
	// variables or setup the variables as var's and call if result,err = BasicFunctionDualReturn(2, 3)...
	fmt.Printf("\nDemonstrating dual returns:\n")
	if result, err := BasicFunctionDualReturn(2, 3); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Result of calling BasicFunctionDualReturn with arguments 2 and 3 was %v\n", result)
	}

	fmt.Printf("\nDemonstrating nested anonymous functions\n")
	NestedAnonymous(10, 5)

	fmt.Printf("\nDemonstrating struct methods\n")
	// First we create a variable of type DemoStruct and fill its elements. Note we are creating a pointer
	// here - and the below line is the equivalent of doing demoVar := new(DemoStruct) and then filling
	// in the elements later.
	demoVar := &DemoStruct{a: 1, b: 2}
	fmt.Printf("Since DemoStruct has a string method we can print its contents like this:\n%v\n", demoVar.String())
	fmt.Printf("Multiplying the elements contained in demoVar as a result of: %d\n", demoVar.Multiply())

	// This next section creates a slice an empty slice of integers and then appends some numbers to it
	// For this - we make an empty slice, and then tell the compiler this is of type SliceReceiver
	fmt.Printf("\nDemonstrating slice receivers\n")
	var a = SliceReceiver(make([]int, 0))
	(&a).Append(1, 2, 3, 4)
	fmt.Printf("Our slice now contains: %v\n", a)
	fmt.Printf("Deleting element 2 from our slice\n")
	if err := (&a).DeleteElementNoOrder(2); err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("Slice element deleted, slice now contains: %v\n", a)
	}

}

// Various binary level routines for demonstration purposes
// Andrew Alston

package main

import (
	"fmt"
	"strings"
)

// LogicalAnd performs a logical AND between num1 and num2 and prints the results.
// A Logical and states that if two bits are set, the resultant bit will be set,
// otherwise the resultant bit will be unset.  We use unsigned 16 bit integers here
// purely to make printing and reading easier
func LogicalAnd(num1, num2 uint16) {
	fmt.Printf("Performing Logical AND between %d and %d\n", num1, num2)
	fmt.Printf("Binary [%2d]: %016b\n", num1, num1)
	fmt.Printf("Binary [%2d]: %016b\n", num2, num2)
	fmt.Printf("Result [%2d]: %016b\n", num1&num2, num1&num2)
}

// ExclusiveOr performs a XOR between two numbers and prints the results.
// When an exclusive OR is performed, the result will be 0 if both bits are
// the same (0 ^ 0 = 0, 1 ^ 1 = 0) or 1 if they weren't (0^1 = 1)
func ExclusiveOr(num1, num2 uint16) {
	fmt.Printf("Performing Exclusive OR between %d and %d\n", num1, num2)
	fmt.Printf("Binary [%2d]: %016b\n", num1, num1)
	fmt.Printf("Binary [%2d]: %016b\n", num2, num2)
	fmt.Printf("Result [%2d]: %016b\n", num1^num2, num1^num2)
}

// LogicalOr performs a Logical OR between two numbers and prints the results.
// When a logical OR is performed, the result will be 1 if either of the bits
// were 1, otherwise the result will be 0
func LogicalOr(num1, num2 uint16) {
	fmt.Printf("Performing Logical OR between %d and %d\n", num1, num2)
	fmt.Printf("Binary [%2d]: %016b\n", num1, num1)
	fmt.Printf("Binary [%2d]: %016b\n", num2, num2)
	fmt.Printf("Result [%2d]: %016b\n", num1|num2, num1|num2)
}

// LeftShift shifts num1 left by the number of places specified in num2
// A Left Shift shifts all bits to the left by pushing zeros in from the right
// and dropping the left most bits.  This effectively doubles the num in question
// for each bit shifted left
func LeftShift(num1, num2 uint16) {
	fmt.Printf("Left shifting %d by %d bits\n", num1, num2)
	fmt.Printf("Binary [%3d]: %016b\n", num1, num1)
	fmt.Printf("Result [%3d]: %016b\n", num1<<num2, num1<<num2)
}

// RightShift shifts num1 right by the number of places specified in num2
// A Right Shift shifts all bits to the right by pushing zeros in from the left
// and dropping the right most bits.
func RightShift(num1, num2 uint16) {
	fmt.Printf("Right shifting %d by %d bits\n", num1, num2)
	fmt.Printf("Binary [%2d]: %016b\n", num1, num1)
	fmt.Printf("Result [%2d]: %016b\n", num1>>num2, num1>>num2)
}

// TestBit sets to see if a binary bit is set in the given number.  It assumes
// that numbers are numbered from the right starting at 0 and returns true
// if the bit is set or false if the bit is unset.  This works by shifting the
// given number to the right by x places, setting all but the right most to zero
// and then verifying that the result is 1.  For this, to cater for different
// size integers we use an interface type and do some assertion
func TestBit(num1 interface{}, bit uint8) (bool, error) {
	switch num1.(type) {
	case uint8:
		fmt.Printf("Binary [%3d] [%08b] >> %d == [%3d] %08b [%v]\n",
			num1, num1, 7-bit, num1.(uint8)>>(7-bit), num1.(uint8)>>(7-bit), (num1.(uint8)>>(7-bit))&1 == 1)
		return (num1.(uint8)>>(7-bit))&1 == 1, nil
	case uint16:
		fmt.Printf("Binary [%3d] [%08b] >> %d == [%3d] %08b [%v]\n",
			num1, num1, 15-bit, num1.(uint16)>>(15-bit), num1.(uint16)>>(15-bit), (num1.(uint16)>>(15-bit))&1 == 1)
		return (num1.(uint16)>>uint16(15-bit))&1 == 1, nil
	case uint32:
		fmt.Printf("Binary [%3d] [%08b] >> %d == [%3d] %08b [%v]\n",
			num1, num1, 31-bit, num1.(uint32)>>(31-bit), num1.(uint32)>>(31-bit), (num1.(uint32)>>(31-bit))&1 == 1)
		return (num1.(uint32)>>uint32(31-bit))&1 == 1, nil
	case uint64:
		fmt.Printf("Binary [%3d] [%08b] >> %d == [%3d] %08b [%v]\n",
			num1, num1, 63-bit, num1.(uint64)>>(63-bit), num1.(uint64)>>(63-bit), (num1.(uint64)>>(63-bit))&1 == 1)
		return (num1.(uint64)>>uint64(61-bit))&1 == 1, nil
	default:
		return false, fmt.Errorf("error, unrecognized type")
	}
}

// Varadic OR takes a varadic argument of a number of integers, performs a logical OR against all
// of those integers and returns the result
func VaradicOr(input ...int) int {
	var res int
	// strSlice exists just so we can make a nice output string
	var strSlice = make([]string, len(input))
	for i, in := range input {
		// We put our input integers into our string slice so we can print them nicely with a join
		strSlice[i] = fmt.Sprintf("%d", in)
		res = res | in
	}
	fmt.Printf("Performing %s = %d\n", strings.Join(strSlice, "|"), res)
	return res
}

// CombinedContains tests if a number is logically part of another number at a binary level. As an example:
// Given numbers 1, 2, 4, 8, if a logical OR is performed, they will produce a single number with the
// left 4 most bits set.  If we then do a logical AND of any of those against the result, we will have a
// non-zero return.
func CombinedContains(num1, num2 int) bool {
	fmt.Printf("Input [%d] [%08b] & [%d] [%08b] == [%d] [%08b] [%v]\n",
		num1,
		num1,
		num2,
		num2,
		num1&num2,
		num1&num2,
		!(num1&num2 == 0))
	return !(num1&num2 == 0)
}

func main() {
	fmt.Printf("Logical AND:\n")
	LogicalAnd(50, 60)
	fmt.Printf("\nExclusive OR:\n")
	ExclusiveOr(50, 60)
	fmt.Printf("\nLogical OR:\n")
	LogicalOr(50, 60)
	fmt.Printf("\nLeft Shift:\n")
	LeftShift(50, 4)
	fmt.Printf("\nRight Shift:\n")
	RightShift(50, 2)
	fmt.Printf("\nBit Testing:\n")
	_, _ = TestBit(uint8(50), 7)
	_, _ = TestBit(uint8(50), 6)
	_, _ = TestBit(uint8(50), 5)
	_, _ = TestBit(uint8(50), 4)
	fmt.Printf("\nVaradic OR:\n")
	res := VaradicOr(1, 2, 4, 8, 16, 32)
	fmt.Printf("%d\n", res)
	fmt.Printf("\nCombinedContains [Match]:\n")
	CombinedContains(res, 4)
	fmt.Printf("\nCombinedContains [No Match]:\n")
	CombinedContains(res, 64)
}

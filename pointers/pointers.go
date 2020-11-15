package main

import (
	"fmt"
	"unsafe"
)

// PointerNoReturn takes 3 arguments, and returns the result of a multiplied by b into
// the integer pointed to by c
func PointerNoReturn(a, b int, c *int) {
	*c = a * b
}

// PointerByteSwap takes a single uint16 pointer swaps the underlying byte order of the
// uint16 pointed to by the input number
func PointerByteSwap(a *uint16) {
	*a = (*a << 8) | (*a >> 8)
}

// ReturnByte returns a single byte at position b from one of the four bytes that comprise
// the uint32 passed in as a
func ReturnByte(a uint32, b uint8) (uint8, error) {
	if b >= 3 {
		return 0, fmt.Errorf("byte out of range")
	}
	return *(*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(&a)) + uintptr(b))), nil
}

// ReturnUint32FromSlice returns a single uint32 comprising of 4 sequential bytes
// in slice a, starting at reference position b
func ReturnUint32FromSlice(a []uint8, b int) (uint32, error) {
	if len(a) <= 3 {
		return 0, fmt.Errorf("insufficient bytes in slice")
	}
	return *(*uint32)(unsafe.Pointer(&a[b])), nil
}

// ModifySlicePointer appends the variadic b to the slice referenced by pointer a
func ModifySlicePointer(a *[]uint8, b ...uint8) {
	*a = append(*a, b...)
}

// IndexSlicePointer returns the byte at position b from the slice pointed to by a
func IndexSlicePointer(a *[]uint8, b int) (uint8, error) {
	if len(*a)-1 <= b {
		return 0, fmt.Errorf("byte not within slice range")
	}
	return (*a)[b], nil
}

func main() {
	// Various test variables
	var IntVar int
	var Uint16Var uint16
	var Uint32Var uint32 = 0xFFFEFDFC
	var ByteSlice = []uint8{1, 2, 3, 4}
	var ByteSlice2 = []uint8{5, 6, 7, 8}

	// Set IntVar by passing the pointer to the PointerNoReturn Function
	PointerNoReturn(300, 300, &IntVar)
	fmt.Printf("%d * %d placed in IntVar, result was %d\n", 300, 300, IntVar)

	// Swap the byte order of IntVar
	Uint16Var = uint16(IntVar)
	fmt.Printf("Swapping byte order of %0X\n", Uint16Var)
	PointerByteSwap(&Uint16Var)
	fmt.Printf("Uint16Var is now %0X\n", Uint16Var)

	x, _ := ReturnByte(Uint32Var, 2)
	fmt.Printf("Byte 2 of %0X is %0X\n", Uint32Var, x)

	// As a bonus question for discussion - note the byte order of the generated uint32
	y, _ := ReturnUint32FromSlice([]uint8{1, 2, 3, 4}, 0)
	fmt.Printf("Uint32 generated from byte slice %v at offset 0 is %0X\n", ByteSlice, y)

	fmt.Printf("Appending %v to %v...\n", ByteSlice, ByteSlice2)
	ModifySlicePointer(&ByteSlice, ByteSlice2...)
	fmt.Printf("ByteSlice is now %v\n", ByteSlice)

	ret, _ := IndexSlicePointer(&ByteSlice, 2)
	fmt.Printf("Byte at position 2 of ByteSlice is %d\n", ret)
}

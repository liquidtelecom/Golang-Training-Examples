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

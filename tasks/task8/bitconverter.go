// Package task8 Дана переменная int64.
// Написать программу которая устанавливает i-й бит в 1 или 0.
package task8

import (
	"fmt"
	"unsafe"
)

var (
	number   int64 = -5
	position uint  = 0
)

// BitConverter ...
func BitConverter(n int64, pos uint) int64 {
	// var mult int64 = 1
	// if n < 0 {
	// 	mult *= -1
	// }
	if isBitSetted(n, pos) {
		return clearBit(n, pos)
	}
	return setBit(n, pos)
}

func setBit(n int64, pos uint) int64 {

	n |= (1 << pos)

	return n

	// Ex
	// n = 5, pos = 0
	// 101 |= 001
	// 101 -> 5

	// n = 5, pos = 1
	// 101 |= 010
	// 111 -> 7

	// n = 5, pos = 2
	// 101 |= 100
	// 101 -> 5

	// n = 5, pos = 3
	// 0101 |= 1000
	// 1101 ->  13
}

// Clears the bit at pos in n.
func clearBit(n int64, pos uint) int64 {
	mask := ^(1 << pos)
	n &= int64(mask)
	return n
	// Ex
	//	// n = 5 , pos = 0
	// mask = ^1 = 0
	// 101 &= 0
	//	// 100 -> 4

	//	// n = 5, pos = 1
	// mask = ^10 = 01
	// 101 & = 01
	//	// 101

	// n = 5, pos = 2
	// mask = ^100 = 011
	// 101 &= 011
	// 001 -> 1

	// n = 5, pos = 3
	// mask = ^1000 = 0111
	// 101 &= 111
	// 101 -> 5
}

func isBitSetted(n int64, pos uint) bool {
	val := n & (1 << pos)
	return (val > 0)
	// Ex n = 5 pos = 0
	// 101 & (001)
	// 001 -> 1

	// Ex n = 5 pos = 1
	// 101 & (010)
	// 000 -> 0

	// Ex n = 5 pos = 2
	// 101 & (100)
	// 100 -> 4

	// Ex n = 5 pos = 3
	// 0101 & (1000)
	// 0000 -> 0

}

// GetBytesImplementationAsOneByte function for stdout testing
func GetBytesImplementationAsOneByte(n int64) [1]byte {
	return *(*[1]byte)(unsafe.Pointer(&n))
}

// Start ...
func Start() {
	fmt.Printf("Origin bytes : %b Origin val : %d\n",
		GetBytesImplementationAsOneByte(number),
		number,
	)
	fmt.Println("Position to change: ", position)
	res := BitConverter(number, position)

	fmt.Printf("Calculated bytes : %b, Calculated value : %v\n",
		GetBytesImplementationAsOneByte(res),
		res,
	)
}

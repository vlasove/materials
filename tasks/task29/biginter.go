// Package task29 Написать программу,
// которая перемножает,
// делит, складывает, вычитает 2 числовых переменных a,b,
// значение которых > 2^20.
package task29

import (
	"fmt"
	"math/big"
	"math/rand"
)

// GenerateBigIng ...
func GenerateBigInt(x int) *big.Int {
	b := big.NewInt(1<<x + int64(rand.Intn(15)))
	return b
}

// Add ...
func Add(x, y *big.Int) *big.Int {
	res := big.NewInt(0)
	res.Add(x, y)
	return res
}

// Mult ...
func Mult(x, y *big.Int) *big.Int {
	res := big.NewInt(0)
	res.Mul(x, y)
	return res
}

// Sub ...
func Sub(x, y *big.Int) *big.Int {
	res := big.NewInt(0)
	res.Sub(x, y)
	return res
}

// Div ...
func Div(x, y *big.Int) *big.Int {
	res := big.NewInt(0)
	res.Div(x, y)
	return res
}

// Start ...
func Start() {
	a, b := GenerateBigInt(39), GenerateBigInt(27)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("Sum is:", Add(a, b))
	fmt.Println("Mult:", Mult(a, b))
	fmt.Println("Sub:", Sub(a, b))
	fmt.Println("Div:", Div(a, b))
}

// Package task10 Дана последовательность температурных колебаний
// (-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5).
// Объединить данные значения в группы с шагом в 10 градусов.
// Последовательность в подмножностве не важна.
package task10

import (
	"errors"
	"fmt"
	"log"
)

const (
	step = 10
)

var (
	temperatures  = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	errZeroLength = errors.New("can not calc with zero length")
)

// Start ...
func Start() {
	minT, err := min(temperatures)
	if err != nil {
		log.Fatal(err)
	}
	maxT, err := max(temperatures)
	if err != nil {
		log.Fatal(err)
	}
	minB, maxB := borders(minT, maxT)
	chunk := buildChunck(minB, maxB)
	res := fillChunk(temperatures, chunk)
	fmt.Printf("%v\n", res)
}

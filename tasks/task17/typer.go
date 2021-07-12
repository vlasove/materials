// Package task17 Написать программу, которая в рантайме способна
// определить тип переменной — int, string, bool, channel из переменной типа interface{}.
package task17

import (
	"fmt"
	"reflect"
)

// WhichType ...
func WhichType(v interface{}) {
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Int:
		fmt.Println("It's int")
	case reflect.String:
		fmt.Println("It's string")
	case reflect.Bool:
		fmt.Println("It's bool")
	case reflect.Chan:
		fmt.Println("It's chan")
	default:
		fmt.Println("I don't know that is it")
	}
}

// Start ...
func Start() {

	var (
		intVal  = 10
		strVal  = "bob"
		boolVal = true
		chVal   = make(chan int)
	)

	WhichType(intVal)
	WhichType(strVal)
	WhichType(boolVal)
	WhichType(chVal)

	// var v interface{} = intVal
	// val, ok := v.(int)
	// fmt.Println(val, ok)

}

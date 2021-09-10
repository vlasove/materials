package main

import (
	"log"
	"strings"
)

type Object interface{}

type Collection []Object

func NewCollection(size int) Collection {
	return make(Collection, size)
}

type Callback func(current, currentKey, src Object) Object

func Map(c Collection, cb Callback) Collection {
	if c == nil {
		return Collection{}
	}
	if cb == nil {
		return c
	}
	result := NewCollection(len(c))
	for id, val := range c {
		result[id] = cb(val, id, c)
	}
	return result
}

func main() {
	transform10 := func(curVal, _, _ Object) Object {
		return curVal.(int) * 10
	}
	result := Map(Collection{1, 2, 3, 4, 5}, transform10)
	log.Println(result)

	transformUpper := func(curVal, _, _ Object) Object {
		return strings.ToUpper(curVal.(string))
	}
	result = Map(Collection{"alice", "bob", "cindy"}, transformUpper)
	log.Println(result)

}

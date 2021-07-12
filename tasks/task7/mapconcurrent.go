// Package task7 Реализовать
// конкурентную запись в map.
package task7

import (
	"fmt"
	"math/rand"
	"sync"
)

// Storage for concerrent writing
type Storage struct {
	mu      *sync.RWMutex
	storage map[string]string
}

// sample for generation random strings
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandStringRunes helper for generating random strings
// with fixed length
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// Start ...
func Start() {
	s := &Storage{&sync.RWMutex{}, make(map[string]string)}
	wg := new(sync.WaitGroup)

	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, storage *Storage, lenStr int) {
			storage.mu.Lock()
			storage.storage[RandStringRunes(lenStr+1)] = RandStringRunes(lenStr + 1)
			storage.mu.Unlock()
			wg.Done()
		}(wg, s, i)
	}
	wg.Wait()

	for k, v := range s.storage {
		fmt.Println(k, ":", v)
	}
}

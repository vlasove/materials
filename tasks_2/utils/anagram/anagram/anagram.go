package anagram

import (
	"fmt"
	"sort"
	"strings"
)

// Anagram ...
type Anagram struct {
	Words  []string
	Normal string
}

// SortNormal ...
func SortNormal(word string) string {
	parts := strings.Split(word, "")
	sort.Strings(parts)
	return strings.Join(parts, "")
}

// Find ...
func Find(words []string) []*Anagram {
	buckets := map[string][]string{}

	for _, w := range words {
		normal := SortNormal(w)
		buckets[normal] = append(buckets[normal], w)
	}

	anas := []*Anagram{}
	for _, ws := range buckets {
		if len(ws) == 1 {
			continue
		}

		a := &Anagram{
			Words:  ws,
			Normal: ws[0],
		}
		sort.Strings(a.Words)
		anas = append(anas, a)
	}

	return anas
}

// Set ...
type Set map[string][]string

// NewSet ...
func NewSet(anagrams []*Anagram) Set {
	as := make(map[string][]string)
	for _, val := range anagrams {
		as[val.Normal] = val.Words
	}
	return as
}
func (as Set) String() string {
	var res strings.Builder
	for key, val := range as {
		res.WriteString(
			fmt.Sprintf("%v : %v\n", key, val),
		)
	}
	return res.String()
}

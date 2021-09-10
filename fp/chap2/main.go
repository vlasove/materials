package main

import (
	"fmt"
	"strings"

	"github.com/vlasove/materials/fp/chap2/iterator"
)

type WordSize int

const (
	ZERO WordSize = 6 * iota
	SMALL
	MEDIUM
	LARGE
	XLARGE
	XXLARGE   WordSize = 50
	SEPARATOR          = ", "
)

type ChainLink struct {
	Data []string
}

func (cl *ChainLink) Value() []string {
	return cl.Data
}

type stringFunc func(string) string

func (cl *ChainLink) Map(fn stringFunc) *ChainLink {
	var mapped []string
	orig := *cl
	for _, s := range orig.Data {
		mapped = append(mapped, fn(s))
	}
	cl.Data = mapped
	return cl
}

func (cl *ChainLink) Filter(max WordSize) *ChainLink {
	filtered := []string{}
	orig := *cl
	for _, s := range orig.Data {
		if len(s) <= int(max) {
			filtered = append(filtered, s)
		}
	}
	cl.Data = filtered
	return cl
}

func main() {
	var iter iterator.CarIterator = iterator.NewCollection([]string{"CRV", "IS250", "Blazer"})
	value, ok := iter.Next()
	for ok {
		fmt.Println(value)
		value, ok = iter.Next()
	}

	constants := `
** Constants ***
ZERO: %v
SMALL: %d
MEDIUM: %d
LARGE: %d
XLARGE: %d
XXLARGE: %d
`
	fmt.Printf(constants, ZERO, SMALL, MEDIUM, LARGE, XLARGE, XXLARGE)

	words := []string{
		"tiny",
		"marathon",
		"philanthropinist",
		"supercalifragilisticexpialidocious"}

	data := ChainLink{words}
	fmt.Printf("unfiltered: %#v\n", data.Value())

	filtered := data.Filter(SMALL)
	fmt.Printf("filtered: %#v\n", filtered)

	fmt.Printf("filtered and mapped (<= SMALL sized words): %#v\n",
		filtered.Map(strings.ToUpper).Value())

	data = ChainLink{words}
	fmt.Printf("filtered and mapped (<= Up to MEDIUM sized words): %#v\n",
		data.Filter(MEDIUM).Map(strings.ToUpper).Value())

	data = ChainLink{words}
	fmt.Printf("filtered twice and mapped (<= Up to LARGE sized words): %#v\n",
		data.Filter(XLARGE).Map(strings.ToUpper).Filter(LARGE).Value())

	data = ChainLink{words}
	val := data.Map(strings.ToUpper).Filter(XXLARGE).Value()
	fmt.Printf("mapped and filtered (<= Up to XXLARGE sized words): %#v\n", val)

}

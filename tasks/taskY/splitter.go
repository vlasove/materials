package tasky

import (
	"fmt"
	"strings"
)

func sanitize(input []string) []string {
	res := []string{}
	for _, world := range input {
		if len(world) != 0 {
			res = append(res, world)
		}
	}
	return res
}

func Start() {
	msg := "Hello      world   this. is...B       ob!"
	res := strings.Join(
		sanitize(strings.Split(msg, " ")),
		" ",
	)
	fmt.Println(res)
}

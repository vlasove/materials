package main

import (
	"fmt"
	"log"

	"github.com/vlasove/materials/bytest/f10/whos"
)

var (
	urls = []string{"https://stackoverflow.com/", "https://github.com/", "https://yandex.ru/", "https://quii.gitbook.io/learn-go-with-tests/"}
)

func main() {
	url, err := whos.WhosFasterURL(urls...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("fastest url:", url)
}

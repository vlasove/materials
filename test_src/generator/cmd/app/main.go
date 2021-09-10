package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/vlasove/materials/test_src/generator/internal/app/filemanager"
	"github.com/vlasove/materials/test_src/generator/internal/app/generator"
)

var (
	n        int
	filePath string
)

func init() {
	flag.StringVar(&filePath, "path", "data.json", "path to output json file")
	flag.IntVar(&n, "n", 40, "amount of orders to generate")
}

func main() {
	flag.Parse()
	os := generator.GenerateRandomOrders(n)
	fm := filemanager.New(filePath)
	if err := fm.Dump(os); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully dumped random data to:", filePath)
}

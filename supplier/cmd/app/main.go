package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/vlasove/materials/supplier/internal/app/generator"
	"github.com/vlasove/materials/supplier/internal/app/publisher"
)

var (
	generatorMode   bool
	fileOutputPath  string
	countToGenerate int
	publisherMode   bool
	configPath      string
	sourcePath      string
)

func init() {
	flag.BoolVar(&generatorMode, "g", false, "Generator mode (generate some data to output file)")
	flag.StringVar(&fileOutputPath, "f", "./data.json", "Path to output .json file in generator mode")
	flag.IntVar(&countToGenerate, "c", 10, "Amout of samples in generator mode")
	flag.BoolVar(&publisherMode, "p", false, "Publisher mode (send data to NATS)")
	flag.StringVar(&configPath, "conf", "configs/nats.toml", "Path to config file in publisher mode")
	flag.StringVar(&sourcePath, "s", "./data.json", "Path to source file in publisher mode")
}

func main() {
	flag.Parse()
	if !generatorMode && !publisherMode {
		log.Fatal("should choose one of this : g OR p")
		os.Exit(1)
	}

	if generatorMode {
		if err := generator.Start(countToGenerate, fileOutputPath); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Successfully dumped random data to:", fileOutputPath)
	}

	if publisherMode {
		if err := publisher.Start(configPath, sourcePath); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Successfully sended data to NATS server")
	}
}

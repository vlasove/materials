package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/vlasove/materials/test_src/publisher/internal/app/helpers"
	"github.com/vlasove/materials/test_src/publisher/internal/app/natspub"
)

var (
	configPath     string
	sourceDataPath string
)

func init() {
	flag.StringVar(&configPath, "c", "configs/server.toml", "path to config .toml file")
	flag.StringVar(&sourceDataPath, "s", "source/data.json", "path to source .json data")
}

func main() {
	flag.Parse()
	natsConfig := natspub.NewNatsConfig()
	_, err := toml.DecodeFile(configPath, natsConfig)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Current Nats Urls:", natsConfig.NatsUrls)

	log.Println("Trying to build connection....")
	api := natspub.New(natsConfig)
	if err := api.BuildConnection(); err != nil {
		log.Fatal(err)
	}
	defer api.CloseConnection()
	log.Println("Successfully connected to:", api.GetConnect().Opts.Name, "by", api.GetConnect().ConnectedUrl())

	data, err := helpers.Load(sourceDataPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range data.Orders {
		byteData, err := helpers.UnmarhsallAndSleep(d)
		if err != nil {
			log.Println(err)
			continue
		}
		if err := api.Publish(byteData); err != nil {
			log.Println(err)
			continue
		}
	}
	log.Println("All data sended. Turning off...")
}

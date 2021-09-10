package main

import (
	"flag"
	"log"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/vlasove/materials/test_src/subscriber/internal/app/client"
)

var (
	configPath       string
	recievedDataPath string
)

func init() {
	flag.StringVar(&configPath, "config", "configs/client.toml", "path to .toml config file")
	flag.StringVar(&recievedDataPath, "recv", "source/data.json", "path to .json file for saving data")
}

func main() {
	flag.Parse()
	config := client.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	client := client.NewAPINatsConnector(config)

	log.Println("Trying to build connection ...")
	if err := client.BuildConnection(); err != nil {
		log.Fatal(err)
	}
	//defer client.CloseConnection()
	log.Println(
		"Successfully connected to URL:",
		client.GetConnection().ConnectedUrl(),
		"by:", client.GetConnection().Opts.Name,
	)

	log.Println("Trying to subscribe to", config.NatsSubject)
	if err := client.Subscribe(client.PrintAdapter); err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully subscribed to", config.NatsSubject)

	runtime.Goexit()

	// termChan := make(chan os.Signal, 1)
	// signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTSTP)

}

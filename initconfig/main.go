package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hashicorp/consul/api"
)

type Consul struct {
	Address string
}

func main() {
	flag.Parse()

	consulbytes, err := ioutil.ReadFile(consulFile)
	if err != nil {
		log.Fatal(err)
	}

	var consul Consul

	err = json.Unmarshal(consulbytes, &consul)
	if err != nil {
		log.Fatal(err)
	}

	client, err := api.NewClient(&api.Config{Scheme: "http", Address: consul.Address})

	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	kvPair := &api.KVPair{Key: "future", Value: bytes}

	_, err = client.KV().Put(kvPair, nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("init done!")
}

var consulFile string

var configFile string

func init() {
	flag.StringVar(&consulFile, "consul", "consul.json", "consul address file path")
	flag.StringVar(&configFile, "config", "appconfig.json", "config file path")
}

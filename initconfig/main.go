package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/liuyuexclusive/utils/configutil"

	"github.com/hashicorp/consul/api"
)

type Consul struct {
	Address string
}

func main() {
	flag.Parse()

	client, err := configutil.Client()

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

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "appconfig.json", "config file path")
}

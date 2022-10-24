package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/Evheniq/http-rest-api/internal/app/apiserver"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config path")
}

func main() {
	flag.Parse()

	Config := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, Config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(Config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}

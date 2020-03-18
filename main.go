package main

import (
	"flag"
	"log"
)

func main() {
	flag.StringVar(&Config.ConfigPath, "config", "./config.toml", "The config file for verteilzentrum.")
	flag.StringVar(&Config.DataDir, "datadir", "./", "The location where all persistent data is stored.")
	flag.Parse()

	if err := ReadConfig(); err != nil {
		log.Fatal(err)
	}

	if err := InitDatabase(); err != nil {
		log.Fatal(err)
	}

	InitServer()
}

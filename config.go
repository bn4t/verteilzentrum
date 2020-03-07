package main

import (
	"flag"
	"github.com/BurntSushi/toml"
)

type list struct {
	Name      string
	Whitelist []string
	Blacklist []string
}

type verteilzentrumConfig struct {
	BinTo           string `toml:"bind_to"`
	Hostname        string
	ReadTimeout     int `toml:"read_timeout"`
	WriteTimeout    int `toml:"write_timeout"`
	MaxMessageBytes int `toml:"max_message_bytes"`
}

type configuration struct {
	Verteilzentrum verteilzentrumConfig
	Lists          []list `toml:"list"`
}

var Config configuration

func ReadConfig() error {
	var fileName string

	flag.StringVar(&fileName, "config", "./config.toml", "The config file for verteilzentrum.")
	flag.Parse()

	_, err := toml.DecodeFile(fileName, &Config)
	return err
}

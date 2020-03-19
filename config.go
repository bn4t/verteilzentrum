package main

import (
	"github.com/BurntSushi/toml"
)

type list struct {
	Name       string
	Whitelist  []string
	Blacklist  []string
	CanPublish []string `toml:"can_publish"`
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
	ConfigPath     string
	DataDir        string
}

var Config configuration

func ReadConfig() error {
	_, err := toml.DecodeFile(Config.ConfigPath, &Config)
	return err
}

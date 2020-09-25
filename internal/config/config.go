/*
 *     verteilzentrum
 *     Copyright (C) 2020  bn4t
 *
 *     This program is free software: you can redistribute it and/or modify
 *     it under the terms of the GNU General Public License as published by
 *     the Free Software Foundation, either version 3 of the License, or
 *     (at your option) any later version.
 *
 *     This program is distributed in the hope that it will be useful,
 *     but WITHOUT ANY WARRANTY; without even the implied warranty of
 *     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *     GNU General Public License for more details.
 *
 *     You should have received a copy of the GNU General Public License
 *     along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package config

import (
	"errors"
	"github.com/BurntSushi/toml"
	"os"
	"strings"
)

type list struct {
	Name       string   `toml:"name"`
	Whitelist  []string `toml:"whitelist"`
	Blacklist  []string `toml:"blacklist"`
	CanPublish []string `toml:"can_publish"`
}

type verteilzentrumConfig struct {
	Logging         string `toml:"logging"`
	BindTo          string `toml:"bind_to"`
	BindToTls       string `toml:"bind_to_tls"`
	Hostname        string `toml:"hostname"`
	ReadTimeout     int    `toml:"read_timeout"`
	WriteTimeout    int    `toml:"write_timeout"`
	MaxMessageBytes int    `toml:"max_message_bytes"`
	TlsCertFile     string `toml:"tls_cert_file"`
	TlsKeyFile      string `toml:"tls_key_file"`
	DataDir         string `toml:"data_dir"`
	MtaAddress      string `toml:"mta_address"`
	MtaAuthMethod   string `toml:"mta_auth_method"`
	MtaUsername     string `toml:"mta_username"`
	MtaPassword     string `toml:"mta_password"`
}

type configuration struct {
	Verteilzentrum verteilzentrumConfig
	Lists          []list `toml:"list"`
	ConfigPath     string
}

var Config = &configuration{}

// ReadConfig reads in the config and does some basic config validation
func ReadConfig() error {
	if _, err := toml.DecodeFile(Config.ConfigPath, Config); err != nil {
		return err
	}

	// some basic config validation
	if Config.Verteilzentrum.MtaAuthMethod != "PLAIN" && Config.Verteilzentrum.MtaAuthMethod != "ANONYMOUS" {
		return errors.New("invalid mta_auth_method specified. Valid auth methods are 'PLAIN' and 'ANONYMOUS'")
	}

	fi, err := os.Stat(Config.Verteilzentrum.DataDir)
	if err != nil {
		return errors.New("specified data_dir could not be found: " + err.Error())
	}

	if !fi.IsDir() {
		return errors.New("specified data_dir is not a directory")
	}

	// append a trailing slash to the data dir if it doesn't have one
	if !strings.HasSuffix(Config.Verteilzentrum.DataDir, "/") {
		Config.Verteilzentrum.DataDir += "/"
	}

	// the constants defined in the logging package can't be used here because it would lead to an import cycle
	if Config.Verteilzentrum.Logging != "DEBUG" &&
		Config.Verteilzentrum.Logging != "INFO" &&
		Config.Verteilzentrum.Logging != "ERROR" {
		return errors.New("invalid logging level specified. Valid logging levels are: DEBUG, INFO and ERROR")
	}
	return nil
}

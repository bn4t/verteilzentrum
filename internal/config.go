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

package internal

import (
	"github.com/BurntSushi/toml"
)

type list struct {
	Name       string   `toml:"name"`
	Whitelist  []string `toml:"whitelist"`
	Blacklist  []string `toml:"blacklist"`
	CanPublish []string `toml:"can_publish"`
}

type verteilzentrumConfig struct {
	BindTo          string `toml:"bind_to"`
	BindToTls       string `toml:"bind_to_tls"`
	Hostname        string `toml:"hostname"`
	ReadTimeout     int    `toml:"read_timeout"`
	WriteTimeout    int    `toml:"write_timeout"`
	MaxMessageBytes int    `toml:"max_message_bytes"`
	TlsCertFile     string `toml:"tls_cert_file"`
	TlsKeyFile      string `toml:"tls_key_file"`
	DataDir         string `toml:"data_dir"`
}

type configuration struct {
	Verteilzentrum verteilzentrumConfig
	Lists          []list `toml:"list"`
	ConfigPath     string
}

var Config *configuration

func ReadConfig() error {
	_, err := toml.DecodeFile(Config.ConfigPath, Config)
	return err
}

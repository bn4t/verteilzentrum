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

package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"verteilzentrum/internal"
	"verteilzentrum/internal/config"
	"verteilzentrum/internal/logging"
)

func main() {
	flag.StringVar(&config.Config.ConfigPath, "config", "./config.toml", "The config file for verteilzentrum.")
	flag.Parse()

	if err := config.ReadConfig(); err != nil {
		log.Fatal(err)
	}

	if err := internal.InitDatabase(); err != nil {
		log.Fatal(err)
	}

	internal.InitServer()

	logging.LogMsg("verteilzentrum successfully started", logging.LogLvlInfo)

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	<-sigc
	logging.LogMsg("Stopping verteilzentrum gracefully...", logging.LogLvlInfo)
	for _, v := range internal.Servers {
		v.Close()
	}
	logging.LogMsg("Goodbye...", logging.LogLvlInfo)
}

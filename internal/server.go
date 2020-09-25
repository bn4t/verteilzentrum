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
	"github.com/emersion/go-smtp"
	"log"
	"math/rand"
	"time"
	"verteilzentrum/internal/config"
	"verteilzentrum/internal/logging"
)

var Servers []*smtp.Server

func InitServer() {
	rand.Seed(time.Now().UTC().UnixNano())

	logging.LogMsg("starting message queue...", logging.LogLvlInfo)
	go startMsgQueueRunner()

	// if tls options are set start tls listener
	if config.Config.Verteilzentrum.TlsCertFile != "" && config.Config.Verteilzentrum.TlsKeyFile != "" {
		go func() {
			logging.LogMsg("starting tls listener at "+config.Config.Verteilzentrum.BindToTls, logging.LogLvlInfo)
			s := createNewServer()
			if err := s.ListenAndServeTLS(); err != nil {
				log.Fatal(err)
			}
		}()
	}
	go func() {
		logging.LogMsg("starting plaintext listener at "+config.Config.Verteilzentrum.BindTo, logging.LogLvlInfo)
		s := createNewServer()
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
}

func createNewServer() *smtp.Server {
	be := new(Backend)

	s := smtp.NewServer(be)

	s.Domain = config.Config.Verteilzentrum.Hostname
	s.WriteTimeout = time.Duration(config.Config.Verteilzentrum.WriteTimeout) * time.Millisecond
	s.ReadTimeout = time.Duration(config.Config.Verteilzentrum.ReadTimeout) * time.Millisecond
	s.MaxMessageBytes = config.Config.Verteilzentrum.MaxMessageBytes
	s.AuthDisabled = true

	// add the tls config also to the non-tls listener to support STARTTLS
	if config.Config.Verteilzentrum.TlsCertFile != "" && config.Config.Verteilzentrum.TlsKeyFile != "" {
		var err error
		s.Addr = config.Config.Verteilzentrum.BindToTls
		if s.TLSConfig, err = loadTLSCertificate(); err != nil {
			log.Fatal(err)
		}
	} else {
		s.Addr = config.Config.Verteilzentrum.BindTo
	}

	Servers = append(Servers, s)
	return s
}

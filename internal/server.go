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
)

var Servers []*smtp.Server

func InitServer() {
	rand.Seed(time.Now().UTC().UnixNano())

	log.Print("Starting message queue...")
	go StartMsgQueueRunner()

	// if tls options are set start tls listener
	if Config.Verteilzentrum.TlsCertFile != "" && Config.Verteilzentrum.TlsKeyFile != "" {
		go func() {
			log.Print("Starting TLS listener...")
			s := createNewServer()
			if err := s.ListenAndServeTLS(); err != nil {
				log.Fatal(err)
			}
		}()
	}
	go func() {
		log.Print("Starting plaintext listener...")
		s := createNewServer()
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

}

func createNewServer() *smtp.Server {
	be := new(Backend)

	s := smtp.NewServer(be)

	s.Domain = Config.Verteilzentrum.Hostname
	s.WriteTimeout = time.Duration(Config.Verteilzentrum.WriteTimeout) * time.Millisecond
	s.ReadTimeout = time.Duration(Config.Verteilzentrum.ReadTimeout) * time.Millisecond
	s.MaxMessageBytes = Config.Verteilzentrum.MaxMessageBytes
	s.MaxRecipients = 1
	s.AuthDisabled = true

	// add the tls config also to the non-tls listener to support STARTTLS
	if Config.Verteilzentrum.TlsCertFile != "" && Config.Verteilzentrum.TlsKeyFile != "" {
		var err error
		if s.TLSConfig, err = LoadTLSCertificate(); err != nil {
			log.Fatal(err)
		}
	}

	Servers = append(Servers, s)

	return s
}

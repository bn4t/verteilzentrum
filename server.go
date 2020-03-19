package main

import (
	"github.com/emersion/go-smtp"
	"log"
	"math/rand"
	"time"
)

func InitServer() {
	rand.Seed(time.Now().UTC().UnixNano())
	be := &Backend{}

	s := smtp.NewServer(be)

	s.Domain = Config.Verteilzentrum.Hostname
	s.WriteTimeout = time.Duration(Config.Verteilzentrum.WriteTimeout) * time.Millisecond
	s.ReadTimeout = time.Duration(Config.Verteilzentrum.ReadTimeout) * time.Millisecond
	s.MaxMessageBytes = Config.Verteilzentrum.MaxMessageBytes
	s.MaxRecipients = 1
	s.AuthDisabled = true

	log.Print("Starting message queue...")
	go StartMsgQueueRunner()

	// if tls options are set start tls listener
	if Config.Verteilzentrum.TlsCertFile != "" && Config.Verteilzentrum.TlsKeyFile != "" {
		var err error
		if s.TLSConfig, err = LoadTLSCertificate(); err != nil {
			log.Fatal(err)
		}
		go func() {
			log.Print("Starting TLS listener...", s.Addr)
			if err := s.ListenAndServeTLS(); err != nil {
				log.Fatal(err)
			}
		}()

	}
	log.Print("Starting plaintext listener...", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}

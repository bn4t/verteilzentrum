package main

import (
	"github.com/emersion/go-smtp"
	"log"
	"time"
)

func InitServer() {
	be := &Backend{}

	s := smtp.NewServer(be)

	s.Addr = Config.Verteilzentrum.BinTo
	s.Domain = Config.Verteilzentrum.Hostname
	s.WriteTimeout = time.Duration(Config.Verteilzentrum.WriteTimeout) * time.Millisecond
	s.ReadTimeout = time.Duration(Config.Verteilzentrum.ReadTimeout) * time.Millisecond
	s.MaxMessageBytes = Config.Verteilzentrum.MaxMessageBytes
	s.MaxRecipients = 1
	s.AuthDisabled = true

	log.Println("Starting server at", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"crypto/tls"
	"github.com/emersion/go-smtp"
	"net"
	"strings"
)

func ForwardMail(data []byte, from string, to string) error {
	domain := strings.Split(to, "@")[1]
	mx, err := net.LookupMX(domain)
	if err != nil {
		return err
	}

	// TODO: support preferences
	c, err := smtp.Dial(mx[0].Host + ":25")
	if err != nil {
		return err
	}

	if err := c.Hello(Config.Verteilzentrum.Hostname); err != nil {
		return err
	}

	// ignore starttls errors
	_ = c.StartTLS(&tls.Config{})

	// Set the sender and recipient first
	if err := c.Mail(from, nil); err != nil {
		return err
	}
	if err := c.Rcpt(to); err != nil {
		return err
	}

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		return err
	}

	_, err = wc.Write(data)
	if err != nil {
		return err
	}
	err = wc.Close()
	if err != nil {
		return err
	}

	// Send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		return err
	}
	return nil
}

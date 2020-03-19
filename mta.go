package main

import (
	"crypto/tls"
	"errors"
	"github.com/emersion/go-smtp"
	"net"
	"sort"
	"strings"
)

func SendMail(data []byte, from string, to string) error {
	var c *smtp.Client
	domain := strings.Split(to, "@")[1]
	mx, err := net.LookupMX(domain)
	if err != nil {
		return err
	}

	// sort all mx records according to preference
	sort.Slice(mx, func(i, j int) bool {
		return mx[i].Pref < mx[j].Pref
	})

	// try all mx records
	for _, v := range mx {
		c, err = smtp.Dial(v.Host + ":25")

		// if the connection succeeds exit the loop
		if err == nil {
			break
		}
	}

	// if the client is null all mx records failed
	if c == nil {
		return errors.New("Failed to connect to mailserver for " + to)
	}

	if err := c.Hello(Config.Verteilzentrum.Hostname); err != nil {
		return err
	}

	// ignore startTLS errors
	_ = c.StartTLS(&tls.Config{})

	// set the envelope sender and recipient
	if err := c.Mail(from, nil); err != nil {
		return err
	}
	if err := c.Rcpt(to); err != nil {
		return err
	}

	// send the email body.
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

	// send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		return err
	}

	return nil
}

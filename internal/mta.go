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

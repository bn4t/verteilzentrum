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
	"log"
	"time"
)

func SendSubscribeNotif(receiver string, list string) error {
	msg := "From: " + list + "\r\n" +
		"To: " + receiver + "\r\n" +
		"Subject: " + list + " subscription\r\n" +
		"Mime-Version: 1.0\r\n" +
		"Date: " + time.Now().UTC().Format(time.RFC1123) + "\r\n" +
		"Content-Type: text/plain\r\n" +
		"Message-Id: " + GenerateMessageId(receiver) + "\r\n" +
		"\r\n" +
		"Hi,\r\n\r\nYou are now successfully subscribed to " + list + ".\r\n\r\n" +
		"You can unsubscribe at any time by sending an email to unsubscribe+" + list + ".\r\n"

	if err := SendMail([]byte(msg), "bounce+"+list, receiver); err != nil {
		if err = AddToMsgQueue(receiver, "bounce+"+list, msg); err != nil {
			log.Print("Error while adding message to message queue:")
			log.Print(err)
		}
		return err
	}
	return nil
}

func SendUnsubscribeNotif(receiver string, list string) error {
	msg := "From: " + list + "\r\n" +
		"To: " + receiver + "\r\n" +
		"Subject: " + list + " subscription\r\n" +
		"Mime-Version: 1.0\r\n" +
		"Date: " + time.Now().UTC().Format(time.RFC1123) + "\r\n" +
		"Content-Type: text/plain\r\n" +
		"\r\n" +
		"Hi,\r\n\r\nYou are now successfully unsubscribed from " + list + ".\r\n\r\n" +
		"You can resubscribe at any time by sending an email to subscribe+" + list + ".\r\n"

	if err := SendMail([]byte(msg), "bounce+"+list, receiver); err != nil {
		if err = AddToMsgQueue(receiver, "bounce+"+list, msg); err != nil {
			log.Print("Error while adding message to message queue:")
			log.Print(err)
		}
		return err
	}
	return nil
}

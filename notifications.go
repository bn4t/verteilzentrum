package main

import "time"

func SendSubscribeNotif(receiver string, list string) error {
	msg := "From: " + list + "\r\n" +
		"To: " + receiver + "\r\n" +
		"Subject: " + list + " subscription\r\n" +
		"Mime-Version: 1.0\r\n" +
		"Date: " + time.Now().UTC().Format(time.RFC1123) + "\r\n" +
		"Content-Type: text/plain\r\n" +
		"\r\n" +
		"Hi,\r\n\r\nYou are now successfully subscribed to " + list + ".\r\n\r\n" +
		"You can unsubscribe at any time by sending an email to unsubscribe+" + list + ".\r\n"
	return SendMail([]byte(msg), "bounce+"+list, receiver)
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
	return SendMail([]byte(msg), "bounce+"+list, receiver)
}

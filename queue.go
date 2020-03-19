package main

import (
	"log"
	"time"
)

// start a loop that tries to send out queued mails every 10 minutes
func StartMsgQueueRunner() {
	for {

		time.Sleep(20 * time.Minute)

		rows, err := DbCon.Query("SELECT receiver,list, data, failed_retries from msg_queue;")
		if err != nil {
			log.Print("Error getting message queue from database:")
			log.Print(err)
			continue
		}

		for rows.Next() {
			var data string
			var recv string
			var list string
			var failedRetries int
			err = rows.Scan(&recv, &list, &data, &failedRetries)
			if err != nil {
				log.Print("Error getting message queue from database:")
				log.Print(err)
				continue
			}

			// if the message failed to deliver for more than 10 times delete it from the queue
			if failedRetries > 10 {
				if _, err := DbCon.Exec("DELETE FROM msg_queue where receiver = $1 and list = $2 and data = $3", recv, list, data); err != nil {
					log.Print("Error deleting message from message queue:")
					log.Print(err)
				}
				continue
			}

			// if the mail could be sent successfully delete it from the queue
			if err := SendMail([]byte(data), list, data); err == nil {
				if _, err := DbCon.Exec("DELETE FROM msg_queue where receiver = $1 and list = $2 and data = $3", recv, list, data); err != nil {
					log.Print("Error deleting message from message queue:")
					log.Print(err)
				}
			} else {
				// increment failed retries
				if _, err := DbCon.Exec("UPDATE msg_queue SET failed_retries = failed_retries + 1 where receiver = $1 and list = $2 and data = $3", recv, list, data); err != nil {
					log.Print("Error incrementing failed retries fro message in message queue:")
					log.Print(err)
				}
			}
		}

		if err := rows.Close(); err != nil {
			log.Print(err)
		}
	}
}

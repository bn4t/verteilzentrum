package main

import (
	"log"
	"time"
)

// start a loop that tries to send out queued mails every 10 minutes
func StartMsgQueueRunner() {
	for {

		time.Sleep(10 * time.Minute)

		rows, err := DbCon.Query("SELECT receiver,list, data from msg_queue;")
		if err != nil {
			log.Print("Error getting message queue from database:")
			log.Print(err)
			continue
		}

		for rows.Next() {
			var data string
			var recv string
			var list string
			err = rows.Scan(&recv, &list, &data)
			if err != nil {
				log.Print("Error getting message queue from database:")
				log.Print(err)
				continue
			}

			// if the mail could be sent successfully delete it from the queue
			if err := ForwardMail([]byte(data), list, data); err == nil {
				_, err := DbCon.Exec("DELETE FROM msg_queue where receiver = $1 and list = $2 and data = $3", recv, list, data)
				log.Print("Error deleting message from message queue:")
				log.Print(err)
				continue
			}
		}

		if err := rows.Close(); err != nil {
			log.Print(err)
		}
	}
}

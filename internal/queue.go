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
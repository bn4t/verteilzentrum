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
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"path"
)

var DbCon *sql.DB

func InitDatabase() error {
	var err error
	DbCon, err = sql.Open("sqlite3", path.Join(Config.Verteilzentrum.DataDir+"verteilzentrum.db"))
	if err != nil {
		return err
	}

	_, err = DbCon.Exec("CREATE TABLE if not exists subscriber (id integer not null primary key autoincrement, email text not null, list text, unique(email, list));")
	if err != nil {
		return err
	}

	_, err = DbCon.Exec("CREATE TABLE if not exists msg_queue (id integer not null primary key autoincrement, receiver text not null, list text not null, data text not null, failed_retries integer not null default 1);")
	if err != nil {
		return err
	}

	return nil
}

func GetSubscribers(list string) ([]string, error) {
	var subscribers []string

	rows, err := DbCon.Query("SELECT email from subscriber where subscriber.list = $1;", list)
	if err != nil {
		return []string{}, err
	}

	defer rows.Close()
	for rows.Next() {
		var s string
		err = rows.Scan(&s)
		if err != nil {
			return []string{}, err
		}
		subscribers = append(subscribers, s)
	}

	return subscribers, nil
}

func Subscribe(email string, list string) error {

	_, err := DbCon.Exec("INSERT OR IGNORE INTO subscriber (email, list) VALUES ($1, $2);",
		email, list)
	if err != nil {
		return err
	}
	return nil
}

func Unsubscribe(email string, list string) error {

	_, err := DbCon.Exec("DELETE FROM subscriber WHERE email = $1 and list= $2;",
		email, list)
	if err != nil {
		return err
	}
	return nil
}

func AddToMsgQueue(recv string, list string, data string) error {

	_, err := DbCon.Exec("INSERT INTO msg_queue (receiver, list, data) VALUES ($1, $2, $3);",
		recv, list, data)
	if err != nil {
		return err
	}
	return nil
}

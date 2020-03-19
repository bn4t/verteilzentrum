package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"path"
)

var DbCon *sql.DB

func InitDatabase() error {
	var err error
	DbCon, err = sql.Open("sqlite3", path.Join(Config.DataDir+"verteilzentrum.db"))
	if err != nil {
		return err
	}

	_, err = DbCon.Exec("CREATE TABLE if not exists subscriber (id integer not null primary key autoincrement, email text not null, list text, unique(email, list));")
	if err != nil {
		return err
	}

	_, err = DbCon.Exec("CREATE TABLE if not exists msg_queue (id integer not null primary key autoincrement, receiver text not null, list text not null, data text not null);")
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

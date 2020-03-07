package main

import (
	"log"
)

func main() {
	if err := ReadConfig(); err != nil {
		log.Fatal(err)
	}

	if err := InitDatabase(); err != nil {
		log.Fatal(err)
	}

	InitServer()
}

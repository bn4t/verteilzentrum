package main

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strconv"
)

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// check if a provided mailing list exists
func ListExists(list string) bool {
	for _, b := range Config.Lists {
		if b.Name == list {
			return true
		}
	}
	return false
}

func GenerateMessageId(receiver string) string {
	var randnums string
	for i := 0; i < 20; i++ {
		randnums += strconv.Itoa(rand.Int())
	}

	hash := sha256.Sum256([]byte(receiver + randnums))
	return "<" + hex.EncodeToString(hash[:32]) + "@" + Config.Verteilzentrum.Hostname + ">"
}

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

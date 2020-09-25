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
	"bytes"
	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
	"verteilzentrum/internal/config"
)

// SendMail sends the specified message to the given sender using the mta credentials from the config
func SendMail(data []byte, from string, to string) error {
	var auth sasl.Client

	switch config.Config.Verteilzentrum.MtaAuthMethod {
	case "PLAIN":
		auth = sasl.NewPlainClient("", config.Config.Verteilzentrum.MtaUsername, config.Config.Verteilzentrum.MtaPassword)
		break
	case "ANONYMOUS":
		auth = sasl.NewAnonymousClient("verteilzentrum")
		break
	}

	return smtp.SendMail(config.Config.Verteilzentrum.MtaAddress, auth, from, []string{to}, bytes.NewReader(data))
}

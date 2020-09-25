package logging

import (
	"log"
	"verteilzentrum/internal/config"
)

const LogLvlErr = "ERROR"
const LogLvlInfo = "INFO"
const LogLvlDebug = "DEBUG"

func LogMsg(msg string, lvl string) {
	if config.Config.Verteilzentrum.Logging == LogLvlDebug {
		logMsg(msg, lvl)
		return
	}
	if config.Config.Verteilzentrum.Logging == LogLvlInfo && (lvl == LogLvlInfo || lvl == LogLvlErr) {
		logMsg(msg, lvl)
	} else if config.Config.Verteilzentrum.Logging == LogLvlErr && lvl == LogLvlErr {
		logMsg(msg, lvl)
	}
}

func logMsg(msg string, lvl string) {
	log.Printf("%s: %s", lvl, msg)
}
